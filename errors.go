package main

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SocketError struct {
	Error error `json:"error"`
}

func internalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
	logrus.WithField("Error", err.Error()).Error("Internal Server Error")
}

func socketError(conn *websocket.Conn, error error) error {
	return conn.WriteJSON(SocketError{Error: error})
}
