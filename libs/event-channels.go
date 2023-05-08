package libs

import (
	"fmt"
	"strings"
)

type GameEventChannel string

const (
	//Private GameEventChannel = "private"
	Game   GameEventChannel = "game"
	Global GameEventChannel = "global"
)

func (gec GameEventChannel) TopicName() string {
	eventName := string(gec)
	return fmt.Sprintf("cydea_%s", strings.ReplaceAll(eventName, "-", "_"))
}
