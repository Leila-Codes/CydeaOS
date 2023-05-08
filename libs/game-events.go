package libs

type GameEventHeader struct {
	Channel GameEventChannel
	Type    EventType
}

// GameEvent - A generic game event payload container.
type GameEvent struct {
	Channel  GameEventChannel
	Type     EventType
	GameCode *string
	Player   *Player
	Result   *string
	Success  bool
	Error    error
}

func (ge *GameEvent) HasGameCode() bool {
	return ge.GameCode != nil
}
