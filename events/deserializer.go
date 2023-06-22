package events

import (
	"encoding/json"
	"errors"
)

var (
	ErrUnknownEventType = errors.New("unrecognised event type")
)

// Deserialize a JSON byte array into an appropriate event struct
func Deserialize(data []byte) (Event, error) {
	var header EventHeader
	err := json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	switch header.Type {
	case GameManagementType:
		var gameManagementEvent GameManagementEvent

		err := json.Unmarshal(data, &gameManagementEvent)
		if err != nil {
			return nil, err
		}

		return &gameManagementEvent, nil
	default:
		return nil, ErrUnknownEventType
	}
}
