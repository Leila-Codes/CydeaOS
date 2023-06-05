package main

import (
	"cydeaos/libs"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

func wsConnHandler(w http.ResponseWriter, r *http.Request) {
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
	// TODO: go socketReceiver(socket)
}
