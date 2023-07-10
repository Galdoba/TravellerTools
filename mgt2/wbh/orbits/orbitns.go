package orbitns

import "github.com/Galdoba/TravellerTools/pkg/dice"

type orbitn struct {
	orbN         uint
	distanse     float64
	difference   float64
	eccentricity float64
	min          float64
	max          float64
}

type au float64

type orbN uint

func (on *orbN) full() float64 {
	return float64(*on)
}

func NewOrbit(fo int, dice *dice.Dicepool) *orbitn {
	orb := orbitn{}
	orb.fullOrbit = fo
	return &orb
}

func OR2MKM(orbit float64) float64 {
	/*
			4.3 => 1.96
		    4=> 1.6
			0.3(3) => 1.2 * 0.4 =>

	*/
	key := int(orbit)
	frac := orbit - float64(key)
	auDistance := tableBaseDistance(key) + (tableDifference(key) * frac)
	return auDistance
}

func (o *orbitn) centerpoint() int {
	return int(o.orbN / 100000000)
}

func (o *orbitn) AU() float64 {
	return -1
}
