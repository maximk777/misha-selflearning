package kafka

import "context"

// Consumer demonstrates the idempotent-consumer boundary: persist event ID with the side effect.
type Consumer struct{}

func (Consumer) Consume(context.Context, string) error { return nil }
