package tradecodes

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

func FromUWPstr(uwpCode string) ([]string, error) {
	uwp, err := uwp.FromString(uwpCode)
	if err != nil {
		return []string{}, err
	}
	return parceCodes(uwp)
}

func FromUWP(prof uwp.UWP) ([]string, error) {
	return parceCodes(prof)
}

func extractCodes(prof uwp.UWP) ([]string, error) {

	return nil, nil
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
