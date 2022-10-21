package calculations

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/struct/star"
)

func GenerateNewStellar(systemName string) string {
	//Declaration

	dp := dice.New().SetSeed(systemName)
	stellar := ""
	sysComp := newSystemComposition(dp.SetSeed(systemName + "systemComp"))
	sp := ""
	dec := ""
	sz := ""
	//Process
	primaryFluxSp := dp.Flux()
	primaryFluxSz := dp.Flux()
	code := ""
	for _, st := range sysComp {
		try := 0
		dec = strconv.Itoa(dp.Roll("1d10").Sum() - 1)
		accepted := false
		for !accepted {
			try++
			code = ""
			switch st {
			case star.Category_Primary:
				sp = spectral(primaryFluxSp)
				index := primaryFluxSz
				sz = size(sp, index)
			default:
				spIndex := primaryFluxSp + dp.Roll("1d6").Sum() - 1
				szIndex := primaryFluxSz + dp.Roll("1d6").Sum() + 2
				sp = spectral(spIndex)
				sz = size(sp, szIndex)
			}
			code = star.EncodeStellar(sp, dec, sz)

			code = star.FixCode(code)
			if star.CodeValid(code) {
				accepted = true
			}
		}
		stellar = stellar + code + " "
	}
	stellar = strings.TrimSuffix(stellar, " ")
	return stellar
}

func newSystemComposition(dp *dice.Dicepool) []int {
	sysComp := []int{}
	sysComp = append(sysComp, star.Category_Primary)
	if dp.Flux() >= 3 {
		sysComp = append(sysComp, star.Category_PrimaryCompanion)
	}
	for _, categ := range []int{star.Category_Close, star.Category_Near, star.Category_Far} {
		if dp.Flux() >= 3 {
			sysComp = append(sysComp, categ)
			if dp.Flux() >= 3 {
				sysComp = append(sysComp, categ+1)
			}
		}
	}
	return sysComp
}

func spectral(index int) string {
	switch {
	case index < -6:
		index = -6
	case index > 8:
		index = 8
	}
	sp := []string{"OB", "A", "A", "F", "F", "G", "G", "K", "K", "M", "M", "M", "BD", "BD", "BD"}
	return sp[index+6]
}

func size(spec string, index int) string {
	switch {
	case spec == "BD":
		return ""
	case index < -6:
		index = -6
	case index > 8:
		index = 8

	}
	sizeMap := make(map[string][]string)
	sizeMap["O"] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "V", "V", "V", "IV", "D", "IV", "IV", "IV"}
	sizeMap["B"] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "III", "V", "V", "IV", "D", "IV", "IV", "IV"}
	sizeMap["A"] = []string{"Ia", "Ia", "Ib", "II", "III", "IV", "V", "V", "V", "V", "V", "D", "V", "V", "V"}
	sizeMap["F"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	sizeMap["G"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	sizeMap["K"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	sizeMap["M"] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	return sizeMap[spec][index+6]
}
