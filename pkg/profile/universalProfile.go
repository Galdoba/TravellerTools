package profile

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	PROFILE_PERSONALITY     = "Personality"
	PROFILE_WORLD           = "World"
	Planetary_Physical_Data = "Physical Data"
)

// "HZvar",    //h [0-F] (... Bo Ho Tr [8] Tu Co Fr... F=Ds)
// 	"SatCode",  // [0-2] h" Planet, Close Sat, Far Sat
// 	"SatOrbit", // [0-F] h" постфикс латинской буквы
// 	"Climate",  // [0-6] h"
// 	"Hydr",
// 	"LifeFactor",        // [0-A] h"
// 	"LifeCompatability", // [0-A] h (количество совпадений в геноме рандомного софонта с человеком минус 2)"

// 	"Pops", // [0-F]"
// 	"Govr", // [0-F]"
// 	"Laws", // [0-J]"
// 	"Tech", // [0-N] s"

// 	"PopDigit",   // [0-9] h"
// 	"OtherBelts", // [0-3] h"
// 	"GasG [0-5]", // h"

type universalProfile struct {
	data    map[string]ehex.Ehex
	comment string
}

type Profile interface {
	Data(string) ehex.Ehex
	Inject(string, interface{})
	Map() map[string]ehex.Ehex
	//Create(string)            //C
	//Read(string) ehex.Ehex    //R
	//Update(string, ehex.Ehex) //U
	Delete(string) //D
}

func (up *universalProfile) Data(k string) ehex.Ehex {
	if v, ok := up.data[k]; ok {
		return v
	}
	return nil
}

func (up *universalProfile) Map() map[string]ehex.Ehex {
	// prMap := make(map[string]ehex.Ehex)
	// for k, v := range up.data {
	// 	prMap[k] = v
	// }
	// return prMap
	return up.data
}

func (up *universalProfile) Delete(s string) {
	delete(up.data, s)
}

// func (up *universalProfile) Format(f int) string {
// 	switch f {
// 	default:
// 		return "Format func Not Implemented"
// 	}
// }

func (up *universalProfile) Inject(k string, data interface{}) {
	switch data.(type) {
	default:
		up.data[k] = ehex.New().Set(data)
	case ehex.Ehex:
		up.data[k] = data.(ehex.Ehex)
	}
	if _, ok := up.data[k]; ok != true {
		panic(fmt.Errorf("injection failed [%v:%v]", k, data))
	}
}

// func (up *universalProfile) InjectAll(profile string) error {
// 	data := strings.Split(profile, "")
// 	//	input := len(strings.ReplaceAll(profile, "-", ""))

// 	separatorMod := 0
// 	for i, val := range data {
// 		pos := i - separatorMod
// 		if val == "-" {
// 			separatorMod++
// 			continue
// 		}
// 		dp := up.data[pos]
// 		up.data[pos] = DataPoint{dp.Key, ehex.New().Set(val), dp.Hidden, dp.Separated}
// 	}
// 	return nil
// }

func (up *universalProfile) String() string {
	str := fmt.Sprintf("profile contains %v points of data:\n", len(up.data))
	for k, v := range up.data {

		str += k + ": '" + v.Code() + "'\n"
	}
	return str
}

func New() *universalProfile {
	up := universalProfile{}
	up.data = make(map[string]ehex.Ehex)
	up.comment = "universal"
	return &up
}

func Merge(oldPr, newPr Profile) Profile {
	outPr := New()
	for k, v := range oldPr.Map() {
		outPr.data[k] = v
	}
	for k, v := range newPr.Map() {
		if _, ok := outPr.data[k]; !ok {
			outPr.data[k] = v
		}
	}
	return outPr
}

func Append(oldPr, newPr Profile) Profile {
	outPr := New()
	for k, v := range oldPr.Map() {
		outPr.data[k] = v
	}
	for k, v := range newPr.Map() {
		if _, ok := outPr.data[k]; !ok {
			outPr.data[k] = v
		}
	}
	return outPr
}

func (pr *universalProfile) Update(newPr Profile) {
	for k, v := range newPr.Map() {
		pr.data[k] = v
	}
}

func (pr *universalProfile) UpdateSafe(newPr Profile) {
	for k, v := range newPr.Map() {
		if _, ok := pr.data[k]; !ok {
			pr.data[k] = v
		}
	}
}

type prequisite struct {
	body       string
	definition string
	val        int
	more       bool
	less       bool
}

func (pr *universalProfile) PrequisiteMet(preq string) bool {
	//func preqIsMet(preq string, assets ...Asset) bool {
	if preq == "" {
		return true
	}
	p := prequisite{}
	splitter := " "
	switch {
	case strings.Contains(preq, "-"):
		preq = strings.TrimSuffix(preq, "-")
		p.less = true
	case strings.Contains(preq, "+"):
		preq = strings.TrimSuffix(preq, "+")
		p.more = true
	case strings.Contains(preq, "="):
		splitter = "="
	}
	prArr := strings.Split(preq, splitter)
	p.body = prArr[0]
	if len(prArr) == 1 {
		prArr = append(prArr, prArr[0])
	}
	p.definition = prArr[1]
	p.val, _ = strconv.Atoi(p.definition)
	for k, v := range pr.data {
		if p.body != k {
			continue
		}
		val := v.Value()
		if p.less {
			if val <= p.val {
				return true
			}
		}
		if p.more {
			if val >= p.val {
				return true
			}
		}
		if val == p.val {
			return true
		}
		return false
	}
	return false
}

// func validKeys(e string) []string {
// 	switch e {
// 	default:
// 		return []string{}
// 	case PROFILE_PERSONALITY:
// 		return []string{KEY_C1, KEY_C2, KEY_C3, KEY_C4, KEY_C5, KEY_C6, KEY_CP, KEY_CS}
// 	case PROFILE_WORLD:
// 		return []string{KEY_PORT, KEY_SIZE, KEY_ATMO, KEY_HYDR, KEY_POPS, KEY_GOVR, KEY_LAWS, KEY_TL}
// 	}
// }
