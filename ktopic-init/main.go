package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"os"
	"strconv"
)

func main() {
	brokerURL := os.Getenv("BROKER_URL")
	topic := os.Getenv("TOPIC")
	partitions, _ := strconv.Atoi(os.Getenv("PARTITIONS"))
	replicationFactor, _ := strconv.Atoi(os.Getenv("REPLICATION_FACTOR"))

	client := kafka.Client{
		Addr:    kafka.TCP(brokerURL),
		Timeout: 60,
	}

	_, err := client.CreateTopics(context.TODO(), &kafka.CreateTopicsRequest{
		Topics: []kafka.TopicConfig{{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		}},
		ValidateOnly: false,
	})

	if err != nil {
		panic(err)
	}
}
