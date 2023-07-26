package star

import (
	"fmt"
	"strconv"
	"strings"

	orbitns "github.com/Galdoba/TravellerTools/mgt2/wbh/orbits"
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
	tableSubTypeNumeric
	tableSubTypePrimaryM

	starType                    = "Star Type"
	special                     = "Special"
	hot                         = "Hot"
	gigants                     = "Gigants"
	peculiar                    = "Peculiar"
	BlackHole                   = "Black Hole"
	Pulsar                      = "Pulsar"
	NeutronStar                 = "Neutron Star"
	Nebula                      = "Nebula"
	Protostar                   = "Protostar"
	Starcluster                 = "Star Cluster"
	Primordial                  = "Primordial System"
	anomaly                     = "Anomaly"
	TypeO                       = "Type O"
	TypeB                       = "Type B"
	TypeA                       = "Type A"
	TypeF                       = "Type F"
	TypeG                       = "Type G"
	TypeK                       = "Type K"
	TypeM                       = "Type M"
	TypeL                       = "Type L"
	TypeT                       = "Type T"
	TypeY                       = "Type Y"
	ClassIa                     = "Class Ia"
	ClassIb                     = "Class Ib"
	ClassII                     = "Class II"
	ClassIII                    = "Class III"
	ClassIV                     = "Class IV"
	ClassV                      = "Class V"
	ClassVI                     = "Class VI"
	ClassBD                     = "Class BD"
	ClassD                      = "Class D"
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

type Star struct {
	StType        string
	Class         string
	SubType       string
	Specialcase   string
	Designation   string
	Determination string
	Mass          float64
	Temperature   int
	IsPrimary     bool
	Diameter      float64
	Luminocity    float64
	Age           float64 //Gyrs
	Orbit         *orbitns.OrbitN
}

func DefineStarPresence(st Star, dice *dice.Dicepool) []string {
	dm := 0
	switch st.Class {
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV:
		dm++
	case ClassV, ClassVI:
		switch st.StType {
		case TypeO, TypeB, TypeA, TypeF:
			dm++
		case TypeM:
			dm--
		}
	case ClassBD, ClassD:
		dm--
	case Pulsar, NeutronStar, BlackHole:
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

func DefineStarDetermination(primary Star, targetDesig string, dice *dice.Dicepool) (string, string) {
	dm := 0
	switch primary.Class {
	case ClassIII, ClassIV:
		dm--
	case ClassBD:
		return determinationSibling, "Aa"
	}
	secondary := []string{determinationOther, determinationOther, determinationRandom, determinationRandom, determinationRandom, determinationLesser, determinationLesser, determinationSibling, determinationSibling, determinationTwin, determinationTwin}
	companion := []string{determinationOther, determinationOther, determinationRandom, determinationRandom, determinationLesser, determinationLesser, determinationSibling, determinationSibling, determinationTwin, determinationTwin, determinationTwin}
	poststellar := []string{determinationOther, determinationOther, determinationRandom, determinationRandom, determinationRandom, determinationRandom, determinationRandom, determinationLesser, determinationLesser, determinationTwin, determinationTwin}
	other := []string{NeutronStar, ClassD, ClassD, ClassD, ClassD, ClassD, ClassBD, ClassBD, ClassBD, ClassBD, ClassBD}
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
	switch primary.Class {
	case ClassD, Pulsar, NeutronStar, BlackHole:
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

func New(dice *dice.Dicepool, TypeTableVariant, starGenerationMethod int, designationCode, determination string, contextStars ...Star) (Star, error) {
	st := Star{}
	if determination != determinationPrimary && len(contextStars) < 1 {
		return st, fmt.Errorf("can not create non primary star without context")
	}
	switch determination {
	default:
		st.StType = determination
	case determinationPrimary:
		st.StType, st.Class, st.Specialcase = starTypeAndClass(dice, TypeTableVariant, starGenerationMethod)
		st.SubType = starSubtype(dice, st)
	case determinationTwin:
		st.StType = contextStars[0].StType
		st.Class = contextStars[0].Class
		st.SubType = contextStars[0].SubType
	case determinationSibling:
		st.StType = contextStars[0].StType
		st.Class = contextStars[0].Class
		st.SubType = contextStars[0].SubType
		st.StType, st.SubType, st.Class = makeSibling(st, dice)
	case determinationLesser:
		st.StType = lowerType(contextStars[0].StType)
		st.Class = contextStars[0].Class
		st.SubType = starSubtype(dice, st)
	case determinationRandom:
		st.StType, st.Class, st.Specialcase = starTypeAndClass(dice, TypeTableVariant, starGenerationMethod)
		st.SubType = starSubtype(dice, st)
		if valOfStar(st) >= valOfStar(contextStars[0]) {
			st.StType = lowerType(contextStars[0].StType)
			st.Class = contextStars[0].Class
			st.SubType = starSubtype(dice, st)
		}
	case ClassBD, ClassD:
		for st.Class != determination {
			st.StType, st.Class, st.Specialcase = starTypeAndClass(dice, TypeTableVariant, starGenerationMethod)
			st.SubType = starSubtype(dice, st)
		}
	}
	st.Mass = massOf(st, dice)
	if st.Class == ClassBD {
		st.StType, st.SubType = evaluateBDClassData(st.Mass)
	}

	st.Age = ageOf(st, dice)
	if st.Age < 0.1 {
		st.Specialcase = Primordial
	}
	if st.Mass < 4.7 && st.Age < 0.01 {
		st.Specialcase = Protostar
	}
	st.Temperature = temperatureOf(st, dice)
	st.Diameter = diameterOf(st, dice)
	st.Luminocity = luminocityOf(st)

	return st, nil
}

func makeSibling(st Star, dice *dice.Dicepool) (string, string, string) {
	sType, subType, Class := st.StType, st.SubType, st.Class
	switch Class {
	case ClassD, Pulsar, NeutronStar, BlackHole:
		return sType, subType, Class
	}
	subInt, _ := strconv.Atoi(subType)

	subInt = subInt + dice.Sroll("1d6")
	if subInt > 9 {
		sType = lowerType(sType)
		subInt = subInt - 10
	}
	subType = fmt.Sprintf("%v", subInt)
	return sType, subType, Class
}

func lowerType(sType string) string {
	switch sType {
	default:
		panic("not a Class:" + sType)
	case "":
		return ""
	case TypeO:
		return TypeB
	case TypeB:
		return TypeA
	case TypeA:
		return TypeF
	case TypeF:
		return TypeG
	case TypeG:
		return TypeK
	case TypeK:
		return TypeM
	case TypeM:
		return TypeM
	case TypeL:
		return TypeT
	case TypeT:
		return TypeY
	case TypeY:
		return TypeY
	case BlackHole:
		return NeutronStar
	case NeutronStar:
		return ClassD
	case ClassD:
		return ClassBD
	}
}

func ShortStarDescription(st Star) string {
	descr := st.StType + st.SubType + " " + st.Class
	if st.Class == ClassBD {
		descr = st.StType + st.SubType
	}
	if st.Class == ClassD {
		descr = st.Class + st.StType
	}
	switch st.StType {
	case Nebula, Protostar, NeutronStar, Pulsar, BlackHole, Starcluster, anomaly:
		return st.StType
	}
	descr = strings.ReplaceAll(descr, "Class ", "")
	descr = strings.ReplaceAll(descr, "Type ", "")
	return descr
}

func subTypeInt(stp string) int {
	switch stp {
	default:
		return -1
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		i, _ := strconv.Atoi(stp)
		return i
	}
}

func rollTable(dice *dice.Dicepool, table, TypeTableVariant, method int, mods ...int) string {
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
		table = selectTableBy(starType, method, TypeTableVariant)
	}
	tableRollResult := determinationTable(table)[r-2]
	switch tableRollResult {
	case starType, hot, special, gigants, peculiar:
		return rollTable(dice, selectTableBy(tableRollResult, method, TypeTableVariant), TypeTableVariant, method, mods...)
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
		return []string{special, TypeM, TypeM, TypeM, TypeM, TypeK, TypeK, TypeG, TypeG, TypeF, hot}
	case tableStarTypeRealistic:
		return []string{special, TypeM, TypeM, TypeM, TypeM, TypeM, TypeM, TypeK, TypeG, TypeF, hot}
	case tableHot:
		return []string{TypeA, TypeA, TypeA, TypeA, TypeA, TypeA, TypeA, TypeA, TypeB, TypeB, TypeO}
	case tableSpecial:
		return []string{ClassIV, ClassIV, ClassIV, ClassIV, ClassVI, ClassVI, ClassVI, ClassIII, ClassIII, gigants, gigants}
	case tableUnusual:
		return []string{peculiar, ClassVI, ClassIV, ClassBD, ClassBD, ClassBD, ClassD, ClassD, ClassD, ClassIII, gigants}
	case tableGiants:
		return []string{ClassIII, ClassIII, ClassIII, ClassIII, ClassIII, ClassIII, ClassIII, ClassII, ClassII, ClassIb, ClassIa}
	case tablePecuilar:
		return []string{BlackHole, Pulsar, NeutronStar, Nebula, Nebula, Protostar, Protostar, Protostar, Starcluster, anomaly, anomaly}
	case tableSubTypeNumeric:
		return []string{"0", "1", "3", "5", "7", "9", "8", "6", "4", "2", "0"}
	case tableSubTypePrimaryM:
		return []string{"8", "6", "5", "4", "0", "2", "1", "3", "5", "7", "9"}
	}
}

func hotter(a, b Star) Star {
	if valOfStar(a) > valOfStar(b) {
		return a
	}
	return b
}

func valOfStar(s Star) int {
	val := 0
	switch s.Class {
	case ClassIa:
		val += 9000
	case ClassIb:
		val += 8000
	case ClassII:
		val += 7000
	case ClassIII:
		val += 6000
	case ClassIV:
		val += 5000
	case ClassV:
		val += 4000
	case ClassVI:
		val += 3000
	case ClassBD:
		val += 2000
	case ClassD:
		val += 1000
	}
	switch s.StType {
	case TypeO:
		val += 100
	case TypeB:
		val += 90
	case TypeA:
		val += 80
	case TypeF:
		val += 70
	case TypeG:
		val += 60
	case TypeK:
		val += 50
	case TypeM:
		val += 40
	case TypeL:
		val += 30
	case TypeT:
		val += 20
	case TypeY:
		val += 10
	}
	v, _ := strconv.Atoi(s.SubType)
	val += v
	return val
}

func DesignationCodes() []string {
	return []string{"Aa", "Ab", "Ba", "Bb", "Ca", "Cb", "Da", "Db"}
}

func (st *Star) normalizeValues() {
	switch st.Class {
	default:
		st.Class = "????????" + st.Class + "??????????"
		panic(st.Class)
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV, ClassV, ClassVI, ClassBD, ClassD:
		st.Class = strings.ReplaceAll(st.Class, "Class ", "")
	}
	st.StType = strings.ReplaceAll(st.StType, "Type ", "")
	st.Mass = float64(int(st.Mass*1000)) / 1000
	st.Diameter = float64(int(st.Diameter*1000)) / 1000
	st.Luminocity = float64(int(st.Luminocity*1000)) / 1000
	st.Age = float64(int(st.Age*1000)) / 1000

}

func AUof(st Star) float64 {
	if st.Orbit == nil {
		return 0
	}
	return st.Orbit.AU
}
