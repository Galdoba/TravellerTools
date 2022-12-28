package belts

import "github.com/Galdoba/TravellerTools/pkg/dice"

type beltData struct {
	beltZone             string //N M C R
	predominantDiameter  int    //в метрах
	maximumPlanetoidSize int    //в километрах
	width                float64
	asteroidPersentage   []int
	parentStar           string
	star                 int
	orbit                int
	satelite             int
}

func Generate(dice *dice.Dicepool) []*beltData {
	beltNum := dice.Sroll("1d6-3")
	if beltNum <= 0 {
		return nil
	}
	bData := []*beltData{}
	b := beltData{}
	for i := 0; i < beltNum; i++ {
		b.star = -1
		b.orbit = -1
		b.satelite = -1
		b.predominantDiameter = bodyDiam(dice.Sroll("2d6-2"))
		b.maximumPlanetoidSize = maxBodyDiam(dice.Sroll("1d6"))
		if b.maximumPlanetoidSize*1000 < b.predominantDiameter {
			b.maximumPlanetoidSize = b.predominantDiameter / 1000
		}
		bData = append(bData, &b)
	}
	return bData
}

func bodyDiam(i int) int {
	return []int{1, 5, 10, 25, 50, 100, 300, 1000, 5000, 50000, 500000}[i]
}

func maxBodyDiam(i int) int {
	return []int{0, 0, 1, 10, 100, 1000}[i]
}

func (b *beltData) SetStar(i int) {
	b.star = i
}

func (b *beltData) SetOrbit(i int) {
	b.orbit = i
}

//ORBITER interface

func (b *beltData) SystemPosition() (int, int, int) {
	return b.star, b.orbit, b.satelite
}
