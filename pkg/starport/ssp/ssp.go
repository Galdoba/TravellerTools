package ssp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	presense_planetary = iota
	presense_orbital
	presense_system
	stance
	classification_Corrupt
	classification_Covert
	classification_Factionalised
	classification_Focussed
	classification_Impersonal
	classification_Militarised
	classification_Pervasive
	classification_Technological
	classification_Volunteer
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
	name       string
	wsp        string
	value      map[int]int
	codes      []string
	balkanised bool
}

type SecurityProfile interface {
	String() string
	Describe() string
}

func NewSecurityProfile(world spData) (*securityProfile, error) {
	sp := securityProfile{}
	sp.name = "UNSET"
	sp.wsp = "UNSET"
	sp.value = make(map[int]int)
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
	//fmt.Println(st, atm, size, gov, law, dp)
	/////////////////////////////////////////
	sp.name = world.MW_Name()
	if pops == 0 {
		return noProfile(), nil
	}
	if gov == 0 || law == 0 {
		return individualResponsibility(), nil
	}

	ppDM := planetaryPresenceDM(size, gov, tc)
	sp.value[presense_planetary] = dp.Roll("2d6").DM(law + ppDM - 7).Sum()
	if sp.value[presense_planetary] < 0 {
		sp.value[presense_planetary] = 0
	}

	opDM := orbitalPresenseDM(st, world.Bases(), size, gov, tc)
	sp.value[presense_orbital] = dp.Roll("2d6").DM(law + opDM - 7).Sum()
	if st == "X" {
		sp.value[presense_orbital] = 0
	}

	if sp.value[presense_orbital] < 0 {
		sp.value[presense_orbital] = 0
	}

	spDM := systemPresenseDM(st, gov, tc, ggPresent)
	sp.value[presense_system] = dp.Roll("2d6").DM(sp.value[presense_orbital] + spDM - 7).Sum()
	if sp.value[presense_system] < 0 {
		sp.value[presense_system] = 0
	}

	stDM := stanceDM(st, atm, gov, tc)
	sp.value[stance] = dp.Roll("2d6").DM(law + stDM - 7).Sum()
	if sp.value[stance] < 0 {
		sp.value[stance] = 0
	}

	secCodes := confirmSecurityCodes(gov, pops, tl, sp.value[presense_planetary], tc, dp)
	govSuff := ""
	if gov == 7 {
		govSuff = "B"
		sp.balkanised = true
	}
	sp.wsp = "S" + ehex.New().Set(sp.value[presense_planetary]).Code() + ehex.New().Set(sp.value[presense_orbital]).Code() +
		ehex.New().Set(sp.value[presense_system]).Code() + "-" + ehex.New().Set(sp.value[stance]).Code() + govSuff
	for _, code := range secCodes {
		switch code {
		case classification_Corrupt:
			sp.wsp += " Cr"
			sp.codes = append(sp.codes, "Cr")
		case classification_Covert:
			sp.wsp += " Co"
			sp.codes = append(sp.codes, "Co")
		case classification_Factionalised:
			sp.wsp += " Fa"
			sp.codes = append(sp.codes, "Fa")
		case classification_Focussed:
			sp.wsp += " Fo"
			sp.codes = append(sp.codes, "Fo")
		case classification_Impersonal:
			sp.wsp += " Ip"
			sp.codes = append(sp.codes, "Ip")
		case classification_Militarised:
			sp.wsp += " Mi"
			sp.codes = append(sp.codes, "Mi")
		case classification_Pervasive:
			sp.wsp += " Pe"
			sp.codes = append(sp.codes, "Pe")
		case classification_Technological:
			sp.wsp += " Te"
			sp.codes = append(sp.codes, "Te")
		case classification_Volunteer:
			sp.wsp += " Vo"
			sp.codes = append(sp.codes, "Vo")
		}
	}
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

func orbitalPresenseDM(st, bases string, size, gov int, tc []string) int {
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

func systemPresenseDM(st string, gov int, tc []string, ggPresent bool) int {
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

func stanceDM(st string, atm, gov int, tc []string) int {
	stDM := 0
	switch st {
	case "X":
		stDM = +2
	}
	atmDM := 0
	switch atmDM {
	case 1, 10:
		atmDM = 1
	case 0, 11, 12:
		atmDM = 2
	}
	govDm := 0
	switch gov {
	case 2, 12:
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
		case "Hi":
			tcDM = tcDM - 2
		case "Ht":
			tcDM = tcDM - 1
		case "Lt":
			tcDM = tcDM + 1
		}
	}
	return stDM + atmDM + govDm + tcDM
}

func confirmSecurityCodes(gov, pop, tl, pp int, tc []string, dp *dice.Dicepool) []int {
	govArrayMap := make(map[int][]int)
	govArrayMap[classification_Corrupt] = []int{1, 3, 5, 6, 7, 8, 9, 11, 13, 14, 15}
	govArrayMap[classification_Covert] = []int{}
	govArrayMap[classification_Factionalised] = []int{4, 5, 6, 9, 11, 12, 13, 14, 15}
	govArrayMap[classification_Focussed] = []int{1, 6, 9, 10, 11, 12, 13, 14, 15}
	govArrayMap[classification_Impersonal] = []int{1, 3, 6, 9, 13, 14, 15}
	govArrayMap[classification_Militarised] = []int{3, 5, 6, 7, 11, 15}
	govArrayMap[classification_Pervasive] = []int{1, 5, 6, 8, 9, 11, 13, 14, 15}
	govArrayMap[classification_Technological] = []int{}
	govArrayMap[classification_Volunteer] = []int{2, 3, 4, 7, 10, 12}
	popArrayMap := make(map[int][]int)
	popArrayMap[classification_Corrupt] = []int{4}
	popArrayMap[classification_Covert] = []int{6}
	popArrayMap[classification_Factionalised] = []int{5}
	popArrayMap[classification_Focussed] = []int{8}
	popArrayMap[classification_Impersonal] = []int{5}
	popArrayMap[classification_Militarised] = []int{4}
	popArrayMap[classification_Pervasive] = []int{1, 9}
	popArrayMap[classification_Technological] = []int{}
	popArrayMap[classification_Volunteer] = []int{1, 2}
	tlArrayMap := make(map[int][]int)
	tlArrayMap[classification_Corrupt] = []int{}
	tlArrayMap[classification_Covert] = []int{}
	tlArrayMap[classification_Factionalised] = []int{}
	tlArrayMap[classification_Focussed] = []int{}
	tlArrayMap[classification_Impersonal] = []int{}
	tlArrayMap[classification_Militarised] = []int{}
	tlArrayMap[classification_Pervasive] = []int{}
	tlArrayMap[classification_Technological] = []int{12}
	tlArrayMap[classification_Volunteer] = []int{}
	ppArrayMap := make(map[int][]int)
	ppArrayMap[classification_Corrupt] = []int{1, 5}
	ppArrayMap[classification_Covert] = []int{1, 5}
	ppArrayMap[classification_Factionalised] = []int{5}
	ppArrayMap[classification_Focussed] = []int{1, 6}
	ppArrayMap[classification_Impersonal] = []int{}
	ppArrayMap[classification_Militarised] = []int{}
	ppArrayMap[classification_Pervasive] = []int{7}
	ppArrayMap[classification_Technological] = []int{}
	ppArrayMap[classification_Volunteer] = []int{}
	tnArrayMap := make(map[int][]int)
	tnArrayMap[classification_Corrupt] = []int{12}
	tnArrayMap[classification_Covert] = []int{10}
	tnArrayMap[classification_Factionalised] = []int{10}
	tnArrayMap[classification_Focussed] = []int{}
	tnArrayMap[classification_Impersonal] = []int{10}
	tnArrayMap[classification_Militarised] = []int{10}
	tnArrayMap[classification_Pervasive] = []int{}
	tnArrayMap[classification_Technological] = []int{}
	tnArrayMap[classification_Volunteer] = []int{5}
	tcArrayMap := make(map[int][]string)
	tcArrayMap[classification_Corrupt] = []string{"Po", "Ri"}
	tcArrayMap[classification_Covert] = []string{}
	tcArrayMap[classification_Factionalised] = []string{}
	tcArrayMap[classification_Focussed] = []string{}
	tcArrayMap[classification_Impersonal] = []string{}
	tcArrayMap[classification_Militarised] = []string{}
	tcArrayMap[classification_Pervasive] = []string{}
	tcArrayMap[classification_Technological] = []string{}
	tcArrayMap[classification_Volunteer] = []string{}
	validCodes := []int{}
	for _, code := range []int{classification_Corrupt, classification_Covert, classification_Factionalised, classification_Focussed, classification_Impersonal, classification_Militarised, classification_Pervasive, classification_Technological, classification_Volunteer} {
		criteria := []bool{}
		criteria = append(criteria, meetGovCriteria(govArrayMap[code], gov))
		criteria = append(criteria, meetIntCriteria(popArrayMap[code], pop))
		criteria = append(criteria, meetIntCriteria(tlArrayMap[code], tl))
		criteria = append(criteria, meetIntCriteria(ppArrayMap[code], pp))
		criteria = append(criteria, meetStringCriteria(tcArrayMap[code], tc))
		r := dp.Roll("2d6").Sum()
		if code == classification_Impersonal && gov == 9 {
			r = r + 5
		}
		met := true
		criteria = append(criteria, meetIntCriteria(tnArrayMap[code], r))
		for _, check := range criteria {
			met = met && check
		}
		if met {
			validCodes = append(validCodes, code)
		}
	}
	return validCodes
}

func meetGovCriteria(govArray []int, gov int) bool {
	if len(govArray) == 0 {
		return true
	}
	for _, meets := range govArray {
		if gov == meets {
			return true
		}
	}
	return false
}

func meetIntCriteria(minmax []int, pop int) bool {
	switch len(minmax) {
	default:
		return false
	case 0:
		return true
	case 1:
		if pop >= minmax[0] {
			return true
		}
	case 2:
		if pop >= minmax[0] && pop <= minmax[1] {
			return true
		}
	}
	return false
}

func meetStringCriteria(array []string, tc []string) bool {
	if len(array) == 0 {
		return true
	}
	for _, c1 := range array {
		for _, c2 := range tc {
			if c1 == c2 {
				return true
			}
		}
	}
	return false
}

func noProfile() *securityProfile {
	return &securityProfile{
		wsp: "No Security Profile",
		value: map[int]int{
			presense_planetary: 0,
			presense_orbital:   0,
			presense_system:    0,
			stance:             0,
		},
	}
}

func individualResponsibility() *securityProfile {
	return &securityProfile{
		wsp: "S000-0 (Individual Responsibility)",
		value: map[int]int{
			presense_planetary: 0,
			presense_orbital:   0,
			presense_system:    0,
			stance:             0,
		},
	}
}

func (sp *securityProfile) String() string {
	return sp.wsp
}

func (sp *securityProfile) Describe() string {
	if sp.wsp == "No Security Profile" {
		return "No Security Profile"
	}
	str := "Security Profile: " + sp.String() + "\n"
	str += fmt.Sprintf("• Planetary presence: %v\n", sp.value[presense_planetary])
	str += fmt.Sprintf("• Orbital presence  : %v\n", sp.value[presense_orbital])
	str += fmt.Sprintf("• System presence   : %v\n", sp.value[presense_system])
	str += fmt.Sprintf("• Security stance   : %v\n", sp.value[stance])
	if sp.balkanised {
		str += "Security Profile is valid for the area around the starport on a balkanised world\n"
	}
	if len(sp.codes) > 0 {
		str += fmt.Sprintf("Security Codes:\n")
	}
	for _, code := range sp.codes {
		switch code {
		case "Cr":
			str += fmt.Sprintf(" Corrupt (%v): Graft, bribery, and self-interest are extremely common in the ranks of the security officers. Travellers should expect fair treatment only if it benefits the officers – or if they can pay for it.\n", strings.TrimPrefix(code, " "))

		case "Co":
			str += fmt.Sprintf(" Covert (%v): Whilst most worlds have small covert security forces, this world's security is predominantly hidden and consists of extensive surveillance, and in some societies a network of citizen informants.\n", strings.TrimPrefix(code, " "))

		case "Fa":
			str += fmt.Sprintf(" Factionalised (%v): Security forces are numerous and often hold very specific mandates. This can lead to inefficiency and bureaucratic infighting that can inconvenience (or be exploited by) the Travellers.\n", strings.TrimPrefix(code, " "))

		case "Fo":
			str += fmt.Sprintf(" Focussed (%v): The strongest security and enforcement is found around key locations and people, with the rest of the world or system having much less. High Presence values with the Focussed code can mean extensive passive monitoring, with significant resources available when needed.\n", strings.TrimPrefix(code, " "))

		case "Ip":
			str += fmt.Sprintf(" Impersonal (%v): The security forces are less concerned with individual rights and justice, and more with the laws themselves and public order. A Difficult (10+) Advocate check can reverse the negative DM on sentencing rolls on these worlds, as Travellers use the letter of the law to their favour.\n", strings.TrimPrefix(code, " "))

		case "Mi":
			str += fmt.Sprintf(" Militarised (%v): All key security forces are military in nature. Typically more heavily armed and armoured than civilian security forces, they will normally be granted significant latitude by the government.\n", strings.TrimPrefix(code, " "))

		case "Pe":
			str += fmt.Sprintf(" Pervasive (%v): Security apparatus is wide-ranging and common. This can vary from constant data-mining of computer networks, to a panopticon of cameras and gunshot sensors, to guards on every door, depending on the Tech Level. Pervasive security may be limited to the planet alone, or reach beyond it.\n", strings.TrimPrefix(code, " "))

		case "Te":
			str += fmt.Sprintf(" Technological (%v): Main security functions are automated, or heavily reliant on hardware and software. Fewer officers will be present, but cameras, drones, and other devices will be very common.\n", strings.TrimPrefix(code, " "))

		case "Vo":
			str += fmt.Sprintf(" Volunteer (%v): Security forces are made up of volunteers, perhaps led by one or two paid full-time officer(s). They will typically be less well-trained but are dedicated to their community.\n", strings.TrimPrefix(code, " "))

		}
	}
	str += "--------------------------------------------------------------------------------"

	return str
}
