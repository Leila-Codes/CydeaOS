package kafka

import (
	"context"
	"cydeaos/libs"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type EventHandler struct {
	reader        *kafka.Reader
	listening     bool
	subscriptions map[libs.EventType]chan libs.GameEvent
}

func NewEventHandler(topicName, consumerName string) *EventHandler {
	return &EventHandler{
		reader:        kReader(topicName, consumerName),
		subscriptions: make(map[libs.EventType]chan libs.GameEvent),
	}
}

func (h *EventHandler) Subscribe(eventType libs.EventType) (handler chan libs.GameEvent) {
	handler = make(chan libs.GameEvent)
	h.subscriptions[eventType] = handler

	// start listening if not already listening
	if !h.listening {
		go h.Listen()
		h.listening = true
	}

	return handler
}

func (h *EventHandler) Unsubscribe(eventType libs.EventType) {
	delete(h.subscriptions, eventType)
}

func (h *EventHandler) Listen() {
	for {
		m, err := h.reader.ReadMessage(context.TODO())
		if err != nil {
			logrus.Fatal(err)
		}

		// partially decode the json to get the headers
		var header libs.GameEventHeader
		err = json.Unmarshal(m.Value, &header)
		if err != nil {
			logrus.Fatal(err)
		}

		if listener, hasListener := h.subscriptions[header.Type]; hasListener {
			var event libs.GameEvent
			err := json.Unmarshal(m.Value, &event)
			if err != nil {
				logrus.Fatal(err)
			}

			listener <- event
		}
	}
}
