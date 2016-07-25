package drudsub

import "google.golang.org/cloud/pubsub"

// Topic is a drudsub topic to interact with.
type Topic struct {
	Name       string
	Connection Connection
	Topic      *pubsub.Topic
}

// Create a topic.
func (t Topic) Create() error {
	// get topic, create it if it does not exist
	exists, err := pubsub.TopicExists(t.Connection.Context, t.Name)
	if err != nil {
		return err
	}

	if exists {
		t.Topic = t.Connection.Client.Topic(t.Name)
	} else {
		t.Topic, err = t.Connection.Client.NewTopic(t.Connection.Context, t.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

// Publish messages to a topic.
func (t Topic) Publish(m []Message) ([]string, error) {
	var messages []*pubsub.Message

	for _, v := range m {
		messages = append(messages, &pubsub.Message{
			Data:       v.Data,
			Attributes: v.Attributes,
		})
	}
	return t.Topic.Publish(t.Connection.Context, messages...)
}
