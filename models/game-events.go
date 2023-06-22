package models

import "encoding/json"

type GameEventHeader struct {
	Channel GameEventChannel
	Type    EventType
}

type GameEventDetails struct {
	GameCode *string `json:"gameCode,omitempty"`
	Player   *Player `json:"player,omitempty"`
}

type GameEventResponse struct {
	Result  *string `json:"result,omitempty"`
	Success bool    `json:"success"`
	Error   error   `json:"error,omitempty"`
}

type GameEvent struct {
	*GameEventHeader
	*GameEventDetails
	*GameEventResponse
	Payload json.RawMessage `json:"Payload,omitempty"`
}

func (ge *GameEvent) SetPayload(payload interface{}) error {
	_, err := json.Marshal(payload)
	return err
}

func (ge *GameEvent) GetPayload(target interface{}) error {
	return json.Unmarshal(ge.Payload, target)
}

// NewGameResponse - constructs a new GameEvent with the given payload and success status
func NewGameResponse(
	originalEvent GameEvent,
	payload interface{},
	success bool,
	errors ...error,
) (GameEvent, error) {
	event := &GameEvent{
		GameEventHeader:  originalEvent.GameEventHeader,
		GameEventDetails: originalEvent.GameEventDetails,
		GameEventResponse: &GameEventResponse{
			Success: success,
		},
	}

	err := event.SetPayload(payload)
	if err != nil {
		return GameEvent{}, err
	}

	if len(errors) > 0 {
		event.Error = errors[0]
	}

	return *event, nil
}
