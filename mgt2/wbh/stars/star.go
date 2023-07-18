package stars

import (
	"fmt"
	"strconv"
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
	tableSubtypeNumeric
	tableSubtypePrimaryM

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
	primordial  = "Primordial System"
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
	sttype      string
	class       string
	subtype     string
	specialcase string
	mass        float64
	temperature int
	isPrimary   bool
	diameter    float64
	luminocity  float64
	age         float64 //Gyrs
}

func NewStarSystem(dice *dice.Dicepool, starGenerationMethod, tableVariant int) (*starsystem, error) {
	ss := starsystem{}
	switch starGenerationMethod {
	case GenerationMethodUnusual, GenerationMethodSpecial:
	default:
		return &ss, fmt.Errorf("starGenerationMethod unknown (%v)", starGenerationMethod)
	}
	primary, err := NewStar(dice, tableVariant, starGenerationMethod, true)
	if err != nil {
		return &ss, err
	}
	ss.primary = primary
	return &ss, nil
}

func NewStar(dice *dice.Dicepool, typeTableVariant, starGenerationMethod int, isPrimary bool) (star, error) {
	st := star{}
	st.isPrimary = isPrimary
	st.sttype, st.class, st.specialcase = starTypeAndClass(dice, typeTableVariant, starGenerationMethod)
	st.subtype = starSubtype(dice, st)
	st.mass = massOf(st, dice)
	if st.class == classBD {
		st.sttype, st.subtype = evaluateBDclassData(st.mass)
	}
	st.temperature = temperatureOf(st, dice)
	st.diameter = diameterOf(st, dice)
	st.luminocity = luminocityOf(st)
	st.age = ageOf(st, dice)
	if st.age < 0.1 {
		st.specialcase = primordial
	}
	if st.mass < 4.7 && st.age < 0.01 {
		st.specialcase = protostar
	}

	return st, nil
}

func shortStarDescription(st star) string {
	descr := st.sttype + st.subtype + " " + st.class
	if st.class == classBD {
		return "BD"
	}
	if st.class == classD {
		return "D"
	}
	switch st.sttype {
	case nebula, protostar, neutronStar, pulsar, blackHole, starcluster, anomaly:
		return st.sttype
	}
	descr = strings.ReplaceAll(descr, "Class ", "")
	descr = strings.ReplaceAll(descr, "Type ", "")
	return descr
}

func subtypeInt(stp string) int {
	switch stp {
	default:
		return -1
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		i, _ := strconv.Atoi(stp)
		return i
	}
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

func determinationTable(table int) []string {
	switch table {
	default:
		panic(fmt.Sprintf("table with key %v was not provided", table))
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
	case tableSubtypeNumeric:
		return []string{"0", "1", "3", "5", "7", "9", "8", "6", "4", "2", "0"}
	case tableSubtypePrimaryM:
		return []string{"8", "6", "5", "4", "0", "2", "1", "3", "5", "7", "9"}
	}
}
