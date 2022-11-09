package starport

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/starport/ssp"
)

type World interface {
	MW_Name() string
	MW_UWP() string
	Hex() string
	TravelZone() string
	PBG() string
	Bases() string
	Stellar() string
	Sector() string
	MW_Remarks() string
}

type Starport struct {
	MW_Name    string
	MW_UWP     string
	Sector     string
	Hex        string
	PBG        string
	Stellar    string
	TradeCodes []string
	Remarks    []string

	FacilityPresent map[int]bool
	SecurityProfile ssp.SecurityProfile
}

func New(w World) (*Starport, error) {
	st := Starport{}
	st.MW_Name = w.MW_Name()
	st.MW_UWP = w.MW_UWP()
	st.Hex = w.Hex()
	st.Sector = w.Sector()
	st.Stellar = w.Stellar()
	st.PBG = w.PBG()
	tcT5 := strings.Fields(w.MW_Remarks())
	tcmgt, err := tradecodes.FromUWPstr(w.MW_UWP())
	if err != nil {
		return &st, err
	}
	st.TradeCodes = tcmgt
	for _, rem := range tcT5 {
		if !contained(st.TradeCodes, rem) {
			st.Remarks = append(st.Remarks, rem)
		}
	}
	secProf, err := ssp.NewSecurityProfile(w)
	if err != nil {
		return &st, err
	}
	st.SecurityProfile = secProf
	return &st, nil
}

func contained(sl []string, elem string) bool {
	for _, e := range sl {
		if e == elem {
			return true
		}
	}
	return false
}

func (st *Starport) String() string {
	str := fmt.Sprintf("Name: %v (%v %v)\n", st.MW_Name, st.Sector, st.Hex)
	str += fmt.Sprintf("UWP : %v\n", st.MW_UWP)
	str += fmt.Sprintf("Stellar    : %v      %v\n", st.Stellar, st.PBG)
	str += fmt.Sprintf("Trade Codes: %v\n", st.TradeCodes)
	if len(st.Remarks) > 0 {
		str += fmt.Sprintf("Remarks    : %v\n", st.Remarks)
	}

	str += "--------------------------------------------------------------------------------\n"
	if st.SecurityProfile != nil {
		str += st.SecurityProfile.Describe()
	}
	return str
}

/*
Name: Hope (Reaver's Deep 2526)
UWP : E65778B-4
Stellar    : F4 V        200
Trade Codes: Ag Ga
Remarks    : Pz
Coordinates: (1,-1,0)
-------------------------
NIL: none
Population: 70000+
-------------------------

*/
