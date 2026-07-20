package main

import (
	"fmt"
	"runtime"
)

type payload struct {
	bytes [1024]byte
}

func makePayload() *payload {
	return &payload{}
}

func main() {
	var before, after runtime.MemStats
	runtime.ReadMemStats(&before)
	_ = makePayload()
	runtime.GC()
	runtime.ReadMemStats(&after)
	fmt.Printf("alloc before=%d after=%d\n", before.Alloc, after.Alloc)
}
