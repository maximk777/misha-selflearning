package contextlab

import (
	"context"
	"testing"
	"time"
)

func TestWaitStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	started := make(chan struct{})
	done := make(chan error, 1)
	go func() { done <- Wait(ctx, started) }()
	<-started
	cancel()
	if err := <-done; err != context.Canceled {
		t.Fatalf("got %v", err)
	}
}
