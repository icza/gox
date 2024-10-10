package gox

import (
	"context"
	"sync"
	"time"
)

// Evictable defines a single Evict() method.
// [OpCache] has Evict().
type Evictable interface {
	Evict()
}

// RunEvictor should be run as a goroutine, it evicts expired cache entries from the listed OpCaches.
// Returns only if ctx is cancelled.
//
// [OpCache] has Evict() method, so any OpCache can be listed (does not depend on the type parameter).
func RunEvictor(ctx context.Context, evictorPeriod time.Duration, opCaches ...Evictable) {
	ticker := time.NewTicker(evictorPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}

		for _, oc := range opCaches {
			oc.Evict()
		}
	}
}

type evictableItem struct {
	opCache        Evictable
	evictionPeriod time.Duration
	nextEvictAt    time.Time
}

var (
	globalEvictorMu      sync.Mutex
	globalEvictableItems []*evictableItem
)

// addToGlobalEvictor adds the given opCache to the global evictor.
// The global evictor is started on demand and is never stopped.
func addToGlobalEvictor(opCache Evictable, evictionPeriodMinutes int) {
	evictionPeriod := time.Duration(evictionPeriodMinutes)*time.Minute - 5*time.Second // -5 sec to make sure we don't skip an eviction due to imprecise timing
	item := &evictableItem{
		opCache:        opCache,
		evictionPeriod: evictionPeriod,
		nextEvictAt:    time.Now().Add(evictionPeriod),
	}

	globalEvictorMu.Lock()
	defer globalEvictorMu.Unlock()

	if len(globalEvictableItems) == 0 {
		// This is the first evictable opCache, launch global evictor:
		go func() {
			for now := range time.Tick(time.Minute) { // Every minute
				globalEvictorMu.Lock()
				var items []*evictableItem
				for _, item := range globalEvictableItems {
					if now.After(item.nextEvictAt) {
						items = append(items, item)
						item.nextEvictAt = now.Add(item.evictionPeriod)
					}
				}
				globalEvictorMu.Unlock()

				for _, item := range items {
					item.opCache.Evict()
				}
			}
		}()
	}

	globalEvictableItems = append(globalEvictableItems, item)
}
