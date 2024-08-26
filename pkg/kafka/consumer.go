package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func NewConsumer(brokers string, groupId string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}
	return consumer, nil
}
