package traffic

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	Freight_MGT1_MP   = "_fr_mgt1_mp_"
	Passenger_MGT1_MP = "_ps_mgt1_mp_"
)

func validFactorInstruction(instr string) bool {
	for _, check := range []string{
		Freight_MGT1_MP,
		Passenger_MGT1_MP,
	} {
		if check == instr {
			return true
		}
	}
	return false
}

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

func trafficTradeCodes() []string {
	return []string{"Ag", "As", "Ba", "De", "Fl", "Ga", "Hi", "Ic", "In", "Lo", "Na", "Ni", "Po", "Ri", "Wa", "A", "R"}
}

func getFreightFactorsMap() map[string]int {
	trafficTCmap := make(map[string]int)
	//source freight
	trafficTCmap["s_fr_mgt1_mp_Ag"] = 2
	trafficTCmap["s_fr_mgt1_mp_As"] = -3
	trafficTCmap["s_fr_mgt1_mp_Ba"] = -1000
	trafficTCmap["s_fr_mgt1_mp_De"] = -3
	trafficTCmap["s_fr_mgt1_mp_Fl"] = -3
	trafficTCmap["s_fr_mgt1_mp_Ga"] = 2
	trafficTCmap["s_fr_mgt1_mp_Hi"] = 2
	trafficTCmap["s_fr_mgt1_mp_Ic"] = -3
	trafficTCmap["s_fr_mgt1_mp_In"] = 3
	trafficTCmap["s_fr_mgt1_mp_Lo"] = -5
	trafficTCmap["s_fr_mgt1_mp_Na"] = -3
	trafficTCmap["s_fr_mgt1_mp_Ni"] = -3
	trafficTCmap["s_fr_mgt1_mp_Po"] = -3
	trafficTCmap["s_fr_mgt1_mp_Ri"] = 2
	trafficTCmap["s_fr_mgt1_mp_Wa"] = -3
	trafficTCmap["s_fr_mgt1_mp_A"] = 5
	trafficTCmap["s_fr_mgt1_mp_R"] = -5
	//destination freight
	trafficTCmap["d_fr_mgt1_mp_Ag"] = 1
	trafficTCmap["d_fr_mgt1_mp_As"] = 1
	trafficTCmap["d_fr_mgt1_mp_Ba"] = -5
	trafficTCmap["d_fr_mgt1_mp_De"] = 0
	trafficTCmap["d_fr_mgt1_mp_Fl"] = 0
	trafficTCmap["d_fr_mgt1_mp_Ga"] = 1
	trafficTCmap["d_fr_mgt1_mp_Hi"] = 0
	trafficTCmap["d_fr_mgt1_mp_Ic"] = 0
	trafficTCmap["d_fr_mgt1_mp_In"] = 2
	trafficTCmap["d_fr_mgt1_mp_Lo"] = 0
	trafficTCmap["d_fr_mgt1_mp_Na"] = 1
	trafficTCmap["d_fr_mgt1_mp_Ni"] = 1
	trafficTCmap["d_fr_mgt1_mp_Po"] = -3
	trafficTCmap["d_fr_mgt1_mp_Ri"] = 2
	trafficTCmap["d_fr_mgt1_mp_Wa"] = 0
	trafficTCmap["d_fr_mgt1_mp_A"] = -5
	trafficTCmap["d_fr_mgt1_mp_R"] = -1000
	//source passengers
	trafficTCmap["s_ps_mgt1_mp_Ag"] = 0
	trafficTCmap["s_ps_mgt1_mp_As"] = 1
	trafficTCmap["s_ps_mgt1_mp_Ba"] = -5
	trafficTCmap["s_ps_mgt1_mp_De"] = -1
	trafficTCmap["s_ps_mgt1_mp_Fl"] = 0
	trafficTCmap["s_ps_mgt1_mp_Ga"] = 2
	trafficTCmap["s_ps_mgt1_mp_Hi"] = 0
	trafficTCmap["s_ps_mgt1_mp_Ic"] = 1
	trafficTCmap["s_ps_mgt1_mp_In"] = 2
	trafficTCmap["s_ps_mgt1_mp_Lo"] = 0
	trafficTCmap["s_ps_mgt1_mp_Na"] = 0
	trafficTCmap["s_ps_mgt1_mp_Ni"] = 0
	trafficTCmap["s_ps_mgt1_mp_Po"] = -2
	trafficTCmap["s_ps_mgt1_mp_Ri"] = -1
	trafficTCmap["s_ps_mgt1_mp_Wa"] = 0
	trafficTCmap["s_ps_mgt1_mp_A"] = 2
	trafficTCmap["s_ps_mgt1_mp_R"] = 4
	//destination passengers
	trafficTCmap["d_ps_mgt1_mp_Ag"] = 0
	trafficTCmap["d_ps_mgt1_mp_As"] = -1
	trafficTCmap["d_ps_mgt1_mp_Ba"] = -5
	trafficTCmap["d_ps_mgt1_mp_De"] = -1
	trafficTCmap["d_ps_mgt1_mp_Fl"] = 0
	trafficTCmap["d_ps_mgt1_mp_Ga"] = 2
	trafficTCmap["d_ps_mgt1_mp_Hi"] = 4
	trafficTCmap["d_ps_mgt1_mp_Ic"] = -1
	trafficTCmap["d_ps_mgt1_mp_In"] = 1
	trafficTCmap["d_ps_mgt1_mp_Lo"] = -4
	trafficTCmap["d_ps_mgt1_mp_Na"] = 0
	trafficTCmap["d_ps_mgt1_mp_Ni"] = -1
	trafficTCmap["d_ps_mgt1_mp_Po"] = -1
	trafficTCmap["d_ps_mgt1_mp_Ri"] = 2
	trafficTCmap["d_ps_mgt1_mp_Wa"] = 0
	trafficTCmap["d_ps_mgt1_mp_A"] = -2
	trafficTCmap["d_ps_mgt1_mp_R"] = -4
	return trafficTCmap
}

