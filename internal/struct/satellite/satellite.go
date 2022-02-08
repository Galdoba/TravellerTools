package satellite

import "fmt"

type Satellite struct {
	orbitSuffix string
	orbit       int
	multiplier  int
	locked      bool
	comment     string
	uwp         string
}

const (
	Orbit_Ay = iota
	Orbit_Bee
	Orbit_Cee
	Orbit_Dee
	Orbit_Ee
	Orbit_Eff
	Orbit_Gee
	Orbit_Aitch
	Orbit_Eye
	Orbit_Jay
	Orbit_Kay
	Orbit_Ell
	Orbit_Em
	Orbit_En
	Orbit_Oh
	Orbit_Pee
	Orbit_Que
	Orbit_Arr
	Orbit_Ess
	Orbit_Tee
	Orbit_Yu
	Orbit_Vee
	Orbit_Dub
	Orbit_Ex
	Orbit_Wye
	Orbit_Zee
)

/*
орбита спутника напрямую зависит от родителя
*/

func New(uwp string, orbit int) (Satellite, error) {
	so := Satellite{}
	so.orbitSuffix = suffixMap(orbit)
	if so.orbitSuffix == "" {
		return so, fmt.Errorf("orbit index '%v' invalid", orbit)
	}
	so.multiplier = multiplierMap(orbit)
	so.locked = lockedMap(orbit)
	so.comment = commentMap(orbit)
	so.uwp = uwp
	return so, nil
}

func (s *Satellite) Suffix() string {
	return s.orbitSuffix
}

func (s *Satellite) Orbit() int {
	return s.orbit
}

func (s *Satellite) UWP() string {
	return s.uwp
}

func suffixMap(orbit int) string {
	sm := make(map[int]string)
	sm[Orbit_Ay] = "Ay"
	sm[Orbit_Bee] = "Bee"
	sm[Orbit_Cee] = "Cee"
	sm[Orbit_Dee] = "Dee"
	sm[Orbit_Ee] = "Ee"
	sm[Orbit_Eff] = "Eff"
	sm[Orbit_Gee] = "Gee"
	sm[Orbit_Aitch] = "Aitch"
	sm[Orbit_Eye] = "Eye"
	sm[Orbit_Jay] = "Jay"
	sm[Orbit_Kay] = "Kay"
	sm[Orbit_Ell] = "Ell"
	sm[Orbit_Em] = "Em"
	sm[Orbit_En] = "En"
	sm[Orbit_Oh] = "Oh"
	sm[Orbit_Pee] = "Pee"
	sm[Orbit_Que] = "Que"
	sm[Orbit_Arr] = "Arr"
	sm[Orbit_Ess] = "Ess"
	sm[Orbit_Tee] = "Tee"
	sm[Orbit_Yu] = "Yu"
	sm[Orbit_Vee] = "Vee"
	sm[Orbit_Dub] = "Dub"
	sm[Orbit_Ex] = "Ex"
	sm[Orbit_Wye] = "Wye"
	sm[Orbit_Zee] = "Zee"
	return sm[orbit]
}

func multiplierMap(orbit int) int {
	sm := make(map[int]int)
	sm[Orbit_Ay] = 1
	sm[Orbit_Bee] = 2
	sm[Orbit_Cee] = 3
	sm[Orbit_Dee] = 4
	sm[Orbit_Ee] = 5
	sm[Orbit_Eff] = 6
	sm[Orbit_Gee] = 8
	sm[Orbit_Aitch] = 10
	sm[Orbit_Eye] = 20
	sm[Orbit_Jay] = 30
	sm[Orbit_Kay] = 40
	sm[Orbit_Ell] = 50
	sm[Orbit_Em] = 60
	sm[Orbit_En] = 70
	sm[Orbit_Oh] = 80
	sm[Orbit_Pee] = 100
	sm[Orbit_Que] = 150
	sm[Orbit_Arr] = 200
	sm[Orbit_Ess] = 250
	sm[Orbit_Tee] = 300
	sm[Orbit_Yu] = 400
	sm[Orbit_Vee] = 500
	sm[Orbit_Dub] = 600
	sm[Orbit_Ex] = 700
	sm[Orbit_Wye] = 800
	sm[Orbit_Zee] = 1000
	return sm[orbit]
}

func lockedMap(orbit int) bool {
	sm := make(map[int]bool)
	sm[Orbit_Ay] = true
	sm[Orbit_Bee] = true
	sm[Orbit_Cee] = true
	sm[Orbit_Dee] = true
	sm[Orbit_Ee] = true
	sm[Orbit_Eff] = true
	sm[Orbit_Gee] = true
	sm[Orbit_Aitch] = true
	sm[Orbit_Eye] = true
	sm[Orbit_Jay] = true
	sm[Orbit_Kay] = true
	sm[Orbit_Ell] = true
	sm[Orbit_Em] = true
	return sm[orbit]
}

func commentMap(orbit int) string {
	sm := make(map[int]string)
	sm[Orbit_Ay] = "Ring System or Size < 2"
	sm[Orbit_Bee] = "Ring System or Size < 2"
	sm[Orbit_Cee] = "Ring System or Size < 2"
	sm[Orbit_Arr] = "If Primary is a White Dwarf Size=D this region is a Habitable Zone"
	sm[Orbit_Ess] = "If Primary is a White Dwarf Size=D this region is a Habitable Zone"
	sm[Orbit_Tee] = "If Primary is a White Dwarf Size=D this region is a Habitable Zone"
	sm[Orbit_Yu] = "If Primary is a White Dwarf Size=D this region is a Habitable Zone"
	return sm[orbit]
}
