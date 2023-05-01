package main

import (
	"CydeaOS/libs"
	"fmt"
)

func wrapError(err error) libs.GameEventPayload {
	return libs.GameEventPayload{
		GameEventChannel: libs.CoreSystemEvent,
		Error:            err,
	}
}

func processor(input <-chan GameManagementPayload, output chan<- interface{}) {
	for {
		event := <-input

		switch event.Type {
		case GetGameType:
			game, err := GetGame(event.GameCode)
			if err != nil {
				wrapError(err)
				continue
			}
			output <- game

		case GameCreationType:
			game := NewGame(event.GameConfig, libs.Player{ID: -1})
			output <- game
		case GameJoinedType:
			game, err := GetGame(event.GameCode)
			if err != nil {
				wrapError(err)
			}
			// TODO: check if player is already in game. update player socket if so.
			game.Players = append(game.Players, event.Player)

			output <- libs.GameEventPayload{
				GameEventChannel: libs.GameManagementEvent,
				GameCode:         game.GameCode,
				Success:          true,
			}
		case GameLeftType:
			// TODO: check if player is in game. remove player from game if so.
			output <- wrapError(fmt.Errorf("not implemented"))
		case GameDeletionType:
			err := DeleteGame(event.GameCode)
			if err != nil {
				wrapError(err)
				continue
			}
			output <- libs.GameEventPayload{
				GameEventChannel: libs.GameManagementEvent,
				GameCode:         event.GameCode,
				Success:          true,
			}
		case GameStartedType:
			output <- wrapError(fmt.Errorf("not implemented"))
			//output <- startGame(event)
		case GameStoppedType:
			//output <- stopGame(event)
		}
	}
}
