package main

import "cydeaos/libs"

type GameManagementPayload struct {
	libs.GameEvent
	GameInfo   GameObject        `json:"gameInfo"`
	GameConfig GameConfiguration `json:"gameConfig"`
}
