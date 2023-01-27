package profile2

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	PROFILE_PERSONALITY     = "Personality"
	PROFILE_WORLD           = "World"
	Planetary_Physical_Data = "Physical Data"
)

const (
	Starport = iota
	Size
	Atmo
	HZ
	Hydr
	Life
	Resources
	Pops
	Govr
	Laws
	Tech
)

type universalProfile struct {
	dice       *dice.Dicepool
	data       map[int]ehex.Ehex
	hiddenData map[int]bool
	entityType string
}

func New(entity string) (*universalProfile, error) {
	up := universalProfile{}
	up.data = make(map[int]ehex.Ehex)
	up.hiddenData = make(map[int]bool)
	switch entity {
	default:
		return nil, fmt.Errorf("unknown entity type '%v'", entity)
	case PROFILE_WORLD:
		for _, val := range expectedData(entity) {
			up.data[val] = ehex.New().Set("?")
		}
		for _, val := range hiddenData(entity) {
			up.hiddenData[val] = true
		}
	}
	return &universalProfile{}, nil
}

func expectedData(e string) []int {
	switch e {
	default:
		return []int{}
	case PROFILE_WORLD:
		return []int{
			Starport,
			Size,
			Atmo,
			HZ,
			Hydr,
			Life,
			Resources,
			Pops,
			Govr,
			Laws,
			Tech,
		}
	}
}

func hiddenData(e string) []int {
	switch e {
	default:
		return []int{}
	case PROFILE_WORLD:
		return []int{
			HZ,
			Life,
			Resources,
		}
	}
}
