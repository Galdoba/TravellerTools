package systemgeneration

import "fmt"

func (gs *GenerationState) Step10() error {
	fmt.Println("START Step 10")
	if gs.NextStep != 10 {
		return fmt.Errorf("not actual step")
	}
	gs.System.Belts = gs.Dice.Roll("1d6").DM(-3).Sum()
	if gs.System.Belts < 0 {
		gs.System.Belts = 0
	}
	fmt.Println("Belts:", gs.System.Belts)
	gs.ConcludedStep = 10
	gs.NextStep = 11
	switch gs.NextStep {
	case 11:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 10")
	//test
	return nil
}
