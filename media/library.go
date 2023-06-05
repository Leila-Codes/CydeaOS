package media

import (
	"cydeaos/config"
	"cydeaos/libs/media"
	"cydeaos/log"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"regexp"
)

var (
	definitions map[uuid.UUID]Entry
	groupByMood map[media.Mood][]Entry
	logger      *logrus.Logger

	reSupportedMedia = regexp.MustCompile(`(?i)\.(mp3|wav|ogg)$`)
)

func init() {
	definitions = make(map[uuid.UUID]Entry)
	groupByMood = make(map[media.Mood][]Entry)

	logger = log.GetLogger()
}

func LoadLibrary() error {
	if config.MediaDir == "" {
		config.MediaDir = "./media"
		logger.Warn("MEDIA_DIR not set, defaulting to ", config.MediaDir)
	}

	logger.WithField("dir", config.MediaDir).Debug("Scanning MEDIA_DIR...")

	files, err := os.ReadDir(config.MediaDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			mood := media.Mood(file.Name())
			groupByMood[mood] = make([]Entry, 0)

			logger.Infof("Discovering media files for mood %s", mood)

			potentialMediaFiles, err := os.ReadDir(path.Join(config.MediaDir, file.Name()))
			if err != nil {
				return err
			}

			for _, mediaFile := range potentialMediaFiles {
				if mediaFile.IsDir() {
					continue
				} else if reSupportedMedia.MatchString(mediaFile.Name()) {
					entry := NewEntry(mediaFile.Name(), path.Join(config.MediaDir, file.Name(), mediaFile.Name()), mood)
					definitions[entry.ID] = *entry
					groupByMood[mood] = append(groupByMood[mood], *entry)
					logger.WithFields(logrus.Fields{
						"mood": mood,
						//"name": entry.Name,
						"url": entry.URL,
					}).Debug("Discovered media file")
				}
			}
			continue
		}
	}

	if len(definitions) == 0 {
		logger.Warn("No media files discovered")
		return nil
	}

	return nil
}
