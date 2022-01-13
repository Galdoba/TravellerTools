package stellar

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/struct/star"
)

type StarSystem struct {
	Sun       star.Star
	Companion star.Star
	Body      []PlanetaryBody
}

type PlanetaryBody interface {
	Orbit() int //скорее всего да
	Name() string
	//Distance() float64 //скорее всего нет
}

//Distance - расчитывает расстояние тела от центра массы главной звезды
func Distance(pb PlanetaryBody) (float64, error) {
	orb := pb.Orbit()
	switch {
	case orb < 0:
		return -1.0, fmt.Errorf("orbit is negaive")
	case orb > 20:
		return -1.0, fmt.Errorf("orbit is in another hex")
	default:
		return pb.decimalOrbit(), nil
	}
}

func decimalOrbit(pb *PlanetaryBody) float64 {
	dp := dice.New().SetSeed(pb.Name())
	fl := dp.Flux() + 5
	switch pb.Orbit() {
	case 0:

	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:

	case 9:

	case 10:

	case 11:

	case 12:

	case 13:

	case 14:

	case 15:

	case 16:

	case 17:

	case 18:

	case 19:

	case 20:

	}
}

/*
Планетарным телом может быть:

-тело
--звезда-компаньён
--Газовый Гигант
--Обычная
--Астеройдный Пояс

stellar.PlanetaryPosition(Star (Mass), Body (Distance), Date.Day())


*/
