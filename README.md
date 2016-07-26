# drudsub

```golang

package main

import (
	"log"

	"github.com/drud/drudsub"
)

var topicName = "test-topic"
var subName = "sub-name"

func main() {
	connection := drudsub.Connection{}
	connection.Connect()

	topic := drudsub.Topic{
		Name:       topicName,
		Connection: connection,
	}
	err := topic.Create()

	if err != nil {
		log.Println("Could not create topic.")
		log.Fatal(err)
	}

	var messages []drudsub.Message
	attributes := make(map[string]string)
	attributes["foo"] = "foo"
	attributes["bar"] = "bar"

	messages = append(messages, drudsub.Message{
		Data:       []byte("Sample Message"),
		Attributes: attributes,
	})

	_, err = topic.Publish(messages)
	if err != nil {
		log.Println("could not send messages")
		log.Fatal(err)
	}

	sub := drudsub.Subscription{
		Name:       subName,
		Topic:      topic,
		Connection: connection,
	}

	subChan, err := sub.Subscribe(true)

	if err != nil {
		log.Println("Could not create subscription")
		log.Fatal(err)
	}

	for {
		select {
		case msg := <-subChan:
			log.Printf("%s\n", string(msg.Data))
			log.Printf("%v\n", msg.Attributes)
		}
	}
}


```