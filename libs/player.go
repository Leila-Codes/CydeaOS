package libs

type PlayerState int

const (
	Offline PlayerState = iota
	Disconnected
	Online
	InGame
)

type Player struct {
	ID            int         `json:"id"`
	Username      string      `json:"username"`
	State         PlayerState `json:"state"`
	socketID      string      `json:"-"`
	currentTarget string      `json:"-"`
}

func (p *Player) SocketID() string {
	return p.socketID
}

func (p *Player) SetSocketID(socketID string) {
	p.socketID = socketID
}

func (p *Player) IsOnline() bool {
	return p.State >= Offline
}
