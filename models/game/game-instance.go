package game

import "cydeaos/models"

type State int

const (
	WaitingForPlayers State = iota
	Running
	Stopped
)

type Instance struct {
	Code    Code             `json:"gameCode"`
	Config  *Config          `json:"config"`
	State   State            `json:"state"`
	Players []*models.Player `json:"-"`
	Host    *models.Player   `json:"-"`
}

func (i *Instance) AddPlayer(player *models.Player) {
	i.Players = append(i.Players, player)
}

func (i *Instance) RemovePlayer(player *models.Player) {
	for index, p := range i.Players {
		if p == player {
			i.Players = append(i.Players[:index], i.Players[index+1:]...)
		}
	}
}
