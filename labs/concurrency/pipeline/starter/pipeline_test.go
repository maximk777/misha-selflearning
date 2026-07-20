package pipeline

import (
	"context"
	"testing"
)

func TestDouble(t *testing.T) {
	got := []int{}
	for value := range Double(context.Background(), []int{1, 2, 3}) {
		got = append(got, value)
	}
	if len(got) != 3 || got[2] != 6 {
		t.Fatal(got)
	}
}
