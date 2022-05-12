package traffic

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	Passage_Low = iota
	Passage_Basic
	Passage_Middle
	Passage_High
)

type mWorld interface {
	MW_UWP() string
	TravelZone() string
	MW_Remarks() string
	hexagon.Hex
	hexagon.Cube
}

type PassengerTrafficData struct {
	Passage map[int]int
}

func BasePassengerFactor_MGT2_Core(source, destination mWorld) (int, error) {
	factor := -1000
	sUWP, err := uwp.FromString(source.MW_UWP())
	if err != nil {
		return factor, err
	}
	dUWP, err := uwp.FromString(destination.MW_UWP())
	if err != nil {
		return factor, err
	}
	fMod := 0
	if source.CoordX() == destination.CoordX() && source.CoordY() == destination.CoordY() {
		return factor, fmt.Errorf("sorce and destination can not have same coordinates")
	}
	//applying uwp factors for noth S and D:
	if sUWP.Pops() < 2 {
		fMod -= 4
	}
	if dUWP.Pops() < 2 {
		fMod -= 4
	}
	if sUWP.Pops() == 6 || sUWP.Pops() == 7 {
		fMod += 1
	}
	if dUWP.Pops() == 6 || dUWP.Pops() == 7 {
		fMod += 1
	}
	if sUWP.Pops() > 7 {
		fMod += 3
	}
	if dUWP.Pops() > 7 {
		fMod += 3
	}
	if sUWP.Starport() == "A" {
		fMod += 2
	}
	if dUWP.Starport() == "A" {
		fMod += 2
	}
	if sUWP.Starport() == "B" || sUWP.Starport() == "F" {
		fMod += 1
	}
	if dUWP.Starport() == "B" || dUWP.Starport() == "F" {
		fMod += 1
	}
	if sUWP.Starport() == "E" || sUWP.Starport() == "H" {
		fMod -= 1
	}
	if dUWP.Starport() == "E" || dUWP.Starport() == "H" {
		fMod -= 1
	}
	if sUWP.Starport() == "X" {
		fMod -= 3
	}
	if dUWP.Starport() == "X" {
		fMod -= 3
	}
	if source.TravelZone() == "A" {
		fMod += 1
	}
	if destination.TravelZone() == "A" {
		fMod += 1
	}
	if source.TravelZone() == "R" {
		fMod -= 4
	}
	if destination.TravelZone() == "R" {
		fMod -= 4
	}
	dist := hexagon.DistanceHex(source, destination)
	if dist > 1 {
		fMod = fMod - (dist - 1)
	}
	factor = fMod
	return factor, nil
}

func PassengerTrafficValues_MGT2_Core(ptv int) (dice, add int) {
	switch ptv {
	default:
		if ptv > 19 {
			return 10, 0
		}
		return 0, 0
	case 2, 3:
		return 1, 0
	case 4, 5, 6:
		return 2, 0
	case 7, 8, 9, 10:
		return 3, 0
	case 11, 12, 13:
		return 4, 0
	case 14, 15:
		return 5, 0
	case 16:
		return 6, 0
	case 17:
		return 7, 0
	case 18:
		return 8, 0
	case 19:
		return 9, 0
	}
}

func PassengerTrafficValues_MGT1_MP(ptv, pType int) (dice, add int) {
	switch pType {
	case Passage_Low:
		switch ptv {
		default:
			if ptv > 15 {
				return 10, 0
			}
			return 0, 0
		case 1:
			return 1, 0
		case 2, 3, 4, 5:
			return 2, 0
		case 6, 7:
			return 3, 0
		case 8, 9:
			return 4, 0
		case 10, 11:
			return 5, 0
		case 12, 13:
			return 6, 0
		case 14:
			return 7, 0
		case 15:
			return 8, 0
		}
	case Passage_Middle, Passage_Basic:
		switch ptv {
		default:
			if ptv > 15 {
				return 6, 0
			}
			return 0, 0
		case 2, 3, 4, 5:
			return 1, 0
		case 6, 7, 8, 9:
			return 2, 0
		case 10:
			return 3, 0
		case 11, 12, 13:
			return 4, 0
		case 14, 15:
			return 5, 0
		}
	case Passage_High:
		switch ptv {
		default:
			if ptv > 15 {
				return 5, 0
			}
			return 0, 0
		case 5, 6, 7:
			return 1, 0
		case 8, 9, 10:
			return 2, 0
		case 11, 12:
			return 3, 0
		case 13, 14, 15:
			return 4, 0
		}
	}
	return 0, 0
}

