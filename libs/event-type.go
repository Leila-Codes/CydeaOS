package libs

import "strings"

type EventType string

func (ev EventType) Topic() string {
	return "cydea_" + strings.ReplaceAll(string(ev), "-", "_")
}
