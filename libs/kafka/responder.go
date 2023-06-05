package kafka

import (
	"context"
	"cydeaos/libs"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	bootstrapServers string
)

func init() {
	bootstrapServers = os.Getenv("BROKER_URL")
	if bootstrapServers == "" {
		bootstrapServers = "localhost:9092"
	}
}

type EventResponder struct {
	writer *kafka.Writer
	Out    chan libs.GameEvent
}

func NewEventResponder() *EventResponder {
	return &EventResponder{
		writer: &kafka.Writer{
			Addr:       kafka.TCP(bootstrapServers),
			BatchBytes: 1e3,
			Async:      true,
			//BatchSize: 1,
		},
		Out: make(chan libs.GameEvent, 1_000),
	}
}

func (r *EventResponder) Respond(event libs.GameEvent) {
	r.Out <- event
}

func (r *EventResponder) Responder() {
	for {
		msg := <-r.Out

		data, err := json.Marshal(msg)
		if err != nil {
			logrus.Fatal(err)
		}

		err = r.writer.WriteMessages(
			context.TODO(),
			kafka.Message{
				Topic: msg.Channel.TopicName(),
				Value: data,
			},
		)

		if err != nil {
			logrus.Fatal(err)
		}
	}
}
