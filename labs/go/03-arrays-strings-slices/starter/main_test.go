package starter

import "testing"

func TestFirstRuneDoesNotReturnFirstUTF8Byte(t *testing.T) {
	if got := FirstRune("Жора"); got != 'Ж' {
		t.Fatalf("FirstRune()=%q, want %q", got, 'Ж')
	}
}

func TestAppendMarkerCanReuseBackingArray(t *testing.T) {
	base := make([]int, 1, 2)
	base[0] = 1
	withMarker := AppendMarker(base, 9)
	if base[:2][1] != 9 || withMarker[1] != 9 {
		t.Fatalf("base=%v result=%v, want shared marker", base[:2], withMarker)
	}
}
