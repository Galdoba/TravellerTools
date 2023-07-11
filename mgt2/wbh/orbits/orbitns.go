package orbitns

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type orbitn struct {
	orbN         uint
	distanse     float64
	difference   float64
	eccentricity float64
	min          float64
	max          float64
}

func NewOrbit(fl float64, dice *dice.Dicepool) (*orbitn, error) {
	orb := orbitn{}
	on, err := encodeUINT(fl)
	if err != nil {
		return &orb, fmt.Errorf("encodeUINT(%v): %v", fl, err.Error())
	}
	orb.orbN = on
	orb.distanse = OR2MKM(orb.orbN)
	return &orb, nil
}

func OR2MKM(orbit uint) float64 {
	/*
			4.3 => 1.96
		    4=> 1.6
			0.3(3) => 1.2 * 0.4 =>

	*/
	_, key, frac, err := decodeUINT(orbit)
	if err != nil {
		return 0
	}
	auDistance := tableBaseDistance(key) + (tableDifference(key) * frac / 1000000)
	auDistFloat := float64(auDistance) / 1000000
	return auDistFloat
}

func (o *orbitn) centerpoint() int {
	return int(o.orbN / 100000000)
}

func (o *orbitn) AU() float64 {
	return -1
}
