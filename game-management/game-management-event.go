package main

import "CydeaOS/libs"

type GameManagementEventType string

const (
	GetGameType      GameManagementEventType = "get-game"
	GameCreationType GameManagementEventType = "game-creation"
	GameJoinedType   GameManagementEventType = "game-joined"
	GameLeftType     GameManagementEventType = "game-left"
	GameDeletionType GameManagementEventType = "game-deletion"
	GameStartedType  GameManagementEventType = "game-started"
	GameStoppedType  GameManagementEventType = "game-stopped"
)

type GameManagementPayload struct {
	Type       GameManagementEventType `json:"type"`
	GameInfo   GameObject              `json:"gameInfo"`
	GameConfig GameConfiguration       `json:"gameConfig"`
	libs.GameEventPayload
}
