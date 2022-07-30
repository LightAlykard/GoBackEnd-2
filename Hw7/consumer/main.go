package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/natspubsub"

	//_ "gocloud.dev/pubsub/rabbitpubsub"
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
	subs, err := natspubsub.OpenSubscription(
		natsConn,
		"example.mysubject",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	defer subs.Shutdown(ctx)

	//subs, err := WithRabbit(ctx)
	// subs, err := WithNATS(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer subs.Shutdown(ctx)

	for {
		msg, err := subs.Receive(ctx)
		if err != nil {
			log.Printf("Receiving message: %v", err)
			break
		}

		fmt.Printf("Received message: %s\n", string(msg.Body))
		fmt.Printf("Metadata: %v\n", msg.Metadata)

		msg.Ack()
	}
}

func WithRabbit(ctx context.Context) (*pubsub.Subscription, error) {
	// Under the hood - processing connect RABBIT_SERVER_URL=amqp://guest:guest@localhost:5672
	// UI for RabbitMQ - http://localhost:15672/
	// Another useful library - github.com/streadway/amqp
	return pubsub.OpenSubscription(ctx, "rabbit://MQ")
}

func WithKafka(ctx context.Context) (*pubsub.Subscription, error) {
	// Under the hood - processing connect KAFKA_BROKERS=localhost:29092
	// UI for Kafka - http://localhost:8082/
	// Another useful library - github.com/Shopify/sarama
	return pubsub.OpenSubscription(ctx, "kafka://MYG?topic=MY")
}

// func WithNATS(ctx context.Context) (*pubsub.Subscription, error) {
// 	// Under the hood - processing connect KAFKA_BROKERS=localhost:29092
// 	// UI for Kafka - http://localhost:8082/
// 	// Another useful library - github.com/Shopify/sarama
// 	return pubsub.OpenSubscription(ctx, "example.mysubject")
// }
