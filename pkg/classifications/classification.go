package classifications

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/planets"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	// KEY_uwp         = "uwp"
	// KEY_hz          = "hz"
	// KEY_mw          = "mw"
	// KEY_sw          = "sw"
	// KEY_travel      = "travel"
	// KEY_sector      = "sector"
	// KEY_manual      = "manual"
	// uwpApproval     = 0
	// hzApproval      = 1
	// mwApproval      = 2
	// swApproval      = 3
	// travellApproval = 4
	// sectorApproval  = 5
	// manualApproval  = 6
	NONE = iota
	manual
	As
	De
	Fl
	Ga
	He
	Ic
	Oc
	Va
	Wa
	Di
	Ba
	Lo
	Ni
	Ph
	Hi
	Pa
	Ag
	Na
	Px
	Pi
	In
	Po
	Pr
	Ri
	Lt
	Ht
	Fr
	Ho
	Co
	Lk
	Tr
	Tu
	Tz
	Fa
	Mi
	Mr
	Pe
	Re
	Cp
	Cs
	Cx
	Cy
	Sa
	Fo
	Pz
	Da
	Ab
	An
	Ts
	Bo
	Ds
	TypePlanetary  = "Planetary"
	TypePopulation = "Population"
	TypeEconomic   = "Economic"
	TypeClimate    = "Climate"
	TypeSecondary  = "Secondary"
	TypePolitical  = "Political"
	TypeSpecial    = "Special"
)

func ListAll() []string {
	list := []string{}
	for i := As; i <= Bo; i++ {
		cls := Call(i)
		list = append(list, cls.code)
	}
	return list
}

type classificationCode struct {
	val            int
	code           string
	classification string
	tcType         string
	description    string
	sourceBook     string
}

type Classification interface {
}

func Manual(code, classification, tcType, description, sourceBook string) Classification {
	return classificationCode{manual, code, classification, tcType, description, sourceBook}
}

