package portsec

import (
	"fmt"
	"strconv"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	org_No                 = "Not organised force"
	org_SmallPT            = "Small, part-time security force"
	org_SmallProf          = "Small professional security force"
	org_Modest             = "Modest-sized professional security force"
	org_Large              = "Large professional security force"
	org_Huge               = "Huge professional security force"
	org_Enormous           = "Enormous security force"
	fund_No                = "No significant funding"
	fund_VeryUnderfunded   = "Grossly underfunded"
	fund_Underfunded       = "Underfunded "
	fund_Normal            = "Normal level of funding"
	fund_Well              = "Well-funded"
	fund_Lavishly          = "Lavishly funded"
	equip_None             = "None provided"
	equip_Minimal          = "Minimal"
	equip_Basic            = "Basic"
	equip_Standard         = "Standard"
	equip_Well             = "Well Equipped"
	equip_Heavily          = "Heavily Equipped"
	equip_Lavishly         = "Lavishly Equipped"
	competenceShambolic    = "Shambolic competence"
	competenceChaotic      = "Chaotic competence"
	competenceDisorganised = "Disorganised competence"
	competenceLow          = "Low competence"
	competenceNormal       = "Normal competence"
	competenceHigh         = "High competence"
	competenceVHigh        = "Very High competence"
	competenceFlawless     = "Near-Flawless competence"
	corruptionWracked      = "Total corruption"
	corruptionSevere       = "Severe corruption"
	corruptionSome         = "Some corruption"
	corruptionNormal       = "Normal corruption"
	corruptionLow          = "Low corruption"
	corruptionVLow         = "Very Low corruption"
	corruptionNone         = "Virtually no corruption"
	responseAversion       = "Aversion"
	responseVReluctant     = "Very Reluctant"
	responseReluctant      = "Reluctant"
	responseNormal         = "Normal"
	responseRobust         = "Robust"
	responseAgressive      = "Agressive"
	responseExtreme        = "Trigger-Happy"
)

type StarportSecurityForces struct {
	portName        string
	worldPopulation float64
	basePersonal    float64
	organisation    string
	funding         string
	equipment       string
	competence      string
	corruption      string
	response        string
	checksDM        int
	fiascoTN        int
	rout            int
	resp            int
	elite           int
}

type Port interface {
	MW_Name() string
	MW_UWP() string
	PBG() string
}

func GenerateSecurityForces(port Port) (*StarportSecurityForces, error) {
	ssf := StarportSecurityForces{}
	seed := port.MW_Name() + port.MW_UWP() + port.PBG() + port.MW_UWP()
	dp := dice.New().SetSeed(seed)
	ssf.portName = port.MW_Name()
	ssf.fiascoTN = -1
	ssf.checksDM = -5
	population := 1.0

	err := fmt.Errorf("err value was not adressed")
	uwpData, err := uwp.FromString(port.MW_UWP())
	if err != nil {
		return &ssf, err
	}
	pops := uwpData.Pops()
	law := uwpData.Laws()
	st := uwpData.Starport()
	for i := 0; i < pops; i++ {
		population = population * 10
	}
	pbg := port.PBG()
	p := string(pbg[0])
	pi, err := strconv.Atoi(p)
	ssf.worldPopulation = population * float64(pi)

	org, mult := defineOrganisation(dp, pops, law)
	if st == "X" {
		org = org_No
	}
	ssf.organisation = org
	ssf.basePersonal = (ssf.worldPopulation / 500) * mult
	fund, equipDM, sep := defineFunding(dp, pops, st)
	equip := defineEquipment(dp, equipDM)
	if org == org_No {
		fund = fund_No
	}
	if fund == fund_No {
		equip = equip_None
	}
	ssf.funding = fund
	ssf.equipment = equip
	ssf.defineCompetence(dp)
	ssf.corruption, ssf.response = defineCorruptionAndresponce(dp)
	ssf.rout = int(ssf.basePersonal) * sep[0] / 100
	ssf.resp = int(ssf.basePersonal) * sep[1] / 100
	ssf.elite = int(ssf.basePersonal) * sep[2] / 100

	return &ssf, err
}

