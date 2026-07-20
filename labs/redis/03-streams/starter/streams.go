package streams

type Entry struct{ ID, Body string }

// Ack is intentionally explicit: unacked entries are pending and can be retried.
func Ack(pending map[string]Entry, id string) { delete(pending, id) }
