package kafka

import (
	"fmt"
	"service-bus/pkg/models"
	"service-bus/pkg/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewProducer(brokers string) (*kafka.Producer, error) {
	pdcr, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return nil, err
	}
	return pdcr, nil
}

func ProduceOrderEvent(producer *kafka.Producer, topic string, order models.Order) error {
	orderBytes, err := utils.ToJSON(order)
	if err != nil {
		return fmt.Errorf("failure serializing order %v", err)
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(orderBytes),
	}, nil)

	// wait for delivery confirmation
	e := <-producer.Events()
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("failure delivering message %v", m.TopicPartition.Error)
	}

	return nil
}
