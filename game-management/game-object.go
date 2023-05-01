package main

import "cydeaos/libs"

type GameState int

const (
	WaitingForPlayers GameState = iota
	Running
	Stopped
)

type GameObject struct {
	GameCode string            `json:"gameCode"`
	Config   GameConfiguration `json:"config"`
	Players  []libs.Player     `json:"-"`
	State    GameState         `json:"state"`
	Host     libs.Player       `json:"-"`
}
