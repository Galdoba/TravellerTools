package education

import (
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic/charset"
)

const (
	NIL = iota
	CHAR_EDU
	CHAR_TRA
	CHAR_INS
)

type educationalProcess struct {
	CHAR_BASE       int
	Characteristics charset.CharSet
	Waivers         int
}

type institution struct {
	validCHAR int
}

func New() {
	ep := educationalProcess{}
	ep.Characteristics.String()
}
