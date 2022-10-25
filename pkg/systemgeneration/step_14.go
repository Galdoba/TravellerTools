package systemgeneration

import (
	"fmt"
	"strings"
)

const (
	habZoneInner     = "Inner"
	habZoneHabitable = "Habitable"
	habZoneOuter     = "Outer"
)

func (gs *GenerationState) Step14() error {
	if gs.NextStep != 14 {
		return fmt.Errorf("not actual step")
	}
	for _, star := range gs.System.Stars {
		gs.System.body = append(gs.System.body, star) //   .orbit[star.distanceFromPrimaryAU] = star.orbitData()
	}
	gs.System.RockyPlanets = gs.Dice.Roll("2d6").DM(0).Sum()
	for i := 1; i < gs.System.RockyPlanets+1; i++ {
		gs.System.body = append(gs.System.body, &rockyPlanet{num: i})
	}
	for i := 1; i < gs.System.Belts+1; i++ {
		gs.System.body = append(gs.System.body, &belt{num: i})
	}

	gg := gs.System.GasGigants
	canPutGGIn := gs.canPlaceGGin()
	if gg > len(canPutGGIn) {
		gg = len(canPutGGIn)
	}

	if err := gs.placeGG(canPutGGIn); err != nil {
		return err
	}
	if err := gs.placePlanets(); err != nil {
		return err
	}

	if err := gs.placeBelts(); err != nil {
		return err
	}

	gs.ConcludedStep = 14
	gs.NextStep = 15
	switch gs.NextStep {
	case 15:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func (gs *GenerationState) placePlanets() error {
	pl := gs.System.RockyPlanets
	planetMarkers := gs.canPlaceBodies()
	if pl > freeSlots(planetMarkers) {
		pl = freeSlots(planetMarkers)
	}
	for i := 0; i < pl; i++ {
		st := gs.Dice.Roll(fmt.Sprintf("1d%v", len(gs.System.Stars))).DM(-1).Sum()
		star := gs.System.Stars[st]
		if len(star.orbitDistances) < 1 {
			i--
			continue
		}

		try := 0
		//placed := false
		for {
			orb := gs.Dice.Roll(fmt.Sprintf("1d%v", len(star.orbitDistances))).DM(-1).Sum()
			if orb-try < 0 {
				i--
				break
			}
			dist := star.orbitDistances[orb-try]
			if v, ok := star.orbit[dist]; ok == true {
				if !strings.Contains(v.Describe(), "empty orbit") {
					try++
					continue
				}
				planet := &rockyPlanet{num: i + 1, star: star.Describe(), orbit: dist, eccentricity: 0.0, comment: "Rocky Planet"}
				if gs.Dice.Roll("1d6").Sum() < 4 {
					roll := gs.Dice.Roll("2d10").Sum()
					planet.eccentricity = eccentricity(roll)
				}
				planet.habZone = star.habZone(dist)
				gs.System.Stars[st].orbit[dist] = planet
				//	placed = true
				break
			}
		}
	}
	return nil
}

func (star *star) habZone(dist float64) string {
	habZone := "UNDEFINED"
	switch {
	default:
	case dist-star.habitableLow < 0:
		habZone = habZoneInner
	case (dist-star.habitableLow >= 0) && (dist-star.habitableHigh) < 0:
		habZone = habZoneHabitable
	case (dist - star.habitableHigh) >= 0:
		habZone = habZoneOuter
	}
	return habZone
}

func eccentricity(roll int) float64 {
	switch roll {
	default:
		return 0.0
	case 2:
		return 0.002
	case 3:
		return 0.003
	case 4:
		return 0.004
	case 5:
		return 0.005
	case 6:
		return 0.006
	case 7:
		return 0.007
	case 8:
		return 0.008
	case 9:
		return 0.009
	case 10:
		return 0.010
	case 11:
		return 0.020
	case 12:
		return 0.030
	case 13:
		return 0.040
	case 14:
		return 0.050
	case 15:
		return 0.070
	case 16:
		return 0.100
	case 17:
		return 0.125
	case 18:
		return 0.150
	case 19:
		return 0.200
	case 20:
		return 0.250
	}
}

func (gs *GenerationState) canPlaceBodies() [][]orbMarker {
	om := [][]orbMarker{}
	for s, star := range gs.System.Stars {
		om = append(om, []orbMarker{})
		//TODO: переписать чтобы шло по orbitDistance
		for _, orb := range star.orbitDistances {
			if v, ok := star.orbit[orb]; ok == true {
				if strings.Contains(v.Describe(), "empty orbit") {
					om[s] = append(om[s], orbMarker{s, orb})
				}
			}
		}
	}
	return om
}

func freeSlots(om [][]orbMarker) int {
	sl := 0
	for _, arr := range om {
		sl = sl + len(arr)
	}
	return sl
}

func (gs *GenerationState) placeGG(markers []orbMarker) error {
	for n, gg := range gs.System.GG {
		if n >= len(markers) {
			continue
		}
		placed := false
		try := 0
		for !placed {
			try++
			r := gs.Dice.Roll("1d" + fmt.Sprintf("%v", len(gs.System.Stars))).DM(-1).Sum()
			for _, m := range markers {

				if m.starPos != r {
					continue
				}
				if !strings.Contains(gs.System.Stars[r].orbit[m.orbRad].Describe(), " ggPossible") {
					continue
				}
				gg.spawnedAtAU = m.orbRad
				if gg.migratedToAU != 0.0 {
					starInner := gs.System.Stars[r].innerLimit
					placeTo := gg.spawnedAtAU - gg.migratedToAU
					if placeTo < starInner {
						placeTo = starInner
					}
					gg.spawnedAtAU = roundFloat(placeTo, 2)
					//TODO: переписать чтобы шло по orbitDistance
					for i := 0; i < 10000000; i++ {
						orbFl := float64(i) / 1000
						if body, ok := gs.System.Stars[r].orbit[orbFl]; ok == true {
							if orbFl < starInner {
								continue
							}
							if orbFl > m.orbRad {
								break
							}
							if !strings.Contains(body.Describe(), "Mainworld") {
								delete(gs.System.Stars[r].orbit, orbFl)
								markers = gs.canPlaceGGin()
							}
						}
					}
				}
				gs.System.Stars[r].orbit[gg.spawnedAtAU] = gg
				placed = true
				break
			}
		}
	}
	return nil
}

func removeMarker(markers []orbMarker, n int) []orbMarker {
	return append(markers[:n], markers[n+1:]...)
}

func (gs *GenerationState) canPlaceGGin() []orbMarker {
	cp := []orbMarker{}
	for i, star := range gs.System.Stars {
		for _, k := range star.orbitDistances {
			if v, ok := star.orbit[k]; ok == true {
				if strings.Contains(v.Describe(), " ggPossible") {
					cp = append(cp, orbMarker{i, k})
				}
			}

		}
	}
	return cp
}

type orbMarker struct {
	starPos int
	orbRad  float64
}
