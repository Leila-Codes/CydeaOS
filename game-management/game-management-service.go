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

// JoinGame - Adds a player to a game (matching the game code).
// TODO: check if player is already in game. update player socket if so.
func JoinGame(gameCode string, player libs.Player) error {
	game, err := GetGame(gameCode)
	if err != nil {
		return err
	}

	game.Players = append(game.Players, player)

	return nil
}

// LeaveGame - Removes a player from a game (matching the game code).
func LeaveGame(gameCode string, player libs.Player) error {
	game, err := GetGame(gameCode)
	if err != nil {
		return err
	}

	for i, p := range game.Players {
		if p == player {
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("player not found in game")
}

// DeleteGame - Stops the game immediately and deletes all associated data.
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
