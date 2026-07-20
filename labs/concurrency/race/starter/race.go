package race

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add()       { c.mu.Lock(); c.value++; c.mu.Unlock() }
func (c *Counter) Value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.value }

// UnsafeAdd is deliberately unsynchronised and is called only by the opt-in race demo.
func UnsafeAdd(value *int) { *value++ }
