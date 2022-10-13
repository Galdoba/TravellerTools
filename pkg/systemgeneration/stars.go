package systemgeneration

import (
	"fmt"
	"sort"
	"strings"
)

type star struct {
	class                 string
	num                   int
	size                  string
	rank                  string
	generated             bool
	distanceType          string
	distanceFromPrimaryAU float64
	temperature           int
	mass                  float64
	luminocity            float64
	innerLimit            float64
	habitableLow          float64
	habitableHigh         float64
	snowLine              float64
	outerLimit            float64
	orbit                 map[float64]StellarBody
	orbitDistances        []float64
}

func (s *star) Describe() string {
	return fmt.Sprintf("%v: %v", s.rank, s.Code())
}

func (s *star) Code() string {
	code := fmt.Sprintf("%v%v %v", s.class, s.num, s.size)
	code = strings.TrimSpace(code)
	return code
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

func (s *star) cleanOrbitDistances() {
	s.orbitDistances = []float64{}
	for i := -5; i < 9999999; i++ {
		orb := roundFloat(float64(i)/100, 2)
		if body, ok := s.orbit[orb]; ok == true {
			switch body.(type) {
			case *rockyPlanet, *ggiant, *belt, *jumpZoneBorder:
				s.orbitDistances = append(s.orbitDistances, orb)
			}

		}
	}
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
