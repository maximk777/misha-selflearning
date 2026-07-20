package kafka

import (
	"context"

	"github.com/maximkirienkov/misha-backend-lab/internal/task"
)

// Publisher is a compileable port adapter placeholder for the outbox relay checkpoint.
type Publisher struct{}

func (Publisher) Publish(context.Context, task.Event) error { return task.ErrNotConfigured }
