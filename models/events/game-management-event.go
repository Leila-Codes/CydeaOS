package events

import (
	"cydeaos/models"
)

const (
	GameGet    models.EventType = "get-game"
	GameCreate models.EventType = "create-game"
	GameJoin   models.EventType = "join-game"
	GameLeave  models.EventType = "leave-game"
	GameStart  models.EventType = "start-game"
	GameStop   models.EventType = "stop-game"
	GameDelete models.EventType = "delete-game"
)
