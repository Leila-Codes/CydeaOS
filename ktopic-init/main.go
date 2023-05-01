package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	brokerURL := os.Getenv("BROKER_URL")
	topic := os.Getenv("TOPIC")
	partitions, _ := strconv.Atoi(os.Getenv("PARTITIONS"))
	replicationFactor, _ := strconv.Atoi(os.Getenv("REPLICATION_FACTOR"))

	logrus.WithFields(logrus.Fields{
		"Broker":            brokerURL,
		"Topic":             topic,
		"Partitions":        partitions,
		"ReplicationFactor": replicationFactor,
	}).Info("Creating topic")

	client := kafka.Client{
		Addr:    kafka.TCP(brokerURL),
		Timeout: 30 * time.Second,
	}

	_, err := client.CreateTopics(context.Background(), &kafka.CreateTopicsRequest{
		Topics: []kafka.TopicConfig{{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		}},
		ValidateOnly: false,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Topic created")
}