func Call(i int) *classificationCode {
	tc := classificationCode{}
	tc.val = i
	tc.sourceBook = "T5 Book 2"
	switch tc.val {
	default:
		tc.code = "??"
		tc.classification = "[Unknown]"
		tc.tcType = TypeSpecial
		tc.description = "[No Description]"
		return &tc
	case As:
		tc.code = "As"
		tc.classification = "Asteroid Belt"
		tc.tcType = TypePlanetary
		tc.description = "The world is an asteroid belt which is the primary world or mainworld in the system. It is a producer of raw materials and semi-finished goods, especially ores, metals, and minerals."
	case De:
		tc.code = "De"
		tc.classification = "Desert World"
		tc.tcType = TypePlanetary
		tc.description = "The world has no open or standing water. This lack of water significantly reduces the level of agricultural development."
	case Fl:
		tc.code = "Fl"
		tc.classification = "Fluid Oceans"
		tc.tcType = TypePlanetary
		tc.description = "The world's oceans are not composed of water. Non-water oceans may be valuable sources of raw materials for industry."
	case Ga:
		tc.code = "Ga"
		tc.classification = "Garden World"
		tc.tcType = TypePlanetary
		tc.description = "The world is hospitable to most sophonts. Its size, atmosphere, and hydrographic make it an extremely attractive world. A Garden World has a safe environment which does not require protective equipment for humans and sophonts which share the human environment."
	case He:
		tc.code = "He"
		tc.classification = "Hellworld"
		tc.tcType = TypePlanetary
		tc.description = "The world is inhospitable to most sophonts. Its size, atmosphere, and hydrographic make it an extremely unattractive world."
	case Ic:
		tc.code = "Ic"
		tc.classification = "Ice capped"
		tc.tcType = TypePlanetary
		tc.description = "The world's water is locked in ice caps."
	case Oc:
		tc.code = "Oc"
		tc.classification = "Ocean World"
		tc.tcType = TypePlanetary
		tc.description = "The world surface is covered with very deep seas. There is no (less than a hundredth) land above sea level."
	case Va:
		tc.code = "Va"
		tc.classification = "Vacuum World"
		tc.tcType = TypePlanetary
		tc.description = "The world has no atmosphere."
	case Wa:
		tc.code = "Wa"
		tc.classification = "Water World"
		tc.tcType = TypePlanetary
		tc.description = "The world surface is covered with water; there is very little land (less than 10%) above the water surface."
	case Di:
		tc.code = "Di"
		tc.classification = "Die-Back"
		tc.tcType = TypePopulation
		tc.description = "The world was once extensively settled and developed, but at some time in the last thousand years its inhabiting sophonts died out leaving behind the remnants of their civilization."
	case Ba:
		tc.code = "Ba"
		tc.classification = "Barren World"
		tc.tcType = TypePopulation
		tc.description = "The world has no population, government, or law level. It has never been developed; it has no local infrastructure beyond the starport (if that)."
	case Lo:
		tc.code = "Lo"
		tc.classification = "Low Population"
		tc.tcType = TypePopulation
		tc.description = "The world has a non-zero-population less than 10,000. Low Population fluctuates wildly and may change significantly on a yearly (or less) basis.\nLocals are Transients: merchants, corporate employees, military, security, or research personnel."
	case Ni:
		tc.code = "Ni"
		tc.classification = "Non-Industrial"
		tc.tcType = TypePopulation
		tc.description = "The world has a non-zero population (more than 10,000 and less than one million). The TC Non-Industrial remains constant and reflects an expected population level.\nInhabitants of a Non-Industrial world are Settlers: part of a permanent settlement not yet a Colony."
	case Ph:
		tc.code = "Ph"
		tc.classification = "Pre-High"
		tc.tcType = TypePopulation
		tc.description = "The world is a candidate for elevation to the High Population trade classification; its population level is just below the requirements for High."
	case Hi:
		tc.code = "Hi"
		tc.classification = "High Population"
		tc.tcType = TypePopulation
		tc.description = "The world's population is one billion or more (Pop = 9 or A or more). High population worlds, because of the economy of scale for production, produce quality inexpensive trade goods."
	case Pa:
		tc.code = "Pa"
		tc.classification = "Pre-Agricultural"
		tc.tcType = TypeEconomic
		tc.description = "The world is a candidate for the Agricultural trade classification; its population is just outside the requirement for Agricultural."
	case Ag:
		tc.code = "Ag"
		tc.classification = "Agricultural"
		tc.tcType = TypeEconomic
		tc.description = "The world has climate and conditions which promote farming and ranching. It is a producer of inexpen sive foodstuffs. It also is a source of unusual, exotic, or strange delicacies."
	case Na:
		tc.code = "Na"
		tc.classification = "Non-Agricultural"
		tc.tcType = TypeEconomic
		tc.description = "The world is unable to produce enough food agriculturally to feed its population; synthetic food production generally meets basic food needs."
	case Px:
		tc.code = "Px"
		tc.classification = "Prison. Exile Camp."
		tc.tcType = TypeEconomic
		tc.description = "The non-mainworld population consists of criminals or undesirables transported here from other worlds."
	case Pi:
		tc.code = "Pi"
		tc.classification = "Pre-Industrial"
		tc.tcType = TypeEconomic
		tc.description = "The world is a candidate for the Industrial trade classification; its population is just below the requirements."
	case In:
		tc.code = "In"
		tc.classification = "Industrial"
		tc.tcType = TypeEconomic
		tc.description = "The world has a strong manufacturing infrastructure and is a producer of many types of goods."
	case Po:
		tc.code = "Po"
		tc.classification = "Poor"
		tc.tcType = TypeEconomic
		tc.description = "The world has poor grade living conditions: a scarcity of water and a relatively sparse atmosphere."
	case Pr:
		tc.code = "Pr"
		tc.classification = "Pre-Rich"
		tc.tcType = TypeEconomic
		tc.description = "The world is a candidate for the Rich trade classification; its population is just outside the criteria for Rich."
	case Ri:
		tc.code = "Ri"
		tc.classification = "Rich"
		tc.tcType = TypeEconomic
		tc.description = "The world has an untainted atmosphere which is comfortable and attractive for most sophonts, and has a population suitable as a workforce."
	case Lt:
		tc.code = "Lt"
		tc.classification = "Low Tech"
		tc.tcType = TypeEconomic
		tc.description = "The world is pre-industrial and cannot produce advanced goods."
		tc.sourceBook = "MGT2 CRB"
	case Ht:
		tc.code = "Ht"
		tc.classification = "High Tech"
		tc.tcType = TypeEconomic
		tc.description = "The world is among the most technologicaly advanced in Charted Space."
		tc.sourceBook = "MGT2 CRB"
	case Ts:
		tc.code = "Ts"
		tc.classification = "Temperature Swing"
		tc.tcType = TypeClimate
		tc.description = "The world's temperature swings from roasting during the day to frozen at night."
		tc.sourceBook = "MGT2 CRB"
	case Fr:
		tc.code = "Fr"
		tc.classification = "Frozen"
		tc.tcType = TypeClimate
		tc.description = "The world's environmental temperatures are well below the freezing point of many gases."
	case Co:
		tc.code = "Co"
		tc.classification = "Cold"
		tc.tcType = TypeClimate
		tc.description = "The world is at the lower temperature range of human endurance. Little liquid water, extencive ice caps, few clouds."
	case Lk:
		tc.code = "Lk"
		tc.classification = "Locked"
		tc.tcType = TypeClimate
		tc.description = "The world is a satellite (in orbits Ay through Em) which is locked to the planet it orbits. A Locked satellite does not have a Twilight Zone; its day length equals the time it takes to orbit its planet."
	case Tu:
		tc.code = "Tu"
		tc.classification = "Tundra"
		tc.tcType = TypeClimate
		tc.description = "The world is relatively colder than normal (although it is considered habitable). The world has a Cold climate (at the lower limits of human temperature endurance)."
	case Ho:
		tc.code = "Ho"
		tc.classification = "Hot"
		tc.tcType = TypeClimate
		tc.description = "The world is at the upper temperature range of human endurance. Small or no ice caps, little liquid water. Most water in the form of clouds."
	case Tr:
		tc.code = "Tr"
		tc.classification = "Tropic"
		tc.tcType = TypeClimate
		tc.description = "The world is relatively warmer than normal (although it is considered habitable)."
	case Bo:
		tc.code = "Bo"
		tc.classification = "Boiling"
		tc.tcType = TypeClimate
		tc.description = "Boiling world. No ice caps, little liquid water."
		tc.sourceBook = "MGT2 Core"
	case Tz:
		tc.code = "Tz"
		tc.classification = "Twilight Zone"
		tc.tcType = TypeClimate
		tc.description = "The world is tidally locked with a Temperate band at the Twilight Zone, plus a Hot region (hemisphere) facing the Primary and a Cold region (hemisphere) away from the Primary."
	case Fa:
		tc.code = "Fa"
		tc.classification = "Farming"
		tc.tcType = TypeSecondary
		tc.description = "Farming The world has climate and conditions which promote farming and ranching. In addition, it is in the Habitable Zone and not a Mainworld."
	case Mi:
		tc.code = "Mi"
		tc.classification = "Mining"
		tc.tcType = TypeSecondary
		tc.description = "The world is the site of extensive mineral resource exploitation. It is not a Mainworld and is located in a star system with an Industrial Mainworld."
	case Mr:
		tc.code = "Mr"
		tc.classification = "Military Rule"
		tc.tcType = TypeSecondary
		tc.description = "The non-Mainworld is ruled by the military from a nearby world."
	case Pe:
		tc.code = "Pe"
		tc.classification = "Penal Colony"
		tc.tcType = TypeSecondary
		tc.description = "The world is a dumping ground for individuals who will not / do not / cannot conform to standards of behavior."
	case Re:
		tc.code = "Re"
		tc.classification = "Reserve"
		tc.tcType = TypeSecondary
		tc.description = "The world has been set aside (by the highest levels of government) to preserve indigenous life forms, to delay resource development, or to frustrate inquiry into local conditions."
	case Cp:
		tc.code = "Cp"
		tc.classification = "Subsector Capital"
		tc.tcType = TypePolitical
		tc.description = "The world is the political center of a group of tens or dozens of star systems (typically a subsector)."
	case Cs:
		tc.code = "Cs"
		tc.classification = "Sector Capital"
		tc.tcType = TypePolitical
		tc.description = "The world is the political center of a group of hundreds of star systems (typically a sector)."
	case Cx:
		tc.code = "Cx"
		tc.classification = "Capital"
		tc.tcType = TypePolitical
		tc.description = "The world is the overall political center of an interstellar government controlling thousands of star systems."
	case Cy:
		tc.code = "Cy"
		tc.classification = "Colony"
		tc.tcType = TypePolitical
		tc.description = "The world is a colony Owned by the Most Important, Highest Population, Highest TL world within 6 hexes. Add the remark O:[hex] (=hex of owning world)."
	case Sa:
		tc.code = "Sa"
		tc.classification = "Satelite"
		tc.tcType = TypeSpecial
		tc.description = "The world is the satellite of a planet (or gas giant) in the system."
	case Fo:
		tc.code = "Fo"
		tc.classification = "Forbidden"
		tc.tcType = TypeSpecial
		tc.description = "Some conditions, customs, laws, life forms, climate, economics, or other circumstance presents an active threat to the health and well-being of individuals. The world is a TAS Red Zone."
	case Pz:
		tc.code = "Pz"
		tc.classification = "Puzzle"
		tc.tcType = TypeSpecial
		tc.description = "Some aspect of the world (conditions, customs, laws, life forms, climate, economics, or other) is not well or easily understood by typical visitors. The world is a TAS Amber Zone."
	case Da:
		tc.code = "Da"
		tc.classification = "Dangerous"
		tc.tcType = TypeSpecial
		tc.description = "Some aspect of the world (conditions, customs, laws, life forms, climate, economics, or other) is not well understood or easily understood by typical visitors, and it presents a danger. The world is a TAS Amber Zone."
	case Ab:
		tc.code = "Ab"
		tc.classification = "Data Repository"
		tc.tcType = TypeSpecial
		tc.description = "The world has a centralized collection point for information and data. Organizations and governments deposit records of their transactions and output in this collection point. The TC refers to AAB, the Imperial designation for data repositories."
	case An:
		tc.code = "An"
		tc.classification = "Antient Site"
		tc.tcType = TypeSpecial
		tc.description = "The world (or the star system) includes one or more locations identified as the ruins of the long-dead race called the Ancients. Ancient Sites are exploited for the Artifact remains of this long dead technological civilization."

	}
	return &tc
}

