package main

import "github.com/tjarratt/babble"

var babbler = babble.NewBabbler()

func NewGameCode() string {
	babbler.Count = 3
	babbler.Separator = "-"

	return babbler.Babble()
}
