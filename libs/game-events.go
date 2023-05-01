package libs

type GameEventPayload struct {
	GameEventChannel `json:"event"`
	GameCode         string `json:"gameCode"`
	Player           Player `json:"player"`
	Success          bool   `json:"success"`
	Error            error  `json:"error"`
}

func (ge *GameEventPayload) HasGameCode() bool {
	return len(ge.GameCode) > 0
}
