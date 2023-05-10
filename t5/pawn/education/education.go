package education

import (
	"github.com/Galdoba/TravellerTools/pkg/profile"
)

const (
	NIL = iota
	CHAR_EDU
	CHAR_TRA
	CHAR_INS
)

type educationalProcess struct {
	CHAR_BASE int
	Character profile.Profile
	Waivers   int
}

type institution struct {
	validCHAR int
}

func New(char profile.Profile) {
	ep := educationalProcess{}
	ep.Character = char
}
