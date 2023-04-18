package systemgeneration

import "fmt"

func (gs *GenerationState) Step02() error {
	if gs.NextStep != 2 {
		return fmt.Errorf("not actual step")
	}
	typeRoll := gs.Dice.Roll("1d100").Sum()
	switch {
	case typeRoll <= 80:
		gs.System.ObjectType = ObjectStar
		gs.NextStep = 4
	case typeRoll <= 88:
		gs.System.ObjectType = ObjectBrownDwarf
		gs.NextStep = 3
	case typeRoll <= 94:
		gs.System.ObjectType = ObjectRoguePlanet
		gs.NextStep = 15
	case typeRoll <= 97:
		gs.System.ObjectType = ObjectRogueGasGigant
		gs.NextStep = 13
	case typeRoll <= 98:
		gs.System.ObjectType = ObjectNeutronStar
		gs.NextStep = 18
	case typeRoll <= 99:
		gs.System.ObjectType = ObjectNebula
		gs.NextStep = 18
	case typeRoll <= 100:
		gs.System.ObjectType = ObjectBlackHole
		gs.System.Stars = append(gs.System.Stars, BlackHole())
		gs.NextStep = 20
	}
	switch gs.System.ObjectType {
	case ObjectBlackHole:
		//return fmt.Errorf("Stop %v", gs.System.ObjectType)
		//gs.System.Stars = append(gs.System.Stars, &star{})
		//gs.setOrbitSpots()
	case ObjectRoguePlanet, ObjectRogueGasGigant, ObjectNebula:
		return fmt.Errorf("Stop %v", gs.System.ObjectType)
		//gs.System.Stars = append(gs.System.Stars, &star{})
		//gs.setOrbitSpots()
	}
	gs.ConcludedStep = 2
	switch gs.NextStep {
	case 4, 3, 15, 13, 18, 20:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}