func (tc *classificationCode) String() string {
	return tc.code
}

func (tc *classificationCode) TcType() string {
	return tc.tcType
}

func (tc *classificationCode) Classification() string {
	return tc.classification
}

func (tc *classificationCode) Description() string {
	return tc.description
}

func Prints(tc []classificationCode) string {
	s := ""
	for _, t := range tc {
		s += t.String() + " "
	}
	s = strings.TrimSuffix(s, " ")
	return s
}

func Values(tc []classificationCode) []int {
	sl := []int{}
	for _, t := range tc {
		sl = append(sl, t.val)
	}
	return sl
}

/////////////////////////////

func FromUWP(prof uwp.UWP) ([]*classificationCode, error) {
	return parceCodes(prof)
}

func parceCodes(data uwp.UWP) ([]*classificationCode, error) {
	parsedTradeCodes := []*classificationCode{}
	//S S      A     H   P  G  L  - T
	codes := tradeCodeDemandsPPE()
	for codeNum, tg := range codes {
		//fmt.Println("GO ", codeNum, tg)
		tcMatch := true
		values := strings.Split(tg, " ")
		for valNum, d := range values {
			if !tcMatch {
				break
			}
			if d == "--" || d == "-" {
				continue
			}
			switch valNum {
			case 1:
				if !strings.Contains(d, ehex.New().Set(data.Size()).Code()) {
					tcMatch = tcMatch && false
				}
			case 2:
				if !strings.Contains(d, ehex.New().Set(data.Atmo()).Code()) {
					tcMatch = tcMatch && false
				}
			case 3:
				if !strings.Contains(d, ehex.New().Set(data.Hydr()).Code()) {
					tcMatch = tcMatch && false
				}
			case 4:
				if !strings.Contains(d, ehex.New().Set(data.Pops()).Code()) {
					tcMatch = tcMatch && false
				}
			case 5:
				if !strings.Contains(d, ehex.New().Set(data.Govr()).Code()) {
					tcMatch = tcMatch && false
				}
			case 6:
				if !strings.Contains(d, ehex.New().Set(data.Laws()).Code()) {
					tcMatch = tcMatch && false
				}
			case 8:
				if !strings.Contains(d, ehex.New().Set(data.TL()).Code()) {
					tcMatch = tcMatch && false
				}
			}

		}
		if tcMatch {
			//fmt.Println("add code", codeNum)
			clss := Call(codeNum)
			if clss.code != "??" {
				parsedTradeCodes = append(parsedTradeCodes, clss)
			}

		}
		//fmt.Println(codeNum, "-------------------")

	}
	return parsedTradeCodes, nil
}

func tradeCodeDemandsPPE() []string {
	return []string{
		"-- 0 0 0 -- -- -- - --",                                //As
		"-- -- 23456789 0 -- -- -- - --",                        //De
		"-- -- ABC 123456789A -- -- -- - --",                    //Fl
		"-- 678 568 567 -- -- -- - --",                          //Ga
		"-- 3456789ABCDEF 2479ABC 012 -- -- -- - --",            //He
		"-- -- 01 123456789A -- -- -- - --",                     //Ic
		"-- ABCDEF 3456789DEF A -- -- -- - --",                  //Oc
		"-- -- 0 -- -- -- -- - --",                              //Va
		"-- 3456789 3456789DEF A -- -- -- - --",                 //Wa
		"-- -- -- -- 0 0 0 - 123456789ABCDEFGHJKLMNPQRSTUVWXYZ", //Di

		"EXHY -- -- -- 0 0 0 - --",                    //Ba
		"-- -- -- -- 123 -- -- - --",                  //Lo
		"-- -- -- -- 456 -- -- - --",                  //Ni
		"-- -- -- -- 8 -- -- - --",                    //Ph
		"-- -- -- -- 9ABCDEF -- -- - --",              //Hi
		"-- -- 456789 45678 48 -- -- - --",            //Pa
		"-- -- 456789 45678 567 -- -- - --",           //Ag
		"-- -- 0123 0123 6789ABCDEF -- -- - --",       //Na
		"-- -- 23AB 12345 3456 -- 6789ABCDEFGHJ - --", //Px
		"-- -- 012479ABC -- 9ABCDEF -- -- - --",       //In

		"-- -- 2345 0123 -- -- -- - --",                 //Po
		"-- -- 68 -- 59 -- -- - --",                     //Pr
		"-- -- 68 -- 678 -- -- - --",                    //Ri
		"-- -- -- -- -- -- -- - 12345",                  //Lt
		"-- -- -- -- -- -- -- - CDEFGHJKLMNPQRSTUVWXYZ", //Ht

		//"-- 23456789ABCDEF -- 123456789A -- -- -- - --", //Fr
		//"-- 6789 456789 34567 -- -- -- - --",            //Tr
		//"-- 6789 456789 34567 -- -- -- - --",            //Tu
	}
}

// type classifications struct {
// 	confirmedByUWP []string

// 	excludeMWonly bool
// 	excludeSWonly bool
// 	errors        []error
// }

// type tcInput struct {
// 	key, val string
// }

// func Input(key, val string) tcInput {
// 	return tcInput{key, val}
// }

// type World interface {
// 	UWP() string
// 	Stellar() string
// 	HZ() int
// }

