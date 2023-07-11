package stars

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	defauiltVal = iota
	GenerationMethodSpecial
	GenerationMethodUnusual
	TypeVariantTraditional
	TypeVariantRealistic
	tableStarTypeUnselected
	tableStarTypeTraditional
	tableStarTypeRealistic
	tableHot
	tableSpecial
	tableUnusual
	tableGiants
	tablePecuilar

	starType    = "Star Type"
	special     = "Special"
	hot         = "Hot"
	gigants     = "Gigants"
	peculiar    = "Peculiar"
	blackHole   = "Black Hole"
	pulsar      = "Pulsar"
	neutronStar = "Neutron Star"
	nebula      = "Nebula"
	protostar   = "Protostar"
	starcluster = "Star Cluster"
	anomaly     = "Anomaly"
	typeO       = "Type O"
	typeB       = "Type B"
	typeA       = "Type A"
	typeF       = "Type F"
	typeG       = "Type G"
	typeK       = "Type K"
	typeM       = "Type M"
	classIa     = "Class Ia"
	classIb     = "Class Ib"
	classII     = "Class II"
	classIII    = "Class III"
	classIV     = "Class IV"
	classV      = "Class V"
	classVI     = "Class VI"
	classBD     = "Class BD"
	classD      = "Class D"
)

type starsystem struct {
	starGenerationMethod int
	typeTableVariant     int
	primary              star
}

type star struct {
	sttype  string
	class   string
	subtype int
}

func NewStarSystem(dice *dice.Dicepool, starGenerationMethod, tableVariant int) (*starsystem, error) {
	ss := starsystem{}
	switch starGenerationMethod {
	case GenerationMethodUnusual, GenerationMethodSpecial:
	default:
		return &ss, fmt.Errorf("starGenerationMethod unknown (%v)", starGenerationMethod)
	}
	ss.starGenerationMethod = starGenerationMethod
	ss.typeTableVariant = tableVariant

	stRoll := rollTable(dice, tableStarTypeUnselected, ss.typeTableVariant, ss.starGenerationMethod)
	switch stRoll {
	case typeO, typeB, typeA, typeF, typeG, typeK, typeM:
		ss.primary.sttype = stRoll
		ss.primary.class = classV
	case classIV, classVI, classBD, classD:
		ss.primary.class = stRoll
		for !strings.HasPrefix(stRoll, "Type ") {
			stRoll = rollTable(dice, tableStarTypeUnselected, ss.typeTableVariant, ss.starGenerationMethod)
		}
		ss.primary.sttype = stRoll
	case classIa, classIb, classII, classIII:
		ss.primary.class = stRoll
		dm := append([]int{}, 1)
		for !strings.HasPrefix(stRoll, "Type ") {
			stRoll = rollTable(dice, tableStarTypeUnselected, ss.typeTableVariant, ss.starGenerationMethod, dm...)
		}
		ss.primary.sttype = stRoll
	case blackHole, pulsar, neutronStar, nebula, protostar, starcluster, anomaly:
		ss.primary.sttype = stRoll
	default:
		panic(stRoll)
	}
	switch ss.primary.class {
	case classBD:
		ss.primary.sttype = ""
	}
	return &ss, nil
}

func rollTable(dice *dice.Dicepool, table, typeTableVariant, method int, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	if r < 2 {
		r = 2
	}
	if r > 12 {
		r = 12
	}
	if table == tableStarTypeUnselected {
		table = selectTableBy(starType, method, typeTableVariant)
	}
	tableRollResult := determinationTable(table)[r-2]
	switch tableRollResult {
	case starType, hot, special, gigants, peculiar:
		return rollTable(dice, selectTableBy(tableRollResult, method, typeTableVariant), typeTableVariant, method, mods...)
	}
	return tableRollResult
}

