package systemgeneration

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

func (gs *GenerationState) Step08() error {
	fmt.Println("START Step 08")
	if gs.NextStep != 8 {
		return fmt.Errorf("not actual step")
	}
	switch gs.System.starPopulation {
	default:
		return fmt.Errorf("imposible population at step 08")

	case StarPopulationSolo, StarPopulationBinary, StarPopulationTrinary, StarPopulationQuatenary, StarPopulationQuintenary:
		for i, star := range gs.System.Stars {
			switch i {
			case 0:
				star.rank = "Primary"
			case 1:
				star.rank = "Secondary"
			case 2:
				star.rank = "Tretiary"
			case 3:
				star.rank = "Quatenary"
			case 4:
				star.rank = "Quintenary"
			}
			if i == 0 {
				gs.System.Stars[0].distanceType = "Primary"
				gs.System.Stars[0].distanceFromPrimaryAU = 0.0
				fmt.Println(gs.System.Stars[i])
				continue
			}
			for star.distanceFromPrimaryAU <= gs.System.Stars[i-1].distanceFromPrimaryAU {
				dist, au := rollDistance(gs.Dice)
				star.distanceType = dist
				star.distanceFromPrimaryAU = au
				if dist == StarDistanceContact {
					star.distanceFromPrimaryAU = gs.System.Stars[i-1].distanceFromPrimaryAU
				}

			}
			fmt.Println(star)
		}
		gs.ConcludedStep = 8
		gs.NextStep = 9
	}
	switch gs.NextStep {
	case 9:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 08")
	return nil
}

func rollDistance(dp *dice.Dicepool) (string, float64) {
	r1 := dp.Roll("1d100").Sum()
	r2 := dp.Roll("1d100").Sum()
	dist := distChart(r1)
	au := auChart(r2, dist)
	return dist, au
}

func distChart(i int) string {
	distChart := []int{10, 30, 50, 80, 100}
	switch {
	case i <= distChart[0]:
		return StarDistanceContact
	case i <= distChart[1]:
		return StarDistanceClose
	case i <= distChart[2]:
		return StarDistanceNear
	case i <= distChart[3]:
		return StarDistanceFar
	case i <= distChart[4]:
		return StarDistanceDistant
	}
	return "DISTANCE UNDEFINED"
}

func auChart(i int, dType string) float64 {
	distChart := []float64{}
	switch dType {
	case StarDistanceContact:
		return -1
	case StarDistanceClose:
		distChart = []float64{0.5, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0, 5.5}
	case StarDistanceNear:
		distChart = []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	case StarDistanceFar:
		distChart = []float64{100, 150, 200, 250, 300, 350, 400, 450, 500, 550}
	case StarDistanceDistant:
		distChart = []float64{600, 750, 1000, 1500, 2000, 2500, 3000, 4000, 5000, 6000}
	}
	switch {
	case i <= 9:
		return distChart[0]
	case i <= 19:
		return distChart[1]
	case i <= 29:
		return distChart[2]
	case i <= 39:
		return distChart[3]
	case i <= 49:
		return distChart[4]
	case i <= 59:
		return distChart[5]
	case i <= 69:
		return distChart[6]
	case i <= 79:
		return distChart[7]
	case i <= 89:
		return distChart[8]
	case i <= 100:
		return distChart[9]
	}
	return 0.0
}
