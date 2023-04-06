package profile2

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	PROFILE_PERSONALITY     = "Personality"
	PROFILE_WORLD           = "World"
	Planetary_Physical_Data = "Physical Data"
)

const (
	KEY_C1   = "C1"
	KEY_C2   = "C2"
	KEY_C3   = "C3"
	KEY_C4   = "C4"
	KEY_C5   = "C5"
	KEY_C6   = "C6"
	KEY_CP   = "CP"
	KEY_CS   = "CS"
	KEY_PORT = "Starport"
	KEY_SIZE = "Size"
	KEY_ATMO = "Atmo"
	KEY_HYDR = "Hydr"
	KEY_POPS = "Pops"
	KEY_GOVR = "Govr"
	KEY_LAWS = "Laws"
	KEY_SEP  = "separator"
	KEY_TL   = "Tech"
)

type universalProfile struct {
	data    []dataPoint
	comment string
}

func (up *universalProfile) Data(k string) ehex.Ehex {
	for _, dp := range up.data {
		if dp.key == k {
			return dp.val
		}
	}
	return nil
}

func (up *universalProfile) Inject(k string, data interface{}) error {
	pos := -1
	hidden := false
	separated := false
	for i, dp := range up.data {
		if dp.key == k {
			if pos != -1 {
				return fmt.Errorf("conflicting data found:\n%v\n%v", up.data[pos], up.data[i])
			}
			pos = i
			hidden = dp.hidden
			separated = dp.separated
		}
	}
	if pos == -1 {
		return fmt.Errorf("no data with key '%v'", k)
	}
	up.data[pos] = dataPoint{k, ehex.New().Set(data), hidden, separated}
	return nil
}

func (up *universalProfile) InjectAll(profile string) error {
	data := strings.Split(profile, "")
	//	input := len(strings.ReplaceAll(profile, "-", ""))

	separatorMod := 0
	for i, val := range data {
		pos := i - separatorMod
		if val == "-" {
			separatorMod++
			continue
		}
		dp := up.data[pos]
		up.data[pos] = dataPoint{dp.key, ehex.New().Set(val), dp.hidden, dp.separated}
	}
	return nil
}

func (up *universalProfile) String() string {
	str := ""
	for _, dp := range up.data {
		if dp.hidden {
			continue
		}
		if dp.separated {
			str += "-"
		}
		str += dp.val.Code()
	}
	return str
}

func (up *universalProfile) StringFull() string {
	str := ""
	for _, dp := range up.data {
		if dp.separated {
			str += "-"
		}
		str += dp.val.Code()
	}
	return str
}

type dataPoint struct {
	key       string
	val       ehex.Ehex
	hidden    bool
	separated bool
}

func newDataPoint(k string) dataPoint {
	switch k {
	case KEY_C1, KEY_C2, KEY_C3, KEY_C4, KEY_C5, KEY_C6:
		return dataPoint{k, ehex.New().Set(0), false, false}
	case KEY_CP, KEY_CS:
		return dataPoint{k, ehex.New().Set(0), true, false}
	case KEY_PORT, KEY_SIZE, KEY_ATMO, KEY_HYDR, KEY_POPS, KEY_GOVR, KEY_LAWS:
		return dataPoint{k, ehex.New().Set(0), false, false}
	case KEY_TL:
		return dataPoint{k, ehex.New().Set(0), false, true}
	}
	return dataPoint{"unknown", ehex.New().Set("?"), false, false}
}

func New(entity string) *universalProfile {
	keys := validKeys(entity)
	if len(keys) == 0 {
		return nil
	}
	up := universalProfile{}
	up.comment = entity
	for _, k := range keys {
		up.data = append(up.data, newDataPoint(k))
	}
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
