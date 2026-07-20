package pubsub

// Message models an ephemeral Pub/Sub event: delivery requires a live subscriber.
type Message struct{ Topic, Body string }
