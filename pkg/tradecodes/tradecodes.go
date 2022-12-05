package tradecodes

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	KEY_uwp         = "uwp"
	KEY_hz          = "hz"
	KEY_mw          = "mw"
	KEY_sw          = "sw"
	KEY_travel      = "travel"
	KEY_sector      = "sector"
	KEY_manual      = "manual"
	uwpApproval     = 0
	hzApproval      = 1
	mwApproval      = 2
	swApproval      = 3
	travellApproval = 4
	sectorApproval  = 5
	manualApproval  = 6
)

type classifications struct {
	confirmedByUWP []string

	excludeMWonly bool
	excludeSWonly bool
	errors        []error
}

type tcInput struct {
	key, val string
}

func Input(key, val string) tcInput {
	return tcInput{key, val}
}

type World interface {
	UWP() string
	Stellar() string
	HZ() int
}

func Analize(input ...tcInput) *classifications {
	cls := classifications{}
	for _, data := range input {
		switch data.key {
		case KEY_uwp:
			u, err := uwp.FromString(data.val)
			if err != nil {
				cls.errors = append(cls.errors, err)
				continue
			}
			parcedTC, err := parceCodesFull(u)
			if err != nil {
				cls.errors = append(cls.errors, err)
				continue
			}
			cls.confirmedByUWP = append(cls.confirmedByUWP, parcedTC...)

		}
	}
	return &cls
}

func (cls *classifications) excludeMW() {

}

type clssfctn struct {
	code             string
	approvalExpected []int
	//uwpApproval      int
	//hzApproval       int
	//mwApproval       int
	//swApproval       int
	//travellApproval  int
	//sectorApproval   int
	//manualApproval   int
}

