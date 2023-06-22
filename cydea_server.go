package main

import (
	"cydeaos/config"
	"cydeaos/events"
	"cydeaos/log"
	"cydeaos/media"
	game_management "cydeaos/services/game-management"
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

	// configure services
	configureServices()

	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router)
	if err != nil {
		logrus.Fatal(err)
	}
}

// configureEndpoints creates mux.Router with all the endpoints configured.
func configureEndpoints() *mux.Router {
	r := mux.NewRouter()

	// ##### CONFIGURE ENDPOINTS #####

	r.HandleFunc("/media/{mediaID}", media.Service).Methods("GET")

	logger.WithField("endpoint", "/media/:mediaID").Debug("Configured endpoint")

	// websocket handler
	r.HandleFunc("/service", wsConnHandler)
	logger.WithField("endpoint", "/service").Debug("Configured WebSocket service endpoint")

	// ### END CONFIGURE ENDPOINTS ###

	return r
}

func configureService(
	cb func(event events.Event) (events.Event, error),
	eventType events.EventType,
) {
	events.Subscribe(eventType, &cb)

	logger.WithFields(logrus.Fields{
		"event-type": eventType,
		"handler":    cb,
	}).Info("Configured backend service")
}

// configureServices configures all the services.
func configureServices() {
	// ##### CONFIGURE SERVICES #####

	configureService(game_management.GMHandler, events.GameManagementType)
	configureService(media.GameEventHandler, events.GameManagementType)

	// ### END CONFIGURE SERVICES ###
}
