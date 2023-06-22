package main

import (
	"cydeaos/events"
	"github.com/gorilla/websocket"
)

func wsHandler(conn *websocket.Conn) {
	for {
		t, data, err := conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			continue
		}

		switch t {
		case websocket.TextMessage:
			event, err := events.Deserialize(data)
			if err != nil {
				logger.Error(err)
				continue
			}

			response, err := events.Broadcast(event)
			if err != nil {
				if err == events.ErrNoSubscribers {
					logger.WithField("type", event.Header().Type).Debug("No subscribers for event type")
					continue
				}

				logger.Error(err)
			}

			for _, r := range response {
				if r == nil {
					continue
				}

				// TODO: resend message to all other concerned clients
				err := conn.WriteJSON(r)
				if err != nil {
					logger.Error(err)
				}
			}
		default:
			logger.WithField("type", t).Debug("Received unknown message type")
		}
	}
}
