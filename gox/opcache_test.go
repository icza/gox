package gox

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestOpCache(t *testing.T) {
	expiration := 10 * time.Millisecond
	cfg := OpCacheConfig{
		ResultExpiration:      expiration,
		ResultGraceExpiration: expiration,
	}

	opc := NewOpCache[string, int](cfg)

	resultCh := make(chan int)
	go func() {
		for i := 1; i < 5; i++ {
			resultCh <- i
		}
	}()

	operation := func() (out int, err error) {
		return <-resultCh, nil
	}

	cases := []struct {
		name      string
		delay     time.Duration
		key       string
		result    int
		resultErr error
	}{
		{"0: not cached, operation() called", 0, "1", 1, nil},
		{"1: cached, valid", 0, "1", 1, nil},
		{"2: not cached, operation()  called", 0, "2", 2, nil},
		{"3: cached, valid", 0, "2", 2, nil},
		{"4: grace-valid, operation() called in background", 3 * expiration / 2, "1", 1, nil},
		{"5: value loaded in background", expiration / 2, "1", 3, nil},
		{"6: invalid, operation() called", 2 * expiration, "1", 4, nil},
	}

	for _, c := range cases {
		time.Sleep(c.delay)

		result, err := opc.Get(
			c.key,
			func() (result int, err error) {
				return operation()
			},
		)

		if result != c.result || err != c.resultErr {
			t.Errorf("[%s] Expected (%v, %v), got (%v, %v)", c.name, c.result, c.resultErr, result, err)
		}
	}
}

func TestOpCacheError(t *testing.T) {
	var (
		errNotSpecial = errors.New("err-not-special")
		errToDiscard  = errors.New("err-to-discard")
		errShortCache = errors.New("err-short-cache")
	)

	var (
		expiration = 10 * time.Millisecond
		expShort   = expiration / 2
	)

	cfg := OpCacheConfig{
		ResultExpiration:      expiration,
		ResultGraceExpiration: expiration,
		ErrorExpiration: func(err error) (discard bool, expiration, graceExpiration *time.Duration) {
			if errors.Is(err, errToDiscard) {
				return true, nil, nil
			}
			if errors.Is(err, errShortCache) {
				return false, &expShort, &expShort
			}
			return
		},
	}

	opc := NewOpCache[string, int](cfg)

	resultCh := make(chan int)
	go func() {
		for i := 1; i < 9; i++ {
			resultCh <- i
		}
	}()
	resultErrCh := make(chan error)
	go func() {
		for i := 1; i < 9; i++ {
			switch i {
			case 2, 3:
				resultErrCh <- errNotSpecial
			case 4, 5:
				resultErrCh <- errToDiscard
			case 6, 7, 8:
				resultErrCh <- errShortCache
			default:
				resultErrCh <- nil
			}
		}
	}()

	operation := func() (out int, err error) {
		return <-resultCh, <-resultErrCh
	}

	cases := []struct {
		name      string
		delay     time.Duration
		key       string
		result    int
		resultErr error
	}{
		{"0: not cached, operation() called", 0, "1", 1, nil},
		{"1: cached, valid", 0, "1", 1, nil},
		{"2: not cached, operation() called", 0, "2", 2, errNotSpecial},
		{"3: cached, valid", 0, "2", 2, errNotSpecial},
		{"4: invalid, operation() called", 3 * expiration, "1", 3, errNotSpecial},
		{"5: not cached, operation() called", 0, "3", 4, errToDiscard},
		{"6: not cached, operation() called", 0, "3", 5, errToDiscard},
		{"7: not cached, operation() called", 0, "4", 6, errShortCache},
		{"8: cached, valid", 0, "4", 6, errShortCache},
		{"9: grace-valid, operation() called in background", 3 * expShort / 2, "4", 6, errShortCache},
		{"10: value loaded in background", expShort / 2, "4", 7, errShortCache},
		{"11: invalid, operation() called", 3 * expShort / 2, "4", 8, errShortCache},
	}

	for _, c := range cases {
		time.Sleep(c.delay)

		result, err := opc.Get(
			c.key,
			func() (result int, err error) {
				return operation()
			},
		)

		if result != c.result || err != c.resultErr {
			t.Errorf("[%s] Expected (%v, %v), got (%v, %v)", c.name, c.result, c.resultErr, result, err)
		}
	}
}

func TestOpCacheExecOpOnce(t *testing.T) {
	expiration := 10 * time.Millisecond

	cfg := OpCacheConfig{
		ResultExpiration:      expiration,
		ResultGraceExpiration: expiration,
	}

	var counter int64
	operation := func() (out int, err error) {
		time.Sleep(expiration / 2)
		return int(atomic.AddInt64(&counter, 1)), nil
	}

	opc := NewOpCache[string, int](cfg)

	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ { // 5 concurrent OpCache.Get() calls
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			got, _ := opc.Get("1", operation)
			if exp := 1; got != exp {
				t.Errorf("[%d] Expected: %v, got: %v", i, exp, got)
			}
		}()
	}
	wg.Wait()

	time.Sleep(cfg.ResultExpiration + cfg.ResultGraceExpiration)

	for i := 0; i < 5; i++ { // 5 concurrent OpCache.Get() calls
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			got, _ := opc.Get("1", operation)
			if exp := 2; got != exp {
				t.Errorf("[%d] Expected: %v, got: %v", i, exp, got)
			}
		}()
	}
	wg.Wait()
}
