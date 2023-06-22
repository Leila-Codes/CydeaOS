package kafka

import (
	"context"
	"cydeaos/models"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type EventHandler struct {
	reader        *kafka.Reader
	listening     bool
	subscriptions map[models.EventType]chan models.GameEvent
}

func NewEventHandler(topicName, consumerName string) *EventHandler {
	return &EventHandler{
		reader:        kReader(topicName, consumerName),
		subscriptions: make(map[models.EventType]chan models.GameEvent),
	}
}

func (h *EventHandler) Subscribe(eventType models.EventType) (handler chan models.GameEvent) {
	handler = make(chan models.GameEvent)
	h.subscriptions[eventType] = handler

	// start listening if not already listening
	if !h.listening {
		go h.Listen()
		h.listening = true
	}

	return handler
}

func (h *EventHandler) Unsubscribe(eventType models.EventType) {
	delete(h.subscriptions, eventType)
}

func (h *EventHandler) Listen() {
	for {
		m, err := h.reader.ReadMessage(context.TODO())
		if err != nil {
			logrus.Fatal(err)
		}

		// partially decode the json to get the headers
		var header models.GameEventHeader
		err = json.Unmarshal(m.Value, &header)
		if err != nil {
			logrus.Fatal(err)
		}

		if listener, hasListener := h.subscriptions[header.Type]; hasListener {
			var event models.GameEvent
			err := json.Unmarshal(m.Value, &event)
			if err != nil {
				logrus.Fatal(err)
			}

			listener <- event
		}
	}
}
