package stringsx

import "sync"

// Pool implements a string pool safe for concurrent use.
//
// For details, see https://stackoverflow.com/a/51983331/1705598
type Pool struct {
	mu    sync.Mutex
	cache map[string]string
}

// NewPool returns a new, empty string pool.
func NewPool() *Pool {
	return &Pool{
		cache: map[string]string{},
	}
}

// Interned returns the string instance from the pool that is equal to s.
// If s is not yet in the pool, it is put into it and returned.
func (p *Pool) Interned(s string) string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if s2, ok := p.cache[s]; ok {
		return s2
	}

	// New string, store it
	p.cache[s] = s
	return s
}
