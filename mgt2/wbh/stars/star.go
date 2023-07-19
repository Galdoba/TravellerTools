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

	starType                    = "Star Type"
	special                     = "Special"
	hot                         = "Hot"
	gigants                     = "Gigants"
	peculiar                    = "Peculiar"
	blackHole                   = "Black Hole"
	pulsar                      = "Pulsar"
	neutronStar                 = "Neutron Star"
	nebula                      = "Nebula"
	protostar                   = "Protostar"
	starcluster                 = "Star Cluster"
	primordial                  = "Primordial System"
	anomaly                     = "Anomaly"
	typeO                       = "Type O"
	typeB                       = "Type B"
	typeA                       = "Type A"
	typeF                       = "Type F"
	typeG                       = "Type G"
	typeK                       = "Type K"
	typeM                       = "Type M"
	typeL                       = "Type L"
	typeT                       = "Type T"
	typeY                       = "Type Y"
	classIa                     = "Class Ia"
	classIb                     = "Class Ib"
	classII                     = "Class II"
	classIII                    = "Class III"
	classIV                     = "Class IV"
	classV                      = "Class V"
	classVI                     = "Class VI"
	classBD                     = "Class BD"
	classD                      = "Class D"
	designationPrimary          = "Aa"
	designationPrimaryCompanion = "Ab"
	designationClose            = "Ba"
	designationCloseCompanion   = "Bb"
	designationNear             = "Ca"
	designationNearCompanion    = "Cb"
	designationFar              = "Da"
	designationFarCompanion     = "Db"
	determinationPrimary        = "Prime"
	determinationRandom         = "Random"
	determinationLesser         = "Lesser"
	determinationSibling        = "Sibling"
	determinationTwin           = "Twin"
	determinationOther          = "Other"
)

type starsystem struct {
	starGenerationMethod int
	typeTableVariant     int
	primary              star
	star                 map[string]star
}

type star struct {
	sttype        string
	class         string
	subtype       string
	specialcase   string
	designation   string
	determination string
	mass          float64
	temperature   int
	isPrimary     bool
	diameter      float64
	luminocity    float64
	age           float64 //Gyrs
}

func NewStarSystem(dice *dice.Dicepool, starGenerationMethod, tableVariant int) (*starsystem, error) {
	ss := starsystem{}
	switch starGenerationMethod {
	case GenerationMethodUnusual, GenerationMethodSpecial:
	default:
		return &ss, fmt.Errorf("starGenerationMethod unknown (%v)", starGenerationMethod)
	}
	ss.star = make(map[string]star)

	primary, err := NewStar(dice, tableVariant, starGenerationMethod, designationPrimary, determinationPrimary)
	if err != nil {
		return &ss, err
	}
	ss.star[designationPrimary] = primary
	designations := defineStarPresence(ss.star[designationPrimary], dice)
	for _, desig := range designations {
		if _, ok := ss.star[desig]; ok {
			continue
		}
		determ, context := defineStarDetermination(primary, desig, dice)
		star, err := NewStar(dice, tableVariant, starGenerationMethod, desig, determ, ss.star[context])
		if err != nil {
			return &ss, fmt.Errorf("secondary star %v creation: %v", desig, err.Error())
		}
		ss.star[desig] = star
	}
	ss.ageResetIfRequired(dice)
	return &ss, nil
}

func (ss *starsystem) ageResetIfRequired(dice *dice.Dicepool) {
	switch ss.star["Aa"].class {
	case classIa, classIb, classII, classIII, classIV, classV, classVI, classBD:
		for _, v := range ss.star {
			switch v.class {
			case classD, pulsar, neutronStar, blackHole:
				primary := ss.star["Aa"]
				primary.age = v.age
				ss.star["Aa"] = primary

				// primary.age = starFinalAge(v.mass, dice)
				// if primary.age < was {
				// 	fmt.Println("set new age")
				// 	primary.age = was
				// }
				// if primary.age > 13.5 {
				// 	fmt.Println("set age border", primary)

				// 	primary.age = 13.5
				// }
				// ss.star["Aa"] = primary
			}
		}
	}
}

func defineStarPresence(st star, dice *dice.Dicepool) []string {
	dm := 0
	switch st.class {
	case classIa, classIb, classII, classIII, classIV:
		dm++
	case classV, classVI:
		switch st.sttype {
		case typeO, typeB, typeA, typeF:
			dm++
		case typeM:
			dm--
		}
	case classBD, classD:
		dm--
	case pulsar, neutronStar, blackHole:
		dm--
	}
	defined := []string{"A"}
	for _, new := range []string{"B", "C", "D"} {
		if dice.Sroll("2d6")+dm >= 10 {
			defined = append(defined, new)
		}
	}
	result := []string{}
	for _, d := range defined {
		result = append(result, d+"a")
		if dice.Sroll("2d6")+dm >= 10 {
			result = append(result, d+"b")
		}
	}
	return result
}

