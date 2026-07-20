package task

import (
	"context"
	"fmt"
	"strings"
)

type Service struct {
	repository Repository
	publisher  Publisher
	cache      Cache
}

func NewService(repository Repository, publisher Publisher, caches ...Cache) *Service {
	service := &Service{repository: repository, publisher: publisher}
	if len(caches) > 0 {
		service.cache = caches[0]
	}
	return service
}

func (service *Service) Create(ctx context.Context, title, idempotencyKey string) (Task, bool, error) {
	if strings.TrimSpace(title) == "" {
		return Task{}, false, ErrInvalidTitle
	}
	if strings.TrimSpace(idempotencyKey) == "" {
		return Task{}, false, fmt.Errorf("idempotency key: %w", ErrInvalidTitle)
	}
	task, created, err := service.repository.Create(ctx, title, idempotencyKey)
	if err != nil || !created || service.publisher == nil {
		return task, created, err
	}
	if err := service.publisher.Publish(ctx, Event{ID: "event-" + task.ID, Type: "task.created", TaskID: task.ID}); err != nil {
		return Task{}, false, err
	}
	return task, true, nil
}

func (service *Service) Get(ctx context.Context, id string) (Task, error) {
	if service.cache != nil {
		if value, ok, err := service.cache.Get(ctx, id); err == nil && ok {
			return value, nil
		}
	}
	value, err := service.repository.Get(ctx, id)
	if err == nil && service.cache != nil {
		_ = service.cache.Set(ctx, value)
	}
	return value, err
}

func (service *Service) Complete(ctx context.Context, id string) (Task, error) {
	task, err := service.repository.Complete(ctx, id)
	if err == nil && service.cache != nil {
		_ = service.cache.Delete(ctx, id)
	}
	if err != nil || service.publisher == nil {
		return task, err
	}
	if err := service.publisher.Publish(ctx, Event{ID: "event-" + task.ID + "-completed", Type: "task.completed", TaskID: task.ID}); err != nil {
		return Task{}, err
	}
	return task, nil
}

type MemoryPublisher struct {
	events []Event
}

func (publisher *MemoryPublisher) Publish(_ context.Context, event Event) error {
	publisher.events = append(publisher.events, event)
	return nil
}

func (publisher *MemoryPublisher) Count() int { return len(publisher.events) }
