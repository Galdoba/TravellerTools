package ssp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	presense_planetary = iota
	presense_orbital
	presense_system
	stance
)

type spData interface {
	MW_Name() string
	MW_UWP() string
	TravelZone() string
	PBG() string
	Bases() string
	MW_Remarks() string
}

type securityProfile struct {
	name  string
	wsp   string
	value map[int]int
	codes []string
}

func NewSecurityProfile(world spData) (*securityProfile, error) {
	sp := securityProfile{}
	sp.name = "UNSET"
	sp.wsp = "UNSET"
	sp.value = make(map[int]int)
	err := fmt.Errorf("not implemented")
	tc := strings.Fields(world.MW_Remarks())
	uwp, err := uwp.FromString(world.MW_UWP())
	if err != nil {
		return &sp, fmt.Errorf("invalid UWP data: %v", err.Error())
	}
	st := uwp.Starport()
	size := uwp.Size()
	atm := uwp.Atmo()
	pops := uwp.Pops()
	gov := uwp.Govr()
	law := uwp.Laws()
	tl := uwp.TL()
	if tl >= 12 {
		tc = append(tc, "Ht")
	}
	if tl <= 5 && pops >= 1 {
		tc = append(tc, "Lt")
	}
	ggPresent := false
	ggData := strings.Split(world.PBG(), "")
	if len(ggData) != 3 {
		return &sp, fmt.Errorf("invalid PBG data: %v", ggData)
	}
	ggNum, err := strconv.Atoi(ggData[2])
	if err != nil {
		return &sp, fmt.Errorf("invalid PBG data on Gas Gigant: %v", ggData)
	}
	if ggNum > 0 {
		ggPresent = true
	}
	dp := dice.New().SetSeed(world.MW_Name() + world.MW_Remarks() + world.Bases() + world.TravelZone() + world.PBG() + world.MW_UWP())
	fmt.Println(st, atm, size, gov, law, dp)
	/////////////////////////////////////////
	sp.name = world.MW_Name()
	ppDM := planetaryPresenceDM(size, gov, tc)
	sp.value[presense_planetary] = dp.Roll("2d6").DM(law + ppDM - 7).Sum()

	opDM := orbitalPresense(st, world.Bases(), size, gov, tc)
	sp.value[presense_orbital] = dp.Roll("2d6").DM(law + opDM - 7).Sum()
	if st == "X" {
		sp.value[presense_orbital] = 0
	}

	spDM := systemPresense(st, gov, tc, ggPresent)
	sp.value[presense_system] = dp.Roll("2d6").DM(sp.value[presense_orbital] + spDM - 7).Sum()

	return &sp, err
}

func planetaryPresenceDM(size, gov int, tc []string) int {
	dmSize := 0
	switch size {
	case 0, 1:
		dmSize = 2
	case 2, 3:
		dmSize = 1
	}
	if size >= 8 {
		dmSize = -1
	}
	dmGov := 0
	switch gov {
	case 2, 12:
		dmGov = -2
	case 7, 10:
		dmGov = -1
	case 1, 5, 11:
		dmGov = 1
	case 6, 13, 14, 15:
		dmGov = 2
	}
	tcDM := 0
	for _, code := range tc {
		switch code {
		case "Hi":
			tcDM = tcDM - 2
		case "Lo":
			tcDM = tcDM - 1
		case "Ht", "Ri":
			tcDM = tcDM + 1
		}
	}
	return dmSize + dmGov + tcDM
}

func orbitalPresense(st, bases string, size, gov int, tc []string) int {
	stDM := 0
	switch st {
	case "E":
		stDM = -2
	case "D":
		stDM = -1
	case "B":
		stDM = 1
	case "A":
		stDM = 2
	}
	sizeDM := 0
	switch size {
	case 3, 4:
		sizeDM = 1
	case 0, 1, 2:
		sizeDM = 2
	}
	if size >= 10 {
		sizeDM = -1
	}
	govDm := 0
	switch gov {
	case 2, 7, 12:
		govDm = -2
	case 10:
		govDm = -1
	case 1, 5, 11:
		govDm = 1
	case 6, 13, 14, 15:
		govDm = 2
	}
	tcDM := 0
	for _, code := range tc {
		switch code {
		case "Lo", "Lt":
			tcDM = tcDM - 2
		case "Po":
			tcDM = tcDM - 1
		case "Ag", "In", "Ht":
			tcDM = tcDM + 1
		case "Ri":
			tcDM = tcDM + 2
		}
	}
	navBaseDM := 0
	navyBasePresent := false
	for _, base := range strings.Split(bases, "") {
		if navyBasePresent {
			break
		}
		for _, naval := range strings.Split("ABFGHJKLNOPRTUZ", "") {
			if base == naval {
				navyBasePresent = true
				break
			}
		}
	}
	if navyBasePresent {
		navBaseDM = 1
	}
	return stDM + sizeDM + govDm + tcDM + navBaseDM
}

func systemPresense(st string, gov int, tc []string, ggPresent bool) int {
	stDM := 0
	switch st {
	case "E":
		stDM = -2
	case "C", "D":
		stDM = -1
	case "A":
		stDM = 1
	}
	govDm := 0
	switch gov {
	case 7:
		govDm = -2
	case 1, 9, 10, 12:
		govDm = -1
	case 6:
		govDm = 2
	}
	tcDM := 0
	for _, code := range tc {
		switch code {
		case "Lo", "Po":
			tcDM = tcDM - 2
		case "Lt", "Ni":
			tcDM = tcDM - 1
		case "Ri":
			tcDM = tcDM + 1
		}
	}
	ggDM := 0
	if ggPresent {
		ggDM = -2
	}
	return stDM + govDm + tcDM + ggDM
}
