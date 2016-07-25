package drudsub

// Message represents a Pub/Sub/NATS message.
type Message struct {
	Data       []byte
	Attributes map[string]string
}

// Send a message to a topic.
func (m Message) Send(t Topic) error {
	return nil
}
