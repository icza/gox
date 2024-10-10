package gox

import (
	"sync"
	"time"
)

const DefaultEvictPeriodMinutes = 15

// OpCacheConfig holds configuration options for an [OpCache].
type OpCacheConfig struct {
	// Operation results are valid for this long after creation.
	ResultExpiration time.Duration

	// Expired results are still usable for this long after expiration.
	// Tip: if this field is 0, grace period and thus background
	// op execution will be disabled.
	ResultGraceExpiration time.Duration

	// ErrorExpiration is an optional function.
	// If provided, it will be called for non-nil operation errors.
	// Return discard=true if you do not want to cache an error result.
	// If expiration or graceExpiration is provided (non-nil), they will
	// override the cache expiration for the given error result
	//
	// If provided, this function is only called once for the result error of a single operation execution
	// (regardless of how many times it is accessed from the OpCache).
	ErrorExpiration func(err error) (discard bool, expiration, graceExpiration *time.Duration)

	// AutoEvictPeriodMinutes tells how frequently should expired entries be checked and evicted from the cache.
	// If 0, DefaultEvictPeriodMinutes will be used. Removal is currently not supported.
	//
	// If a negative value is given, the op cache is not added to the internal auto-evictor, and manual eviction
	// should be taken care of with e.g. using the RunEvictor() function.
	AutoEvictPeriodMinutes int
}

// OpCache implements a general value cache. It can be used to cache results of arbitrary operations.
//
// Cached values are tied to a key that should be derived from the operation's arguments.
// If the operation has multiple arguments, a wrapper struct is ideal (such as [Struct2], [Struct3] etc.),
// or fmt.Sprint() will also do as an alternative (with string being the key type).
//
// Cached values have an expiration time and also a grace period during which the cached value
// is considered usable, but getting a cached value during the grace period triggers a reload
// that will happen in the background (the cached value is returned immediately, without waiting).
//
// Operations are captured by a function that returns a value of a certain type (T) and an error.
// If an operation has multiple results beside the error, they must be wrapped in a composite type
// (like a struct or slice).
//
// If multiple input arguments are available (for multiple operation execution),
// operations can often be executed more efficiently if all inputs are handed as a batch
// than executing the operation for each input argument individually.
// A tipical example is loading records by ID from a database: running a query with a condition like "id=?"
// for each ID individually can be significantly slower than running a single query with a condition like
// "id in ?". These operations can take advantage of the [OpCache.MultiGet] method. MultiGet will ensure that
// only the minimal required subset of the arguments is passed in the multi-operation execution
// if some of them are already cached, and [OpCache.Get] methods will also take advantage of entries cached by MultiGet.
type OpCache[K comparable, T any] struct {
	cfg OpCacheConfig

	keyResultsMu sync.RWMutex
	keyResults   map[K]*opResult[T]
}

// NewOpCache creates a new OpCache.
func NewOpCache[K comparable, T any](cfg OpCacheConfig) *OpCache[K, T] {
	opCache := &OpCache[K, T]{
		cfg:        cfg,
		keyResults: map[K]*opResult[T]{},
	}

	if cfg.AutoEvictPeriodMinutes >= 0 {
		epMins := cfg.AutoEvictPeriodMinutes
		if epMins == 0 {
			epMins = DefaultEvictPeriodMinutes
		}
		addToGlobalEvictor(opCache, epMins)
	}

	return opCache
}

func (oc *OpCache[K, T]) getCachedOpResult(key K) *opResult[T] {
	oc.keyResultsMu.RLock()
	defer oc.keyResultsMu.RUnlock()

	return oc.keyResults[key]
}

func (oc *OpCache[K, T]) setCachedOpResult(key K, opResults *opResult[T]) {
	oc.keyResultsMu.Lock()
	oc.keyResults[key] = opResults
	oc.keyResultsMu.Unlock()
}

// Evict checks all cached entries, and removes invalid ones.
func (oc *OpCache[K, T]) Evict() {
	oc.keyResultsMu.Lock()
	defer oc.keyResultsMu.Unlock()

	for key, opResult := range oc.keyResults {
		if !opResult.graceValid() { // Delete if not even grace-valid
			delete(oc.keyResults, key)
		}
	}
}

