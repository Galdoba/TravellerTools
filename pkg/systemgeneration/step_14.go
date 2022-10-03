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
	gs.System.RockyPlanets = gs.Dice.Roll("1d6").DM(1).Sum()
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

func (gs *GenerationState) placeGG(markers []orbMarker) error {
	if len(gs.System.GG) > len(markers) {
		return fmt.Errorf("can't place all GGs %v/%v", len(gs.System.GG), len(markers))
	}

	for p, gg := range gs.System.GG {
		fmt.Println("Placing", p, ". . .")
		placed := false
		try := 0
		for !placed {
			try++
			fmt.Println("try", try)
			r := gs.Dice.Roll("1d" + fmt.Sprintf("%v", len(gs.System.Stars))).DM(-1).Sum()
			for n, m := range markers {
				fmt.Println("GO:", m)
				if m.starPos != r {
					fmt.Println("WRONG STAR")
					continue
				}
				fmt.Println(gs.System.Stars[r].orbit[m.orbRad].Describe())
				if !strings.Contains(gs.System.Stars[r].orbit[m.orbRad].Describe(), " ggPossible") {
					fmt.Println("ORBIT NOT FREE")
					continue
				}
				gg.spawnedAtAU = m.orbRad
				gs.System.Stars[r].orbit[m.orbRad] = gg
				fmt.Println("PLACE!")
				fmt.Println(markers[n])
				//markers = removeMarker(markers, n)
				placed = true
				break
			}
		}
		fmt.Println("done")
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
