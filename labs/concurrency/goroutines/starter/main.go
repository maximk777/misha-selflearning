package main

import (
	"flag"
	"fmt"
)

func work(iterations int) {
	fmt.Println("work: started")
	value := uint64(1)
	for i := 0; i < iterations; i++ {
		value = value*1664525 + 1013904223
	}
	fmt.Printf("work: finished checksum=%d\n", value)
}

func main() {
	mode := flag.String("mode", "sync", "sync or async")
	iterations := flag.Int("work", 10_000_000, "CPU loop iterations")
	flag.Parse()

	switch *mode {
	case "sync":
		work(*iterations)
	case "async":
		go work(*iterations)
	default:
		fmt.Println("mode must be sync or async")
		return
	}

	fmt.Println("main: returning")
}
