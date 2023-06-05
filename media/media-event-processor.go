package media

import (
	"errors"
)

var (
	ErrGameNotFound = errors.New("game not found")
)

/*
func gameEventProcessor() {
	listener := kafka.NewEventHandler("game-events", "media")
	creationChannel := listener.Subscribe("create-game")
	deletionChannel := listener.Subscribe("delete-game")

	for {
		select {
		case creationEvent := <-creationChannel:
			handleGameCreation(creationEvent)

		case deletionEvent := <-deletionChannel:

			handleGameDeletion(deletionEvent)
		}
	}
}*/

/*func mediaEventProcessor() {
	listener := kafka.NewEventHandler("media-events", "media")
	responder := kafka.NewEventResponder()

	for {
		select {
		case message := <-listener.Subscribe(GetMood):
			event := libs.GameEvent{}

			if m, hasManager := gameInstances[*event.GameCode]; hasManager {
				response, _ := libs.NewGameResponse(
					event,
					MediaEventPayload{
						Mood: m.CurrentMood(),
					},
					true,
				)

				responder.Respond(response)
			} else {
				response, _ := libs.NewGameResponse(
					event,
					MediaEventPayload{
						Mood: m.CurrentMood(),
					},
					true,
				)

				responder.Respond(response)
			}
		}
	}

}*/
