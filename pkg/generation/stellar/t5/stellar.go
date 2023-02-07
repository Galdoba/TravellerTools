package t5

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type star struct {
	starType        string // буква определяющая спектр OBAFGKM
	spectralDecimal int    //число определяющее близость к типу 0123456789
	sizeClass       string //римское число определяющее размер Ia Ib II III IV V VI D (D - определяет белого карлика) BD (определяет коричнегого карлика)
}

const (
	posUndefined = iota
	posPrimary
	posPrimaryComp
	posClose
	posCloseComp
	posNear
	posNearComp
	posFar
	posFarComp
	posWRONG
)

/*
t5.NewStellar(knownData ...dataFeed) *stellarConstruct
*/

func StarTypeAndSize(dice *dice.Dicepool, primeFlux1, primeFlux2 int, pos int) (string, int, string) {
	spec, size := "?", "?"
	dec := -1
	switch pos {
	default:
		return spec, dec, size
	case posPrimary:
		spec = spectralType(primeFlux1)
		if spec == "OB" {
			spec = strings.Split(spec, "")[dice.Sroll("1d2-1")]
		}
		dec = dice.Sroll("1d10-1")
		size = starSize(primeFlux2, spec)
	case posPrimaryComp, posClose, posCloseComp, posNear, posNearComp, posFar, posFarComp:
		spec = spectralType(primeFlux1 + dice.Sroll("1d6-1"))
		if spec == "OB" {
			spec = strings.Split(spec, "")[dice.Sroll("1d2-1")]
		}
		dec = dice.Sroll("1d10-1")
		size = starSize(primeFlux2+dice.Sroll("1d6+2"), spec)
	}
	return spec, dec, size
}

func spectralType(i int) string {
	spec := "???"
	if i < -6 {
		i = -6
	}
	if i > 8 {
		i = 8
	}
	spec = []string{"OB", "A", "A", "F", "F", "G", "G", "K", "K", "M", "M", "M", "BD", "BD", "BD"}[i+6]

	return spec
}

func starSize(i int, spec string) string {
	if i < -6 {
		i = -6
	}
	if i > 8 {
		i = 8
	}
	sizeArr := []string{}
	switch spec {
	default:
		return "?"
	case "O":
		sizeArr = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "V", "V", "V", "IV", "D", "IV", "IV", "IV"}
	case "B":
		sizeArr = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "III", "V", "V", "IV", "D", "IV", "IV", "IV"}
	case "A":
		sizeArr = []string{"Ia", "Ia", "Ib", "II", "III", "IV", "V", "V", "V", "V", "V", "D", "V", "V", "V"}
	case "F", "G", "K":
		sizeArr = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	case "M":
		sizeArr = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	case "BD", "L", "T", "Y":
		sizeArr = []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	}
	return sizeArr[i+6]
}

func StarOrbit(dice *dice.Dicepool, pos int) (int, int) {
	roll := dice.Sroll("1d6")
	so := -999 //StarOrbit
	mho := 0   //MaxHighOrbit
	switch pos {
	case posPrimaryComp, posCloseComp, posNearComp, posFarComp:
		return 0, 0
	case posPrimary:
		so, mho = -1, 18
		return so, mho
	case posClose:
		so = roll - 1
	case posNear:
		so = roll + 5
	case posFar:
		so = roll + 11
	}
	mho = so - 3
	if mho < 0 {
		mho = 0
	}
	return so, mho
}
