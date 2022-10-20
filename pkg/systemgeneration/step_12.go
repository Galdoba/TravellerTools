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
	}
	// //fmt.Printf("%v\n", mwuwp.String())
	// switch mwuwp.Size() {
	// case 0:
	// 	fmt.Println("Designate MW as asteriod or worldlet")
	// default:
	// 	fmt.Println("Designate MW as planet or moon")
	// }
	fmt.Println(gs.System.Stars[0].orbitDistances)
	fmt.Println(gs.System.Stars[0].habitableLow, gs.System.Stars[0].habitableHigh)
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

func (gs *GenerationState) setOrbitSpots() error {
	for _, star := range gs.System.Stars {
		orb := 0
		star.orbit = make(map[float64]StellarBody)
		// gs.debug(fmt.Sprintf("Star %v", i))
		// gs.debug(fmt.Sprintf("-------"))
		// gs.debug(fmt.Sprintf("%v", star.innerLimit))
		// gs.debug(fmt.Sprintf("%v", star.habitableLow))
		// gs.debug(fmt.Sprintf("%v", star.habitableHigh))
		// gs.debug(fmt.Sprintf("%v", star.snowLine))
		// gs.debug(fmt.Sprintf("%v", star.outerLimit))
		// gs.debug(fmt.Sprintf("-------"))
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
