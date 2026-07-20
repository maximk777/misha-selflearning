package hashring

import "testing"

func TestDeterministic(t *testing.T) {
	r := New([]string{"a", "b", "c"})
	if r.Node("order-1") != r.Node("order-1") {
		t.Fatal("unstable")
	}
}
func TestEmpty(t *testing.T) {
	if New(nil).Node("x") != "" {
		t.Fatal("want empty")
	}
}
