package main

import "fmt"

func main() {
	name := "Миша"
	for attempt := 1; attempt <= 2; attempt++ {
		fmt.Printf("%s, запуск №%d\n", name, attempt)
	}
}
