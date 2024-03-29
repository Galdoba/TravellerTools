package systemgeneration

import (
	"fmt"
	"sort"
)

func (gs *GenerationState) Step13() error {
	if gs.NextStep != 13 {
		return fmt.Errorf("not actual step")
	}
	if gs.System.Stars[0].snowLine == -999 {
		gs.System.GasGigants = 0
	}
	for i := 0; i < gs.System.GasGigants; i++ {
		gg := ggiant{}
		switch gs.Dice.Roll("1d6").Sum() {
		case 1, 2, 3:
			gg.descr = GasGigantNeptunian
		case 4, 5, 6:
			gg.descr = GasGigantJovian
		}
		gg.size = sizeOfGG(gs.Dice, gg.descr)
		gg.comment = "Gas Gigant"
		gs.System.GG = append(gs.System.GG, &gg)
	}
	gs.System.GG = sortGGiantsBySize(gs.System.GG)
	if len(gs.System.GG) > 0 && gs.Dice.Roll("1d6").Sum() == 6 {
		gs.System.GG[0].migratedToAU = float64(gs.Dice.Roll("1d100").Sum())
		gs.System.GG[0].descr = "Hot " + gs.System.GG[0].descr
		gs.System.GG[0].comment = "Migrated"
	}
	for i, gg := range gs.System.GG {
		gg.num = i + 1
		gs.System.body = append(gs.System.body, gs.System.GG[i])
	}

	gs.ConcludedStep = 13
	gs.NextStep = 14
	switch gs.NextStep {
	case 14:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func sortGGiantsBySize(gg []*ggiant) []*ggiant {
	presentSizes := []int{}
	sortedGG := []*ggiant{}
	last := 0
	for _, gg := range gg {
		if gg.size == last {
			continue
		}
		last = gg.size
		presentSizes = append(presentSizes, gg.size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(presentSizes)))
	for _, s := range presentSizes {
		for _, gg := range gg {
			if gg.size == s {
				sortedGG = append(sortedGG, gg)
			}
		}
	}

	return sortedGG
}
