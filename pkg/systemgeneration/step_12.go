package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

func (gs *GenerationState) Step12() error {
	if gs.NextStep != 12 {
		return fmt.Errorf("not actual step")
	}
	gs.setOrbitSpots()
	//ЗАРЕЗЕРВИРОВАНО
	if err := gs.callImport("MW_UWP"); err != nil {
		return err
	}
	mwuwp := uwp.New()
	if gs.System.MW_UWP != "" {
		importedUWP, err := uwp.FromString(gs.System.MW_UWP)
		if err != nil {
			return fmt.Errorf("imported uwp error: %v", err.Error())
		}
		mwuwp = importedUWP
		gs.System.populationType = PopulationON
	}
	wp := gs.SuggestWorldPosition()
	if wp.star != -1 {
		orb := wp.orbit
		star := gs.System.Stars[wp.star]
		star.orbit[orb] = &rockyPlanet{star: star.Describe(), orbit: orb, eccentricity: 0.0, comment: "Rocky Planet Mainworld", habZone: habZoneHabitable, uwpStr: mwuwp.String()}
	}
	gs.ConcludedStep = 12
	gs.NextStep = 13
	switch gs.NextStep {
	case 13:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func (gs *GenerationState) injectMW_UWP(uwp string) error {
	data := strings.Split(uwp, "")
	if len(data) != 9 {
		return fmt.Errorf("cannot inject UWP data: len(%v) != 9", data)
	}
	gs.System.MW_UWP = uwp
	return nil
}

func (gs *GenerationState) injectMW_Seed(seed string) error {
	gs.Dice.SetSeed(seed)
	return nil
}

func (gs *GenerationState) setOrbitSpots() error {
	for _, star := range gs.System.Stars {
		orb := 0
		star.orbit = make(map[float64]StellarBody)
		currentPoint := star.innerLimit
		for currentPoint < star.outerLimit {
			au := roundFloat(currentPoint, 2)
			star.orbit[au] = &bodyHolder{fmt.Sprintf("empty orbit %v", orb)}
			orb++
			d := gs.Dice.Flux()
			multiplicator := 1.0 + float64(d+5)/10
			currentPoint = currentPoint * multiplicator
		}
		star.markClosestToSnowLine()
		star.markPossibleGG()
		star.updateOrbitDistances()
	}
	return nil
}
