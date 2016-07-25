package drudsub

import (
	"log"
	"time"

	"google.golang.org/cloud/pubsub"
)

// Subscription to a drudsub topic.
type Subscription struct {
	// The Name of the Subscription.
	Name string

	// The Topic the subscription is for.
	Topic Topic

	// The pubsub subscription object - should only be used internally.
	Subscription *pubsub.Subscription

	// The drudsub connection.
	Connection Connection

	// Channel from which to receive subscription messages.
	Channel chan Message

	// The ticker duration for the message channel. Will default to 5 seconds.
	tickInterval time.Duration
}

// Subscribe to a topic.
func (s *Subscription) Subscribe(create bool) (<-chan Message, error) {
	if create {
		s.Topic.Create()
	}

	exists, err := pubsub.SubExists(s.Connection.Context, s.Name)
	if err != nil {
		return s.Channel, err
	}

	if exists {
		s.Subscription = s.Connection.Client.Subscription(s.Name)
	} else {
		s.Subscription, err = s.Connection.Client.NewSubscription(s.Connection.Context, s.Name, s.Topic.Topic, 0, nil)
		if err != nil {
			return s.Channel, err
		}
	}

	go s.read()
	return s.Channel, nil
}

func (s *Subscription) read() error {
	// Pull() returns an iterator which handles requesting of messages in the background
	it, err := s.Subscription.Pull(s.Connection.Context)
	if err != nil {
		log.Fatalln(err)
	}
	defer it.Stop()

	if s.tickInterval == 0 {
		s.tickInterval = 5
	}

	ticker := time.NewTicker(time.Second * s.tickInterval).C
	for {
		select {
		case <-ticker:
			if err == pubsub.Done {
				msg, err := it.Next()
				if err == pubsub.Done {
					s.Channel = nil
					return nil
				}
				if err != nil {
					return err
				}

				s.Channel <- Message{
					Data:       msg.Data,
					Attributes: msg.Attributes,
				}
			}
		}
	}
}

// Create a new subscription
func (s *Subscription) Create() error {
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
