package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
		logrus.WithField("PORT", port).Warn("Invalid port specified. Defaulting to 8080.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.WriteHeader(200)
		w.Write([]byte("Hello, world!"))
	})

	logrus.Info("Starting server.")

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
