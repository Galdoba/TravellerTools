package stars

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func starTypeAndClass(dice *dice.Dicepool, typeTableVariant, starGenerationMethod int) (string, string) {
	sttype, class := "", ""
	stRoll := rollTable(dice, tableStarTypeUnselected, typeTableVariant, starGenerationMethod)
	switch stRoll {
	case typeO, typeB, typeA, typeF, typeG, typeK, typeM:
		sttype = stRoll
		class = classV
	case classBD, classD:
		class = stRoll
		for !strings.HasPrefix(stRoll, "Type ") {
			stRoll = rollTable(dice, tableStarTypeUnselected, typeTableVariant, starGenerationMethod)
		}
		sttype = stRoll
	case classIa, classIb, classII, classIII, classIV, classVI:
		class = stRoll
		dm := append([]int{}, 1)
		for !strings.HasPrefix(stRoll, "Type ") {
			stRoll = rollTable(dice, tableStarTypeUnselected, typeTableVariant, starGenerationMethod, dm...)
			switch class {
			case classIV:
				if stRoll == typeO {
					stRoll = typeB
				}
				if stRoll == typeM {
					stRoll = "rejected"
				}
			case classVI:
				if stRoll == typeF {
					stRoll = typeG
				}
				if stRoll == typeA {
					stRoll = typeB
				}
			}
		}
		sttype = stRoll
	case blackHole, pulsar, neutronStar, nebula, protostar, starcluster, anomaly:
		sttype = stRoll
	default:
		panic(stRoll)
	}
	switch class {
	case classBD:
		sttype = ""

	}
	return sttype, class
}

func starSubtype(dice *dice.Dicepool, st star) string {
	table := tableSubtypeNumeric
	if st.isPrimary && st.sttype == typeM {
		table = tableSubtypePrimaryM
	}
	subtype := ""
	specialCaseResolved := false
	for !specialCaseResolved {
		subtypeRollResult := determinationTable(table)[dice.Sroll("2d6")-2]
		switch st.class {
		default:
			specialCaseResolved = true
			subtype = subtypeRollResult
		case classIV:
			if st.sttype == typeK && subtypeInt(subtypeRollResult) >= 5 {
				continue
			}
		case classBD, classD:
			return ""
		}
		specialCaseResolved = true
		subtype = subtypeRollResult
	}
	return subtype
}
