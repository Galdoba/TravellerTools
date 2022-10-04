package systemgeneration

import (
	"fmt"
	"sort"
	"strings"
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
		gg = len(canPutGGIn)
	}
	fmt.Println(gg, "|", canPutGGIn)
	gs.placeGG(canPutGGIn)
	pl := gs.System.RockyPlanets
	planetMarkers := gs.canPlacePlanets()
	fmt.Println(pl, "|", planetMarkers)
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
	return nil
}

func (gs *GenerationState) canPlacePlanets() [][]orbMarker {
	om := [][]orbMarker{}
	for s, star := range gs.System.Stars {
		om = append(om, []orbMarker{})
		for i := int(star.innerLimit * 1000); i < int(star.outerLimit*1000); i-- {
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
	if len(gs.System.GG) > len(markers) {
		return fmt.Errorf("can't place all GGs %v/%v", len(gs.System.GG), len(markers))
	}
	gs.debug(fmt.Sprintf("Placing %v Gas Gigants...", len(gs.System.GG)))
	for n, gg := range gs.System.GG {
		gs.debug(fmt.Sprintf("Placing Gas Gigant %v...", n+1))
		placed := false
		try := 0
		for !placed {
			try++
			r := gs.Dice.Roll("1d" + fmt.Sprintf("%v", len(gs.System.Stars))).DM(-1).Sum()
			for _, m := range markers {
				fmt.Println("GO:", m)
				if m.starPos != r {
					continue
				}
				fmt.Println(gs.System.Stars[r].orbit[m.orbRad].Describe())
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
						}
					}
				}
				gs.System.Stars[r].orbit[gg.spawnedAtAU] = gg
				placed = true
				break
			}
		}
	}
	for s, _ := range gs.System.Stars {
		for k, v := range gs.System.Stars[s].orbit {
			if !strings.Contains(v.Describe(), "empty orbit") {
				continue
			}
			gs.System.Stars[s].orbit[k] = &bodyHolder{strings.TrimSuffix(v.Describe(), " ggPossible")}
		}
	}
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
	}
	return nil
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
		for _, v := range star.OrbitsSorted() {
			if strings.Contains(star.orbit[v.orbRad].Describe(), " ggPossible") {
				cp = append(cp, orbMarker{i, v.orbRad})
			}
		}
	}
	return cp
}

func (s *star) OrbitsSorted() []orbMarker {
	unsorted := []float64{}
	for k, _ := range s.orbit {
		unsorted = append(unsorted, k)
	}
	sort.Float64s(unsorted)
	sorted := []orbMarker{}
	for _, orb := range unsorted {
		for k, _ := range s.orbit {
			if orb == k {
				sorted = append(sorted, orbMarker{s.num, k})
			}
		}
	}
	return sorted
}

type orbMarker struct {
	starPos int
	orbRad  float64
}
