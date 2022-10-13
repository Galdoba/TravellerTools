package systemgeneration

import "fmt"

func (gs *GenerationState) Step06() error {
	if gs.NextStep != 6 {
		return fmt.Errorf("not actual step")
	}
	lumRoll1 := gs.Dice.Roll("1d100").Sum()
	lumRoll2 := gs.Dice.Roll("1d10").Sum()

	for i, star := range gs.System.Stars {
		switch star.class {
		case "L", "T", "Y":
			continue
		}
		if star.size != "" {
			continue
		}
		switch {
		case lumRoll1 <= 90:
			gs.System.Stars[i].size = "V"
		case lumRoll1 <= 94:
			gs.System.Stars[i].size = "IV"
		case lumRoll1 <= 96:
			gs.System.Stars[i].size = "D"
		case lumRoll1 <= 99:
			gs.System.Stars[i].size = "III"
		case lumRoll1 <= 100:
			switch {
			case lumRoll2 <= 4:
				gs.System.Stars[i].size = "II"
			case lumRoll2 <= 6:
				gs.System.Stars[i].size = "VI"
			case lumRoll2 <= 8:
				gs.System.Stars[i].size = "Ia"
			case lumRoll2 <= 10:
				gs.System.Stars[i].size = "Ib"
			}
		}
		if gs.System.Stars[i].size == "D" {
			gs.System.Stars[i].class = "D"
			gs.System.Stars[i].size = ""
		}
		// if gs.System.Stars[i].size != "" {
		// 	continue
		// }

		switch gs.System.Stars[i].size {
		case "O", "B", "A", "F":
			if gs.System.Stars[i].size == "VI" {
				gs.System.Stars[i].size = "V"
			}
		}
	}
	gs.ConcludedStep = 6
	gs.NextStep = 7
	switch gs.NextStep {
	case 7:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}
