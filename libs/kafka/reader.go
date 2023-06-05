package kafka

import (
	"github.com/segmentio/kafka-go"
	"time"
)

func kReader(topicName, consumerName string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          topicName,
		GroupID:        consumerName,
		CommitInterval: 5 * time.Second,
	})
}
