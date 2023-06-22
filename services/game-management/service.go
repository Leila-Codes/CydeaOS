package game_management

import (
	"cydeaos/models"
	"cydeaos/models/game"
	"errors"
	"fmt"
)

var (
	ErrGameNotFound       = errors.New("game not found")
	ErrGameAlreadyStarted = errors.New("game already started")
	ErrGameAlreadyStopped = errors.New("game already stopped")
)

var (
	service = &gameManagementService{games: make(map[game.Code]*game.Instance)}
)

type gameManagementService struct {
	games map[game.Code]*game.Instance
}

// GetGame returns a game instance matching the game code
func (s *gameManagementService) GetGame(code game.Code) (*game.Instance, error) {
	if game, hasGame := s.games[code]; hasGame {
		return game, nil
	}

	return nil, ErrGameNotFound
}

// CreateGame creates a new game instance and returns the game code
func (s *gameManagementService) CreateGame(config *game.Config, hostInfo *models.Player) *game.Instance {
	code := game.NewCode()

	s.games[code] = &game.Instance{
		Code:    code,
		Config:  config,
		Players: make([]*models.Player, 0),
		Host:    hostInfo,
	}

	return s.games[code]
}

// JoinGame joins a player to a game instance
func (s *gameManagementService) JoinGame(code *game.Code, player *models.Player) error {
	gameInstance, ok := s.games[*code]
	if !ok {
		return ErrGameNotFound
	}

	gameInstance.AddPlayer(player)

	return nil
}

// LeaveGame - Removes a player from a game (matching the game code).
func (s *gameManagementService) LeaveGame(code game.Code, player *models.Player) error {
	game, err := s.GetGame(code)
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
func (s *gameManagementService) DeleteGame(code game.Code) error {
	gameRef, err := s.GetGame(code)
	if err != nil {
		return err
	}

	if gameRef.State == game.Stopped {
		return ErrGameAlreadyStopped
	}

	gameRef.State = game.Stopped

	// TODO: delete game data.
	delete(s.games, code)

	return nil
}

// StartGame - Starts the game.
func (s *gameManagementService) StartGame(code game.Code) error {
	gameRef, err := s.GetGame(code)
	if err != nil {
		return err
	}

	if gameRef.State != game.WaitingForPlayers {
		return ErrGameAlreadyStarted
	}

	gameRef.State = game.Stopped

	return nil
}
