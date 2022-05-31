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
		{"Htalrea", "E767610-1", "103"},
		{"Ea", "C7586AA-7", "214"},
		{"Oihoiei", "A8558A8-C", "214"},
		{"Carrill", "A0009AE-E", "613"},
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
		if ssf.basePersonal < 1 {
			t.Errorf("base personal not defined")
		}
		if ssf.organisation == "" {
			t.Errorf("organisation not defined")
		}
		if ssf.funding == "" {
			t.Errorf("funding not defined")
		}
		if ssf.equipment == "" {
			t.Errorf("equipment not defined")
		}
		if ssf.competence == "" {
			t.Errorf("competence not defined")
		}
		if ssf.response == "" {
			t.Errorf("response not defined")
		}
		if ssf.checksDM == -5 {
			t.Errorf("checksDM not defined")
		}
		if ssf.fiascoTN == -1 {
			t.Errorf("fiascoTN not defined")
		}
	}
}
