package main

import (
	"context"
	"cydeaos/libs"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
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

func globalTransmitter() {
	reader := kReader(libs.Global.TopicName(), "cydea-all-responder")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(err)
		}

		for _, client := range clients {
			err = client.WriteMessage(websocket.TextMessage, m.Value)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func gameTransmitter() {
	reader := kReader(libs.Game.TopicName(), "cydea-game-responder")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			logrus.Error(err)
		}

		event := libs.GameEvent{}
		err = json.Unmarshal(m.Value, &event)
		if err != nil {
			logrus.Fatal(err)
		}

		if event.GameCode != nil {
			logrus.Warnf("game message cannot be sent - missing game code - %v", event)
			continue
		}

		if room, ok := rooms[*event.GameCode]; ok {
			for _, player := range room {
				sid := player.SocketID()
				if client, hasClient := clients[sid]; hasClient {
					err := client.WriteMessage(websocket.TextMessage, m.Value)
					if err != nil {
						logrus.Warnf("cannot send message to client '%s' - %v", sid, err)
					}
				} else {
					logrus.Warnf("cannot find client socket '%s' - may have disconnected", sid)
				}
			}
		}
	}
}
