package systemgeneration

import (
	"fmt"
	"sort"
	"strings"
)

const (
	habZoneInner     = "Inner"
	habZoneHabitable = "Habitable"
	habZoneOuter     = "Outer"
)

func (gs *GenerationState) Step14() error {
	fmt.Println("START Step 14")
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
	gs.setOrbitSpots()
	gs.System.printSystemSheet()

	fmt.Println("РАССТАВЛЯЕ ГГ ТУТ")
	gg := gs.System.GasGigants
	canPutGGIn := gs.canPlaceGGin()
	if gg > len(canPutGGIn) {
		fmt.Println("отказываемся от газовых гигантов")
		fmt.Println(gg, "|", canPutGGIn)
		gg = len(canPutGGIn)
	}

	if err := gs.placeGG(canPutGGIn); err != nil {
		return err
	}
	//pl := gs.System.RockyPlanets
	//planetMarkers := gs.canPlacePlanets()
	gs.placePlanets()

	gs.ConcludedStep = 14
	gs.NextStep = 15
	switch gs.NextStep {
	case 15:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	//fmt.Println("END Step 12")
	return nil
}

func (gs *GenerationState) placePlanets() error {
	pl := gs.System.RockyPlanets
	planetMarkers := gs.canPlacePlanets()
	if pl > freeSlots(planetMarkers) {
		return fmt.Errorf("cannot suport so many planets %v | %v", pl, len(planetMarkers))
	}
	for i := 0; i < pl; i++ {
		st := gs.Dice.Roll(fmt.Sprintf("1d%v", len(gs.System.Stars))).DM(-1).Sum()
		star := gs.System.Stars[st]
		if len(star.orbitDistances) < 1 {
			i--
			continue
		}
		// fmt.Println("DEBUG LOG:")
		// fmt.Println(gs.SystemName)
		// fmt.Println(gs.System.Stars[st])
		// fmt.Printf("pl=%v\ni=%v\nst=%v\nlen(gs.System.Stars[st].orbitDistances)=%v\n", pl, i, st, len(gs.System.Stars[st].orbitDistances))
		// fmt.Println(gs.System.Stars[st].orbitDistances)

		try := 0
		//placed := false
		for {
			orb := gs.Dice.Roll(fmt.Sprintf("1d%v", len(star.orbitDistances))).DM(-1).Sum()
			//gs.debug(fmt.Sprintf("try %v orb %v...", try, orb))
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
				switch {
				default:
					planet.habZone = "UNDEFINED"
				case dist-star.habitableLow < 0:
					planet.habZone = habZoneInner
				case (dist-star.habitableLow >= 0) && (dist-star.habitableHigh) < 0:
					planet.habZone = habZoneHabitable
				case (dist - star.habitableHigh) >= 0:
					planet.habZone = habZoneOuter
				}

				gs.System.Stars[st].orbit[dist] = planet
				//	placed = true
				break
			}
		}
		gs.debug(fmt.Sprintf("planet %v placed...", i+1))
	}
	return nil
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

func (gs *GenerationState) canPlacePlanets() [][]orbMarker {
	om := [][]orbMarker{}
	for s, star := range gs.System.Stars {
		om = append(om, []orbMarker{})
		//TODO: переписать чтобы шло по orbitDistance
		for i := int(star.innerLimit * 1000); i < int(star.outerLimit*1000); i++ {
			orb := float64(i) / 1000
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
	fmt.Println(gs.System.GG)

	// if len(gs.System.GG) > len(markers) {
	// 	return fmt.Errorf("can't place all GGs %v/%v", len(gs.System.GG), len(markers))
	// }
	gs.debug(fmt.Sprintf("Placing %v Gas Gigants...", len(gs.System.GG)))
	for n, gg := range gs.System.GG {
		if n >= len(markers) {
			gs.debug(fmt.Sprintf("Placing Gas Gigant %v aborted...", n+1))
			continue
		}
		gs.debug(fmt.Sprintf("Placing Gas Gigant %v...", n+1))
		placed := false
		try := 0
		for !placed {
			try++
			r := gs.Dice.Roll("1d" + fmt.Sprintf("%v", len(gs.System.Stars))).DM(-1).Sum()
			for _, m := range markers {

				if m.starPos != r {
					continue
				}
				// if _, ok := gs.System.Stars[r].orbit[m.orbRad]; ok == false {
				// 	continue
				// }
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
						if _, ok := gs.System.Stars[r].orbit[orbFl]; ok == true {
							if orbFl < starInner {
								continue
							}
							if orbFl > m.orbRad {
								break
							}
							delete(gs.System.Stars[r].orbit, orbFl)
							markers = gs.canPlaceGGin()
						}
					}
				}
				gs.System.Stars[r].orbit[gg.spawnedAtAU] = gg
				placed = true
				break
			}
		}
	}
	gs.debug(fmt.Sprintf("Cleaning orbits..."))
	return nil
}

func removeMarker(markers []orbMarker, n int) []orbMarker {
	return append(markers[:n], markers[n+1:]...)
}

func (gs *GenerationState) setOrbitSpots() error {
	for i, star := range gs.System.Stars {
		orb := 0
		star.orbit = make(map[float64]StellarBody)
		gs.debug(fmt.Sprintf("Star %v", i))
		gs.debug(fmt.Sprintf("-------"))
		gs.debug(fmt.Sprintf("%v", star.innerLimit))
		gs.debug(fmt.Sprintf("%v", star.habitableLow))
		gs.debug(fmt.Sprintf("%v", star.habitableHigh))
		gs.debug(fmt.Sprintf("%v", star.snowLine))
		gs.debug(fmt.Sprintf("%v", star.outerLimit))
		gs.debug(fmt.Sprintf("-------"))
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

func (s *star) updateOrbitDistances() {
	s.orbitDistances = nil
	for k := range s.orbit {
		if k < s.innerLimit || k > s.outerLimit {
			continue
		}
		s.orbitDistances = append(s.orbitDistances, k)
	}
	sort.Float64s(s.orbitDistances)
}

func (s *star) markClosestToSnowLine() {
	sl := s.snowLine
	if sl == -999 {
		return
	}
	lowest := 999999.0
	candidate := 0.0
	candidateVal := ""
	for k, v := range s.orbit {
		dist := k - sl
		if dist < 0 {
			dist = dist * -1
		}
		if dist < lowest {
			lowest = dist
			candidate = k
			candidateVal = v.Describe()
		}
	}
	if candidate == 0.0 {
		return
	}
	s.orbit[candidate] = &bodyHolder{candidateVal + " CSN"}
}

func (s *star) markPossibleGG() {
	csn := -1.0
	for k, v := range s.orbit {
		if strings.Contains(v.Describe(), " CSN") {
			csn = k
			break
		}
	}
	if csn == -1.0 {
		return
	}
	for k, v := range s.orbit {
		if k >= csn {
			s.orbit[k] = &bodyHolder{v.Describe() + " ggPossible"}
		}
	}
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