func mapCLS() map[int]clssfctn {
	clMap := make(map[int]clssfctn)
	clMap[0] = clssfctn{"As", []int{1, 0, 1, 0, 0, 0, 0}}
	clMap[1] = clssfctn{"De", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[2] = clssfctn{"Fl", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[3] = clssfctn{"Ga", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[4] = clssfctn{"He", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[5] = clssfctn{"Ic", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[6] = clssfctn{"Oc", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[7] = clssfctn{"Va", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[8] = clssfctn{"Wa", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[9] = clssfctn{"Di", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[10] = clssfctn{"Ba", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[11] = clssfctn{"Lo", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[12] = clssfctn{"Ni", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[13] = clssfctn{"Ph", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[14] = clssfctn{"Hi", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[15] = clssfctn{"Pa", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[16] = clssfctn{"Ag", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[17] = clssfctn{"Na", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[18] = clssfctn{"Px", []int{1, 0, 1, 0, 0, 0, 0}}
	clMap[19] = clssfctn{"In", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[20] = clssfctn{"Po", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[21] = clssfctn{"Pr", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[22] = clssfctn{"Ri", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[23] = clssfctn{"Lt", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[24] = clssfctn{"Ht", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[25] = clssfctn{"Fr", []int{1, 1, 0, 0, 0, 0, 0}}
	clMap[26] = clssfctn{"Ho", []int{0, 1, 0, 0, 0, 0, 0}}
	clMap[27] = clssfctn{"Co", []int{0, 1, 0, 0, 0, 0, 0}}
	clMap[28] = clssfctn{"Lk", []int{0, 0, 0, 0, 0, 0, 1}}
	clMap[29] = clssfctn{"Tr", []int{1, 1, 0, 0, 0, 0, 0}}
	clMap[30] = clssfctn{"Tu", []int{1, 1, 0, 0, 0, 0, 0}}
	clMap[31] = clssfctn{"Tz", []int{0, 0, 0, 0, 0, 0, 1}}
	clMap[32] = clssfctn{"Fa", []int{1, 0, 0, 1, 0, 0, 0}}
	clMap[33] = clssfctn{"Mn", []int{1, 0, 0, 1, 0, 1, 0}}
	clMap[34] = clssfctn{"Mr", []int{1, 0, 0, 1, 0, 0, 0}}
	clMap[35] = clssfctn{"Pe", []int{1, 0, 0, 1, 0, 0, 0}}
	clMap[36] = clssfctn{"Re", []int{1, 0, 0, 0, 0, 0, 0}}
	clMap[37] = clssfctn{"Cp", []int{0, 0, 1, 0, 0, 1, 0}}
	clMap[38] = clssfctn{"Cs", []int{0, 0, 1, 0, 0, 1, 0}}
	clMap[39] = clssfctn{"Cx", []int{0, 0, 1, 0, 0, 1, 0}}
	clMap[40] = clssfctn{"Cy", []int{1, 0, 1, 0, 0, 1, 0}}
	clMap[41] = clssfctn{"Sa", []int{0, 0, 0, 0, 0, 0, 1}}
	clMap[42] = clssfctn{"Fo", []int{0, 0, 0, 0, 0, 1, 0}}
	clMap[43] = clssfctn{"Pz", []int{1, 0, 0, 0, 0, 1, 0}}
	clMap[44] = clssfctn{"Da", []int{1, 0, 0, 0, 0, 1, 0}}
	clMap[45] = clssfctn{"Ab", []int{0, 0, 0, 0, 0, 0, 1}}
	clMap[46] = clssfctn{"An", []int{0, 0, 0, 0, 0, 0, 1}}
	return clMap
}

func fullTCsList() []string {
	return []string{
		"As", "De", "Fl", "Ga", "He", "Ic", "Oc", "Va", "Wa",
		"Di", "Ba", "Lo", "Ni", "Ph", "Hi",
		"Pa", "Ag", "Na", "Px", "Pi", "In", "Po", "Pr", "Ri", "Lt", "Ht",
		"Fr", "Ho", "Co", "Lk", "Tr", "Tu", "Tz",
		"Fa", "Mi", "Mr", "Pe", "Re",
		"Cp", "Cs", "Cx", "Cy",
		"Sa", "Fo", "Pz", "Da", "Ab", "An",
	}
}

func FromUWP(prof uwp.UWP) ([]string, error) {
	return parceCodes(prof)
}

func parceCodes(data uwp.UWP) ([]string, error) {
	parsedTradeCodes := []string{}
	//S S      A     H   P  G  L  - T
	codes := tradeCodeDemands()
	for codeNum, tg := range codes {
		//fmt.Println("GO ", codeNum, tg)
		tcMatch := true
		values := strings.Split(tg, " ")
		for valNum, d := range values {
			if tcMatch == false {
				//	fmt.Println("BREAK")
				break
			}
			if d == "--" || d == "-" {
				//fmt.Println("DEBUG: skip", valNum, d, tcMatch)
				continue
			}
			//fmt.Println("DEBUG: test", valNum, d, tcMatch)
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
				parsedTradeCodes = append(parsedTradeCodes, "Ag")
			case 1:
				parsedTradeCodes = append(parsedTradeCodes, "As")
			case 2:
				parsedTradeCodes = append(parsedTradeCodes, "Ba")
			case 3:
				parsedTradeCodes = append(parsedTradeCodes, "De")
			case 4:
				parsedTradeCodes = append(parsedTradeCodes, "Fl")
			case 5:
				parsedTradeCodes = append(parsedTradeCodes, "Ga")
			case 6:
				parsedTradeCodes = append(parsedTradeCodes, "Hi")
			case 7:
				parsedTradeCodes = append(parsedTradeCodes, "Ht")
			case 8:
				parsedTradeCodes = append(parsedTradeCodes, "Ic")
			case 9:
				parsedTradeCodes = append(parsedTradeCodes, "In")
			case 10:
				parsedTradeCodes = append(parsedTradeCodes, "Lo")
			case 11:
				parsedTradeCodes = append(parsedTradeCodes, "Lt")
			case 12:
				parsedTradeCodes = append(parsedTradeCodes, "Na")
			case 13:
				parsedTradeCodes = append(parsedTradeCodes, "Ni")
			case 14:
				parsedTradeCodes = append(parsedTradeCodes, "Po")
			case 15:
				parsedTradeCodes = append(parsedTradeCodes, "Ri")
			case 16:
				parsedTradeCodes = append(parsedTradeCodes, "Va")
			case 17:
				parsedTradeCodes = append(parsedTradeCodes, "Wa")
			}
		}
		//fmt.Println(codeNum, "-------------------")

	}
	return parsedTradeCodes, nil
}

func parceCodesFull(data uwp.UWP) ([]string, error) {
	parsedTradeCodes := []string{}
	//S S      A     H   P  G  L  - T
	codes := tradeCodeDemandsFull()
	for codeNum, tg := range codes {
		//fmt.Println("GO ", codeNum, tg)
		tcMatch := true
		values := strings.Split(tg, " ")
		for valNum, d := range values {
			if tcMatch == false {
				//	fmt.Println("BREAK")
				break
			}
			if d == "--" || d == "-" {
				//fmt.Println("DEBUG: skip", valNum, d, tcMatch)
				continue
			}
			//fmt.Println("DEBUG: test", valNum, d, tcMatch)
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
				parsedTradeCodes = append(parsedTradeCodes, "As")
			case 1:
				parsedTradeCodes = append(parsedTradeCodes, "De")
			case 2:
				parsedTradeCodes = append(parsedTradeCodes, "Fl")
			case 3:
				parsedTradeCodes = append(parsedTradeCodes, "Ga")
			case 4:
				parsedTradeCodes = append(parsedTradeCodes, "He")
			case 5:
				parsedTradeCodes = append(parsedTradeCodes, "Ic")
			case 6:
				parsedTradeCodes = append(parsedTradeCodes, "Oc")
			case 7:
				parsedTradeCodes = append(parsedTradeCodes, "Va")
			case 8:
				parsedTradeCodes = append(parsedTradeCodes, "Wa")
			case 9:
				parsedTradeCodes = append(parsedTradeCodes, "Di")
			case 10:
				parsedTradeCodes = append(parsedTradeCodes, "Ba")
			case 11:
				parsedTradeCodes = append(parsedTradeCodes, "Lo")
			case 12:
				parsedTradeCodes = append(parsedTradeCodes, "Ni")
			case 13:
				parsedTradeCodes = append(parsedTradeCodes, "Ph")
			case 14:
				parsedTradeCodes = append(parsedTradeCodes, "Hi")
			case 15:
				parsedTradeCodes = append(parsedTradeCodes, "Pa")
			case 16:
				parsedTradeCodes = append(parsedTradeCodes, "Ag")
			case 17:
				parsedTradeCodes = append(parsedTradeCodes, "Na")
			case 18:
				parsedTradeCodes = append(parsedTradeCodes, "Px")
			case 19:
				parsedTradeCodes = append(parsedTradeCodes, "In")
			case 20:
				parsedTradeCodes = append(parsedTradeCodes, "Po")
			case 21:
				parsedTradeCodes = append(parsedTradeCodes, "Pr")
			case 22:
				parsedTradeCodes = append(parsedTradeCodes, "Ri")
			case 23:
				parsedTradeCodes = append(parsedTradeCodes, "Lt")
			case 24:
				parsedTradeCodes = append(parsedTradeCodes, "Ht")
			case 25:
				parsedTradeCodes = append(parsedTradeCodes, "Fr")
			case 26:
				parsedTradeCodes = append(parsedTradeCodes, "Ho")
			case 27:
				parsedTradeCodes = append(parsedTradeCodes, "Co")
			case 28:
				parsedTradeCodes = append(parsedTradeCodes, "Lk")
			case 29:
				parsedTradeCodes = append(parsedTradeCodes, "Tr")
			case 30:
				parsedTradeCodes = append(parsedTradeCodes, "Tu")
			case 31:
				parsedTradeCodes = append(parsedTradeCodes, "Tz")
			case 32:
				parsedTradeCodes = append(parsedTradeCodes, "Fa")
			case 33:
				parsedTradeCodes = append(parsedTradeCodes, "Mn")
			case 34:
				parsedTradeCodes = append(parsedTradeCodes, "Mr")
			case 35:
				parsedTradeCodes = append(parsedTradeCodes, "Pe")
			case 36:
				parsedTradeCodes = append(parsedTradeCodes, "Re")
			case 37:
				parsedTradeCodes = append(parsedTradeCodes, "Cp")
			case 38:
				parsedTradeCodes = append(parsedTradeCodes, "Cs")
			case 39:
				parsedTradeCodes = append(parsedTradeCodes, "Cx")
			case 40:
				parsedTradeCodes = append(parsedTradeCodes, "Cy")
			case 41:
				parsedTradeCodes = append(parsedTradeCodes, "Sa")
			case 42:
				parsedTradeCodes = append(parsedTradeCodes, "Fo")
			case 43:
				parsedTradeCodes = append(parsedTradeCodes, "Pz")
			case 44:
				parsedTradeCodes = append(parsedTradeCodes, "Da")
			case 45:
				parsedTradeCodes = append(parsedTradeCodes, "Ab")
			case 46:
				parsedTradeCodes = append(parsedTradeCodes, "An")
			}
		}
		//fmt.Println(codeNum, "-------------------")

	}
	return parsedTradeCodes, nil
}

func tradeCodeDemands() []string {
	return []string{
		"-- 456789 45678 567 -- -- -- - --",          //Ag
		"-- 0 0 0 -- -- -- - --",                     //As
		"-- -- -- -- 0 0 0 - --",                     //Ba
		"-- -- 23456789 0 -- -- -- - --",             //De
		"-- -- ABCDEF 123456789A -- -- -- - --",      //Fl
		"-- 678 568 567 -- -- -- - --",               //Ga
		"-- -- -- -- 9ABCDEF -- -- - --",             //Hi
		"-- -- -- -- -- -- -- - CDEFGH",              //Ht
		"-- -- 01 123456789A -- -- -- - --",          //Ic
		"-- -- 012479ABC -- 9ABCDEF -- -- - --",      //In
		"-- -- -- -- 123 -- -- - --",                 //Lo
		"-- -- -- -- 123456789ABCDEF -- -- - 012345", //Lt
		"-- -- 0123 0123 6789ABCDEF -- -- - --",      //Na
		"-- -- -- -- 456 -- -- - --",                 //Ni
		"-- -- 2345 0123 -- -- -- - --",              //Po
		"-- -- 68 -- 678 456789 -- - --",             //Ri
		"-- -- 0 -- -- -- -- - --",                   //Va
		"-- -- 3456789DEF A -- -- -- - --",           //Wa

	}
}

func tradeCodeDemandsFull() []string {
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
		"EXHY -- -- -- 0 0 0 - --",                              //Ba
		"-- -- -- -- 123 -- -- - --",                            //Lo
		"-- -- -- -- 456 -- -- - --",                            //Ni
		"-- -- -- -- 8 -- -- - --",                              //Ph
		"-- -- -- -- 9ABCDEF -- -- - --",                        //Hi
		"-- -- 456789 45678 48 -- -- - --",                      //Pa
		"-- -- 456789 45678 567 -- -- - --",                     //Ag
		"-- -- 0123 0123 6789ABCDEF -- -- - --",                 //Na
		"-- -- 23AB 12345 3456 -- 6789ABCDEFGHJ - --",           //Px
		"-- -- 012479ABC -- 9ABCDEF -- -- - --",                 //In
		"-- -- 2345 0123 -- -- -- - --",                         //Po
		"-- -- 68 -- 59 -- -- - --",                             //Pr
		"-- -- 68 -- 678 -- -- - --",                            //Ri
		"-- -- -- -- -- -- -- - 12345",                          //Lt
		"-- -- -- -- -- -- -- - CDEFGHJKLMNPQRSTUVWXYZ",         //Ht
		"-- 23456789 -- 123456789A -- -- -- - --",               //Fr
		"-- -- -- -- -- -- -- - --",                             //Ho
		"-- -- -- -- -- -- -- - --",                             //Co
		"-- -- -- -- -- -- -- - --",                             //Lk
		"-- 6789 456789 34567 -- -- -- - --",                    //Tr
		"-- 6789 456789 34567 -- -- -- - --",                    //Tu
		"-- -- -- -- -- -- -- - --",                             //Tz
		"-- -- 456789 45678 23456 -- -- - --",                   //Fa
		"-- -- -- -- 23456 -- -- - --",                          //Mn
		"-- -- -- -- -- -- -- - --",                             //Mr
		"-- -- 23AB 12345 3456 6 6789ABCDEFGHJ - --",            //Pe
		"-- -- -- -- 1234 6 45 - --",                            //Re
		"A -- -- -- -- -- -- - --",                              //Cp
		"A -- -- -- -- -- -- - --",                              //Cs
		"A -- -- -- -- -- -- - --",                              //Cx
		"-- -- -- -- 56789A 6 0123 - --",                        //Cy
		"-- -- -- -- -- -- -- - --",                             //Sa
		"-- -- -- -- -- -- -- - --",                             //Fo
		"-- -- -- -- 789ABCDEF -- -- - --",                      //Pz
		"-- -- -- -- 0123456 -- -- - --",                        //Da
		"-- -- -- -- -- -- -- - --",                             //Ab
		"-- -- -- -- -- -- -- - --",                             //An
	}
}
