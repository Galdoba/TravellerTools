package starport

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/utils"
)

type DataSource interface {
	Statistics() string
	PBG() string
}

//FuelCost - bridgeMultiplicator not implemented/ default 1.0
func FuelCost(data DataSource, bridgeMultiplicator float64) int {
	g := getG(data)
	w, s := getWS(data)
	p := (500 * g) * w * s * bridgeMultiplicator
	return int(p)
}

func getG(data DataSource) float64 {
	switch len(data.PBG()) {
	case 3:
		pbg := strings.Split(data.PBG(), "")
		switch pbg[2] {
		case "1", "2", "3", "4", "5", "6":
			return 1.0
		default:
			return 2.0
		}
	case 1:
		if data.PBG() == "G" {
			return 1.0
		}
		return 2.0
	}
	return 2.0
}

func getWS(data DataSource) (float64, float64) {
	uwp := uwp.Inject(data.Statistics())
	w := 1.0
	if uwp.Atmo() < 3 {
		w = 2.0
	}
	s := 3.0
	switch uwp.Starport() {
	case "A":
		s = 0.5
	case "B":
		s = 0.75
	case "C":
		s = 1
	case "D":
		s = 1.5
	case "E":
		s = 2
	}
	return w, s
}

func ItemProceAdjustment(data DataSource) float64 {
	s, l := getSL(data)
	x := 1.0 * s * l
	return utils.RoundFloat64(x, 2)
}

func getSL(data DataSource) (float64, float64) {
	s := 3.0
	uwp := uwp.Inject(data.Statistics())
	switch uwp.Starport() {
	case "A":
		s = 1
	case "B":
		s = 1.25
	case "C":
		s = 1.5
	case "D":
		s = 1.75
	case "E":
		s = 2
	}
	l := 2.0
	switch uwp.Laws() {
	case 0:
		l = 1.5
	case 1, 2, 3:
		l = 1.25
	case 4, 5, 6, 7:
		l = 1
	case 8, 9, 10:
		l = 1.5
	}
	return s, l
}
