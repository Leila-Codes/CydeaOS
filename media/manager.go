package media

import (
	"cydeaos/models/media"
	"github.com/google/uuid"
)

type Manager struct {
	queues       map[media.Mood]*media.Queue
	currentMood  media.Mood
	currentTrack uuid.UUID
}

func NewManager() *Manager {
	m := &Manager{
		queues: make(map[media.Mood]*media.Queue),
	}

	for _, mood := range media.InGameMoods {
		// initialize the queue
		m.queues[mood] = media.NewMediaQueue(mood)
		// populate with global media entries
		for _, track := range groupByMood[mood] {
			m.queues[mood].Enqueue(track.ID)
		}
	}

	return m
}

func (m *Manager) CurrentMood() media.Mood {
	return m.currentMood
}

func (m *Manager) CurrentTrack() uuid.UUID {
	return m.currentTrack
}

func (m *Manager) NextTrack(newMood ...[]media.Mood) uuid.UUID {
	if len(newMood) > 0 {
		m.currentMood = newMood[0][0]
	}

	m.currentTrack, _ = m.queues[m.currentMood].Dequeue()
	return m.currentTrack
}