func sliceContains(sl []string, elem string) bool {
	for _, v := range sl {
		if v == elem {
			return true
		}
	}
	return false
}

func BasePassengerFactor_MGT1_MP(source, destination mWorld) (int, error) {
	factor := -1000
	sUWP, err := uwp.FromString(source.MW_UWP())
	if err != nil {
		return factor, err
	}
	dUWP, err := uwp.FromString(destination.MW_UWP())
	if err != nil {
		return factor, err
	}
	sTC := strings.Fields(source.MW_Remarks())
	dTC := strings.Fields(destination.MW_Remarks())
	fMod := 0
	fMod += sUWP.Pops()
	if source.CoordX() == destination.CoordX() && source.CoordY() == destination.CoordY() {
		return factor, fmt.Errorf("source and destination can not have same coordinates")
	}
	//applying uwp factors for noth S and D:
	if sliceContains(sTC, "Ag") {
		fMod += 0
	}
	if sliceContains(dTC, "Ag") {
		fMod += 0
	}
	if sliceContains(sTC, "As") {
		fMod += 1
	}
	if sliceContains(dTC, "As") {
		fMod += -1
	}
	if sliceContains(sTC, "Ba") {
		fMod += -5
	}
	if sliceContains(dTC, "Ba") {
		fMod += -5
	}
	if sliceContains(sTC, "De") {
		fMod += -1
	}
	if sliceContains(dTC, "De") {
		fMod += -1
	}
	if sliceContains(sTC, "Fl") {
		fMod += 0
	}
	if sliceContains(dTC, "Fl") {
		fMod += 0
	}
	if sliceContains(sTC, "Ga") {
		fMod += 2
	}
	if sliceContains(dTC, "Ga") {
		fMod += 2
	}
	if sliceContains(sTC, "Hi") {
		fMod += 0
	}
	if sliceContains(dTC, "Hi") {
		fMod += 4
	}
	if sliceContains(sTC, "Ic") {
		fMod += 1
	}
	if sliceContains(dTC, "Ic") {
		fMod += -1
	}
	if sliceContains(sTC, "In") {
		fMod += 2
	}
	if sliceContains(dTC, "In") {
		fMod += 1
	}
	if sliceContains(sTC, "Lo") {
		fMod += 0
	}
	if sliceContains(dTC, "Lo") {
		fMod += -4
	}
	if sliceContains(sTC, "Na") {
		fMod += 0
	}
	if sliceContains(dTC, "Na") {
		fMod += 0
	}
	if sliceContains(sTC, "Ni") {
		fMod += 0
	}
	if sliceContains(dTC, "Ni") {
		fMod += -1
	}
	if sliceContains(sTC, "Po") {
		fMod += -2
	}
	if sliceContains(dTC, "Po") {
		fMod += -1
	}
	if sliceContains(sTC, "Ri") {
		fMod += -1
	}
	if sliceContains(dTC, "Ri") {
		fMod += 2
	}
	if sliceContains(sTC, "Wa") {
		fMod += 0
	}
	if sliceContains(dTC, "Wa") {
		fMod += 0
	}
	if source.TravelZone() == "A" {
		fMod += -2
	}
	if destination.TravelZone() == "A" {
		fMod += -2
	}
	if source.TravelZone() == "R" {
		fMod += 4
	}
	if destination.TravelZone() == "R" {
		fMod += -4
	}
	sTL := sUWP.TL()
	dTL := dUWP.TL()
	tMod := sTL - dTL
	if tMod > 0 {
		tMod = tMod * -1
	}
	if tMod < -5 {
		tMod = -5
	}
	fMod += tMod

	// dist := astrogation.DistanceRaw(source.CoordX(), source.CoordY(), destination.CoordX(), destination.CoordY())
	// if dist > 1 {
	// 	fMod = fMod - (dist - 1)
	// }
	factor = fMod
	return factor, nil
}
