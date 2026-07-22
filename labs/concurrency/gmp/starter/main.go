// Observation harness for the GMP lesson.
//
// sync.WaitGroup is intentionally hidden inside the harness so every observed
// job completes. The learner is not expected to understand or modify it yet.
package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	jobs := flag.Int("jobs", 4, "number of ready jobs")
	work := flag.Int("work", 1_000_000, "CPU loop iterations per job")
	procs := flag.Int("gomaxprocs", 1, "number of scheduler Ps")
	flag.Parse()

	if *jobs < 1 || *work < 1 || *procs < 1 {
		fmt.Println("jobs, work and gomaxprocs must be positive")
		return
	}

	runtime.GOMAXPROCS(*procs)
	fmt.Printf("config jobs=%d work=%d gomaxprocs=%d\n", *jobs, *work, *procs)

	var wg sync.WaitGroup // harness detail; not part of the learner task
	wg.Add(*jobs)
	for id := 1; id <= *jobs; id++ {
		go func() {
			defer wg.Done()
			fmt.Printf("job %d start\n", id)
			value := uint64(id)
			for i := 0; i < *work; i++ {
				value = value*1664525 + 1013904223
			}
			fmt.Printf("job %d finish checksum=%d\n", id, value)
		}()
	}
	wg.Wait()
	fmt.Println("all jobs finished")
}
