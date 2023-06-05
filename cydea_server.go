package main

import (
	"cydeaos/config"
	"cydeaos/log"
	"cydeaos/media"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	logger *logrus.Logger
)

func main() {
	logger = log.GetLogger()

	logger.WithField("port", config.Port).Info("CydeaOS configured and ready to start")

	// load media library
	err := media.LoadLibrary()
	if err != nil {
		logger.Fatal(err)
	}

	// configure endpoints
	router := configureEndpoints()

	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
	if err != nil {
		logrus.Fatal(err)
	}
}

// configureEndpoints creates mux.Router with all the endpoints configured.
func configureEndpoints() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/media/{mediaID}", media.Service).Methods("GET")
	logger.WithField("endpoint", "/media/:mediaID").Debug("Configured endpoint")

	r.HandleFunc("/service", wsConnHandler)
	logger.WithField("endpoint", "/service").Debug("Configured endpoint")

	return r
}
