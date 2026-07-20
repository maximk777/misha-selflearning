package channels

// SendOnce demonstrates ownership: producer closes the channel after its sends.
func SendOnce(values []int) <-chan int {
	out := make(chan int, len(values))
	for _, value := range values {
		out <- value
	}
	close(out)
	return out
}
