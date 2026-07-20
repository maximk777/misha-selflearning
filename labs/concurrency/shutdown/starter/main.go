package shutdown

import (
	"context"
	"sync"
)

func Drain(ctx context.Context, jobs <-chan func()) error {
	var wg sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		case job, ok := <-jobs:
			if !ok {
				wg.Wait()
				return nil
			}
			wg.Add(1)
			go func() { defer wg.Done(); job() }()
		}
	}
}
