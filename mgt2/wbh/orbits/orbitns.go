package orbitns

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/mgt2/wbh/helper"
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

/*
3 типа орбиты:
1. центр
2. вокруг центра
3. вокруг тела

А. Вокруг точки
Б. Вокруг тела
//Прайм вращается вокруг точки с дистанцией 0
//вторая звезда вращается вокруг центра с дистанцией N
//Компаньён вращается вокруг одной из звезд

дистанция AU - от 0.0 до 20.0
1 - центр
2 - звезда
*/

type OrbitN struct {
	ReferenceCode string
	AU            float64
	OrbitNum      float64
	Distance      int //микроAU (1/1000000)
	Eccentricity  float64
	MinSeparation float64
	MaxSeparation float64
	Period        string
	AsignedBody   string
}

func New(fl float64) *OrbitN {
	orb := OrbitN{}
	orb.Distance = encodeFL2INT(fl)
	orb.AU = orn2au(orb.Distance)
	orb.OrbitNum = float64(int(fl*1000)) / 1000
	return &orb
}

func (orb *OrbitN) SetReference(refCode string) {
	orb.ReferenceCode = refCode
}

func (orb *OrbitN) DetermineEccentrisity(dice *dice.Dicepool, dm int) {
	fr := helper.EnsureMinMax(dice.Sroll("2d6")+dm, 5, 12)
	sRollCode := "1d6"
	base := 0.0
	delim := 0.0
	switch fr {
	case 5:
		base = -0.001
		delim = 1000
	case 6, 7:
		base = 0.000
		delim = 200
	case 8, 9:
		base = 0.03
		delim = 100
	case 10:
		base = 0.05
		delim = 20
	case 11:
		base = 0.05
		delim = 20
		sRollCode = "2d6"
	case 12:
		base = 0.3
		delim = 20
		sRollCode = "2d6"
	}
	sr := float64(dice.Sroll(sRollCode))
	orb.Eccentricity = base + (sr / delim)
	ecc := int(orb.Eccentricity * 1000)
	mns := int(orb.AU*1000) * (1000 - ecc)
	mxs := int(orb.AU*1000) * (1000 + ecc)
	orb.MinSeparation = float64(mns) / 1000
	orb.MaxSeparation = float64(mxs) / 1000
	orb.Eccentricity = float64(int(orb.Eccentricity*1000)) / 1000
	orb.MinSeparation = float64(int(orb.MinSeparation*1000)) / 1000
	orb.MaxSeparation = float64(int(orb.MaxSeparation*1000)) / 1000
}

func orn2au(orbit int) float64 {
	/*
			4.3 => 1.96
		    4=> 1.6
			0.3(3) => 1.2 * 0.4 =>

	*/
	key, frac := decodeINT2FRACK(orbit)
	// _, key, frac, err := decodeUINT(orbit)
	// if err != nil {
	// 	return 0
	// }
	auDistance := tableBaseDistance(key) + (tableDifference(key) * frac / 1000)
	auDistFloat := float64(auDistance) / 1000
	// audisInt := int(auDistance * 1000)
	// auDistFloat = float64(audisInt) / 1000
	return auDistFloat
}

func DetermineStarOrbit(dice *dice.Dicepool, orbCode string) (float64, error) {
	dm := -99
	if orbCode == "Aa" {
		return 0, nil
	}
	if strings.Contains(orbCode, "b") {
		return float64(dice.Sroll("1d6"))/10 + float64(dice.Flux())/100, nil
	}
	if strings.Contains(orbCode, "B") {
		dm = -1
	}
	if strings.Contains(orbCode, "C") {
		dm = 5
	}
	if strings.Contains(orbCode, "D") {
		dm = 11
	}
	r := float64(dice.Sroll("1d6") + dm)
	flux1 := float64(dice.Flux()) / 10
	flux2 := float64(dice.Flux() / 100)
	if dm == -99 {
		return 0, fmt.Errorf("incorrect orbit code '%v'", orbCode)
	}
	if r+flux1+flux2 < 0 {
		return 0.01, nil
	}
	return r + flux1 + flux2, nil
}

func CalculateOrbitalPeriod(au, m1, m2 float64) string {
	y := math.Sqrt(math.Pow(au, 3)/m1 + m2)
	d := 0.0
	h := 0.0
	val := 0.0
	units := "y"
	if y < 1 {
		d = y * 365.25
		units = "d"
	}
	if d != 0 && d < 1 {
		h = y * 8766
		units = "h"
	}
	switch units {
	case "y":
		val = y
	case "d":
		val = d
	case "h":
		val = h
	}
	val = float64(int(val*1000)) / 1000
	return fmt.Sprintf("%v%v", val, units)
}