func (ssf *StarportSecurityForces) String() string {
	str := fmt.Sprintf("Security Forces at %v starport:\n", ssf.portName)
	underline := ""
	for len(underline) < 80 {
		underline += "-"
	}

	str += fmt.Sprintf("Organisation : %v\n", ssf.organisation)
	str += fmt.Sprintf("Funding      : %v\n", ssf.funding)
	str += fmt.Sprintf("Equipment    : %v\n", ssf.equipment)
	str += fmt.Sprintf("Competence   : %v\n", ssf.competence)
	str += fmt.Sprintf("Response     : %v\n", ssf.response)
	str += underline
	routineG, respG, eliteG := ssf.rout, ssf.resp, ssf.elite
	if routineG > 50000 {
		routineG = 50000
	}
	if respG > 7000 {
		respG = 7000
	}
	if eliteG > 3000 {
		eliteG = 3000
	}
	str += fmt.Sprintf("\n Starport of %v is defended by total of %v security personel. ", ssf.portName, routineG+respG+eliteG)
	str += fmt.Sprintf("Local security are mostly equiped with %v ", describe(ssf.equipment))
	str += fmt.Sprintf("\n %v gives [DM: %v] on all checks made by Security Force. ", ssf.competence, ssf.checksDM)
	if ssf.fiascoTN > 0 {
		str += fmt.Sprintf("\n Every check has a potential of a Fiasco (%v+ on strait 2D roll). ", ssf.fiascoTN)
	}
	str += fmt.Sprintf("\n %v indicates that %v", ssf.corruption, describe(ssf.corruption))
	str += fmt.Sprintf("\n If incident arise local Security will feel %v", describe(ssf.response))
	if ssf.resp > 0 {
		str += fmt.Sprintf("\n %v of personnel undertake most normal tasks such as standing guard, patrolling an area, carrying out customs searches and the like. Routine personnel tend to be equipped for the possibility of trouble, such as a patrol officer carrying a sidearm, perhaps with access to more powerful weapons at need.\n", routineG)
		str += fmt.Sprintf(" While %v are equipped to back up their routine colleagues with heavy firepower or specialist capabilities. In a society that has considerable numbers of psions, this might mean the possession of psionic shielding equipment or personnel may be psionically adept in their own right.\n", respG)
		if ssf.elite > 0 {
			str += fmt.Sprintf(" There %v personnel are equipped as best as possible for the worst situations. These may be heavy elements of the ruler's personal guard or a specialist security formation, perhaps even a 'bodyguard' regiment of the planetary army.", eliteG)
		}
	} else {
		str += fmt.Sprintf("\n Security personnel is not separeted on any grades due to lack of professional training.")
	}
	str += "\n" + underline

	return str
}

func defineOrganisation(dp *dice.Dicepool, pops, laws int) (string, float64) {
	dm := 0
	switch pops {
	default:
		dm += 3
	case 0, 1, 2, 3:
		dm += -2
	case 6, 7:
		dm += 1
	case 4, 5:
		//dm += 0
	}
	switch laws {
	default:
		dm += 3
	case 0:
		dm += -4
	case 1, 2, 3:
		dm += -1
	case 4, 5:
		//dm += 0
	case 6, 7, 8:
		dm += 1
	}
	r := dp.Roll("2d6").DM(dm).Sum()
	switch r {
	default:
		if r < 1 {
			return org_No, 0.0
		}
		return org_Enormous, 5.0
	case 1, 2, 3:
		return org_SmallPT, 0.2
	case 4, 5, 6:
		return org_SmallProf, 0.4
	case 7, 8, 9:
		return org_Modest, 1
	case 10, 11, 12:
		return org_Large, 1.5
	case 13, 14, 15:
		return org_Huge, 2.5

	}
}

