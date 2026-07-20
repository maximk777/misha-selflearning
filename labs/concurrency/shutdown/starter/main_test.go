package shutdown

import (
	"context"
	"sync/atomic"
	"testing"
)

func TestDrain(t *testing.T) {
	jobs := make(chan func(), 2)
	var done atomic.Int64
	jobs <- func() { done.Add(1) }
	jobs <- func() { done.Add(1) }
	close(jobs)
	if err := Drain(context.Background(), jobs); err != nil || done.Load() != 2 {
		t.Fatalf("%v %d", err, done.Load())
	}
}