// func Analize(input ...tcInput) *classifications {
// 	cls := classifications{}
// 	for _, data := range input {
// 		switch data.key {
// 		case KEY_uwp:
// 			u, err := uwp.FromString(data.val)
// 			if err != nil {
// 				cls.errors = append(cls.errors, err)
// 				continue
// 			}
// 			parcedTC, err := parceCodesFull(u)
// 			if err != nil {
// 				cls.errors = append(cls.errors, err)
// 				continue
// 			}
// 			cls.confirmedByUWP = append(cls.confirmedByUWP, parcedTC...)

// 		}
// 	}
// 	return &cls
// }

// func (cls *classifications) excludeMW() {

// }

// type clssfctn struct {
// 	code             string
// 	approvalExpected []int
// }

// func mapCLS() map[int]clssfctn {
// 	clMap := make(map[int]clssfctn)
// 	clMap[0] = clssfctn{"As", []int{1, 0, 1, 0, 0, 0, 0}}
// 	clMap[1] = clssfctn{"De", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[2] = clssfctn{"Fl", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[3] = clssfctn{"Ga", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[4] = clssfctn{"He", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[5] = clssfctn{"Ic", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[6] = clssfctn{"Oc", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[7] = clssfctn{"Va", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[8] = clssfctn{"Wa", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[9] = clssfctn{"Di", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[10] = clssfctn{"Ba", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[11] = clssfctn{"Lo", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[12] = clssfctn{"Ni", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[13] = clssfctn{"Ph", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[14] = clssfctn{"Hi", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[15] = clssfctn{"Pa", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[16] = clssfctn{"Ag", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[17] = clssfctn{"Na", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[18] = clssfctn{"Px", []int{1, 0, 1, 0, 0, 0, 0}}
// 	clMap[19] = clssfctn{"In", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[20] = clssfctn{"Po", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[21] = clssfctn{"Pr", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[22] = clssfctn{"Ri", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[23] = clssfctn{"Lt", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[24] = clssfctn{"Ht", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[25] = clssfctn{"Fr", []int{1, 1, 0, 0, 0, 0, 0}}
// 	clMap[26] = clssfctn{"Ho", []int{0, 1, 0, 0, 0, 0, 0}}
// 	clMap[27] = clssfctn{"Co", []int{0, 1, 0, 0, 0, 0, 0}}
// 	clMap[28] = clssfctn{"Lk", []int{0, 0, 0, 0, 0, 0, 1}}
// 	clMap[29] = clssfctn{"Tr", []int{1, 1, 0, 0, 0, 0, 0}}
// 	clMap[30] = clssfctn{"Tu", []int{1, 1, 0, 0, 0, 0, 0}}
// 	clMap[31] = clssfctn{"Tz", []int{0, 0, 0, 0, 0, 0, 1}}
// 	clMap[32] = clssfctn{"Fa", []int{1, 0, 0, 1, 0, 0, 0}}
// 	clMap[33] = clssfctn{"Mn", []int{1, 0, 0, 1, 0, 1, 0}}
// 	clMap[34] = clssfctn{"Mr", []int{1, 0, 0, 1, 0, 0, 0}}
// 	clMap[35] = clssfctn{"Pe", []int{1, 0, 0, 1, 0, 0, 0}}
// 	clMap[36] = clssfctn{"Re", []int{1, 0, 0, 0, 0, 0, 0}}
// 	clMap[37] = clssfctn{"Cp", []int{0, 0, 1, 0, 0, 1, 0}}
// 	clMap[38] = clssfctn{"Cs", []int{0, 0, 1, 0, 0, 1, 0}}
// 	clMap[39] = clssfctn{"Cx", []int{0, 0, 1, 0, 0, 1, 0}}
// 	clMap[40] = clssfctn{"Cy", []int{1, 0, 1, 0, 0, 1, 0}}
// 	clMap[41] = clssfctn{"Sa", []int{0, 0, 0, 0, 0, 0, 1}}
// 	clMap[42] = clssfctn{"Fo", []int{0, 0, 0, 0, 0, 1, 0}}
// 	clMap[43] = clssfctn{"Pz", []int{1, 0, 0, 0, 0, 1, 0}}
// 	clMap[44] = clssfctn{"Da", []int{1, 0, 0, 0, 0, 1, 0}}
// 	clMap[45] = clssfctn{"Ab", []int{0, 0, 0, 0, 0, 0, 1}}
// 	clMap[46] = clssfctn{"An", []int{0, 0, 0, 0, 0, 0, 1}}
// 	return clMap
// }

// func fullTCsList() []string {
// 	return []string{
// 		"As", "De", "Fl", "Ga", "He", "Ic", "Oc", "Va", "Wa",
// 		"Di", "Ba", "Lo", "Ni", "Ph", "Hi",
// 		"Pa", "Ag", "Na", "Px", "Pi", "In", "Po", "Pr", "Ri", "Lt", "Ht",
// 		"Fr", "Ho", "Co", "Lk", "Tr", "Tu", "Tz",
// 		"Fa", "Mi", "Mr", "Pe", "Re",
// 		"Cp", "Cs", "Cx", "Cy",
// 		"Sa", "Fo", "Pz", "Da", "Ab", "An",
// 	}
// }

// func FromUWP(prof uwp.UWP) ([]string, error) {
// 	return parceCodes(prof)
// }

// func parceCodes(data uwp.UWP) ([]string, error) {
// 	parsedTradeCodes := []string{}
// 	//S S      A     H   P  G  L  - T
// 	codes := tradeCodeDemands()
// 	for codeNum, tg := range codes {
// 		//fmt.Println("GO ", codeNum, tg)
// 		tcMatch := true
// 		values := strings.Split(tg, " ")
// 		for valNum, d := range values {
// 			if tcMatch == false {
// 				//	fmt.Println("BREAK")
// 				break
// 			}
// 			if d == "--" || d == "-" {
// 				//fmt.Println("DEBUG: skip", valNum, d, tcMatch)
// 				continue
// 			}
// 			//fmt.Println("DEBUG: test", valNum, d, tcMatch)
// 			switch valNum {
// 			case 1:
// 				if !strings.Contains(d, ehex.New().Set(data.Size()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 2:
// 				if !strings.Contains(d, ehex.New().Set(data.Atmo()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 3:
// 				if !strings.Contains(d, ehex.New().Set(data.Hydr()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 4:
// 				if !strings.Contains(d, ehex.New().Set(data.Pops()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 5:
// 				if !strings.Contains(d, ehex.New().Set(data.Govr()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 6:
// 				if !strings.Contains(d, ehex.New().Set(data.Laws()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 8:
// 				if !strings.Contains(d, ehex.New().Set(data.TL()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			}

