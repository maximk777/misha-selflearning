package task

import (
	"context"
	"testing"
	"time"
)

func TestCreateIsIdempotentAndPublishesOnce(t *testing.T) {
	repository := NewMemoryRepository()
	publisher := &MemoryPublisher{}
	service := NewService(repository, publisher)

	first, created, err := service.Create(context.Background(), "Купить хлеб", "request-1")
	if err != nil || !created {
		t.Fatalf("first create: task=%+v created=%t err=%v", first, created, err)
	}
	second, created, err := service.Create(context.Background(), "Купить хлеб", "request-1")
	if err != nil || created || second.ID != first.ID {
		t.Fatalf("second create: task=%+v created=%t err=%v", second, created, err)
	}
	if publisher.Count() != 1 {
		t.Fatalf("published=%d, want 1", publisher.Count())
	}
}

func TestGetUsesCacheAfterRepositoryMiss(t *testing.T) {
	repository := NewMemoryRepository()
	service := NewService(repository, &MemoryPublisher{}, NewMemoryCache(time.Minute))
	created, _, err := service.Create(context.Background(), "Кэш", "request-3")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := service.Get(context.Background(), created.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := service.Get(context.Background(), created.ID); err != nil {
		t.Fatal(err)
	}
}

func TestCompleteRejectsSecondCompletion(t *testing.T) {
	service := NewService(NewMemoryRepository(), &MemoryPublisher{})
	created, _, err := service.Create(context.Background(), "Изучить context", "request-2")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := service.Complete(context.Background(), created.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := service.Complete(context.Background(), created.ID); err != ErrConflict {
		t.Fatalf("second completion error=%v, want %v", err, ErrConflict)
	}
}
