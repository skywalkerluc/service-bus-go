package main

import (
	"fmt"
	"os"
	"service-bus/pkg/kafka"
	"service-bus/pkg/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading env variables: %v\n", err)
		os.Exit(1)
	}

	kafkaBroker := os.Getenv("KAFKA_BROKER")

	producer, err := kafka.NewProducer(kafkaBroker)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when creating kafka producer: %v\n", err)
		os.Exit(1)
	}
	defer producer.Close()

	order := models.Order{
		OrderID:   "123",
		ProductID: "345",
		Quantity:  2,
	}

	if err := kafka.ProduceOrderEvent(producer, "tp_order_created", order); err != nil {
		fmt.Printf("error when sending order event %v\n", err)
	} else {
		fmt.Println("order event sent successfully")
	}
}
