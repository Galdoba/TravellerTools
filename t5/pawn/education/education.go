package education

import (
	"github.com/Galdoba/TravellerTools/t5/pawn"
)

const (
	NIL = iota
	CHAR_EDU
	CHAR_TRA
	CHAR_INS
)

const (
	CHAR_STRENGHT = iota
	CHAR_DEXTERITY
	CHAR_AGILITY
	CHAR_GRACE
	CHAR_ENDURANCE
	CHAR_STAMINA
	CHAR_VIGOR
	CHAR_INTELLIGENCE
	CHAR_EDUCATION
	CHAR_TRAINING
	CHAR_INSTINCT
	CHAR_SOCIAL
	CHAR_CHARISMA
	CHAR_CASTE
	CHAR_SANITY
	CHAR_PSIONICS
	C1
	C2
	C3
	C4
	C5
	C6
	AUTO
)

type educationalProcess struct {
	Character *pawn.Pawn
	BA        bool
	MA        bool
}

type institution struct {
	name              string
	preRequsite       string
	applyCheck        string
	duration          int //years
	validPassFailCHAR []string
	howManyRolls      int
	provides          string
	graduation        string
	form              string
	/*
		ED5
		Trade School
		Colledge/University


	*/
}

func New(char *pawn.Pawn) {
	ep := educationalProcess{}
	ep.Character = char
}
