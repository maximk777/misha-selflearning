package race

import (
	"os"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() { defer wg.Done(); c.Add() }()
	}
	wg.Wait()
	if c.Value() != 100 {
		t.Fatal(c.Value())
	}
}
func TestDeliberateRace(t *testing.T) {
	if os.Getenv("RACE_DEMO") == "" {
		t.Skip("set RACE_DEMO=1 for the bounded failure demo")
	}
	var value int
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() { defer wg.Done(); UnsafeAdd(&value) }()
	}
	wg.Wait()
}
