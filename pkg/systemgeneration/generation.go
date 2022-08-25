package systemgeneration

import (
	"github.com/Galdoba/TravellerTools/internal/dice"
)

const (
	DefaultValue = iota
	SubsectorEmpty
	SubsectorScattered
	SubsectorDispersed
	SubsectorAverage
	SubsectorCrowded
	SubsectorDense
	ObjectNONE
	ObjectStar
	ObjectBrownDwarf
	ObjectRoguePlanet
	ObjectRogueGasGigant
	ObjectNeutronStar
	ObjectNebula
	ObjectBlackHole
)

type GenerationState struct {
	Dice        *dice.Dicepool
	SystemName  string
	CurrentStep int
	NextStep    int
	System      *StarSystem
}

type Generator interface {
}

func NewGenerator(name string) (*GenerationState, error) {
	gs := GenerationState{}
	gs.Dice = dice.New().SetSeed(name)
	gs.Dice.Vocal()
	gs.SystemName = name
	gs.NextStep = 1
	sts, _ := gs.NewStarSystem(SubsectorDense)
	gs.System = sts
	return &gs, nil
}

func (gs *GenerationState) NewStarSystem(stsType int) (*StarSystem, error) {
	ss := StarSystem{}
	ss.subsectorType = stsType
	return &ss, nil
}

func (gs *GenerationState) Step01() {

}

type StarSystem struct {
	subsectorType int
	ObjectType    int
}

func (gs *GenerationState) RollD100() {
	gs.Dice.Roll("1d100")
}

/*
TESTRUN

*/
