package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

func main() {
	brokerURL := os.Getenv("BROKER_URL")
	topic := os.Getenv("TOPIC")
	responderTopic := os.Getenv("RESPONDER_TOPIC")

	var (
		inputs  = make(chan GameManagementPayload, 1_000)
		outputs = make(chan interface{}, 1_000)
	)

	go responder(brokerURL, responderTopic, outputs)
	go reactor(brokerURL, topic, inputs)

	processor(inputs, outputs)
}

func responder(brokerURL string, topic string, responses <-chan interface{}) {
	writer := kafka.Writer{Addr: kafka.TCP(brokerURL), Topic: topic}

	for {
		response := <-responses
		data, _ := json.Marshal(response)
		writer.WriteMessages(context.Background(), kafka.Message{Value: data})
	}
}

func reactor(brokerURL string, topic string, events chan<- GameManagementPayload) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		Topic:   topic,
		GroupID: "game-management",
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		var event GameManagementPayload
		err = json.Unmarshal(m.Value, &event)
		if err != nil {
			fmt.Println(err)
		}

		events <- event
	}
}