func differenceTL(s, d mWorld) (int, error) {
	sUWP, err := uwp.FromString(s.MW_UWP())
	if err != nil {
		return 0, err
	}
	dUWP, err := uwp.FromString(d.MW_UWP())
	if err != nil {
		return 0, err
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
	return tMod, nil
}

func BaseFactor(source, destination mWorld, FactorType string) (int, error) {
	if !validFactorInstruction(FactorType) {
		return 0, fmt.Errorf("unknown FactorType instruction '%v'", FactorType)
	}
	sTC := strings.Fields(source.MW_Remarks())
	factorsToAplly := []string{}
	for _, tc := range sTC {
		factorsToAplly = append(factorsToAplly, "s"+FactorType+tc)
	}
	dTC := strings.Fields(destination.MW_Remarks())
	for _, tc := range dTC {
		factorsToAplly = append(factorsToAplly, "d"+FactorType+tc)
	}
	fmt.Println("DEBUG: factors =", factorsToAplly)
	baseFactor, err := differenceTL(source, destination)
	if err != nil {
		return baseFactor, err
	}
	trafficTCmap := getFreightFactorsMap()
	for _, fctr := range factorsToAplly {
		baseFactor = baseFactor + trafficTCmap[fctr]
		fmt.Printf("DEBUG: factor '%v' applyed (%v) - total base factor is now %v\n", fctr, trafficTCmap[fctr], baseFactor)
	}
	fmt.Printf("DEBUG: factor total is %v\n", baseFactor)
	return baseFactor, nil
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
		fmt.Println("Applying factor sTC Ag - fMod += 2")
		fMod += 2
	}
	if sliceContains(dTC, "Ag") {
		fmt.Println("Applying factor dTC Ag - fMod += 1")
		fMod += 1
	}
	if sliceContains(sTC, "As") {
		fmt.Println("Applying factor sTC As - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "As") {
		fmt.Println("Applying factor dTC As - fMod += 1")
		fMod += 1
	}
	if sliceContains(sTC, "Ba") {
		fmt.Println("Applying factor sTC Ba - fMod += -100000")
		fMod += -100000
	}
	if sliceContains(dTC, "Ba") {
		fmt.Println("Applying factor dTC Ba - fMod += -5")
		fMod += -5
	}
	if sliceContains(sTC, "De") {
		fmt.Println("Applying factor sTC De - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "De") {
		fmt.Println("Applying factor dTC De - fMod += 0")
		fMod += 0
	}
	if sliceContains(sTC, "Fl") {
		fmt.Println("Applying factor sTC Fl - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "Fl") {
		fmt.Println("Applying factor dTC Fl - fMod += 0")
		fMod += 0
	}
	if sliceContains(sTC, "Ga") {
		fmt.Println("Applying factor sTC Ga - fMod += 2")
		fMod += 2
	}
	if sliceContains(dTC, "Ga") {
		fmt.Println("Applying factor dTC Ga - fMod += 1")
		fMod += 1
	}
	if sliceContains(sTC, "Hi") {
		fmt.Println("Applying factor sTC Hi - fMod += 2")
		fMod += 2
	}
	if sliceContains(dTC, "Hi") {
		fmt.Println("Applying factor dTC Hi - fMod += 0")
		fMod += 0
	}
	if sliceContains(sTC, "Ic") {
		fmt.Println("Applying factor sTC Ic - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "Ic") {
		fmt.Println("Applying factor dTC Ic - fMod += 0")
		fMod += 0
	}
	if sliceContains(sTC, "In") {
		fmt.Println("Applying factor sTC In - fMod += 3")
		fMod += 3
	}
	if sliceContains(dTC, "In") {
		fmt.Println("Applying factor dTC In - fMod += 2")
		fMod += 2
	}
	if sliceContains(sTC, "Lo") {
		fmt.Println("Applying factor sTC Lo - fMod += -5")
		fMod += -5
	}
	if sliceContains(dTC, "Lo") {
		fmt.Println("Applying factor dTC Lo - fMod += 0")
		fMod += 0
	}
	if sliceContains(sTC, "Na") {
		fmt.Println("Applying factor sTC Na - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "Na") {
		fmt.Println("Applying factor dTC Na - fMod += 1")
		fMod += 1
	}
	if sliceContains(sTC, "Ni") {
		fmt.Println("Applying factor sTC Ni - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "Ni") {
		fmt.Println("Applying factor dTC Ni - fMod += 1")
		fMod += 1
	}
	if sliceContains(sTC, "Po") {
		fmt.Println("Applying factor sTC Po - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "Po") {
		fmt.Println("Applying factor dTC Po - fMod += -3")
		fMod += -3
	}
	if sliceContains(sTC, "Ri") {
		fmt.Println("Applying factor sTC Ri - fMod += 2")
		fMod += 2
	}
	if sliceContains(dTC, "Ri") {
		fmt.Println("Applying factor dTC Ri - fMod += 2")
		fMod += 2
	}
	if sliceContains(sTC, "Wa") {
		fmt.Println("Applying factor sTC Wa - fMod += -3")
		fMod += -3
	}
	if sliceContains(dTC, "Wa") {
		fmt.Println("Applying factor dTC Wa - fMod += 0")
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
