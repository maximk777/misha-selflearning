package postgres

import (
	"context"

	"github.com/maximkirienkov/misha-backend-lab/internal/task"
)

// Repository is a teaching stub: wire a pgx pool here after Compose is healthy.
// PostgreSQL remains the source of truth; unit tests intentionally use MemoryRepository.
type Repository struct{}

func (Repository) Create(context.Context, string, string) (task.Task, bool, error) {
	return task.Task{}, false, task.ErrNotConfigured
}
func (Repository) Get(context.Context, string) (task.Task, error) {
	return task.Task{}, task.ErrNotConfigured
}
func (Repository) Complete(context.Context, string) (task.Task, error) {
	return task.Task{}, task.ErrNotConfigured
}
