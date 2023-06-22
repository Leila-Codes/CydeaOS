package events

import (
	"cydeaos/models"
	"cydeaos/models/game"
)

type EventType string

const (
	GameManagementType EventType = "game-management"
)

type Event interface {
	Header() *EventHeader
}

type EventHeader struct {
	Type     EventType      `json:"name"`
	GameCode *game.Code     `json:"gameCode"`
	Player   *models.Player `json:"player"`
	Success  bool           `json:"success"`
	Error    error          `json:"error,omitempty"`
}

func (e *EventHeader) Header() *EventHeader {
	return e
}
