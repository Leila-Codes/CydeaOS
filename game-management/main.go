package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	brokerURL := os.Getenv("BROKER_URL")
	topic := os.Getenv("TOPIC")
	responderTopic := os.Getenv("RESPONDER_TOPIC")

	//ev := make(chan GameManagementPayload, 1_000)

	//libs.Subscribe(brokerURL, libs.GameManagementEvent, ev)

	logrus.WithFields(logrus.Fields{
		"Broker": brokerURL,
		"Topic":  topic,
	}).Info("Consumer connected.")

	logrus.WithFields(logrus.Fields{
		"Broker": brokerURL,
		"Topic":  responderTopic,
	}).Info("Producer connected.")

	var (
		inputs  = make(chan GameManagementPayload, 1_000)
		outputs = make(chan GameManagementPayload, 1_000)
	)

	go responder(brokerURL, responderTopic, outputs)
	go reactor(brokerURL, topic, inputs)

	processor(inputs, outputs)
}

func responder(brokerURL string, topic string, responses <-chan GameManagementPayload) {
	writer := kafka.Writer{Addr: kafka.TCP(brokerURL), Topic: topic}

	for {
		response := <-responses
		data, _ := json.Marshal(response)
		err := writer.WriteMessages(
			context.Background(),
			kafka.Message{ /*Topic: response.Type.Topic(), */ Value: data},
		)
		if err != nil {
			logrus.Warnf("Failed to write message: %s", err.Error())
		}
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