// 		}
// 		if tcMatch {
// 			//fmt.Println("add code", codeNum)
// 			switch codeNum {
// 			default:
// 				return nil, fmt.Errorf("unknown trade code position from tradeCodeDemands()")
// 			case 0:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ag")
// 			case 1:
// 				parsedTradeCodes = append(parsedTradeCodes, "As")
// 			case 2:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ba")
// 			case 3:
// 				parsedTradeCodes = append(parsedTradeCodes, "De")
// 			case 4:
// 				parsedTradeCodes = append(parsedTradeCodes, "Fl")
// 			case 5:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ga")
// 			case 6:
// 				parsedTradeCodes = append(parsedTradeCodes, "Hi")
// 			case 7:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ht")
// 			case 8:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ic")
// 			case 9:
// 				parsedTradeCodes = append(parsedTradeCodes, "In")
// 			case 10:
// 				parsedTradeCodes = append(parsedTradeCodes, "Lo")
// 			case 11:
// 				parsedTradeCodes = append(parsedTradeCodes, "Lt")
// 			case 12:
// 				parsedTradeCodes = append(parsedTradeCodes, "Na")
// 			case 13:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ni")
// 			case 14:
// 				parsedTradeCodes = append(parsedTradeCodes, "Po")
// 			case 15:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ri")
// 			case 16:
// 				parsedTradeCodes = append(parsedTradeCodes, "Va")
// 			case 17:
// 				parsedTradeCodes = append(parsedTradeCodes, "Wa")
// 			}
// 		}
// 		//fmt.Println(codeNum, "-------------------")

// 	}
// 	return parsedTradeCodes, nil
// }

// func parceCodesFull(data uwp.UWP) ([]string, error) {
// 	parsedTradeCodes := []string{}
// 	//S S      A     H   P  G  L  - T
// 	codes := tradeCodeDemandsFull()
// 	for codeNum, tg := range codes {
// 		//fmt.Println("GO ", codeNum, tg)
// 		tcMatch := true
// 		values := strings.Split(tg, " ")
// 		for valNum, d := range values {
// 			if tcMatch == false {
// 				//	fmt.Println("BREAK")
// 				break
// 			}
// 			if d == "--" || d == "-" {
// 				//fmt.Println("DEBUG: skip", valNum, d, tcMatch)
// 				continue
// 			}
// 			//fmt.Println("DEBUG: test", valNum, d, tcMatch)
// 			switch valNum {
// 			case 1:
// 				if !strings.Contains(d, ehex.New().Set(data.Size()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 2:
// 				if !strings.Contains(d, ehex.New().Set(data.Atmo()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 3:
// 				if !strings.Contains(d, ehex.New().Set(data.Hydr()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 4:
// 				if !strings.Contains(d, ehex.New().Set(data.Pops()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 5:
// 				if !strings.Contains(d, ehex.New().Set(data.Govr()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 6:
// 				if !strings.Contains(d, ehex.New().Set(data.Laws()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			case 8:
// 				if !strings.Contains(d, ehex.New().Set(data.TL()).Code()) {
// 					tcMatch = tcMatch && false
// 				}
// 			}

// 		}
// 		if tcMatch {
// 			//fmt.Println("add code", codeNum)
// 			switch codeNum {
// 			default:
// 				return nil, fmt.Errorf("unknown trade code position from tradeCodeDemands()")
// 			case 0:
// 				parsedTradeCodes = append(parsedTradeCodes, "As")
// 			case 1:
// 				parsedTradeCodes = append(parsedTradeCodes, "De")
// 			case 2:
// 				parsedTradeCodes = append(parsedTradeCodes, "Fl")
// 			case 3:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ga")
// 			case 4:
// 				parsedTradeCodes = append(parsedTradeCodes, "He")
// 			case 5:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ic")
// 			case 6:
// 				parsedTradeCodes = append(parsedTradeCodes, "Oc")
// 			case 7:
// 				parsedTradeCodes = append(parsedTradeCodes, "Va")
// 			case 8:
// 				parsedTradeCodes = append(parsedTradeCodes, "Wa")
// 			case 9:
// 				parsedTradeCodes = append(parsedTradeCodes, "Di")
// 			case 10:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ba")
// 			case 11:
// 				parsedTradeCodes = append(parsedTradeCodes, "Lo")
// 			case 12:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ni")
// 			case 13:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ph")
// 			case 14:
// 				parsedTradeCodes = append(parsedTradeCodes, "Hi")
// 			case 15:
// 				parsedTradeCodes = append(parsedTradeCodes, "Pa")
// 			case 16:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ag")
// 			case 17:
// 				parsedTradeCodes = append(parsedTradeCodes, "Na")
// 			case 18:
// 				parsedTradeCodes = append(parsedTradeCodes, "Px")
// 			case 19:
// 				parsedTradeCodes = append(parsedTradeCodes, "In")
// 			case 20:
// 				parsedTradeCodes = append(parsedTradeCodes, "Po")
// 			case 21:
// 				parsedTradeCodes = append(parsedTradeCodes, "Pr")
// 			case 22:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ri")
// 			case 23:
// 				parsedTradeCodes = append(parsedTradeCodes, "Lt")
// 			case 24:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ht")
// 			case 25:
// 				parsedTradeCodes = append(parsedTradeCodes, "Fr")
// 			case 26:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ho")
// 			case 27:
// 				parsedTradeCodes = append(parsedTradeCodes, "Co")
// 			case 28:
// 				parsedTradeCodes = append(parsedTradeCodes, "Lk")
// 			case 29:
// 				parsedTradeCodes = append(parsedTradeCodes, "Tr")
// 			case 30:
// 				parsedTradeCodes = append(parsedTradeCodes, "Tu")
// 			case 31:
// 				parsedTradeCodes = append(parsedTradeCodes, "Tz")
// 			case 32:
// 				parsedTradeCodes = append(parsedTradeCodes, "Fa")
// 			case 33:
// 				parsedTradeCodes = append(parsedTradeCodes, "Mn")
// 			case 34:
// 				parsedTradeCodes = append(parsedTradeCodes, "Mr")
// 			case 35:
// 				parsedTradeCodes = append(parsedTradeCodes, "Pe")
// 			case 36:
// 				parsedTradeCodes = append(parsedTradeCodes, "Re")
// 			case 37:
// 				parsedTradeCodes = append(parsedTradeCodes, "Cp")
// 			case 38:
// 				parsedTradeCodes = append(parsedTradeCodes, "Cs")
// 			case 39:
// 				parsedTradeCodes = append(parsedTradeCodes, "Cx")
// 			case 40:
// 				parsedTradeCodes = append(parsedTradeCodes, "Cy")
// 			case 41:
// 				parsedTradeCodes = append(parsedTradeCodes, "Sa")
// 			case 42:
// 				parsedTradeCodes = append(parsedTradeCodes, "Fo")
// 			case 43:
// 				parsedTradeCodes = append(parsedTradeCodes, "Pz")
// 			case 44:
// 				parsedTradeCodes = append(parsedTradeCodes, "Da")
// 			case 45:
// 				parsedTradeCodes = append(parsedTradeCodes, "Ab")
// 			case 46:
// 				parsedTradeCodes = append(parsedTradeCodes, "An")
// 			}
// 		}
// 		//fmt.Println(codeNum, "-------------------")

