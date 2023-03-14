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
	Starport = "Starport"
	Size     = "Size"
	Atmo     = "Atmo"
	Hydr     = "Hydr"
	Pops     = "Pops"
	Govr     = "Govr"
	Laws     = "Laws"
	sep      = "separator"
	Tech     = "Tech"
)

type universalProfile struct {
	data       map[string]string
	keys       []string
	hiddenData map[string]bool
}

func New(entity string) (*universalProfile, error) {
	up := universalProfile{}
	up.data = make(map[string]string)
	up.hiddenData = make(map[string]bool)
	switch entity {
	default:
		return nil, fmt.Errorf("unknown entity type '%v'", entity)
	case PROFILE_WORLD:
		for _, val := range expectedData(entity) {
			up.data[val] = "?"
			up.keys = append(up.keys, val)
		}
		for _, val := range hiddenData(entity) {
			up.hiddenData[val] = true
		}
	}
	return &up, nil
}

func (up *universalProfile) Inject(values string) error {
	val := strings.Split(values, "")
	if len(val) != len(up.keys) {
		fmt.Println(val)
		fmt.Println(up.keys)
		return fmt.Errorf("values does not match keys")
	}
	v := strings.Split(values, "")
	for n, key := range up.keys {
		up.data[key] = v[n]
	}
	return nil
}

func (up *universalProfile) String() string {
	s := ""
	for _, key := range up.keys {
		if up.hiddenData[key] == true {
			continue
		}
		if key == sep {
			s += "-"
			continue
		}
		s += ehex.New().Set(up.data[key]).Code()
	}
	return s
}

func expectedData(e string) []string {
	switch e {
	default:
		return []string{}
	case PROFILE_WORLD:
		return []string{
			Starport,
			Size,
			Atmo,
			Hydr,
			Pops,
			Govr,
			Laws,
			sep,
			Tech,
		}
	}
}

func hiddenData(e string) []string {
	switch e {
	default:
		return []string{}
	case PROFILE_WORLD:
		return []string{}
	}
}
