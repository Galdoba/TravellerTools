package systemgeneration

import "fmt"

func (gs *GenerationState) Step11() error {
	if gs.NextStep != 11 {
		return fmt.Errorf("not actual step")
	}
	for i := range gs.System.Stars {
		gs.System.Stars[i].LoadValues()
	}
	if err := gs.adjustStarValues(); err != nil {
		return err
	}
	gs.ConcludedStep = 11
	gs.NextStep = 12
	switch gs.NextStep {
	case 12:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func (gs *GenerationState) adjustStarValues() error {
	switch len(gs.System.Stars) {
	default:
		return fmt.Errorf("not implemented %v stars\n%v", len(gs.System.Stars), gs.System.Stars)
	case 0, 1:
		return nil
	case 2, 3, 4, 5:
		for st := 0; st < len(gs.System.Stars); st++ {
			switch gs.System.Stars[st].distanceType {
			default:
				outerSum := starDistanceToClosest(gs.System.Stars, st) * 0.3
				//outerSum := gs.System.Stars[st].distanceFromPrimaryAU * 0.3
				if outerSum < gs.System.Stars[st].outerLimit {
					gs.System.Stars[st].outerLimit = outerSum

				}
			case StarDistanceContact:
				innerSum := gs.System.Stars[st-1].innerLimit + gs.System.Stars[st].innerLimit
				gs.System.Stars[st-1].innerLimit = innerSum
				gs.System.Stars[1].innerLimit = innerSum
				outerSum := gs.System.Stars[st-1].outerLimit + gs.System.Stars[st].outerLimit
				gs.System.Stars[st-1].outerLimit = outerSum
				gs.System.Stars[st].outerLimit = outerSum
				habLow := gs.System.Stars[st-1].habitableLow + gs.System.Stars[st].habitableLow
				gs.System.Stars[st-1].habitableLow = habLow
				gs.System.Stars[st].habitableLow = habLow
				habHi := gs.System.Stars[st-1].habitableHigh + gs.System.Stars[st].habitableHigh
				gs.System.Stars[st-1].habitableHigh = habHi
				gs.System.Stars[st].habitableHigh = habHi
			case StarDistanceClose: //две звезды в близко и разрывают близкие к ним планеты
				innerSum := gs.System.Stars[st].distanceFromPrimaryAU * 2.5
				gs.System.Stars[st-1].innerLimit = innerSum
				gs.System.Stars[st].innerLimit = innerSum
			}
		}
	}
	return nil
}
