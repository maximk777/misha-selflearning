package workerpool

import (
	"context"
	"testing"
)

func TestSquare(t *testing.T) {
	got, err := Square(context.Background(), 2, []int{2, 3, 4})
	if err != nil || len(got) != 3 {
		t.Fatalf("%v %v", got, err)
	}
}
func TestSquareCancels(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Square(ctx, 2, []int{1})
	if err != context.Canceled {
		t.Fatal(err)
	}
}
