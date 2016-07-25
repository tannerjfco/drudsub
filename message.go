package drudsub

// Message represents a Pub/Sub/NATS message.
type Message struct {
	Data       []byte
	Attributes map[string]string
}
