package main

import (
	"cydeaos/libs"
	"fmt"
)

var (
	ErrGameNotFound       = fmt.Errorf("game not found")
	ErrGameAlreadyStopped = fmt.Errorf("game already stopped")

	activeGames map[string]*GameObject
)

func GetGame(gameCode string) (*GameObject, error) {
	if game, hasGame := activeGames[gameCode]; hasGame {
		return game, nil
	}

	return nil, ErrGameNotFound
}

func NewGame(config GameConfiguration, hostInfo libs.Player) GameObject {
	gameCode := NewGameCode()

	activeGames[gameCode] = &GameObject{
		GameCode: gameCode,
		Config:   config,
		Players:  make([]libs.Player, 0),
		State:    WaitingForPlayers,
		Host:     hostInfo,
	}

	// TODO: trigger dependent services to populate game data.

	return *activeGames[gameCode]
}

func DeleteGame(gameCode string) error {
	game, err := GetGame(gameCode)
	if err != nil {
		return err
	}

	if game.State == Stopped {
		return ErrGameAlreadyStopped
	}

	game.State = Stopped

	return nil
}
