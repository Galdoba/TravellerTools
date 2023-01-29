package size

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/table"
)

func CoreType(dp *dice.Dicepool, size, hz int) string {
	dm := 0
	if size <= 5 && valBetween(hz, -999, 2) {
		dm += 3
	}
	if size >= 6 && valBetween(hz, -999, 1) {
		dm -= 2
	}
	if size >= 6 && valBetween(hz, -1, 1) {
		dm -= 4
	}
	if size <= 5 && valBetween(hz, 2, 999) {
		dm += 9
	}
	if size >= 6 && valBetween(hz, 2, 999) {
		dm += 3
	}
	r := dp.Sroll("2d10") + dm
	tbl := *table.DiceChart(
		table.Row("6-", "Molten"),
		table.Row("7-15", "Rocky"),
		table.Row("16+", "Icy"),
	)
	return tbl.Result(r)
}

func valBetween(a, min, max int) bool {
	if a < min {
		return false
	}
	if a > max {
		return false
	}
	return true
}

func Density(dice *dice.Dicepool, coreType string) float64 {
	tbl := table.DiceChart(
		table.Row("2", "0.86 0.50 0.12"),
		table.Row("3", "0.88 0.52 0.14"),
		table.Row("4", "0.90 0.54 0.16"),
		table.Row("5", "0.92 0.56 0.18"),
		table.Row("6", "0.94 0.58 0.20"),
		table.Row("7", "0.96 0.60 0.22"),
		table.Row("8", "0.98 0.62 0.24"),
		table.Row("9", "1.00 0.64 0.26"),
		table.Row("10", "1.00 0.64 0.28"),
		table.Row("11", "1.00 0.66 0.30"),
		table.Row("12", "1.02 0.68 0.32"),
		table.Row("13", "1.04 0.70 0.34"),
		table.Row("14", "1.06 0.72 0.36"),
		table.Row("15", "1.08 0.74 0.38"),
		table.Row("16", "1.10 0.76 0.40"),
		table.Row("17", "1.12 0.78 0.42"),
		table.Row("18", "1.14 0.80 0.44"),
		table.Row("19", "1.16 0.82 0.46"),
		table.Row("20", "1.18 0.84 0.48"),
	)
	res := tbl.Result(dice.Sroll("2d10"))
	data := strings.Split(res, " ")
	val := ""
	switch coreType {
	default:
		return -1
	case "Molten":
		val = data[0]
	case "Rocky":
		val = data[1]
	case "Icy":
		val = data[2]
	}
	fl, _ := strconv.ParseFloat(val, 64)
	return fl
}

func PlanetaryMass(dense float64, size ehex.Ehex)
