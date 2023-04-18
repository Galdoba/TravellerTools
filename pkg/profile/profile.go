package profile

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

type profileData struct {
	Data        map[string]ehex.Ehex
	profileType string
}

type ProfileEncoder interface {
	Decode(string) (string, int, string, error)
	Encode(string, ehex.Ehex) error
}

// New - Creates new abstarct profile
func NewPD(dataType ...string) *profileData {
	pd := profileData{}
	pd.Data = make(map[string]ehex.Ehex)
	for _, data := range dataType {
		pd.Data[data] = ehex.New()
	}
	return &pd
}

// Decode - return all info contained in profile data
func (pd *profileData) Decode(data string) (string, int, string, error) {
	if _, ok := pd.Data[data]; !ok {
		return "", 0, "", fmt.Errorf("data '%v' not present in profile", data)
	}
	code := pd.Data[data].Code()
	val := pd.Data[data].Value()
	comment := pd.Data[data].Meaning()
	return code, val, comment, nil
}

func (pd *profileData) Encode(aspect string, eh ehex.Ehex) error {
	return fmt.Errorf("this is abstract function and should not be used")
}

func (pd *profileData) EncodeString(uwpString string) error {
	return fmt.Errorf("this is abstract function and should not be used")
}
