package traffic

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation"
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
	dist := astrogation.DistanceRaw(source.CoordX(), source.CoordY(), destination.CoordX(), destination.CoordY())
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
	if source.CoordX() == destination.CoordX() && source.CoordY() == destination.CoordY() {
		return factor, fmt.Errorf("sorce and destination can not have same coordinates")
	}
	sTC := strings.Fields(source.MW_Remarks())
	dTC := strings.Fields(destination.MW_Remarks())
	fmt.Println(sTC, dTC)
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
	dist := astrogation.DistanceRaw(source.CoordX(), source.CoordY(), destination.CoordX(), destination.CoordY())
	if dist > 1 {
		fMod = fMod - (dist - 1)
	}
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
