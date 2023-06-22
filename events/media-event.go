package events

import (
	"cydeaos/models"
	"cydeaos/models/media"
	"github.com/google/uuid"
)

const (
	GetTrack   models.EventType = "get-track"
	GetMood    models.EventType = "get-mood"
	SwitchMood models.EventType = "switch-mood"
	NextTrack  models.EventType = "next-track"
	PlayTrack  models.EventType = "play-track"
)

type MediaEventPayload struct {
	MediaID  uuid.UUID  `json:"id"`
	Mood     media.Mood `json:"mood"`
	GameCode *string    `json:"gameCode"`
}
