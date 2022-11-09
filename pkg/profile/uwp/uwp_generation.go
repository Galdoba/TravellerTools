package uwp

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	GenerationRules_MGT2_MW_CG  = "MgT2 Mainworld Sector Construction Guide"
	GenerationRules_MGT2_MW_CRB = "MgT2 Mainworld Core"
)

func (u *uwp) Generate(dp *dice.Dicepool, generationType string) error {
	//TODO: менеджер генерирующий профайлы планет (uwp) исходя из разных правил
	switch generationType {
	case GenerationRules_MGT2_MW_CRB:

	}
	return nil
}

func isValid(u *uwp, dataType string) bool {
	switch dataType {
	default:
		return false
	case Port, Size, Atmo, Hydr, Pops, Govr, Laws, TL:
		if u.Describe(dataType) != "description error" {
			return true
		}
	}
	return false

}
