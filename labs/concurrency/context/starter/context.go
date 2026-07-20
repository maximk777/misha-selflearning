package contextlab

import "context"

func Wait(ctx context.Context, started chan<- struct{}) error {
	close(started)
	<-ctx.Done()
	return ctx.Err()
}
