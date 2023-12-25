package interfaces

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	Unasigned = iota
	ExtraGalactic
	Rift
	Sparce
	Scatered
	Standard
	Dense
	Cluster
	Core
)

type HEX struct {
	RegionType             int
	Population             bool
	populationTestResolved bool
	Rogues                 []string
}

func NewHex() *HEX {
	hex := HEX{}
	hex.RegionType = Unasigned
	return &hex
}

func (h *HEX) Populate(dice *dice.Dicepool) error {
	if h.populationTestResolved {
		return fmt.Errorf("can't populate: test already resolved")
	}
	if h.RegionType == Unasigned {
		region := []int{Unasigned, ExtraGalactic, Rift, Sparce, Scatered, Standard, Dense, Cluster, Core}
		h.RegionType = region[dice.Sroll("1d8")]
	}
	diceCode := []string{"0d6", "3d6", "1d6", "1d6", "1d6", "1d6", "1d6", "1d6", "2d6"}
	tn := []int{0, 3, 2, 1, 2, 3, 4, 5, 11}
	switch h.RegionType {
	default:
		return fmt.Errorf("can't populate: unknown RegionType (%v)", h.RegionType)
	case ExtraGalactic, Rift, Sparce, Scatered, Standard, Dense, Cluster, Core:
		if dice.Sroll(diceCode[h.RegionType]) > tn[h.RegionType] {
			h.populationTestResolved = true
			return nil
		}
		h.Population = true
		h.populationTestResolved = true
	}
	return nil
}