// 	}
// 	return parsedTradeCodes, nil
// }

// func tradeCodeDemands() []string {
// 	return []string{
// 		"-- 456789 45678 567 -- -- -- - --",          //Ag
// 		"-- 0 0 0 -- -- -- - --",                     //As
// 		"-- -- -- -- 0 0 0 - --",                     //Ba
// 		"-- -- 23456789 0 -- -- -- - --",             //De
// 		"-- -- ABCDEF 123456789A -- -- -- - --",      //Fl
// 		"-- 678 568 567 -- -- -- - --",               //Ga
// 		"-- -- -- -- 9ABCDEF -- -- - --",             //Hi
// 		"-- -- -- -- -- -- -- - CDEFGH",              //Ht
// 		"-- -- 01 123456789A -- -- -- - --",          //Ic
// 		"-- -- 012479ABC -- 9ABCDEF -- -- - --",      //In
// 		"-- -- -- -- 123 -- -- - --",                 //Lo
// 		"-- -- -- -- 123456789ABCDEF -- -- - 012345", //Lt
// 		"-- -- 0123 0123 6789ABCDEF -- -- - --",      //Na
// 		"-- -- -- -- 456 -- -- - --",                 //Ni
// 		"-- -- 2345 0123 -- -- -- - --",              //Po
// 		"-- -- 68 -- 678 456789 -- - --",             //Ri
// 		"-- -- 0 -- -- -- -- - --",                   //Va
// 		"-- -- 3456789DEF A -- -- -- - --",           //Wa

// 	}
// }

// func tradeCodeDemandsFull() []string {
// 	return []string{
// 		"-- 0 0 0 -- -- -- - --",                                //As
// 		"-- -- 23456789 0 -- -- -- - --",                        //De
// 		"-- -- ABC 123456789A -- -- -- - --",                    //Fl
// 		"-- 678 568 567 -- -- -- - --",                          //Ga
// 		"-- 3456789ABCDEF 2479ABC 012 -- -- -- - --",            //He
// 		"-- -- 01 123456789A -- -- -- - --",                     //Ic
// 		"-- ABCDEF 3456789DEF A -- -- -- - --",                  //Oc
// 		"-- -- 0 -- -- -- -- - --",                              //Va
// 		"-- 3456789 3456789DEF A -- -- -- - --",                 //Wa
// 		"-- -- -- -- 0 0 0 - 123456789ABCDEFGHJKLMNPQRSTUVWXYZ", //Di
// 		"EXHY -- -- -- 0 0 0 - --",                              //Ba
// 		"-- -- -- -- 123 -- -- - --",                            //Lo
// 		"-- -- -- -- 456 -- -- - --",                            //Ni
// 		"-- -- -- -- 8 -- -- - --",                              //Ph
// 		"-- -- -- -- 9ABCDEF -- -- - --",                        //Hi
// 		"-- -- 456789 45678 48 -- -- - --",                      //Pa
// 		"-- -- 456789 45678 567 -- -- - --",                     //Ag
// 		"-- -- 0123 0123 6789ABCDEF -- -- - --",                 //Na
// 		"-- -- 23AB 12345 3456 -- 6789ABCDEFGHJ - --",           //Px
// 		"-- -- 012479ABC -- 9ABCDEF -- -- - --",                 //In
// 		"-- -- 2345 0123 -- -- -- - --",                         //Po
// 		"-- -- 68 -- 59 -- -- - --",                             //Pr
// 		"-- -- 68 -- 678 -- -- - --",                            //Ri
// 		"-- -- -- -- -- -- -- - 12345",                          //Lt
// 		"-- -- -- -- -- -- -- - CDEFGHJKLMNPQRSTUVWXYZ",         //Ht
// 		"-- 23456789 -- 123456789A -- -- -- - --",               //Fr
// 		"-- -- -- -- -- -- -- - --",                             //Ho
// 		"-- -- -- -- -- -- -- - --",                             //Co
// 		"-- -- -- -- -- -- -- - --",                             //Lk
// 		"-- 6789 456789 34567 -- -- -- - --",                    //Tr
// 		"-- 6789 456789 34567 -- -- -- - --",                    //Tu
// 		"-- -- -- -- -- -- -- - --",                             //Tz
// 		"-- -- 456789 45678 23456 -- -- - --",                   //Fa
// 		"-- -- -- -- 23456 -- -- - --",                          //Mn
// 		"-- -- -- -- -- -- -- - --",                             //Mr
// 		"-- -- 23AB 12345 3456 6 6789ABCDEFGHJ - --",            //Pe
// 		"-- -- -- -- 1234 6 45 - --",                            //Re
// 		"A -- -- -- -- -- -- - --",                              //Cp
// 		"A -- -- -- -- -- -- - --",                              //Cs
// 		"A -- -- -- -- -- -- - --",                              //Cx
// 		"-- -- -- -- 56789A 6 0123 - --",                        //Cy
// 		"-- -- -- -- -- -- -- - --",                             //Sa
// 		"-- -- -- -- -- -- -- - --",                             //Fo
// 		"-- -- -- -- 789ABCDEF -- -- - --",                      //Pz
// 		"-- -- -- -- 0123456 -- -- - --",                        //Da
// 		"-- -- -- -- -- -- -- - --",                             //Ab
// 		"-- -- -- -- -- -- -- - --",                             //An
// 	}
// }
/*
Ab
Data Repository
The world has a centralized collection point for information and data. Organizations and governments deposit records of their transactions and output in this collection point.
The TC refers to AAB, the Imperial designation for data repositories.


An
Ancient Site
The world (or the star system) includes one or more locations identified as the ruins of the long-dead race called the Ancients. Ancient Sites are exploited for the Artifact remains of this long dead technological civilization.

Cp
Subsector Capital
The world is the political center of a group of tens or dozens of star systems (typically a subsector).

Cs
Sector Capital
The world is the political center of a group of hundreds of star systems (typically a sector).

Cx
Imperial Capital
The world is the overall political center of an interstellar government controlling thousands of star systems.

Cy
Colony
The world is a colony Owned by the Most Important, Highest Population, Highest TL world within 6 hexes.
Add the remark O:nnnn (=hex of owning world).

Da
Dangerous
Some aspect of the world (conditions, customs, laws, life forms, climate, economics, or other) is not well understood or easily understood by typical visitors, and it presents a danger. The world is a TAS Amber Zone.



Fa
Farming
The world has climate and conditions which promote farming and ranching. In addition, it is in the Habitable Zone and not a Mainworld.


Fo
Forbidden
Some conditions, customs, laws, life forms, climate, economics, or other circumstance presents an active threat to the health and well-being of individuals. The world is a TAS Red Zone.


Lk
Locked
The world is a satellite (in orbits Ay through Em) which is locked to the planet it orbits. A Locked satellite does not have a Twilight Zone; its day length equals the time it takes to orbit its planet.

Mi
Mining
The world is the site of extensive mineral resource exploitation. It is not a Mainworld and is located in a star system with an Industrial Mainworld.

Mr
Military Rule
The non-Mainworld is ruled by the military from a nearby world.


Pe
Penal Colony
The world is a dumping ground for individuals who will not / do not / cannot conform to standards of behavior.


Po
Poor
The world has poor grade living conditions: a scarcity of water and a relatively sparse atmosphere.





Pz
Puzzle
Some aspect of the world (conditions, customs, laws, life forms, climate, economics, or other) is not well or easily understood by typical visitors.
The world is a TAS Amber Zone.

Re
Reserve
The world has been set aside (by the highest levels of government) to preserve indigenous life forms, to delay resource development, or to frustrate inquiry into local conditions.


Sa
Satellite
The world is the satellite of a planet (or gas giant) in the system.

Tr
Tropic
The world is relatively warmer than normal (although it is considered habitable). Its orbit is at the inner (warmer) edge of the Habitable Zone.




Tz
Twilight Zone
The world is tidally locked with a Temperate band at the Twilight Zone, plus a Hot region (hemisphere) facing the Primary and a Cold region (hemisphere) away from the Primary.


*/

