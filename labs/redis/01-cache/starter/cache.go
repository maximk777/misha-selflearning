package cache

import (
	"sync"
	"time"
)

type entry struct {
	value string
	until time.Time
}
type Cache struct {
	mu   sync.Mutex
	data map[string]entry
}

func New() *Cache { return &Cache{data: map[string]entry{}} }
func (c *Cache) GetOrLoad(key string, ttl time.Duration, load func() string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[key]; ok && time.Now().Before(v.until) {
		return v.value
	}
	value := load()
	c.data[key] = entry{value, time.Now().Add(ttl)}
	return value
}
