package stars

import (
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func starTypeAndClass(dice *dice.Dicepool, TypeTableVariant, starGenerationMethod int) (string, string, string) {
	stType, Class, sscase := "", "", ""
	stRoll := rollTable(dice, tableStarTypeUnselected, TypeTableVariant, starGenerationMethod)
	switch stRoll {
	case TypeO, TypeB, TypeA, TypeF, TypeG, TypeK, TypeM:
		stType = stRoll
		Class = ClassV
	case ClassBD, ClassD:
		Class = stRoll
		for !strings.HasPrefix(stRoll, "Type ") {
			stRoll = rollTable(dice, tableStarTypeUnselected, TypeTableVariant, starGenerationMethod)
		}
		stType = stRoll
	case ClassIa, ClassIb, ClassII, ClassIII, ClassIV, ClassVI:
		Class = stRoll
		dm := append([]int{}, 1)
		for !strings.HasPrefix(stRoll, "Type ") {
			stRoll = rollTable(dice, tableStarTypeUnselected, TypeTableVariant, starGenerationMethod, dm...)
			switch Class {
			case ClassIV:
				if stRoll == TypeO {
					stRoll = TypeB
				}
				if stRoll == TypeM {
					stRoll = "rejected"
				}
			case ClassVI:
				if stRoll == TypeF {
					stRoll = TypeG
				}
				if stRoll == TypeA {
					stRoll = TypeB
				}
			}
		}
		stType = stRoll
	case BlackHole, Pulsar, NeutronStar, Nebula, Protostar, Starcluster, anomaly:
		sscase = stRoll
	default:
		panic(stRoll)
	}
	switch Class {
	case ClassBD:
		stType = ""

	}
	return stType, Class, sscase
}

func starSubtype(dice *dice.Dicepool, st Star) string {
	table := tableSubTypeNumeric
	if st.IsPrimary && st.StType == TypeM {
		table = tableSubTypePrimaryM
	}
	subType := ""
	specialCaseResolved := false
	for !specialCaseResolved {
		subTypeRollResult := determinationTable(table)[dice.Sroll("2d6")-2]
		switch st.Class {
		default:
			specialCaseResolved = true
			subType = subTypeRollResult
		case ClassIV:
			if st.StType == TypeK && subTypeInt(subTypeRollResult) >= 5 {
				continue
			}
		case ClassBD, ClassD:
			return ""
		}
		specialCaseResolved = true
		subType = subTypeRollResult
	}
	return subType
}
