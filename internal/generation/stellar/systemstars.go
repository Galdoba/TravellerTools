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

//Distance - расчитывает расстояние тела от центра массы главной звезды в AU
func Distance(pb PlanetaryBody) (float64, error) {
	orb := pb.Orbit()
	switch {
	case orb < 0:
		return -1.0, fmt.Errorf("orbit is negaive")
	case orb > 20:
		return -1.0, fmt.Errorf("orbit is in another hex")
	default:
		return decimalOrbit(pb), nil
	}
}

func decimalOrbit(pb PlanetaryBody) float64 {
	dp := dice.New().SetSeed(pb.Name())
	fl := dp.Flux() + 5
	switch pb.Orbit() {
	case 0:
		return []float64{0.15, 0.16, 0.17, 0.18, 0.19, 0.2, 0.22, 0.24, 0.26, 0.28, 0.30}[fl]
	case 1:
		return []float64{0.30, 0.32, 0.34, 0.36, 0.38, 0.4, 0.43, 0.46, 0.49, 0.52, 0.55}[fl]
	case 2:
		return []float64{0.55, 0.58, 0.61, 0.64, 0.67, 0.7, 0.73, 0.76, 0.79, 0.82, 0.85}[fl]
	case 3:
		return []float64{0.85, 0.88, 0.91, 0.94, 0.97, 1.0, 1.06, 1.12, 1.18, 1.24, 1.30}[fl]
	case 4:
		return []float64{1.30, 1.36, 1.42, 1.48, 1.54, 1.6, 1.72, 1.84, 1.96, 2.08, 2.20}[fl]
	case 5:
		return []float64{2.20, 2.32, 2.44, 2.56, 2.68, 2.8, 3.04, 3.28, 3.52, 3.76, 4.00}[fl]
	case 6:
		return []float64{4.0, 4.2, 4.4, 4.7, 4.9, 5.2, 5.6, 6.1, 6.6, 7.1, 7.6}[fl]
	case 7:
		return []float64{7.6, 8.1, 8.5, 9.0, 9.5, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0}[fl]
	case 8:
		return []float64{15, 16, 17, 18, 19, 20, 22, 24, 26, 28, 30}[fl]
	case 9:
		return []float64{30, 32, 34, 36, 38, 40, 43, 47, 51, 54, 58}[fl]
	case 10:
		return []float64{58, 62, 65, 69, 73, 77, 84, 92, 100, 107, 115}[fl]
	case 11:
		return []float64{115, 123, 130, 138, 146, 154, 169, 184, 200, 215, 231}[fl]
	case 12:
		return []float64{231, 246, 261, 277, 292, 308, 338, 369, 400, 430, 461}[fl]
	case 13:
		return []float64{461, 492, 522, 553, 584, 615, 676, 738, 799, 861, 922}[fl]
	case 14:
		return []float64{922, 984, 1045, 1107, 1168, 1230, 1352, 1475, 1598, 1721, 1844}[fl]
	case 15:
		return []float64{1844, 1966, 2089, 2212, 2335, 2458, 2703, 2949, 3195, 3441, 3687}[fl]
	case 16:
		return []float64{3687, 3932, 4178, 4424, 4670, 4916, 5407, 5898, 6390, 6881, 7373}[fl]
	case 17:
		return []float64{7373, 7864, 8355, 8847, 9338, 9830, 10797, 11764, 12731, 13698, 14665}[fl]
	}
	return -1
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

func PlanetaryPosition(mass float64, bodyDistance float64, time int64) (float64, int) {
	return 0, 0
}
