package ssp

import (
	"fmt"
	"testing"
)

type input struct {
	name  string
	uwpS  string
	tz    string
	bases string
	tc    string
	pbg   string
}

func (i *input) MW_Name() string {
	return i.name
}
func (i *input) MW_UWP() string {
	return i.uwpS
}
func (i *input) TravelZone() string {
	return i.tz
}
func (i *input) Bases() string {
	return i.bases
}
func (i *input) MW_Remarks() string {
	return i.tc
}
func (i *input) PBG() string {
	return i.pbg
}

func inputList() []input {
	return []input{
		{"", "", "", "", "", ""},
		{"Sarrad", "D88A300-8", "A", "", "Lo Wa Da", "000"},
		{"Earlo", "D542102-7", "A", "", "He Lo Po Da Asla7", "214"},
	}
}

func Test(t *testing.T) {
	for i, planet := range inputList() {
		sp, err := NewSecurityProfile(&planet)
		fmt.Println(i, planet, "==>", sp, err)
		if sp == nil {
			t.Errorf("func returned no object")
			continue
		}
		if err != nil {
			t.Errorf("func returned error: %v", err.Error())
		}
		if sp.name == "UNSET" {
			t.Errorf("sp.name is not addressed ")
		}
		if sp.wsp == "UNSET" {
			t.Errorf("sp.wsp is not addressed ")
		}
		if sp.value == nil {
			t.Errorf("sp.value is not addressed ")
		}
		for i, presence := range []int{presense_planetary, presense_orbital, presense_system, stance} {
			if _, ok := sp.value[presence]; !ok {
				n := ""
				switch i {
				case presense_planetary:
					n = "presense_planetary"
				case presense_orbital:
					n = "presense_orbital"
				case presense_system:
					n = "presense_system"
				case stance:
					n = "stance"
				}
				t.Errorf("sp.value[%v] is not addressed", n)
			}
		}
	}
}
