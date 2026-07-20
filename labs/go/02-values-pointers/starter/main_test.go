package starter

import "testing"

func TestIncrementValueReturnsCopy(t *testing.T) {
	original := Counter{N: 3}
	changed := IncrementValue(original)
	if original.N != 3 || changed.N != 4 {
		t.Fatalf("original=%d changed=%d, want 3 and 4", original.N, changed.N)
	}
}

func TestIncrementPointerChangesOriginal(t *testing.T) {
	counter := Counter{N: 3}
	IncrementPointer(&counter)
	if counter.N != 4 {
		t.Fatalf("counter=%d, want 4", counter.N)
	}
}
