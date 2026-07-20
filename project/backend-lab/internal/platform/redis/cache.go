package redis

import (
	"context"
	"sync"
	"time"

	"github.com/maximkirienkov/misha-backend-lab/internal/task"
)

// Cache is an in-memory stand-in for Redis. A real Redis outage must degrade to repository reads.
type Cache struct {
	mu    sync.Mutex
	items map[string]entry
	ttl   time.Duration
}
type entry struct {
	task    task.Task
	expires time.Time
}

func NewCache(ttl time.Duration) *Cache { return &Cache{items: make(map[string]entry), ttl: ttl} }
func (cache *Cache) Get(_ context.Context, id string) (task.Task, bool, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	item, ok := cache.items[id]
	if !ok || time.Now().After(item.expires) {
		return task.Task{}, false, nil
	}
	return item.task, true, nil
}
func (cache *Cache) Set(_ context.Context, value task.Task) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.items[value.ID] = entry{task: value, expires: time.Now().Add(cache.ttl)}
	return nil
}
func (cache *Cache) Delete(_ context.Context, id string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	delete(cache.items, id)
	return nil
}
