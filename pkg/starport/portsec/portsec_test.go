package portsec

import (
	"fmt"
	"testing"
)

type input struct {
	seed string
	uwp  string
	pbg  string
}

func (i *input) MW_Name() string {
	return i.seed
}

func (i *input) MW_UWP() string {
	return i.uwp
}

func (i *input) PBG() string {
	return i.pbg
}

func setInput() []input {
	return []input{
		{"seed", "uwp", ""},
		{"seed", "C555555-5", "555"},
		{"Oihoiei", "A8558A8-C", "214"},
	}
}

func Test_Portsec(t *testing.T) {
	for _, input := range setInput() {
		fmt.Println("___________________________________________________________")
		fmt.Println("Start test:", input)
		fmt.Println(" ")
		ssf, err := GenerateSecurityForces(&input)
		if ssf == nil {
			t.Errorf("func returned no struct")
			continue
		}
		if err != nil {
			t.Errorf("func returned error: %v", err.Error())
		}
		fmt.Println(ssf)
	}
}
