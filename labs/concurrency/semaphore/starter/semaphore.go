package semaphore

import (
	"context"
	"sync"
)

func Run(ctx context.Context, limit int, jobs []func()) error {
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	for _, job := range jobs {
		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		case sem <- struct{}{}:
		}
		wg.Add(1)
		go func(job func()) { defer wg.Done(); defer func() { <-sem }(); job() }(job)
	}
	wg.Wait()
	return ctx.Err()
}
