package systemgeneration

import "fmt"

func (gs *GenerationState) Step05() error {
	if gs.NextStep != 5 {
		return fmt.Errorf("not actual step")
	}
	numRoll := gs.Dice.Roll("1d10").Sum()
	for i, star := range gs.System.Stars {
		if star.num != -1 {
			continue
		}
		gs.System.Stars[i].num += numRoll
		gs.ConcludedStep = 5
		switch gs.System.Stars[i].class {
		case "O", "B", "A", "F", "G", "K", "M":
			gs.NextStep = 6
		case "L", "T", "Y":
			gs.NextStep = 7
		}
	}
	switch gs.NextStep {
	case 6, 7:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}