func selectTableBy(s string, method, variant int) int {
	switch s {
	case starType:
		switch variant {
		case TypeVariantTraditional:
			return tableStarTypeTraditional
		case TypeVariantRealistic:
			return tableStarTypeRealistic
		}
	case hot:
		return tableHot
	case special:
		switch method {
		case GenerationMethodSpecial:
			return tableSpecial
		case GenerationMethodUnusual:
			return tableUnusual
		}
	case gigants:
		return tableGiants
	case peculiar:
		return tablePecuilar

	}
	return 0
}

func StarTypeRoll(dice *dice.Dicepool, method int, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	if r <= 2 {
		switch method {
		case GenerationMethodSpecial:
			return SpecialRoll(dice, mods...)
		case GenerationMethodUnusual:
			return UnusualRoll(dice, mods...)
		}

	}
	if r <= 6 {
		return typeM
	}
	if r <= 8 {
		return typeK
	}
	if r <= 10 {
		return typeG
	}
	if r <= 11 {
		return typeF
	}
	return HotRoll(dice, mods...)
}

func determinationTable(table int) []string {
	switch table {
	default:
		fmt.Println(table)
		return []string{}
	case tableStarTypeTraditional:
		return []string{special, typeM, typeM, typeM, typeM, typeK, typeK, typeG, typeG, typeF, hot}
	case tableStarTypeRealistic:
		return []string{special, typeM, typeM, typeM, typeM, typeM, typeM, typeK, typeG, typeF, hot}
	case tableHot:
		return []string{typeA, typeA, typeA, typeA, typeA, typeA, typeA, typeA, typeB, typeB, typeO}
	case tableSpecial:
		return []string{classIV, classIV, classIV, classIV, classVI, classVI, classVI, classIII, classIII, gigants, gigants}
	case tableUnusual:
		return []string{peculiar, classVI, classIV, classBD, classBD, classBD, classD, classD, classD, classIII, gigants}
	case tableGiants:
		return []string{classIII, classIII, classIII, classIII, classIII, classIII, classIII, classII, classII, classIb, classIa}
	case tablePecuilar:
		return []string{blackHole, pulsar, neutronStar, nebula, nebula, protostar, protostar, protostar, starcluster, anomaly, anomaly}
	}
}

func StarTypeClassDependetRoll(dice *dice.Dicepool, class string, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	stType := ""
	for r <= 2 {
		dm++
		r = dice.Sroll("2d6") + dm
	}
	if class == classIV && r >= 3 && r <= 8 {
		r = r + 5
	}
	if r <= 8 {
		stType = typeK
	}
	if r <= 10 {
		stType = typeG
	}
	if r <= 11 {
		stType = typeF
	}
	if r >= 12 {
		stType = HotRoll(dice, mods...)
	}
	if class == classIV && stType == typeO {
		stType = typeB
	}
	if class == classVI {
		switch stType {
		default:
		case typeF:
			stType = typeG
		case typeA:
			stType = typeB
		}
	}
	return stType
}

func HotRoll(dice *dice.Dicepool, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	if r <= 9 {
		return typeA
	}
	if r <= 11 {
		return typeB
	}
	return typeO
}

func SpecialRoll(dice *dice.Dicepool, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	if r <= 5 {
		return "Class VI"
	}
	if r <= 8 {
		return "Class IV"
	}
	if r <= 10 {
		return "Class III"
	}
	return GigantsRoll(dice, mods...)
}

func UnusualRoll(dice *dice.Dicepool, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	if r <= 2 {
		return "Peculiar"
	}
	if r <= 3 {
		return classVI
	}
	if r <= 4 {
		return classIV
	}
	if r <= 7 {
		return classBD
	}
	return GigantsRoll(dice, mods...)
}

func GigantsRoll(dice *dice.Dicepool, mods ...int) string {
	dm := 0
	for _, m := range mods {
		dm += m
	}
	r := dice.Sroll("2d6") + dm
	if r <= 8 {
		return classIII
	}
	if r <= 10 {
		return classII
	}
	if r <= 11 {
		return classIb
	}
	return classIa
}
