package profile

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	DT_DEFAULT = iota
	DT_STARPORT
	DT_SIZE
	DT_ATMOSPHERE
	DT_HYDROSPHERE
	DT_TEMPERATURE
	DT_POPULATION
	DT_GOVERMENT
	DT_LAWLEVEL
	DT_TECHLEVEL
	DT_WRONG
)

type uniProfile struct {
	dataGenerationState  int
	dataGenerationDevice *dice.Dicepool
	dataSeed             string
	dataStructure        map[int]string
	data                 []ehex.Ehex
}

func NewProfile(seed string, entityType string) *uniProfile {
	up := uniProfile{}
	up.dataSeed = seed
	up.dataGenerationDevice = dice.New().SetSeed(up.dataSeed)
	return &up
}

func getMapStructure(mapType string) map[int]string {
	mt := make(map[int]string)
	switch mapType {
	default:
		return mt
	case "UPP":

		mt[DT_SIZE] = "Size"
		mt[DT_ATMOSPHERE] = "Atmo"
		mt[DT_HYDROSPHERE] = "Hydr"
		mt[DT_POPULATION] = "Pops"
		mt[DT_GOVERMENT] = "Govr"
		mt[DT_LAWLEVEL] = "Laws"
		mt[DT_STARPORT] = "Port"
	}
	return mt
}
