package main

import (
	"context"
	"cydeaos/api/rooms"
	"cydeaos/libs"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
)

var (
	topicsToListen = []string{
		"cydea_result",
	}
)

func initialise() {
	for _, name := range topicsToListen {
		go responder(name)
	}
}

func responder(topicName string) {
	// TODO: configurable kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topicName,
	})

	for {
		m, err := reader.ReadMessage(context.TODO())
		if err != nil {
			// TODO: configure logrus
			//log.WithError(err).Error("Error reading message from kafka.")
			panic(err)
			continue
		}

		event := libs.GameEventPayload{}
		err = json.Unmarshal(m.Value, &event)
		if err != nil {
			panic(err)
		}

		if len(event.GameCode) > 0 {
			rooms.Broadcast(event.GameCode, string(m.Value))
		} else if len(event.Player.SocketID()) > 0 {
			sockets[event.Player.SocketID()].WriteMessage(websocket.TextMessage, []byte(m.Value))
		} else {
			panic("no game code or player socket id to broadcast to")
		}
	}
}
