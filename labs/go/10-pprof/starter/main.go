package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/work", func(writer http.ResponseWriter, request *http.Request) {
		result := 0
		for number := 0; number < 100_000; number++ {
			result += number
		}
		fmt.Fprintf(writer, "sum=%d\n", result)
	})

	fmt.Println("pprof: http://127.0.0.1:6060/debug/pprof/")
	if err := http.ListenAndServe("127.0.0.1:6060", nil); err != nil {
		panic(err)
	}
}
