package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

var (
	writer kafka.Writer
)

func main() {
	// configure logging
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	// configure application
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
		logrus.WithField("PORT", port).Warn("Invalid port specified. Defaulting to 8080.")
	}

	// initialise writer
	writer = kafka.Writer{Addr: kafka.TCP("localhost:9092")}

	// start I/O ops
	go gameTransmitter()
	go globalTransmitter()

	// configure http routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("CydeaOS Game Services"))
	})

	// websocket server
	http.HandleFunc("/service", wsService)

	// start the server
	logrus.Info("Starting server.")
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
