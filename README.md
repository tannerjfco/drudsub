# drudsub


### Required environment variables

<dl>
<dd>The path to a JWT</dd>
<dt>export DRUDSUB_JWT="/path/to/jwt.json"</dt>
<dd>The google cloud project name</dd>
<dt>export DRUDSUB_PROJECT="bbowman-drud"</dt>
</dl>

### Example implementation code.

```golang

package main

import (
	"log"

	"github.com/drud/drudsub"
)

var topicName = "test-topic"
var subName = "sub-name"

func main() {
    // Create a connection.
	connection := drudsub.Connection{}
	connection.Connect()

    // Create a topic to send messages to.
	topic := drudsub.Topic{
		Name:       topicName,
		Connection: connection,
	}
	err := topic.Create()

	if err != nil {
		log.Println("Could not create topic.")
		log.Fatal(err)
	}

    // Create a drudsu message.
	var messages []drudsub.Message
	attributes := make(map[string]string)
	attributes["foo"] = "foo"
	attributes["bar"] = "bar"

	messages = append(messages, drudsub.Message{
		Data:       []byte("Sample Message"),
		Attributes: attributes,
	})

    // Publish the message to a topic.
	_, err = topic.Publish(messages)
	if err != nil {
		log.Println("could not send messages")
		log.Fatal(err)
	}

    // Create a subscription for a given topic.
	sub := drudsub.Subscription{
		Name:       subName,
		Topic:      topic,
		Connection: connection,
	}

    // Open a subscription channel to this topic.
	subChan, err := sub.Subscribe(true)

	if err != nil {
		log.Println("Could not create subscription")
		log.Fatal(err)
	}  

    
    // Process messages from the channel.
	for {
		select {
		case msg := <-subChan:
			log.Printf("%s\n", string(msg.Data))
			log.Printf("%v\n", msg.Attributes)
		}
	}
}


```

