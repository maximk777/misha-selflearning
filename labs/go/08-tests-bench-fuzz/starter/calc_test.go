package starter

import "testing"

func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, 1, -1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Add(tc.a, tc.b); got != tc.want {
				t.Fatalf("Add(%d, %d)=%d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func FuzzAddZero(f *testing.F) {
	f.Add(0)
	f.Add(7)
	f.Fuzz(func(t *testing.T, value int) {
		if got := Add(value, 0); got != value {
			t.Fatalf("Add(%d, 0)=%d", value, got)
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = Add(40, 2)
	}
}
