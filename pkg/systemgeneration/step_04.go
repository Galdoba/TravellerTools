package systemgeneration

import "fmt"

func (gs *GenerationState) Step04() error {
	if gs.NextStep != 4 {
		return fmt.Errorf("not actual step")
	}
	starType := ""
	starTypeRoll := gs.Dice.Roll("1d100").Sum()
	tn := []int{}
	switch gs.System.starSystemType {
	default:
		return fmt.Errorf("unknown gs.System.starSystemType (%v)", gs.System.starSystemType)
	case StarSystemRealistic:
		tn = []int{80, 88, 94, 97, 98, 99, 100}
	case StarSystemSemiRealistic:
		tn = []int{50, 77, 90, 97, 98, 99, 100}
	case StarSystemFantastic:
		tn = []int{25, 50, 75, 97, 98, 99, 100}
	}
	switch {
	case starTypeRoll <= tn[0]:
		starType = "M"
	case starTypeRoll <= tn[1]:
		starType = "K"
	case starTypeRoll <= tn[2]:
		starType = "G"
	case starTypeRoll <= tn[3]:
		starType = "F"
	case starTypeRoll <= tn[4]:
		starType = "A"
	case starTypeRoll <= tn[5]:
		starType = "B"
	case starTypeRoll <= tn[6]:
		starType = "O"
	}
	str := &star{class: starType, num: -1, size: ""}
	gs.System.Stars = append(gs.System.Stars, str)
	gs.NextStep = 5
	gs.ConcludedStep = 4
	switch gs.NextStep {
	case 5:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}
