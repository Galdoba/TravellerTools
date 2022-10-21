package systemgeneration

import "fmt"

func (gs *GenerationState) Step03() error {
	if gs.NextStep != 3 {
		return fmt.Errorf("not actual step")
	}
	starType := ""
	dwarfTypeRoll := gs.Dice.Roll("1d100").Sum()
	switch {
	case dwarfTypeRoll <= 50:
		starType = "L"
	case dwarfTypeRoll <= 75:
		starType = "T"
	case dwarfTypeRoll <= 100:
		starType = "Y"
	}
	str := &star{class: starType, num: -1, size: ""}
	gs.System.Stars = append(gs.System.Stars, str)
	gs.NextStep = 5
	gs.ConcludedStep = 3
	switch gs.NextStep {
	case 5:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}
