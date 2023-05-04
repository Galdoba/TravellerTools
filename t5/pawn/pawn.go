package pawn

import (
	"fmt"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic/charset"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill/skillset"
	"github.com/Galdoba/devtools/errmaker"
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

type Homeworld interface {
	//UWP() string
	ListTC() []int
}

//func New(geneTemplate genetics.Base, homeworld ) (*pawn, error) {

func New(controller int, gt GeneTemplate, hw Homeworld) (*pawn, error) {
	chr := pawn{}
	chr.controlType = controller
	if err := chr.Characteristics(gt); err != nil {
		return &chr, err
	}
	if err := chr.HomeworldSkills(hw.ListTC()); err != nil {
		return &chr, err
	}

	return &chr, nil
}

func (chr *pawn) Characteristics(genetics GeneTemplate) error {
	chr.chrSet = charset.NewCharSet(dice.New(), genetics.Profile(), genetics.Variations())
	return nil
}

func (chr *pawn) SkillSet() error {
	sklset, err := skillset.NewSkillSet()
	if err != nil {
		return errmaker.ErrorFrom(err)
	}
	chr.sklSet = sklset
	return nil
}

func (chr *pawn) HomeworldSkills(homeworldTradeCodes []int) error {
	set, err := skillset.NewSkillSet()
	if err != nil {
		return errmaker.ErrorFrom(err, homeworldTradeCodes)
	}
	chr.sklSet = set
	for _, tCode := range homeworldTradeCodes {
		idArray := skill.TradeCode2SkillID(tCode)
		for _, id := range idArray {
			if id == skill.ID_NONE {
				continue //skip codes which gives no skill
			}
			if err := chr.sklSet.Increase(id); err != nil {
				for err != nil {
					switch err {
					case skillset.KKSruleNotAllow:
						id, err = ChoseKnowledgeID(chr.controlType, id)
						chr.sklSet.Increase(id)
					case skillset.MustChooseErr:
						id, err = ChooseExactSkillID(chr.controlType, id)
						err = chr.sklSet.Increase(id)
					}
				}
				// if err == skillset.MustChooseErr {

				// 	id, err = ChooseExactSkillID(chr.controlType, id)
				// 	err = chr.sklSet.Increase(id)
				// 	if err == skillset.KKSruleNotAllow {

				// 		id, err = ChoseKnowledgeID(chr.controlType, id)
				// 		chr.sklSet.Increase(id)
				// 	}
				// }

			}
		}
	}
	return nil
}

func ChooseExactSkillID(controler, oldID int) (int, error) {
	idList := []int{}
	switch oldID {
	case skill.One_Art:
		idList = []int{skill.ID_Actor, skill.ID_Artist, skill.ID_Author, skill.ID_Chef, skill.ID_Dancer, skill.ID_Musician}
	case skill.One_Trade:
		//Biologics, Craftsman, Electronics, Fluidics, Gravitics, Magnetics, Mechanic, Photonics, Polymers, Programmer.
		idList = []int{skill.ID_Biologics, skill.ID_Craftsman, skill.ID_Electronics, skill.ID_Fluidics,
			skill.ID_Gravitics, skill.ID_Magnetics, skill.ID_Mechanic, skill.ID_Photonics, skill.ID_Polymers, skill.ID_Programmer}
	}
	dicePool := dice.New()
	switch controler {
	default:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): unknown controlType value", controler, oldID)
	case control_Random:
		return idList[dicePool.Sroll(fmt.Sprintf("1d%v-1", len(idList)))], nil
	case control_PseudoRandom:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType PSEUDORANDOM not implemented", controler, oldID)
	case control_User:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType USER not implemented", controler, oldID)
	}
}

func ChoseKnowledgeID(controller, skill_ID int) (int, error) {
	skl, err := skill.New(skill_ID)
	if err != nil {
		return 0, fmt.Errorf("")
	}
	knowledges := []int{}
	//     Language, Musician
	switch skill_ID {
	default:
		return skill.ID_NONE, fmt.Errorf("skill %v has %v associated knowledges", skl.Name, len(skl.AssociatedKnowledge))
	case skill.ID_Gunnery:
		knowledges = append(knowledges, skill.ID_Bay_Weapons, skill.ID_Ortilery, skill.ID_Screens, skill.ID_Spines, skill.ID_Turrets)
	case skill.ID_Heavy_Weapons:
		knowledges = append(knowledges, skill.ID_Artilery, skill.ID_Launchers, skill.ID_Ordinance, skill.ID_WMD)
	case skill.ID_Fighter:
		knowledges = append(knowledges, skill.ID_Battle_Dress, skill.ID_Beams, skill.ID_Blades, skill.ID_Exotics, skill.ID_Slugs, skill.ID_Sprays, skill.ID_Unarmed)
	case skill.ID_Flyer:
		knowledges = append(knowledges, skill.ID_Flappers, skill.ID_LTA, skill.ID_Rotor, skill.ID_Winged, skill.ID_Grav_f, skill.ID_Aeronautics)
	case skill.ID_Driver:
		knowledges = append(knowledges, skill.ID_ACV, skill.ID_Legged, skill.ID_Mole, skill.ID_Tracked, skill.ID_Wheeled, skill.ID_Grav_d)
	case skill.ID_Engineer:
		knowledges = append(knowledges, skill.ID_Jump, skill.ID_Life_Support, skill.ID_Maneuver, skill.ID_Power)
	case skill.ID_Animals:
		knowledges = append(knowledges, skill.ID_Rider, skill.ID_Teamster, skill.ID_Trainer)
	case skill.ID_Seafarer:
		knowledges = append(knowledges, skill.ID_Aquanautics, skill.ID_Grav_s, skill.ID_Boat, skill.ID_Ship, skill.ID_Sub)
	case skill.ID_Pilot:
		knowledges = append(knowledges, skill.ID_Small_Craft, skill.ID_Spacecraft_ABS, skill.ID_Spacecraft_BCS)
	case skill.ID_Language:
		knowledges = append(knowledges, skill.ID_Language_Kkree, skill.ID_Language_Anglic, skill.ID_Language_Battle,
			skill.ID_Language_Flash, skill.ID_Language_Gonk, skill.ID_Language_Gvegh, skill.ID_Language_Mariel,
			skill.ID_Language_Oynprith, skill.ID_Language_Sagamaal, skill.ID_Language_Tezapet,
			skill.ID_Language_Trokh, skill.ID_Language_Vilani, skill.ID_Language_Zdetl)
	case skill.ID_Musician:
		knowledges = append(knowledges, skill.ID_Instrument_Guitar, skill.ID_Instrument_Banjo, skill.ID_Instrument_Mandolin, skill.ID_Instrument_Keyboard, skill.ID_Instrument_Piano, skill.ID_Instrument_Voice, skill.ID_Instrument_Trumpet, skill.ID_Instrument_Trombone, skill.ID_Instrument_Tuba, skill.ID_Instrument_Violin, skill.ID_Instrument_Viola, skill.ID_Instrument_Cello)
	}
	dicePool := dice.New()
	switch controller {
	default:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): unknown controlType value", controller, skill_ID)
	case control_Random:
		//fmt.Println(dicePool.Sroll("1d100"))
		time.Sleep(time.Nanosecond)
		return knowledges[dicePool.Sroll(fmt.Sprintf("1d%v-1", len(knowledges)))], nil
	case control_PseudoRandom:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType PSEUDORANDOM not implemented", controller, skill_ID)
	case control_User:
		return 0, fmt.Errorf("ChooseExactSkillId(controlType(%v), oldID(%v)): controlType USER not implemented", controller, skill_ID)
	}

}

/*

 */
