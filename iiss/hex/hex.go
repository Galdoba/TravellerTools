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
	nebulaSector        bool
	nebulaLocal         bool
	isEmpty             bool
	isProtostarSystem   bool
	isDeadSystem        bool
	isClusterSystem     bool
	generationCompleted bool
	density             int
	centralObjectKey    int
}

func New(density int, nebS, nebL, empt, proto, dead, clust bool) *hex {
	h := hex{}
	h.density = density
	h.nebulaSector = nebS
	return &h
}

func (h *hex) GenerateDetails(confirmed map[string]string, dice *dice.Dicepool) error {
	switch confirmed[KEY_Primary] {
	default:

	//	return fmt.Errorf("undefined primary key: %v", confirmed[KEY_Primary])
	case "NONE":
		h.isEmpty = true
		return nil
	}
	if !rollStarPresence(dice, h.density) {
		h.isEmpty = true
		return nil
	}
	return fmt.Errorf("hex.GenerateDetails: initial")
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
