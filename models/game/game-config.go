package game

type GameType int

type IPMode int

type MusicPlaybackMode int

const (
	Elimination GameType = iota
)

const (
	IPv4 IPMode = iota
	IPv6
)

const (
	Client MusicPlaybackMode = iota
	Server
)

type Config struct {
	GameType  GameType          `json:"gameType"`
	IPMode    IPMode            `json:"ipType"`
	MusicMode MusicPlaybackMode `json:"musicMode"`
}
