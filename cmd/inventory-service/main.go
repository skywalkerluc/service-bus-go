package main

import (
	"fmt"
	"os"
	"service-bus/internal/inventory"
	"service-bus/pkg/kafka"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading env variables: %v\n", err)
		os.Exit(1)
	}

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	consumerGroup := os.Getenv("KAFKA_CONSUMER_GROUP")

	consumer, err := kafka.NewConsumer(kafkaBroker, consumerGroup)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating kafka consumer: %v\n", err)
		os.Exit(1)
	}
	defer consumer.Close()

	consumer.SubscribeTopics([]string{"tp_order_created"}, nil)

	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("message received: %s\n", string(message.Value))
			inventory.HandleOrder(message.Value)
		} else {
			fmt.Printf("error reading message: %v\n", err)
		}
	}
}
