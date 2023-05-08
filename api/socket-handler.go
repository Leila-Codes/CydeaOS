package main

import (
	"context"
	"cydeaos/libs"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type WebsocketClient struct {
	*websocket.Conn
	socketID string
}

var (
	upgrader = websocket.Upgrader{HandshakeTimeout: 10 * time.Second}

	clients = make(map[string]*WebsocketClient)
	rooms   = make(map[string][]libs.Player)
)

func wsService(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		internalServerError(w, err)
		return
	}

	var (
		socketID = uuid.NewString()
		socket   = &WebsocketClient{socketID: socketID, Conn: conn}
	)

	clients[socketID] = socket
	go wsReceiver(socket)
}

func wsReceiver(conn *WebsocketClient) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}

		// partially decode the json to get the headers
		var header libs.GameEventHeader
		err = json.Unmarshal(data, &header)
		if err != nil {
			panic(err)
		}

		var (
			targetTopic = header.Type.Topic()
		)

		err = writer.WriteMessages(context.TODO(), kafka.Message{
			Topic: targetTopic,
			Value: data,
		})
		if err != nil {
			logrus.Fatal(err)
		}
	}
}
