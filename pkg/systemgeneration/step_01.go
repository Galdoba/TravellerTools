package systemgeneration

import "fmt"

func (gs *GenerationState) Step01() error {
	fmt.Println("START Step 01")
	if gs.NextStep != 1 {
		fmt.Errorf("not actual step")
	}
	tn := 0
	switch gs.System.subsectorType {
	case SubsectorEmpty:
		tn = 5
	case SubsectorScattered:
		tn = 20
	case SubsectorDispersed:
		tn = 35
	case SubsectorAverage:
		tn = 50
	case SubsectorCrowded:
		tn = 60
	case SubsectorDense:
		tn = 75
	}
	gs.System.ObjectType = ObjectNONE
	gs.debug("ObjectType set as NONE")
	presenceRoll := gs.Dice.Roll("1d100").Sum()
	if presenceRoll <= tn {
		gs.System.ObjectType = ObjectPRESENT
		gs.debug("ObjectType set as PRESENT")
	}
	switch gs.System.ObjectType {
	default:
		return fmt.Errorf("system ObjectType is invalid")
	case ObjectNONE:
		gs.debug("ObjectType Is not in the hex: END GENERATION")
		gs.NextStep = 20
	case ObjectPRESENT:
		gs.NextStep = 2
	}
	gs.ConcludedStep = 1
	fmt.Println("END Step 01")
	return nil
}
