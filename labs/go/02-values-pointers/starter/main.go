package starter

type Counter struct {
	N int
}

func IncrementValue(counter Counter) Counter {
	counter.N++
	return counter
}

func IncrementPointer(counter *Counter) {
	counter.N++
}
