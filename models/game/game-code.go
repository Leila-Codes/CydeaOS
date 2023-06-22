package game

import "github.com/tjarratt/babble"

var (
	babbler = babble.NewBabbler()
)

type Code string

func NewCode() Code {
	babbler.Count = 3
	babbler.Separator = "-"

	return Code(babbler.Babble())
}
