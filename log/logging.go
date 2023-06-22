package log

import (
	"cydeaos/config"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		//FullTimestamp: true,

	})

	if config.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.WithField("Level", logger.Level.String()).Info("Logger configured")
}

func GetLogger() *logrus.Logger {
	return logger
}
