package cache

import (
	"testing"
	"time"
)

func TestCacheAside(t *testing.T) {
	c := New()
	calls := 0
	load := func() string { calls++; return "db" }
	c.GetOrLoad("a", time.Hour, load)
	c.GetOrLoad("a", time.Hour, load)
	if calls != 1 {
		t.Fatal(calls)
	}
}
func TestTTL(t *testing.T) {
	c := New()
	calls := 0
	load := func() string { calls++; return "db" }
	c.GetOrLoad("a", time.Nanosecond, load)
	time.Sleep(time.Millisecond)
	c.GetOrLoad("a", time.Nanosecond, load)
	if calls != 2 {
		t.Fatal(calls)
	}
}
