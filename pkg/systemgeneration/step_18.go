package systemgeneration

import (
	"fmt"
	"math"
)

func (gs *GenerationState) Step18() error {
	fmt.Println("START Step 18")
	if gs.NextStep != 18 {
		return fmt.Errorf("not actual step")
	}
	//ЗАРЕЗЕРВИРОВАНО
	d10, d100 := gs.jumpBorders()
	for i, star := range gs.System.Stars {
		d10Border := &jumpZoneBorder{zone: star.Describe() + " D10", orbit: d10[i]}
		d100Border := &jumpZoneBorder{zone: star.Describe() + " D100", orbit: d100[i]}
		gs.System.Stars[i].orbit[d10[i]] = d10Border
		gs.System.Stars[i].orbit[d100[i]] = d100Border
		gs.System.body = append(gs.System.body, d10Border)
		gs.System.body = append(gs.System.body, d100Border)
	}
	gs.ConcludedStep = 18
	gs.NextStep = 19
	switch gs.NextStep {
	case 19:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	//fmt.Println("END Step 18")
	return nil
}

/*
	AU2Megameters        = 149597.9
	SolDiametrMegameters = 1392.77

*/

func (gs *GenerationState) jumpBorders() ([]float64, []float64) {
	d10 := []float64{}
	d100 := []float64{}
	for _, star := range gs.System.Stars {
		tempSolars := roundFloat(float64(star.temperature)/5780.0, 2)
		starRadius := roundFloat(math.Sqrt(star.luminocity)/tempSolars*tempSolars, 2)
		starDiameter := roundFloat(starRadius*2, 2)
		d10Zone := roundFloat(starDiameter*1392.77*10, 2)
		d100Zone := roundFloat(starDiameter*1392.77*100, 2)
		d10AU := roundFloat(d10Zone/149597.9, 2)
		d10 = append(d10, d10AU)
		d100AU := roundFloat(d100Zone/149597.9, 2)
		d100 = append(d100, d100AU)
	}
	return d10, d100
}
