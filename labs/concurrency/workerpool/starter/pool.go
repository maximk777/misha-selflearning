package workerpool

import "context"

func Square(ctx context.Context, workers int, jobs []int) ([]int, error) {
	jobch := make(chan int)
	out := make(chan int, len(jobs))
	done := make(chan struct{})
	for range workers {
		go func() {
			defer func() { done <- struct{}{} }()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobch:
					if !ok {
						return
					}
					out <- job * job
				}
			}
		}()
	}
	go func() {
		defer close(jobch)
		for _, job := range jobs {
			select {
			case <-ctx.Done():
				return
			case jobch <- job:
			}
		}
	}()
	for range workers {
		<-done
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	result := make([]int, 0, len(jobs))
	for range jobs {
		result = append(result, <-out)
	}
	return result, nil
}
