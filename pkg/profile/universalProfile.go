package profile

import (
	"fmt"

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
	//Profile() []DataPoint
	Data(string) ehex.Ehex
	//Format(int) string
	Inject(string, interface{})
}

func (up *universalProfile) Data(k string) ehex.Ehex {
	if v, ok := up.data[k]; ok {
		return v
	}
	return nil
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
	return &up
}

func validKeys(e string) []string {
	switch e {
	default:
		return []string{}
	case PROFILE_PERSONALITY:
		return []string{KEY_C1, KEY_C2, KEY_C3, KEY_C4, KEY_C5, KEY_C6, KEY_CP, KEY_CS}
	case PROFILE_WORLD:
		return []string{KEY_PORT, KEY_SIZE, KEY_ATMO, KEY_HYDR, KEY_POPS, KEY_GOVR, KEY_LAWS, KEY_TL}
	}
}
