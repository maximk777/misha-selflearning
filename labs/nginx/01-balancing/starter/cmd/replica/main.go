package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("REPLICA_NAME")
	if name == "" {
		name = "local"
	}
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, name) })
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
