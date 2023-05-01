package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{HandshakeTimeout: 10 * time.Second}

	clients = make(map[*websocket.Conn]string)
	sockets = make(map[string]*websocket.Conn)
)

func wsService(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		internalServerError(w, err)
		return
	}

	connUUID := uuid.New()
	clients[conn] = connUUID.String()
	sockets[connUUID.String()] = conn
}
