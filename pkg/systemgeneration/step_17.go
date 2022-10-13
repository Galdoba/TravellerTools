package systemgeneration

import "fmt"

func (gs *GenerationState) Step17() error {
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
	return nil
}
