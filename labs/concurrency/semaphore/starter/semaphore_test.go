package semaphore

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunBoundsConcurrency(t *testing.T) {
	var active, max atomic.Int64
	jobs := make([]func(), 8)
	for i := range jobs {
		jobs[i] = func() {
			n := active.Add(1)
			for {
				old := max.Load()
				if n <= old || max.CompareAndSwap(old, n) {
					break
				}
			}
			time.Sleep(time.Millisecond)
			active.Add(-1)
		}
	}
	if err := Run(context.Background(), 2, jobs); err != nil || max.Load() > 2 {
		t.Fatalf("%v max=%d", err, max.Load())
	}
}