// Get gets the result of an operation.
//
// If the result is cached and valid, it is returned immediately.
//
// If the result is cached but not valid, but we're within the grace period,
// execOp() is called in the background to refresh the cache,
// and the cached result is returned immediately.
// Care is taken to only launch a single background worker to refresh the cache even if
// Get() or [OpCache.MultiGet] is called multiple times with the same key before the cache can be refreshed.
//
// Else result is either not cached or we're past even the grace period:
// execOp() is executed, the function waits for its return values, the result is cached,
// and then the fresh result is returned.
func (oc *OpCache[K, T]) Get(
	key K,
	execOp func() (result T, err error),
) (result T, resultErr error) {

	cachedResult := oc.getCachedOpResult(key)

	if cachedResult.valid() {
		return cachedResult.result, cachedResult.resultErr
	}

	// execOpAndCache executes execOp(), caches the result according to the configuration, and returns it
	execOpAndCache := func() (result T, resultErr error) {
		result, resultErr = execOp()
		expiration, graceExpiration := oc.cfg.ResultExpiration, oc.cfg.ResultGraceExpiration
		if resultErr != nil && oc.cfg.ErrorExpiration != nil {
			discard, exp, graceExp := oc.cfg.ErrorExpiration(resultErr)
			if discard {
				// This error result is not to be cached at all, just return:
				return
			}
			if exp != nil {
				expiration = *exp
			}
			if graceExp != nil {
				graceExpiration = *graceExp
			}
		}
		oc.setCachedOpResult(key, newOpResult(result, resultErr, expiration, graceExpiration))
		return
	}

	if !cachedResult.graceValid() {
		// Not valid and not even within grace period: query, cache and return:
		return execOpAndCache()
	}

	// Cached result is within grace period, we can use it:
	result, resultErr = cachedResult.result, cachedResult.resultErr

	// But need to reload, in the background.
	// First use read-lock to check if someone's already doing it:

	cachedResult.reloadMu.RLock()
	reloading := cachedResult.reloading
	cachedResult.reloadMu.RUnlock()
	if reloading {
		// Already reloading, nothing to do
		return
	}

	// Try to take ownership of reloading, needs write-lock:
	cachedResult.reloadMu.Lock()
	if cachedResult.reloading {
		// Someone else got the write-lock first, he'll take care of the reload
		cachedResult.reloadMu.Unlock()
		return
	}
	cachedResult.reloading = true // We'll be the one to do it
	cachedResult.reloadMu.Unlock()

	// reload in new goroutine.
	// Note: we're not using the return values, we're returning the cached (grace-valid) values.
	go execOpAndCache()

	return
}

