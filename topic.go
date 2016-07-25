package drudsub

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/cloud/pubsub"
)

// Topic is a drudsub topic to interact with.
type Topic struct {
	Name    string
	Channel chan Message
	Client  *pubsub.Client
	Context context.Context
	Topic   *pubsub.Topic
}

// Create a topic.
func (t Topic) Create() error {
	// get topic, create it if it does not exist
	exists, err := pubsub.TopicExists(t.Context, t.Name)
	if err != nil {
		return err
	}

	if exists {
		t.Topic = t.Client.Topic(t.Name)
	} else {
		t.Topic, err = t.Client.NewTopic(t.Context, t.Name)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}

// Send to a topic.
func (t Topic) Send(m Message) error {
	return nil
}
