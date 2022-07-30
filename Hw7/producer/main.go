package main

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"gocloud.dev/pubsub/natspubsub"

	"gocloud.dev/pubsub"
	//	_ "gocloud.dev/pubsub/rabbitpubsub"
	//_ "gocloud.dev/pubsub/kafkapubsub"
	_ "gocloud.dev/pubsub/natspubsub"
)

func main() {
	ctx := context.Background()

	natsConn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer natsConn.Close()

	topic, err := natspubsub.OpenTopic(natsConn, "example.mysubject", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer topic.Shutdown(ctx)
	//topic, err := WithRabbit(ctx)
	//topic, err := WithNATS(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer topic.Shutdown(ctx)

	err = topic.Send(ctx, &pubsub.Message{
		Body: []byte("I wonna toch the scy!"), // json

		Metadata: map[string]string{
			"language":   "en",
			"importance": "high",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func WithRabbit(ctx context.Context) (*pubsub.Topic, error) {
	// Under the hood - processing connect RABBIT_SERVER_URL=amqp://guest:guest@localhost:5672
	// UI for RabbitMQ - http://localhost:15672/
	// Another useful library - github.com/streadway/amqp
	return pubsub.OpenTopic(ctx, "rabbit://MY")
}

func WithKafka(ctx context.Context) (*pubsub.Topic, error) {
	// Under the hood - processing connect KAFKA_BROKERS=localhost:29092
	// UI for Kafka - http://localhost:8082/
	// Another useful library - github.com/Shopify/sarama
	return pubsub.OpenTopic(ctx, "kafka://MY")
}

// func WithNATS(ctx context.Context) (*pubsub.Topic, error) {
// 	// Under the hood - processing connect KAFKA_BROKERS=localhost:29092
// 	// UI for Kafka - http://localhost:8082/
// 	// Another useful library - github.com/Shopify/sarama
// 	return pubsub.OpenTopic(ctx, "nats://nats:4222")
// }
