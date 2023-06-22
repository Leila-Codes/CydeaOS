package media

import (
	"cydeaos/events"
	"cydeaos/models/game"
	"cydeaos/models/media"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"math/rand"
)

var (
	gameInstances = make(map[game.Code]*Manager)
)

func getMenuTrack() uuid.UUID {
	trackList := groupByMood[media.MainMenu]
	nextIndex := rand.Intn(len(trackList))
	return trackList[nextIndex].ID
}

func handleGameCreation(gameInfo events.Event) {
	if gameInfo.Header().GameCode == nil {
		return
	}

	if _, exists := gameInstances[*gameInfo.Header().GameCode]; exists {
		return
	}

	logger.Info("Media manager created for game ", *gameInfo.Header().GameCode)

	gameInstances[*gameInfo.Header().GameCode] = NewManager()
}

func handleGameDeletion(gameInfo events.Event) {
	if _, exists := gameInstances[*gameInfo.Header().GameCode]; !exists {
		return
	}

	delete(gameInstances, *gameInfo.Header().GameCode)
}

// GameEventHandler handles game management events
func GameEventHandler(ev events.Event) (events.Event, error) {
	event := ev.(*events.GameManagementEvent)

	logger.WithFields(logrus.Fields{
		"action": event.Action,
		"event":  event,
	}).Info("GameEventHandler [media] rx event")

	switch event.Action {
	case events.CreateGame:
		handleGameCreation(event)
	case events.DeleteGame:
		handleGameDeletion(event)
	}

	return nil, nil
}
