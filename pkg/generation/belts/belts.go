package belts

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type beltData struct {
	beltZone             string //N M C R
	predominantDiameter  int    //в метрах
	maximumPlanetoidSize int    //в километрах
	width                float64
	asteroidPersentage   []int
	//orbitData            orbit.Orbiter
}

func New(dice *dice.Dicepool) *beltData {
	b := beltData{}
	b.predominantDiameter = bodyDiam(dice.Sroll("2d6-2"))
	b.maximumPlanetoidSize = maxBodyDiam(dice.Sroll("1d6"))
	if b.maximumPlanetoidSize*1000 < b.predominantDiameter {
		b.maximumPlanetoidSize = b.predominantDiameter / 1000
	}
	return &b
}

func bodyDiam(i int) int {
	return []int{1, 5, 10, 25, 50, 100, 300, 1000, 5000, 50000, 500000}[i]
}

func maxBodyDiam(i int) int {
	return []int{0, 0, 1, 10, 100, 1000}[i]
}

func OfferBeltOrbit(i int) int {
	return []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}[i]
}
