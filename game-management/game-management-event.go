package main

import "cydeaos/libs"

const (
	GameGet    libs.EventType = "get-game"
	GameCreate libs.EventType = "create-game"
	GameJoin   libs.EventType = "join-game"
	GameLeave  libs.EventType = "leave-game"
	GameStart  libs.EventType = "start-game"
	GameStop   libs.EventType = "stop-game"
	GameDelete libs.EventType = "delete-game"
)

type GameManagementPayload struct {
	libs.GameEvent
	GameInfo   GameObject        `json:"gameInfo"`
	GameConfig GameConfiguration `json:"gameConfig"`
}
