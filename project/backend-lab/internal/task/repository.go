package task

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Repository interface {
	Create(context.Context, string, string) (Task, bool, error)
	Get(context.Context, string) (Task, error)
	Complete(context.Context, string) (Task, error)
}

type Publisher interface {
	Publish(context.Context, Event) error
}

type Cache interface {
	Get(context.Context, string) (Task, bool, error)
	Set(context.Context, Task) error
	Delete(context.Context, string) error
}

type MemoryRepository struct {
	mu          sync.Mutex
	tasks       map[string]Task
	requestToID map[string]string
	next        int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{tasks: make(map[string]Task), requestToID: make(map[string]string)}
}

func (repository *MemoryRepository) Create(_ context.Context, title, idempotencyKey string) (Task, bool, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	if id, ok := repository.requestToID[idempotencyKey]; ok {
		return repository.tasks[id], false, nil
	}
	repository.next++
	task := Task{ID: fmt.Sprintf("task-%d", repository.next), Title: title, Status: StatusOpen, CreatedAt: time.Now().UTC()}
	repository.tasks[task.ID] = task
	repository.requestToID[idempotencyKey] = task.ID
	return task, true, nil
}

func (repository *MemoryRepository) Get(_ context.Context, id string) (Task, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	task, ok := repository.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	return task, nil
}

func (repository *MemoryRepository) Complete(_ context.Context, id string) (Task, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	task, ok := repository.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	if task.Status == StatusCompleted {
		return Task{}, ErrConflict
	}
	task.Status = StatusCompleted
	repository.tasks[id] = task
	return task, nil
}

type MemoryCache struct {
	mu    sync.Mutex
	items map[string]memoryCacheEntry
	ttl   time.Duration
}
type memoryCacheEntry struct {
	task    Task
	expires time.Time
}

func NewMemoryCache(ttl time.Duration) *MemoryCache {
	return &MemoryCache{items: make(map[string]memoryCacheEntry), ttl: ttl}
}
func (cache *MemoryCache) Get(_ context.Context, id string) (Task, bool, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	item, ok := cache.items[id]
	if !ok || time.Now().After(item.expires) {
		return Task{}, false, nil
	}
	return item.task, true, nil
}
func (cache *MemoryCache) Set(_ context.Context, value Task) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.items[value.ID] = memoryCacheEntry{task: value, expires: time.Now().Add(cache.ttl)}
	return nil
}
func (cache *MemoryCache) Delete(_ context.Context, id string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	delete(cache.items, id)
	return nil
}
