package drudsub

import (
	"log"

	"google.golang.org/cloud/pubsub"
)

type Subscription struct {
	Name         string
	Topic        Topic
	Subscription *pubsub.Subscription
	Connection   Connection
}

// Subscribe to a topic.
func (s Subscription) Subscribe(topic string, create bool) (<-chan Message, error) {
	if create {
		s.Topic.Create()
	}

	exists, err := pubsub.SubExists(s.Topic.Context, s.Name)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		s.Subscription = s.Topic.Client.Subscription(s.Name)
	} else {
		s.Subscription, err = s.Topic.Client.NewSubscription(s.Topic.Context, s.Name, s.Topic.Topic, 0, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return s.Topic.Channel, nil
}

// Create a new subscription
func (s Subscription) Create() error {
	exists, err := pubsub.SubExists(s.Connection.Context, s.Name)
	if err != nil {
		return err
	}

	if exists {
		s.Subscription = s.Connection.Client.Subscription(s.Name)
	} else {
		s.Subscription, err = s.Connection.Client.NewSubscription(s.Connection.Context, s.Name, s.Topic.Topic, 0, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
