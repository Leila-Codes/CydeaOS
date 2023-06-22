package events

import (
	"cydeaos/models/game"
)

type GameAction string

const (
	GetGame    GameAction = "get-game"
	CreateGame GameAction = "create-game"
	StartGame  GameAction = "start-game"
	JoinGame   GameAction = "join-game"
	LeaveGame  GameAction = "leave-game"
	StopGame   GameAction = "stop-game"
	DeleteGame GameAction = "delete-game"
)

type GameManagementEvent struct {
	*EventHeader
	Action GameAction     `json:"action"`
	Config *game.Config   `json:"config"`
	Game   *game.Instance `json:"game"`
}

func (e *GameManagementEvent) Header() *EventHeader {
	return e.EventHeader
}

func (e *GameManagementEvent) Error(err error) *GameManagementEvent {
	return &GameManagementEvent{
		EventHeader: &EventHeader{
			Type:    GameManagementType,
			Success: false,
			Error:   err,
		},
		Action: e.Action,
		Config: e.Config,
	}
}

func (e *GameManagementEvent) Success(game *game.Instance) *GameManagementEvent {
	return &GameManagementEvent{
		EventHeader: &EventHeader{
			GameCode: &game.Code,
			Type:     GameManagementType,
			Success:  true,
		},
		Action: e.Action,
		Config: e.Config,
		Game:   game,
	}
}
