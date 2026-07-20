package starter

import "testing"

func TestLookupDistinguishesMissingZero(t *testing.T) {
	value, ok := Lookup(map[string]int{"zero": 0}, "zero")
	if value != 0 || !ok {
		t.Fatalf("existing zero: value=%d ok=%t", value, ok)
	}
	_, ok = Lookup(map[string]int{}, "missing")
	if ok {
		t.Fatal("missing key reported present")
	}
}

func TestNewScoresCanBeWritten(t *testing.T) {
	scores := NewScores()
	scores["Misha"] = 10
	if scores["Misha"] != 10 {
		t.Fatal("map does not accept write")
	}
}
