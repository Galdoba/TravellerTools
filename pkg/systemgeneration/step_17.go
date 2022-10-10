package systemgeneration

import "fmt"

func (gs *GenerationState) Step17() error {
	fmt.Println("START Step 17")
	if gs.NextStep != 17 {
		return fmt.Errorf("not actual step")
	}
	//ЗАРЕЗЕРВИРОВАНО
	gs.ConcludedStep = 17
	gs.NextStep = 18
	switch gs.NextStep {
	case 18:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	//fmt.Println("END Step 17")
	return nil
}