/*
Port
Size
Atmo
Hydr
Pops
Govr
Laws
Tech
worldtype
HZvar
SateliteOrbit
MW
AmberZone
RedZone
*/

type tcRequirements struct {
	Port             string
	Size             string
	Atmo             string
	Hydr             string
	Pops             string
	Govr             string
	Laws             string
	Tech             string
	worldtype        string
	HZvar            string
	PlanetaryOrbit   string
	SateliteOrbit    string
	MW               string
	AmberZone        bool
	RedZone          bool
	MilitaryRule     bool //MT-Refery Manual 29
	ResearchLab      bool
	SubsectorCapital bool
	SectorCapital    bool
	Capital          bool
	Colony           bool
	DataRepository   bool
	AntientSite      bool
}

func (req *tcRequirements) slice() []string {
	sl := []string{}
	sl = append(sl, req.Port)
	sl = append(sl, req.Size)
	sl = append(sl, req.Atmo)
	sl = append(sl, req.Hydr)
	sl = append(sl, req.Pops)
	sl = append(sl, req.Govr)
	sl = append(sl, req.Laws)
	sl = append(sl, req.Tech)
	sl = append(sl, req.worldtype)
	sl = append(sl, req.HZvar)
	sl = append(sl, req.PlanetaryOrbit)
	sl = append(sl, req.SateliteOrbit)
	sl = append(sl, req.MW)
	switch req.AmberZone {
	case true:
		sl = append(sl, "A")
	case false:
		sl = append(sl, "")
	}
	switch req.RedZone {
	case true:
		sl = append(sl, "R")
	case false:
		sl = append(sl, "")
	}
	switch req.MilitaryRule {
	case true:
		sl = append(sl, "Mr")
	case false:
		sl = append(sl, "")
	}
	switch req.ResearchLab {
	case true:
		sl = append(sl, "Rs")
	case false:
		sl = append(sl, "")
	}
	switch req.SubsectorCapital {
	case true:
		sl = append(sl, "Cp")
	case false:
		sl = append(sl, "")
	}
	switch req.SectorCapital {
	case true:
		sl = append(sl, "Cs")
	case false:
		sl = append(sl, "")
	}
	switch req.Capital {
	case true:
		sl = append(sl, "Cx")
	case false:
		sl = append(sl, "")
	}
	switch req.Colony {
	case true:
		sl = append(sl, "Cy")
	case false:
		sl = append(sl, "")
	}
	switch req.DataRepository {
	case true:
		sl = append(sl, "Ab")
	case false:
		sl = append(sl, "")
	}
	switch req.AntientSite {
	case true:
		sl = append(sl, "An")
	case false:
		sl = append(sl, "")
	}

	return sl
}

type world interface {
	Data(string) string
}

func worldClassificationSlice(w world) []string {
	sl := []string{}
	sl = append(sl, w.Data(profile.KEY_PORT))               //  Port
	sl = append(sl, w.Data(profile.KEY_SIZE))               //  Size
	sl = append(sl, w.Data(profile.KEY_ATMO))               //  Atmo
	sl = append(sl, w.Data(profile.KEY_HYDR))               //  Hydr
	sl = append(sl, w.Data(profile.KEY_POPS))               //  Pops
	sl = append(sl, w.Data(profile.KEY_GOVR))               //  Govr
	sl = append(sl, w.Data(profile.KEY_LAWS))               //  Laws
	sl = append(sl, w.Data(profile.KEY_TL))                 //  Tech
	sl = append(sl, w.Data(profile.KEY_WORLDTYPE))          //  worldtype
	sl = append(sl, w.Data(profile.KEY_HABITABLE_ZONE_VAR)) //  HZvar
	sl = append(sl, w.Data(profile.KEY_PLANETARY_ORBIT))    //  PlanetaryOrbit
	sl = append(sl, w.Data(profile.KEY_SATELITE_ORBIT))     //  SateliteOrbit
	sl = append(sl, w.Data(profile.KEY_MAINWORLD))          //  MW
	if w.Data("Amber Zone") != "" {
		sl = append(sl, "A")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Red Zone") != "" {
		sl = append(sl, "R")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Military Rule") != "" {
		sl = append(sl, "Mr")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Research Lab") != "" {
		sl = append(sl, "Rs")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Subsector Capital") != "" {
		sl = append(sl, "Cp")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Sector Capital") != "" {
		sl = append(sl, "Cs")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Capital") != "" {
		sl = append(sl, "Cx")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Colony") != "" {
		sl = append(sl, "Cy")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Data Repository") != "" {
		sl = append(sl, "Ab")
	} else {
		sl = append(sl, "-")
	}
	if w.Data("Antient Site") != "" {
		sl = append(sl, "An")
	} else {
		sl = append(sl, "-")
	}
	for i := range sl {
		if sl[i] == "" {
			sl[i] = "-"
		}
	}
	return sl
}

