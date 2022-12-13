package classifications

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
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
	Mn
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
	TypePlanetary  = "Planetary"
	TypePopulation = "Population"
	TypeEconomic   = "Economic"
	TypeClimate    = "Climate"
	TypeSecondary  = "Secondary"
	TypePolitical  = "Political"
	TypeSpecial    = "Special"
)

type classificationCode struct {
	val            int
	code           string
	classification string
	tcType         string
	description    string
	sourceBook     string
}

func Manual(code, classification, tcType, description, sourceBook string) classificationCode {
	return classificationCode{manual, code, classification, tcType, description, sourceBook}
}

func Call(i int) classificationCode {
	tc := classificationCode{}
	tc.val = i
	switch tc.val {
	default:
		tc.code = "??"
		tc.classification = "[Unknown]"
		tc.tcType = TypeSpecial
		tc.description = "[No Description]"
		return tc
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
		tc.classification = "Ice-Capped"
		tc.tcType = TypePlanetary
		tc.description = "The world's water is locked in ice-caps."
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
	case Ht:
		tc.code = "Ht"
		tc.classification = "High Tech"
		tc.tcType = TypeEconomic
		tc.description = "The world is among the most technologicaly advanced in Charted Space."
	}
	return tc
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

/////////////////////////////

func FromUWP(prof uwp.UWP) ([]classificationCode, error) {
	return parceCodes(prof)
}

func parceCodes(data uwp.UWP) ([]classificationCode, error) {
	parsedTradeCodes := []classificationCode{}
	//S S      A     H   P  G  L  - T
	codes := tradeCodeDemandsPPE()
	for codeNum, tg := range codes {
		//fmt.Println("GO ", codeNum, tg)
		tcMatch := true
		values := strings.Split(tg, " ")
		for valNum, d := range values {
			if tcMatch == false {
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
			switch codeNum {
			default:
				return nil, fmt.Errorf("unknown trade code position from tradeCodeDemands()")
			case 0:
				parsedTradeCodes = append(parsedTradeCodes, Call(As))
			case 1:
				parsedTradeCodes = append(parsedTradeCodes, Call(De))
			case 2:
				parsedTradeCodes = append(parsedTradeCodes, Call(Fl))
			case 3:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ga))
			case 4:
				parsedTradeCodes = append(parsedTradeCodes, Call(He))
			case 5:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ic))
			case 6:
				parsedTradeCodes = append(parsedTradeCodes, Call(Oc))
			case 7:
				parsedTradeCodes = append(parsedTradeCodes, Call(Va))
			case 8:
				parsedTradeCodes = append(parsedTradeCodes, Call(Wa))
			case 9:
				parsedTradeCodes = append(parsedTradeCodes, Call(Di))
			case 10:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ba))
			case 11:
				parsedTradeCodes = append(parsedTradeCodes, Call(Lo))
			case 12:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ni))
			case 13:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ph))
			case 14:
				parsedTradeCodes = append(parsedTradeCodes, Call(Hi))
			case 15:
				parsedTradeCodes = append(parsedTradeCodes, Call(Pa))
			case 16:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ag))
			case 17:
				parsedTradeCodes = append(parsedTradeCodes, Call(Na))
			case 18:
				parsedTradeCodes = append(parsedTradeCodes, Call(Px))
			case 19:
				parsedTradeCodes = append(parsedTradeCodes, Call(Pi))
			case 20:
				parsedTradeCodes = append(parsedTradeCodes, Call(In))
			case 21:
				parsedTradeCodes = append(parsedTradeCodes, Call(Po))
			case 22:
				parsedTradeCodes = append(parsedTradeCodes, Call(Pr))
			case 23:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ri))
			case 24:
				parsedTradeCodes = append(parsedTradeCodes, Call(Lt))
			case 25:
				parsedTradeCodes = append(parsedTradeCodes, Call(Ht))

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

Co
Cold World
The world is at the lower temperature range of human endurance.

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

Fr
Frozen
The world lies substantially beyond the Habitable Zone of the system (HZ+2 or greater) and environmental temperatures are well below the freezing point of many gases.



Ho
Hot World
The world is at the upper temperature range of human endurance; typically in HZ -1.

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


Tu
Tundra
The world is relatively colder than normal (although it is considered habitable). Its orbit is at the outer (colder) edge of the Habitable Zone. The world has a Cold climate (at the lower limits of human temperature endurance).

Tz
Twilight Zone
The world is tidally locked with a Temperate band at the Twilight Zone, plus a Hot region (hemisphere) facing the Primary and a Cold region (hemisphere) away from the Primary.


*/
