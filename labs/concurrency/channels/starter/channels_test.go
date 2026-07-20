package channels

import "testing"

func TestSendOnce(t *testing.T) {
	got := []int{}
	for value := range SendOnce([]int{1, 2, 3}) {
		got = append(got, value)
	}
	if len(got) != 3 || got[2] != 3 {
		t.Fatalf("got %v", got)
	}
}