func defineStarDetermination(primary star, targetDesig string, dice *dice.Dicepool) (string, string) {
	dm := 0
	switch primary.class {
	case classIII, classIV:
		dm--
	case classBD:
		return determinationSibling, "Aa"
	}
	secondary := []string{determinationOther, determinationOther, determinationRandom, determinationRandom, determinationRandom, determinationLesser, determinationLesser, determinationSibling, determinationSibling, determinationTwin, determinationTwin}
	companion := []string{determinationOther, determinationOther, determinationRandom, determinationRandom, determinationLesser, determinationLesser, determinationSibling, determinationSibling, determinationTwin, determinationTwin, determinationTwin}
	poststellar := []string{determinationOther, determinationOther, determinationRandom, determinationRandom, determinationRandom, determinationRandom, determinationRandom, determinationLesser, determinationLesser, determinationTwin, determinationTwin}
	other := []string{neutronStar, classD, classD, classD, classD, classD, classBD, classBD, classBD, classBD, classBD}
	r := dice.Sroll("2d6") - 2 + dm
	if r < 0 {
		r = 0
	}
	result := ""
	design := ""
	if strings.Contains(targetDesig, "a") {
		result = secondary[r]
		design = "Aa"
	}
	if strings.Contains(targetDesig, "b") {
		result = companion[r]
		design = strings.ReplaceAll(targetDesig, "b", "a")
	}
	switch primary.class {
	case classD, pulsar, neutronStar, blackHole:
		result = poststellar[r]
		design = "Aa"
	}
	if result == determinationOther {
		r1 := dice.Sroll("2d6") - 2 + dm
		if r1 < 0 {
			r1 = 0
		}
		result = other[r1]
	}
	return result, design
}

func NewStar(dice *dice.Dicepool, typeTableVariant, starGenerationMethod int, designationCode, determination string, contextStars ...star) (star, error) {
	st := star{}
	if determination != determinationPrimary && len(contextStars) < 1 {
		return st, fmt.Errorf("can not create non primary star without context")
	}
	switch determination {
	default:
		st.sttype = determination
	case determinationPrimary:
		st.sttype, st.class, st.specialcase = starTypeAndClass(dice, typeTableVariant, starGenerationMethod)
		st.subtype = starSubtype(dice, st)
	case determinationTwin:
		st.sttype = contextStars[0].sttype
		st.class = contextStars[0].class
		st.subtype = contextStars[0].subtype
	case determinationSibling:
		st.sttype = contextStars[0].sttype
		st.class = contextStars[0].class
		st.subtype = contextStars[0].subtype
		st.sttype, st.subtype, st.class = makeSibling(st, dice)
	case determinationLesser:
		st.sttype = lowerType(contextStars[0].sttype)
		st.class = contextStars[0].class
		st.subtype = starSubtype(dice, st)
	case determinationRandom:
		st.sttype, st.class, st.specialcase = starTypeAndClass(dice, typeTableVariant, starGenerationMethod)
		st.subtype = starSubtype(dice, st)
		if valOfStar(st) >= valOfStar(contextStars[0]) {
			st.sttype = lowerType(contextStars[0].sttype)
			st.class = contextStars[0].class
			st.subtype = starSubtype(dice, st)
		}
	case classBD, classD:
		for st.class != determination {
			st.sttype, st.class, st.specialcase = starTypeAndClass(dice, typeTableVariant, starGenerationMethod)
			st.subtype = starSubtype(dice, st)
		}
	}
	st.mass = massOf(st, dice)
	if st.class == classBD {
		st.sttype, st.subtype = evaluateBDclassData(st.mass)
	}

	st.age = ageOf(st, dice)
	if st.age < 0.1 {
		st.specialcase = primordial
	}
	if st.mass < 4.7 && st.age < 0.01 {
		st.specialcase = protostar
	}
	st.temperature = temperatureOf(st, dice)
	st.diameter = diameterOf(st, dice)
	st.luminocity = luminocityOf(st)

	return st, nil
}

func makeSibling(st star, dice *dice.Dicepool) (string, string, string) {
	stype, subType, class := st.sttype, st.subtype, st.class
	switch class {
	case classD, pulsar, neutronStar, blackHole:
		return stype, subType, class
	}
	subInt, _ := strconv.Atoi(subType)

	subInt = subInt + dice.Sroll("1d6")
	if subInt > 9 {
		stype = lowerType(stype)
		subInt = subInt - 10
	}
	subType = fmt.Sprintf("%v", subInt)
	return stype, subType, class
}

func lowerType(stype string) string {
	switch stype {
	default:
		panic("not a class:" + stype)
	case "":
		return ""
	case typeO:
		return typeB
	case typeB:
		return typeA
	case typeA:
		return typeF
	case typeF:
		return typeG
	case typeG:
		return typeK
	case typeK:
		return typeM
	case typeM:
		return typeM
	case typeL:
		return typeT
	case typeT:
		return typeY
	case typeY:
		return typeY
	case blackHole:
		return neutronStar
	case neutronStar:
		return classD
	case classD:
		return classBD
	}
}

func shortStarDescription(st star) string {
	descr := st.sttype + st.subtype + " " + st.class
	if st.class == classBD {
		descr = st.sttype + st.subtype
	}
	if st.class == classD {
		descr = st.class + st.sttype
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

func hotter(a, b star) star {
	if valOfStar(a) > valOfStar(b) {
		return a
	}
	return b
}

func valOfStar(s star) int {
	val := 0
	switch s.class {
	case classIa:
		val += 9000
	case classIb:
		val += 8000
	case classII:
		val += 7000
	case classIII:
		val += 6000
	case classIV:
		val += 5000
	case classV:
		val += 4000
	case classVI:
		val += 3000
	case classBD:
		val += 2000
	case classD:
		val += 1000
	}
	switch s.sttype {
	case typeO:
		val += 100
	case typeB:
		val += 90
	case typeA:
		val += 80
	case typeF:
		val += 70
	case typeG:
		val += 60
	case typeK:
		val += 50
	case typeM:
		val += 40
	case typeL:
		val += 30
	case typeT:
		val += 20
	case typeY:
		val += 10
	}
	v, _ := strconv.Atoi(s.subtype)
	val += v
	return val
}
