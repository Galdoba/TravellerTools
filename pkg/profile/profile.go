package profile

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/ehex"
)

type ProfileData struct {
	Data map[string]ehex.Ehex
}

//New - Creates new abstarct profile
func New(dataType ...string) *ProfileData {
	pd := ProfileData{}
	pd.Data = make(map[string]ehex.Ehex)
	for _, data := range dataType {
		pd.Data[data] = ehex.New()
	}
	return &pd
}

//Decode - return all info contained in profile data
func (pd *ProfileData) Decode(data string) (string, int, string, error) {
	if _, ok := pd.Data[data]; !ok {
		return "", 0, "", fmt.Errorf("data '%v' not present in profile", data)
	}
	code := pd.Data[data].Code()
	val := pd.Data[data].Value()
	comment := pd.Data[data].Meaning()
	return code, val, comment, nil
}