func defineFunding(dp *dice.Dicepool, pops int, st string) (string, int, []int) {
	dm := 0
	switch pops {
	case 0, 1, 2:
		dm += -2
	case 3, 4:
		dm += -1
	case 5, 6:
		dm += 0
	case 7, 8:
		dm += 1
	default:
		dm += 2
	}
	switch st {
	case "E":
		dm += -2
	case "D":
		dm += -1
	case "B":
		dm += 1
	case "A":
		dm += 2
	}
	r := dp.Roll("2d6").DM(dm).Sum()
	if r < 2 {
		return fund_No, -1000, []int{100, 0, 0}
	}
	switch r {
	case 2, 3:
		return fund_VeryUnderfunded, -4, []int{100, 0, 0}
	case 4, 5, 6:
		return fund_Underfunded, -2, []int{100, 0, 0}
	case 7, 8, 9:
		return fund_Normal, 0, []int{90, 9, 1}
	case 10, 11, 12:
		return fund_Well, 2, []int{80, 15, 5}
	default:
		return fund_Lavishly, 4, []int{70, 20, 10}
	}
	return "", 0, []int{}
}

func defineEquipment(dp *dice.Dicepool, dm int) string {
	r := dp.Roll("2d6").DM(dm).Sum()
	if r < 1 {
		return equip_None
	}
	switch r {
	case 1, 2, 3:
		return equip_Minimal
	case 4, 5, 6:
		return equip_Basic
	case 7, 8, 9:
		return equip_Standard
	case 10, 11, 12:
		return equip_Well
	case 13, 14, 15:
		return equip_Heavily
	default:
		return equip_Lavishly

	}
}

func (ssf *StarportSecurityForces) defineCompetence(dp *dice.Dicepool) {
	dm := 0
	if ssf.organisation == org_SmallPT {
		dm += -4
	}
	if ssf.funding == fund_Underfunded {
		dm += -1
	}
	if ssf.funding == fund_VeryUnderfunded {
		dm += -3
	}
	if ssf.organisation == org_Huge {
		dm += -1
	}
	if ssf.organisation == org_Enormous {
		dm += -2
	}
	if ssf.equipment == equip_Well || ssf.equipment == equip_Heavily || ssf.equipment == equip_Lavishly {
		dm += 2
	}
	r := dp.Roll("2d6").DM(dm).Sum()
	ssf.fiascoTN = 0
	if r < 1 {
		ssf.competence = competenceShambolic
		ssf.checksDM = -4
		ssf.fiascoTN = 11
	}
	switch r {
	case 1, 2:
		ssf.competence = competenceChaotic
		ssf.checksDM = -3
		ssf.fiascoTN = 12
	case 3, 4:
		ssf.competence = competenceDisorganised
		ssf.checksDM = -2
	case 5, 6:
		ssf.competence = competenceLow
		ssf.checksDM = -1
	case 7, 8:
		ssf.competence = competenceNormal
		ssf.checksDM = 0
	case 9, 10:
		ssf.competence = competenceHigh
		ssf.checksDM = 1
	case 11, 12:
		ssf.competence = competenceVHigh
		ssf.checksDM = 2
	default:
		ssf.competence = competenceFlawless
		ssf.checksDM = 3
	}
}

func defineCorruptionAndresponce(dp *dice.Dicepool) (string, string) {
	c := dp.Roll("2d6").Sum()
	r := dp.Roll("2d6").Sum()
	cor := ""
	res := ""
	switch c {
	case 2:
		cor = corruptionWracked
	case 3, 4:
		cor = corruptionSevere
	case 5, 6:
		cor = corruptionSome
	case 7:
		cor = corruptionNormal
	case 8, 9:
		cor = corruptionLow
	case 10, 11:
		cor = corruptionVLow
	case 12:
		cor = corruptionNone
	}
	switch r {
	case 2:
		res = responseAversion
	case 3, 4:
		res = responseVReluctant
	case 5, 6:
		res = responseReluctant
	case 7:
		res = responseNormal
	case 8, 9:
		res = responseRobust
	case 10, 11:
		res = responseAgressive
	case 12:
		res = responseExtreme
	}
	return cor, res
}