func Evaluate(worldData world) []int {
	tcMatch := []int{}
	worldSl := worldClassificationSlice(worldData)
	for code := As; code <= Bo; code++ {

		req := requirements(code)
		reqSl := req.slice()
		if reqMatch(code, reqSl, worldSl) {
			tcMatch = append(tcMatch, code)
		}
	}
	return tcMatch
}

func reqMatch(code int, reqSl, worldSl []string) bool {
	for i, str := range reqSl {
		if str == "" {
			continue
		}
		if strings.Contains(str, worldSl[i]) {
			continue
		}
		return false
	}
	return true
}

func requirements(code int) tcRequirements {
	req := tcRequirements{}
	switch code {
	default:
		panic(fmt.Sprintf("code %v is invalid", code))
	case As:
		req = tcRequirements{
			Size:      "0",
			Atmo:      "0",
			Hydr:      "0",
			worldtype: planets.WORLDTYPE_Planetoid,
		}
	case De:
		req = tcRequirements{
			Atmo: "23456789",
			Hydr: "0",
		}
	case Fl:
		req = tcRequirements{
			Atmo: "ABC",
			Hydr: "123456789A",
		}
	case Ga:
		req = tcRequirements{
			Size: "678",
			Atmo: "568",
			Hydr: "567",
		}
	case He:
		req = tcRequirements{
			Size: "3456789ABC",
			Atmo: "2479ABC",
			Hydr: "012",
		}
	case Ic:
		req = tcRequirements{
			Atmo: "01",
			Hydr: "123456789A",
		}
	case Oc:
		req = tcRequirements{
			Size: "ABCDEFGHJK",
			Atmo: "3456789DEF",
			Hydr: "A",
		}
	case Va:
		req = tcRequirements{
			Atmo: "0",
		}
	case Wa:
		req = tcRequirements{
			Size: "3456789",
			Atmo: "3456789DEF",
			Hydr: "A",
		}
	case Di:
		req = tcRequirements{
			Pops: "0",
			Govr: "0",
			Laws: "0",
			Tech: "123456789ABCDEFGHJKLMNPQRSTUVWXYZ",
		}
	case Ba:
		req = tcRequirements{
			Port: "EXHY",
			Pops: "0",
			Govr: "0",
			Laws: "0",
		}
	case Lo:
		req = tcRequirements{
			Pops: "123",
		}
	case Ni:
		req = tcRequirements{
			Pops: "456",
		}
	case Ph:
		req = tcRequirements{
			Pops: "8",
		}
	case Hi:
		req = tcRequirements{
			Pops: "9ABCDEF",
		}
	case Pa:
		req = tcRequirements{
			Atmo: "456789",
			Hydr: "45678",
			Pops: "48",
		}
	case Ag:
		req = tcRequirements{
			Atmo: "456789",
			Hydr: "45678",
			Pops: "567",
		}
	case Na:
		req = tcRequirements{
			Atmo: "0123",
			Hydr: "0123",
			Pops: "6789ABCDEF",
		}
	case Px:
		req = tcRequirements{
			Atmo: "23AB",
			Hydr: "12345",
			Pops: "3456",
			Laws: "6789",
			MW:   "1",
		}
	case Pi:
		req = tcRequirements{
			Atmo: "012479",
			Pops: "78",
		}
	case In:
		req = tcRequirements{
			Atmo: "012479ABC",
			Pops: "9ABCDEF",
		}
	case Po:
		req = tcRequirements{
			Atmo: "2345",
			Hydr: "0123",
		}
	case Pr:
		req = tcRequirements{
			Atmo: "68",
			Pops: "59",
		}
	case Ri:
		req = tcRequirements{
			Atmo: "68",
			Hydr: "678",
		}
	case Lt:
		req = tcRequirements{
			Pops: "123456789ABCDEF",
			Tech: "12345",
		}
	case Ht:
		req = tcRequirements{
			Pops: "123456789ABCDEF",
			Tech: "CDEFGHJKLMNPQRSTUVW",
		}
	case Fr:
		req = tcRequirements{
			Size:  "23456789",
			Hydr:  "123456789A",
			HZvar: "CDEFGHJKLMNPQRSTUVW",
		}
	case Ho:
		req = tcRequirements{
			HZvar: "9",
		}
	case Co:
		req = tcRequirements{
			HZvar: "B",
		}
	case Lk:
		req = tcRequirements{
			SateliteOrbit: "ABCDEFGH1JKLM",
		}
	case Tr:
		req = tcRequirements{
			Size:  "6789",
			Atmo:  "456789",
			Hydr:  "34567",
			HZvar: "9",
		}
	case Tu:
		req = tcRequirements{
			Size:  "6789",
			Atmo:  "456789",
			Hydr:  "34567",
			HZvar: "B",
		}
	case Tz:
		req = tcRequirements{
			PlanetaryOrbit: "01",
		}
	case Fa:
		req = tcRequirements{
			Atmo: "456789",
			Hydr: "45678",
			Pops: "23456",
			MW:   "0",
		}
	case Mi:
		req = tcRequirements{
			Pops: "23456", //NOT BY RULES
			Govr: "1",     //NOT BY RULES
			MW:   "0",
		}
	case Mr: //use world.UpdateInContext(MainWorld World)
		req = tcRequirements{
			Port: "L",
		}
	case Pe:
		req = tcRequirements{
			Atmo: "23AB",
			Hydr: "12345",
			Pops: "3456",
			Govr: "6",
			Laws: "6789",
			MW:   "0",
		}
	case Re:
		req = tcRequirements{
			Pops: "1234",
			Govr: "6",
			Laws: "45",
		}
	case Cp:
		req = tcRequirements{
			SubsectorCapital: true,
		}
	case Cs:
		req = tcRequirements{
			SubsectorCapital: true,
		}
	case Cx:
		req = tcRequirements{
			Capital: true,
		}
	case Cy:
		req = tcRequirements{
			Colony: true,
		}
	case Sa:
		req = tcRequirements{
			SateliteOrbit: "N0PQRSTUVWXYZ",
		}
	case Fo:
		req = tcRequirements{
			RedZone: true,
		}
	case Pz:
		req = tcRequirements{
			Pops:      "789ABCDEF",
			AmberZone: true,
		}
	case Da:
		req = tcRequirements{
			Pops:      "0123456",
			AmberZone: true,
		}
	case Ab:
		req = tcRequirements{
			DataRepository: true,
		}
	case An:
		req = tcRequirements{
			AntientSite: true,
		}
	case Ts:
		req = tcRequirements{
			Size:  "23456789ABCDEFGHJK",
			Atmo:  "01",
			HZvar: "9AB",
		}
	case Bo:
		req = tcRequirements{
			Size:  "23456789ABCDEFGHJK",
			HZvar: "012345678",
		}
	}
	return req
}
