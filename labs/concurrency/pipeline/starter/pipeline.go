package pipeline

import "context"

func Double(ctx context.Context, values []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			select {
			case <-ctx.Done():
				return
			case out <- value * 2:
			}
		}
	}()
	return out
}
