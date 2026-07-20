package idempotency

import "testing"

func TestApplyOnce(t *testing.T) {
	s := New()
	if !s.Apply("order-1") || s.Apply("order-1") {
		t.Fatal("duplicate applied")
	}
}
