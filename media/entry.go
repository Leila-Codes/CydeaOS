package media

import (
	"cydeaos/config"
	"cydeaos/libs/media"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type Entry struct {
	ID   uuid.UUID  `json:"-"`
	Path string     `json:"-"`
	Mood media.Mood `json:"mood"`
	Name string     `json:"name"`
	URL  string     `json:"url"`
}

func NewEntry(fileName, path string, mood media.Mood) *Entry {
	id := uuid.New()

	// TODO: This is a bit hacky, but it works for now
	// Trim the file extension
	fileName = strings.TrimSuffix(fileName, ".mp3")
	fileName = strings.TrimSuffix(fileName, ".wav")
	fileName = strings.TrimSuffix(fileName, ".ogg")

	return &Entry{
		ID:   id,
		Mood: mood,
		Path: path,
		Name: fileName,
		URL:  fmt.Sprintf("http://localhost:%d/media/%s", config.Port, id.String()),
	}
}
