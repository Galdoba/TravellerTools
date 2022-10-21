package systemgeneration

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

func (gs *GenerationState) Step20() error {
	if gs.NextStep != 20 {
		return fmt.Errorf("not actual step")
	}
	printSystem(gs)
	switch gs.System.populationType {
	default:
		fmt.Println(gs.System.populationType)
		return fmt.Errorf("gs.System.populationType = %v", gs.System.populationType)
	case PopulationAuto, PopulationON, PopulationOFF:
		if err := gs.PopulateBodies(); err != nil {
			return err
		}
	}

	gs.NextStep = 99
	return nil
}

func populatePlanet(dp *dice.Dicepool, body StellarBody, mwuwp uwp.UWP, popType string) error {
	mwTL := mwuwp.TL()
	switch popType {
	default:
		return fmt.Errorf("Unknown option `%v`", popType)
	case PopulationOFF:
		return nil
	case PopulationON:
		if mwTL < 8 {
			return nil
		}
	}
	return nil
}
