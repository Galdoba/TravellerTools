package traffic

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

func BaseFreightFactor_MGT2_Core(source, destination mWorld) (int, error) {
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
		fMod += 2
	}
	if dUWP.Pops() == 6 || dUWP.Pops() == 7 {
		fMod += 2
	}
	if sUWP.Pops() > 7 {
		fMod += 4
	}
	if dUWP.Pops() > 7 {
		fMod += 4
	}
	if sUWP.TL() < 7 {
		fMod -= 1
	}
	if dUWP.TL() < 7 {
		fMod -= 1
	}
	if sUWP.TL() > 8 {
		fMod += 2
	}
	if dUWP.TL() > 8 {
		fMod += 2
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
		fMod -= 2
	}
	if destination.TravelZone() == "A" {
		fMod -= 2
	}
	if source.TravelZone() == "R" {
		fMod -= 6
	}
	if destination.TravelZone() == "R" {
		fMod -= 6
	}
	dist := hexagon.DistanceHex(source, destination)
	if dist > 1 {
		fMod = fMod - (dist - 1)
	}
	factor = fMod
	return factor, nil
}

func BaseFreightFactor_MGT1_MP(source, destination mWorld) (int, error) {
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
	fMod += dUWP.Pops()
	if source.CoordX() == destination.CoordX() && source.CoordY() == destination.CoordY() {
		return factor, fmt.Errorf("sorce and destination can not have same coordinates")
	}
	sTC := strings.Fields(source.MW_Remarks())
	switch {
	case sUWP.TL() >= 12:
		sTC = append(sTC, "Ht")
	case sUWP.TL() <= 5:
		sTC = append(sTC, "Lt")
	}
	dTC := strings.Fields(destination.MW_Remarks())
	switch {
	case dUWP.TL() >= 12:
		dTC = append(dTC, "Ht")
	case dUWP.TL() <= 5:
		dTC = append(dTC, "Lt")
	}

	//applying uwp factors for noth S and D:
	if sliceContains(sTC, "Ag") {
		fMod += 2
	}
	if sliceContains(dTC, "Ag") {
		fMod += 1
	}
	if sliceContains(sTC, "As") {
		fMod += -3
	}
	if sliceContains(dTC, "As") {
		fMod += 1
	}
	if sliceContains(sTC, "Ba") {
		fMod += -100000
	}
	if sliceContains(dTC, "Ba") {
		fMod += -5
	}
	if sliceContains(sTC, "De") {
		fMod += -3
	}
	if sliceContains(dTC, "De") {
		fMod += 0
	}
	if sliceContains(sTC, "Fl") {
		fMod += -3
	}
	if sliceContains(dTC, "Fl") {
		fMod += 0
	}
	if sliceContains(sTC, "Ga") {
		fMod += 2
	}
	if sliceContains(dTC, "Ga") {
		fMod += 1
	}
	if sliceContains(sTC, "Hi") {
		fMod += 2
	}
	if sliceContains(dTC, "Hi") {
		fMod += 0
	}
	if sliceContains(sTC, "Ic") {
		fMod += -3
	}
	if sliceContains(dTC, "Ic") {
		fMod += 0
	}
	if sliceContains(sTC, "In") {
		fMod += 3
	}
	if sliceContains(dTC, "In") {
		fMod += 2
	}
	if sliceContains(sTC, "Lo") {
		fMod += -5
	}
	if sliceContains(dTC, "Lo") {
		fMod += 0
	}
	if sliceContains(sTC, "Na") {
		fMod += -3
	}
	if sliceContains(dTC, "Na") {
		fMod += 1
	}
	if sliceContains(sTC, "Ni") {
		fMod += -3
	}
	if sliceContains(dTC, "Ni") {
		fMod += 1
	}
	if sliceContains(sTC, "Po") {
		fMod += -3
	}
	if sliceContains(dTC, "Po") {
		fMod += -3
	}
	if sliceContains(sTC, "Ri") {
		fMod += 2
	}
	if sliceContains(dTC, "Ri") {
		fMod += 2
	}
	if sliceContains(sTC, "Wa") {
		fMod += -3
	}
	if sliceContains(dTC, "Wa") {
		fMod += 0
	}
	if source.TravelZone() == "A" {
		fMod += 5
	}
	if destination.TravelZone() == "A" {
		fMod += -5
	}
	if source.TravelZone() == "R" {
		fMod += -5
	}
	if destination.TravelZone() == "R" {
		fMod += -100000
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

func FreightTrafficValues_MGT2_Core(ftv int) (dice, add int) {
	switch ftv {
	default:
		if ftv > 19 {
			return 10, 0
		}
		return 0, 0
	case 2, 3:
		return 1, 0
	case 4, 5:
		return 2, 0
	case 6, 7, 8:
		return 3, 0
	case 9, 10, 11:
		return 4, 0
	case 12, 13, 14:
		return 5, 0
	case 15, 16:
		return 6, 0
	case 17:
		return 7, 0
	case 18:
		return 8, 0
	case 19:
		return 9, 0
	}
}

const (
	Lot_Incidental = iota
	Lot_Minor
	Lot_Major
)

func FreightTrafficValues_MGT1_MP(ftv, lotType int) (dice, add int) {
	minFTV := 0
	switch lotType {
	case Lot_Incidental:
		minFTV = 9
	case Lot_Minor:
		minFTV = 4
	case Lot_Major:
		minFTV = 2
	}
	if ftv < minFTV {
		return 0, 0
	}
	addMod := ftv - minFTV - 4
	return 1, addMod
}
