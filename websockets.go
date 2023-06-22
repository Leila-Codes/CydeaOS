package main

import (
	"cydeaos/config"
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
	upgrader websocket.Upgrader

	clients = make(map[string]*WebsocketClient)
	//rooms   = make(map[string][]*models.Player)
)

func init() {
	upgrader = websocket.Upgrader{
		HandshakeTimeout: 10 * time.Second,
	}

	if config.Debug {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}
}

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
	go wsHandler(conn)
}
