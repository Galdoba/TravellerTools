package systemgeneration

import "fmt"

func (gs *GenerationState) Step09() error {
	if gs.NextStep != 9 {
		return fmt.Errorf("not actual step")
	}
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