// MultiGet gets the results of a multi-operation.
// A multi-operation is an operation that accepts a slice of keys, and can produce results for multiple input parameters
// more efficiently than calling the operation for each input separately.
//
// results and resultErrs will be slices with identical size and elements matching to that of keys.
//
// Each result is taken from the cache if present and valid, or we're within its grace period.
// If there are entries that are either not cached or we're past their grace period,
// execMultiOp() is executed for those keys, the function waits for its return values, the results are cached,
// and the fresh results are returned.
//
// If there are results that are returned because they are cached but not valid but we're within the grace period,
// execMultiOp() is called in the background to refresh them. Care is taken to only launch a single background worker
// to refresh each such entry even if [OpCache.Get] or MultiGet() is called multiple times with the same key(s)
// before the cache can be refreshed.
//
// execMultiOp must return results and errs slices with identical size to that of its keyIndices argument,
// and elements matching to keys designated by keyIndices! Failure to do so is undefined behavior,
// may even result in runtime panic!
//
// Tip: [github.com/icza/gog/slicesx.SelectByIndices] may come handy when implementing execMultiOp.
func (oc *OpCache[K, T]) MultiGet(
	keys []K,
	execMultiOp func(keyIndices []int) (results []T, errs []error),
) (results []T, resultErrs []error) {

	results = make([]T, len(keys))
	resultErrs = make([]error, len(keys))
	cachedResults := make([]*opResult[T], len(keys))

	var (
		invalidKeyIndices    []int // key indices that we must produce and wait for
		graceValidKeyIndices []int // key indices that we may use but must refresh in the background
	)

	for keyIdx, key := range keys {
		cachedResult := oc.getCachedOpResult(key)

		switch {
		case cachedResult.valid():
			results[keyIdx], resultErrs[keyIdx] = cachedResult.result, cachedResult.resultErr
		case cachedResult.graceValid():
			// Cached result is within grace period, we can use it:
			results[keyIdx], resultErrs[keyIdx] = cachedResult.result, cachedResult.resultErr
			graceValidKeyIndices = append(graceValidKeyIndices, keyIdx)
			cachedResults[keyIdx] = cachedResult
		default:
			// Not valid and not even within grace period: query, cache and return:
			invalidKeyIndices = append(invalidKeyIndices, keyIdx)
		}
	}

	// execMultiOpAndCache executes execMultiOp(), caches the results according to the configuration, and returns them
	execMultiOpAndCache := func(keyIndices []int) (results []T, resultErrs []error) {
		results, resultErrs = execMultiOp(keyIndices)
		for i, resultErr := range resultErrs {
			expiration, graceExpiration := oc.cfg.ResultExpiration, oc.cfg.ResultGraceExpiration
			if resultErr != nil && oc.cfg.ErrorExpiration != nil {
				discard, exp, graceExp := oc.cfg.ErrorExpiration(resultErr)
				if discard {
					// This error result is not to be cached at all, just skip:
					continue
				}
				if exp != nil {
					expiration = *exp
				}
				if graceExp != nil {
					graceExpiration = *graceExp
				}
			}
			oc.setCachedOpResult(keys[keyIndices[i]], newOpResult(results[i], resultErr, expiration, graceExpiration))
		}
		return
	}

	if len(invalidKeyIndices) > 0 {
		// Call execMultiOpAndCache and wait for its results!
		mresults, mresultErrs := execMultiOpAndCache(invalidKeyIndices)
		for i, result := range mresults {
			keyIdx := invalidKeyIndices[i]
			results[keyIdx], resultErrs[keyIdx] = result, mresultErrs[i]
		}
	}

	if len(graceValidKeyIndices) > 0 {
		// Launch background goroutine in which call execMultiOp (we're not waiting for its results)!

		// First let's see which elements we do need to process, and if we're the one to do it:
		graceValidKeyIndices2 := make([]int, 0, len(graceValidKeyIndices))
		for _, keyIdx := range graceValidKeyIndices {
			cachedResult := cachedResults[keyIdx]

			// First use read-lock to check if someone's already doing it:

			cachedResult.reloadMu.RLock()
			reloading := cachedResult.reloading
			cachedResult.reloadMu.RUnlock()
			if reloading {
				// Already reloading, nothing to do
				continue
			}

			// Try to take ownership of reloading, needs write-lock:
			cachedResult.reloadMu.Lock()
			if cachedResult.reloading {
				// Someone else got the write-lock first, he'll take care of the reload
				cachedResult.reloadMu.Unlock()
				continue
			}
			cachedResult.reloading = true // We'll be the one to do it
			cachedResult.reloadMu.Unlock()
			graceValidKeyIndices2 = append(graceValidKeyIndices2, keyIdx)
		}
		if len(graceValidKeyIndices2) > 0 {
			// reload in new goroutine.
			// Note: we're not using the return values, we're returning the cached (grace-valid) values.
			go execMultiOpAndCache(graceValidKeyIndices2)
		}
	}

	return
}

// opResult holds the result of an operation.
type opResult[T any] struct {
	expiresAt, graceExpiresAt time.Time

	result    T // If an op has multiple results, this should be a slice (e.g. []any)
	resultErr error

	reloadMu  sync.RWMutex
	reloading bool
}

// newOpResult creates a new OpResult.
func newOpResult[T any](result T, resultErr error, expiration, graceExpiration time.Duration) *opResult[T] {
	now := time.Now()
	return &opResult[T]{
		expiresAt:      now.Add(expiration),
		graceExpiresAt: now.Add(expiration + graceExpiration),
		result:         result,
		resultErr:      resultErr,
	}
}

// valid tells if the result is valid.
func (opr *opResult[T]) valid() bool {
	return opr != nil && time.Now().Before(opr.expiresAt)
}

// graceValid tells if the result is "grace-valid" (valid within the grace expiration beyond the normal expiration).
func (opr *opResult[T]) graceValid() bool {
	return opr != nil && time.Now().Before(opr.graceExpiresAt)
}
