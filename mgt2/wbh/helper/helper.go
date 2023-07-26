package helper

import "github.com/Galdoba/TravellerTools/pkg/dice"

func FluxFactor(dice *dice.Dicepool, factor float64) float64 {
	fl := dice.Flux()
	return factor * float64(fl)
}

func EnsureMinMax(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
