package systemgeneration

import "fmt"

func (gs *GenerationState) Step12() error {
	if gs.NextStep != 12 {
		return fmt.Errorf("not actual step")
	}
	//ЗАРЕЗЕРВИРОВАНО
	gs.ConcludedStep = 12
	gs.NextStep = 13
	switch gs.NextStep {
	case 13:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}
