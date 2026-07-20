package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	deadlock := flag.Bool("deadlock", false, "demonstrate opposite lock ordering")
	flag.Parse()
	if !*deadlock {
		fmt.Println("ok: use -deadlock only for the documented bounded demo")
		return
	}
	var left, right sync.Mutex
	ready := make(chan struct{}, 2)
	var wg sync.WaitGroup
	lock := func(first, second *sync.Mutex) {
		defer wg.Done()
		first.Lock()
		ready <- struct{}{}
		<-ready
		second.Lock()
	}
	wg.Add(2)
	go lock(&left, &right)
	go lock(&right, &left)
	wg.Wait()
}
