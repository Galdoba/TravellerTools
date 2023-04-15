package pawn

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic/charset"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill/skillset"
)

const (
	control_Unknown = iota
	control_Random
	control_PseudoRandom
	control_User
)

type pawn struct {
	controlType int
	name        string
	chrSet      *charset.CharSet
	sklSet      *skillset.SkillSet
}

/*
A Characteristics
B Birthworld
C Education
D Careers
E Muster Out


*/

type GeneTemplate interface { //фактически это 2 стринга. база находится в генетике
	Profile() string
	Variations() string
}

// type Homeworld interface {
// 	UWP() string
// 	TC() []string
// }

//func New(geneTemplate genetics.Base, homeworld ) (*pawn, error) {

func New(controller int, gt GeneTemplate, homeworldTradeCodes []string) (*pawn, error) {
	chr := pawn{}
	chr.controlType = controller
	if err := chr.Characteristics(gt); err != nil {
		return &chr, err
	}
	if err := chr.HomeworldSkills(homeworldTradeCodes); err != nil {
		return &chr, err
	}

	return &chr, nil
}

func (chr *pawn) Characteristics(genetics GeneTemplate) error {
	chr.chrSet = charset.NewCharSet(dice.New(), genetics.Profile(), genetics.Variations())
	return nil
}

func (chr *pawn) HomeworldSkills(homeworldTradeCodes []string) error {
	set, err := skillset.NewSkillSet()
	if err != nil {
		return fmt.Errorf(" HomeworldSkills: %v\n", err.Error())
	}
	chr.sklSet = set
	for _, tCode := range homeworldTradeCodes {
		idArray := skill.TradeCode2SkillID(tCode)
		for _, id := range idArray {
			if id == skill.ID_NONE {
				continue //skip codes which gives no skill
			}
			if err := chr.sklSet.Increase(id); err != nil {
				if err == skillset.MustChooseErr {
					id, err = ChooseExactSkillId(chr.controlType, id)
					err = chr.sklSet.Increase(id)
					chr.sklSet.IncreaseByKKSrule()
				}
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ChooseExactSkillId(controler, oldID int) (int, error) {
	idList := []int{}
	switch oldID {
	case skill.One_Art:
		idList = []int{skill.ID_Actor, skill.ID_Artist, skill.ID_Author, skill.ID_Chef, skill.ID_Dancer, skill.ID_Musician}
	}
	dicePool := dice.New()
	switch controler {
	default:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): unknown controlType value", controler, oldID)
	case control_Random:
		return idList[dicePool.Sroll(fmt.Sprintf("1d%v", len(idList)-1))], nil
		//return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType RANDOM not implemented", controler, oldID)
	case control_PseudoRandom:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType PSEUDORANDOM not implemented", controler, oldID)
	case control_User:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType USER not implemented", controler, oldID)
	}
}

/*

 */

func (chr *pawn) String() string {
	return chr.chrSet.String() + "\n" + chr.sklSet.String()
}
