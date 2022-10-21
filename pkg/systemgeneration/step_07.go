package systemgeneration

import "fmt"

func (gs *GenerationState) Step07() error {
	if gs.NextStep != 7 {
		return fmt.Errorf("not actual step")
	}
	switch gs.System.starPopulation {
	default:
		return fmt.Errorf("star population unexpected")
	case StarPopulationSolo, StarPopulationBinary, StarPopulationTrinary, StarPopulationQuatenary, StarPopulationQuintenary:
		if len(gs.System.Stars) < strSystToNum(gs.System.starPopulation) {
			gs.NextStep = 4
			return nil
		}
		gs.ConcludedStep = 7
		gs.NextStep = 8
		gs.System.Stars = sortStars(gs.System.Stars)
	case StarPopulationUNKNOWN:
		tn := []int{}
		switch gs.System.Stars[0].class {
		case "O", "B", "A":
			tn = []int{10, 90, 98, 99, 100}
		case "F", "G", "K":
			tn = []int{45, 99, 100, 599, 999}
		case "M", "L", "T", "Y", "", "D":
			tn = []int{69, 98, 100, 200, 300}
		}
		strComposRoll := gs.Dice.Roll("1d100").Sum()
		switch {
		case strComposRoll <= tn[0]:
			gs.System.starPopulation = StarPopulationSolo
		case strComposRoll <= tn[1]:
			gs.System.starPopulation = StarPopulationBinary
		case strComposRoll <= tn[2]:
			gs.System.starPopulation = StarPopulationTrinary
		case strComposRoll <= tn[3]:
			gs.System.starPopulation = StarPopulationQuatenary
		case strComposRoll <= tn[4]:
			gs.System.starPopulation = StarPopulationQuintenary
		}
		return nil
	}
	switch gs.NextStep {
	case 4, 8:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func sortStars(stars []*star) []*star {
	strMasses := []float64{}
	for _, str := range stars {
		strMasses = append(strMasses, str.mass)
	}
	newOrder := []*star{}
	for i := 1600001; i > -10; i-- {
		m := roundFloat(float64(i)/100, 2)
		for v, mass := range strMasses {
			if m != mass {
				continue
			}
			newOrder = append(newOrder, stars[v])
		}
	}
	return newOrder
}

func setSize(s star) int {
	ss := 0

	for _, scl := range []string{"L", "T", "Y", "M", "K", "G", "F", "A", "B", "O"} {
		if s.class != scl {
			ss += 10
			continue
		}
		ss -= s.num
		break
	}
	for _, scl := range []string{"", "VI", "V", "IV", "III", "II", "Ib", "Ia"} {
		if s.class != scl {
			ss += 100
			// вса
			continue
		}
		break
	}
	return ss
}

func strSystToNum(str string) int {
	switch str {
	case StarPopulationSolo:
		return 1
	case StarPopulationBinary:
		return 2
	case StarPopulationTrinary:
		return 3
	case StarPopulationQuatenary:
		return 4
	case StarPopulationQuintenary:
		return 5
	default:
		return -1
	}
}
