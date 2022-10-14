package systemgeneration

import (
	"fmt"
	"strconv"
	"strings"
)

func (gs *GenerationState) Step09() error {
	if gs.NextStep != 9 {
		return fmt.Errorf("not actual step")
	}
	if err := gs.callImport("PBG"); err != nil {
		return err
	}
	if gs.System.GasGigants == -1 {
		switch gs.System.Stars[0].class {
		case "O", "B", "A", "F", "G", "K", "M":
			gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-2).Sum()
		case "L":
			gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-4).Sum()
		case "T":
			gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-5).Sum()
		case "Y":
			gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-6).Sum()
		}
	}
	if gs.System.GasGigants < 0 {
		gs.System.GasGigants = 0
	}

	gs.ConcludedStep = 9
	gs.NextStep = 10
	switch gs.NextStep {
	case 10:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func (gs *GenerationState) injectPBG(pbg string) error {
	data := strings.Split(pbg, "")
	if len(data) != 3 {
		return fmt.Errorf("cannot inject PBG data: len(%v) != 3", data)
	}
	for i, d := range data {
		switch i {
		case 1:
			beltNum, err := strconv.Atoi(d)
			if err != nil {
				return fmt.Errorf("cannot inject Belt data: (%v)", d)
			}
			gs.System.Belts = beltNum
		case 2:
			ggNum, err := strconv.Atoi(d)
			if err != nil {
				return fmt.Errorf("cannot inject Gas Gigants data: (%v)", d)
			}
			gs.System.GasGigants = ggNum
		}
	}
	return nil
}
