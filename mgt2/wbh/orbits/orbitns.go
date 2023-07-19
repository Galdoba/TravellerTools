package orbitns

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type orbitn struct {
	OrbN         uint
	Distanse     float64
	Difference   float64
	Eccentricity float64
	Min          float64
	Max          float64
}

func NewOrbit(fl float64, dice *dice.Dicepool) (*orbitn, error) {
	orb := orbitn{}
	on, err := encodeUINT(fl)
	if err != nil {
		return &orb, fmt.Errorf("encodeUINT(%v): %v", fl, err.Error())
	}
	orb.OrbN = on
	orb.Distanse = OR2MKM(orb.OrbN)
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
	return int(o.OrbN / 100000000)
}

func (o *orbitn) AU() float64 {
	return -1
}
