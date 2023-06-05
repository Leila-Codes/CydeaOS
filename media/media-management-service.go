package media

import (
	"cydeaos/libs"
	"cydeaos/libs/media"
	"github.com/google/uuid"
	"math/rand"
)

var (
	gameInstances = make(map[string]*Manager)
)

func getMenuTrack() uuid.UUID {
	trackList := groupByMood[media.MainMenu]
	nextIndex := rand.Intn(len(trackList))
	return trackList[nextIndex].ID
}

func handleGameCreation(gameInfo libs.GameEvent) {
	if _, exists := gameInstances[*gameInfo.GameCode]; exists {
		return
	}

	gameInstances[*gameInfo.GameCode] = NewManager()
}

func handleGameDeletion(gameInfo libs.GameEvent) {
	if _, exists := gameInstances[*gameInfo.GameCode]; !exists {
		return
	}

	delete(gameInstances, *gameInfo.GameCode)
}
