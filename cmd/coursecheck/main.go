package main

import (
	"fmt"
	"os"

	"github.com/maximkirienkov/misha-wzorvaniy/internal/coursecheck"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Использование: coursecheck <корень-курса>")
		os.Exit(2)
	}

	errs := coursecheck.Check(os.Args[1])
	for _, err := range errs {
		fmt.Fprintln(os.Stderr, err)
	}
	if len(errs) != 0 {
		os.Exit(1)
	}
	fmt.Println("Проверка структуры курса пройдена.")
}
