package media

import (
	"cydeaos/libs"
	"cydeaos/libs/media"
	"github.com/google/uuid"
)

const (
	GetTrack   libs.EventType = "get-track"
	GetMood    libs.EventType = "get-mood"
	SwitchMood libs.EventType = "switch-mood"
	NextTrack  libs.EventType = "next-track"
	PlayTrack  libs.EventType = "play-track"
)

type MediaEventPayload struct {
	MediaID  uuid.UUID  `json:"id"`
	Mood     media.Mood `json:"mood"`
	GameCode *string    `json:"gameCode"`
}
