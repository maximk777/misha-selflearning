package starter

import (
	"errors"
	"testing"
)

func TestFindUserWrapsNotFound(t *testing.T) {
	_, err := FindUser("42")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("error %v does not wrap ErrNotFound", err)
	}
}

func TestMustPositivePanicsOnlyForInvalidInput(t *testing.T) {
	if got := MustPositive(2); got != 2 {
		t.Fatalf("MustPositive(2)=%d", got)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("MustPositive(0) did not panic")
		}
	}()
	MustPositive(0)
}
