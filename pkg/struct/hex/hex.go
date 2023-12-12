package hex

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	UNDEFINED = iota
	DENSITY_EXTRA_GALACTIC
	DENSITY_RIFT
	DENSITY_SPARSE
	DENSITY_SCATTERED
	DENSITY_STANDARD
	DENSITY_DENSE
	DENSITY_CLUSTER
	DENSITY_CORE
	DENSITY_ENSURED_debug
	//
	KEY_Primary = "Primary Object"
)

type hex struct {
	Coord_key             string
	Nebula                int
	Density               int
	CentralObjectClassKey string
	Empty                 bool
	ssPresenceLock        bool
}

func New(coord_key string, density int) *hex {
	h := hex{}
	h.Density = density
	h.Coord_key = coord_key
	return &h
}

func (h *hex) RollStarSystemPresence(dice *dice.Dicepool) error {
	if h.ssPresenceLock {
		return fmt.Errorf("star system presense already defined")
	}
	if !rollStarPresence(dice, h.Density) {
		h.Empty = true
	}
	h.ssPresenceLock = true
	return nil
}

func (h *hex) RollCentralObject(dice *dice.Dicepool) error {
	if h.CentralObjectClassKey != "" {
		return fmt.Errorf("central object alredy defined")
	}
	switch h.Empty {
	case true:
		h.CentralObjectClassKey = neutronStarRoll(dice, h.Density)
	case false:
		h.CentralObjectClassKey = "Star"
	}
	return nil
}

func neutronStarRoll(dice *dice.Dicepool, dens int) string {
	r1 := dice.Sroll("3d6")
	switch r1 {
	case 18:
		if rollStarPresence(dice, dens) {
			r2 := dice.Sroll("2d6")
			if r2 >= 11 {
				return "BH"
			}
			return "NS"
		}
	default:
	}
	return whiteDwarfRoll(dice, dens)

}

func whiteDwarfRoll(dice *dice.Dicepool, dens int) string {
	switch dice.Sroll("1d6") {
	case 6:
		if rollStarPresence(dice, dens) {
			return "D"
		}
	default:
	}
	return brownDwarfRoll(dice, dens)
}
func brownDwarfRoll(dice *dice.Dicepool, dens int) string {
	if rollStarPresence(dice, dens) {
		return "BD"
	}
	return "NONE"
}

func rollStarPresence(dice *dice.Dicepool, density int) bool {
	switch density {
	default:
		panic(fmt.Sprintf("star dencity unknown: %v", density))
	case DENSITY_CORE:
		if dice.Sroll("2d6") <= 11 { //91%
			return true
		}
	case DENSITY_CLUSTER:
		if dice.Sroll("1d6") <= 5 { //83%
			return true
		}
	case DENSITY_DENSE:
		if dice.Sroll("1d6") <= 4 { //66%
			return true
		}
	case DENSITY_STANDARD:
		if dice.Sroll("1d6") <= 3 { //50%
			return true
		}
	case DENSITY_SCATTERED:
		if dice.Sroll("1d6") <= 2 { //33%
			return true
		}
	case DENSITY_SPARSE:
		if dice.Sroll("1d6") <= 1 { //17%
			return true
		}
	case DENSITY_RIFT:
		if dice.Sroll("2d6") <= 2 { //3%
			return true
		}
	case DENSITY_EXTRA_GALACTIC:
		if dice.Sroll("3d6") <= 3 { //<1%
			return true
		}
	}
	return false
}
