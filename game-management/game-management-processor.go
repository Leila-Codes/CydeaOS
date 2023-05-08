package main

import (
	"cydeaos/libs"
	"fmt"
	"github.com/sirupsen/logrus"
)

func sendError(originalEvent GameManagementPayload, err error, output chan<- GameManagementPayload) {
	originalEvent.Success = false
	originalEvent.Error = err

	output <- originalEvent
}

// gameBroadcast - Broadcasts a game event to all players of that game.
func gameBroadcast(event GameManagementPayload, game *GameObject, output chan<- GameManagementPayload) {
	for _, player := range game.Players {
		event.Player = &player
		output <- event
	}
}

func processor(input <-chan GameManagementPayload, output chan<- GameManagementPayload) {
	for {
		event := <-input

		logrus.WithFields(logrus.Fields{
			"Type":     event.Type,
			"GameCode": event.GameCode,
		}).Debug("Received event.")

		switch event.Type {
		case GameGet:
			game, err := GetGame(*event.GameCode)
			if err != nil {
				sendError(event, err, output)
				continue
			} else if game == nil {
				sendError(event, fmt.Errorf("game not found"), output)
				continue
			}

			event.Success = true
			event.GameInfo = *game
			event.GameConfig = game.Config

			output <- event

		case GameCreate:
			game := NewGame(event.GameConfig, libs.Player{ID: -1})

			event.Success = true
			event.GameInfo = game
			event.GameConfig = game.Config

			output <- event
		case GameJoin:
			// validate player is defined
			if event.Player == nil {
				sendError(event, fmt.Errorf("no player information detected"), output)
				continue
			}

			err := JoinGame(*event.GameCode, *event.Player)
			if err != nil {
				sendError(event, err, output)
				continue
			}

			event.Success = true
			output <- event

		case GameLeave:
			err := LeaveGame(*event.GameCode, *event.Player)
			if err != nil {
				sendError(event, err, output)
				continue
			}

			event.Success = true
			output <- event

		case GameDelete:
			game, err := GetGame(*event.GameCode)
			if err != nil {
				sendError(event, err, output)
				continue
			}

			// TODO: send STOP GAME event to all players in game.
			gameBroadcast(GameManagementPayload{
				GameEvent: libs.GameEvent{Channel: libs.Game, Type: GameLeave},
			}, game, output)

			err = DeleteGame(*event.GameCode)
			if err != nil {
				sendError(event, err, output)
				continue
			}

			event.Success = true
			output <- event
		case GameStart:
			// TODO: check if player is host. start game if so.
			sendError(event, fmt.Errorf("not implemented"), output)

			game, err := GetGame(*event.GameCode)
			if err != nil {
				sendError(event, err, output)
				continue
			}

			gameBroadcast(GameManagementPayload{
				GameEvent: libs.GameEvent{Channel: libs.Game, Type: GameStart},
			}, game, output)

			event.Success = true
			output <- event
		case GameStop:
			// TODO: validate requester is the host.

			game, err := GetGame(*event.GameCode)
			if err != nil {
				sendError(event, err, output)
				continue
			}
			if game.State != Running {
				sendError(event, fmt.Errorf("game is not running"), output)
			}
			game.State = Stopped

			gameBroadcast(GameManagementPayload{
				GameEvent: libs.GameEvent{Channel: libs.Game, Type: GameStop},
			}, game, output)

			event.Success = true
			output <- event
		default:
			logrus.WithFields(logrus.Fields{
				"Type": event.Type,
			}).Error("Unknown event type.")
		}
	}
}