func describe(val string) string {
	switch val {
	default:
		return "#NO DESCR"
	case org_No:
		return org_No
	case org_SmallPT:
		return org_SmallPT
	case org_SmallProf:
		return org_SmallProf
	case org_Modest:
		return org_Modest
	case org_Large:
		return org_Large
	case org_Huge:
		return org_Huge
	case org_Enormous:
		return org_Enormous
	case fund_No:
	case fund_VeryUnderfunded:
	case fund_Underfunded:
	case fund_Normal:
	case fund_Well:
	case fund_Lavishly:
	case equip_None:
		return "nothig, exept personal weapons and gear"
	case equip_Minimal:
		return "a cheap sidearm, minimal ammunition and a few necessary items such as a flashlight and handcuffs. Minimally equipped personnel usually lack 'official' communications equipment, though they may have their own."
	case equip_Basic:
		return "a standard sidearm and necessary tools such as flashlight, communications equipment and hand-held instruments if the Tech Level permits."
	case equip_Standard:
		return "a full set of law enforcement tools, good communications equipment with access to data-transfer capability (if Tech Level permits) and light body armour. Personnel may have access to additional weapons such as a shotgun in a patrol vehicle."
	case equip_Well:
		return "good body armour and personal weapons, with access to heavier weaponry including possibly light automatics (submachineguns or assault rifles) and specialist weapons such as sniper rifles. These may not be carried all the time but be kept in an accessible location."
	case equip_Heavily:
		return "plentiful supplies of all equipment noted above, some or all of which is of a Tech Level 1-2 above local. A heavily equipped security force may have access to military grade combat armour for some elite units."
	case equip_Lavishly:
		return "very high quality equipment which may be 3-4 Tech Levels above local if appropriate. A heavily equipped security force will have access to military grade combat armour for some response units and may even have a battle dress equipped unit if the Tech Level permits."
	case competenceShambolic:
	case competenceChaotic:
	case competenceDisorganised:
	case competenceLow:
	case competenceNormal:
	case competenceHigh:
	case competenceVHigh:
	case competenceFlawless:
	case corruptionWracked:
		return "Personnel actively seek advantage or profit from all aspects of their duties."
	case corruptionSevere:
		return "Virtually impossible to get anything done without bribery, influence or pursuing someone's agenda."
	case corruptionSome:
		return "DM: 2 on checks that match the corruption type (such as an attempt to bribe an official or prosecute someone the state considers undesirable)"
	case corruptionNormal:
		return "Few organisations are completely clean."
	case corruptionLow:
		return "DM: -2 on checks to bribe officials or attempts to persuade them to take unscrupulous action."
	case corruptionVLow:
		return "DM: -4 on checks that go against the organisation's high moral standards."
	case corruptionNone:
		return "Accepting a bribe or using influence is virtually unthinkable."
	case responseAversion:
		return "Extreme aversion to the use of force. Personnel will wait too long before deploying weapons or using them, and are very reluctant to call for backup."
	case responseVReluctant:
		return "Very reluctant to use force. Personnel will make every effort at a peaceful resolution and may risk their own safety to avoid using weapons."
	case responseReluctant:
		return "Reluctant to threaten or use force, but entirely willing at need."
	case responseNormal:
		return "Normal levels of prudence, balancing negotiation against the threat of force."
	case responseRobust:
		return "Emphasis on officer safety; robust policies on the deployment of weapons in any threatening situation."
	case responseAgressive:
		return "Aggressive, threatening approach to most situations, backed up by weapons and response teams if needed."
	case responseExtreme:
		return "Positively trigger-happy."
	}
	return ""
}
