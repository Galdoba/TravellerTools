package pawn

import (
	"fmt"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/genetics"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
	"github.com/Galdoba/devtools/errmaker"
)

const (
	control_Unknown = iota
	control_Random
	control_PseudoRandom
	control_User
)

const (
	CHAR_STRENGHT = iota
	CHAR_DEXTERITY
	CHAR_AGILITY
	CHAR_GRACE
	CHAR_ENDURANCE
	CHAR_STAMINA
	CHAR_VIGOR
	CHAR_INTELLIGENCE
	CHAR_EDUCATION
	CHAR_TRAINING
	CHAR_INSTINCT
	CHAR_SOCIAL
	CHAR_CHARISMA
	CHAR_CASTE
	CHAR_SANITY
	CHAR_PSIONICS
	C1
	C2
	C3
	C4
	C5
	C6
	CP
	CS
	KEY_VAL_C1     = "C1"
	KEY_VAL_C2     = "C2"
	KEY_VAL_C3     = "C3"
	KEY_VAL_C4     = "C4"
	KEY_VAL_C5     = "C5"
	KEY_VAL_C6     = "C6"
	KEY_VAL_CP     = "CP"
	KEY_VAL_CS     = "CS"
	KEY_GENE_PRF_1 = "GenePrf1"
	KEY_GENE_PRF_2 = "GenePrf2"
	KEY_GENE_PRF_3 = "GenePrf3"
	KEY_GENE_PRF_4 = "GenePrf4"
	KEY_GENE_PRF_5 = "GenePrf5"
	KEY_GENE_PRF_6 = "GenePrf6"
	KEY_GENE_MAP_1 = "GeneMap1"
	KEY_GENE_MAP_2 = "GeneMap2"
	KEY_GENE_MAP_3 = "GeneMap3"
	KEY_GENE_MAP_4 = "GeneMap4"
	KEY_GENE_MAP_5 = "GeneMap5"
	KEY_GENE_MAP_6 = "GeneMap6"
	MajorSkill     = "Major Skill"
	MinorSkill     = "Minor Skill"
)

type Pawn struct {
	generationState int
	controlType     int
	name            string
	profile         profile.Profile
	major           int
	minor           int
	degree          string
}

func (p *Pawn) EducationState() (int, int, string) {
	return p.major, p.minor, p.degree
}

func New(dice *dice.Dicepool, control int, homeworldTC []int) (*Pawn, error) {
	p := Pawn{}
	p.controlType = control
	p.profile = profile.New()
	if err := p.RollCharacteristics(dice); err != nil {
		return nil, errmaker.ErrorFrom(err)
	}
	for _, id := range skill.DefaultSkills() {
		p.CreateSkill(id)
	}
	for _, tCode := range homeworldTC {
		idArray := skill.TradeCode2SkillID(tCode)

		for _, id := range idArray {
			p.IncreaseSkill(id)
		}
	}
	///////////EDUCATION

	return &p, nil
}

func (p *Pawn) String() string {
	str := "UPP: "
	keys := []string{KEY_VAL_C1, KEY_VAL_C2, KEY_VAL_C3, KEY_VAL_C4, KEY_VAL_C5, KEY_VAL_C6}
	for _, k := range keys {
		str += p.profile.Data(k).Code()
	}
	str += "\n"
	sklset := p.Skills()

	for i := skill.ID_NONE; i < skill.ID_END; i++ {
		skl := sklset.Data(skill.NameByID(i))
		if skl != nil {
			str += fmt.Sprintf("%v %v\n", skill.NameByID(i), skl.Value())
		}
	}
	return str
}

func (p *Pawn) RollCharacteristics(dice *dice.Dicepool) error {
	genome, err := p.Genome()
	if err != nil {
		return errmaker.ErrorFrom(err)
	}
	p.InjectGenetics(genome)
	keys := []string{KEY_VAL_C1, KEY_VAL_C2, KEY_VAL_C3, KEY_VAL_C4, KEY_VAL_C5, KEY_VAL_C6}
	mapKeys := []string{KEY_GENE_MAP_1, KEY_GENE_MAP_2, KEY_GENE_MAP_3, KEY_GENE_MAP_4, KEY_GENE_MAP_5, KEY_GENE_MAP_6}
	for i, mKey := range mapKeys {
		diceNbr := p.profile.Data(mKey).Code()
		diceCode := ""
		switch diceNbr {
		case "1":
			diceCode = "1d6"
		case "2":
			diceCode = "2d6"
		case "3":
			diceCode = "3d6"
		case "4":
			diceCode = "2d6+12"
		case "5":
			diceCode = "3d6+12"
		case "6":
			diceCode = "4d6+12"
		case "7":
			diceCode = "5d6+12"
		case "8":
			diceCode = "6d6+12"
		}

		set := dice.Sroll(fmt.Sprintf("%v", diceCode))
		p.profile.Inject(keys[i], set)
	}

	return nil
}

func CharacteristicProfileKeys() []string {
	return []string{
		KEY_GENE_PRF_1,
		KEY_GENE_PRF_2,
		KEY_GENE_PRF_3,
		KEY_GENE_PRF_4,
		KEY_GENE_PRF_5,
		KEY_GENE_PRF_6,
		KEY_GENE_MAP_1,
		KEY_GENE_MAP_2,
		KEY_GENE_MAP_3,
		KEY_GENE_MAP_4,
		KEY_GENE_MAP_5,
		KEY_GENE_MAP_6,
	}
}

func (p *Pawn) InjectGenetics(gp genetics.GeneProfile) error {
	keys := CharacteristicProfileKeys()
	for _, key := range keys {
		p.profile.Inject(key, gp.Data(key).Code())
	}
	return nil
}

/////////////////////////

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
		knowledges = append(knowledges, skill.ID_Small_Craft, skill.ID_Spacecraft_ACS, skill.ID_Spacecraft_BCS)
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

func (p *Pawn) Skills() skill.SkillSet {

	sklset := skill.NewSkillSet(p.profile)
	return sklset
}

func (p *Pawn) CreateSkill(i int) error {
	if i == skill.ID_NONE {
		return nil
	}
	sklKey := skill.NameByID(i)
	if sklKey == "UNDEFINED" {
		return fmt.Errorf("skill [%v] undefined", i)
	}
	if p.profile.Data(sklKey) != nil {
		return fmt.Errorf("skill [%v] present", i)
	}
	p.profile.Inject(sklKey, 0)
	return nil
}

func (p *Pawn) ReadSkill(ID int) *skill.Skill {
	sklKey := skill.NameByID(ID)
	if sklKey == "UNDEFINED" {
		return nil
	}
	skl, err := skill.New(ID)
	if err != nil {
		return nil
	}
	sklHex := p.profile.Data(sklKey)
	if sklHex == nil {
		return nil
	}
	skl.ValueInt = sklHex.Value()
	return skl
}

func (p *Pawn) UpdateSkill(ID, newVal int) error {
	sklKey := skill.NameByID(ID)
	if sklKey == "UNDEFINED" {
		return fmt.Errorf("cann't update skill")
	}
	skl, err := skill.New(ID)
	if err != nil {
		return err
	}

	switch skl.SType() {
	case skill.TYPE_SKILL:
		if newVal > 15 {
			return fmt.Errorf("skill cann't be higher than 15")
		}
	case skill.TYPE_KNOWLEDGE:
		if newVal > 6 {
			return fmt.Errorf("knowledge cann't be higher than 6")
		}
	case skill.TYPE_TALENT:
		if newVal > 6 {
			return fmt.Errorf("talent cann't be higher than 6")
		}
	}
	p.profile.Inject(sklKey, newVal)
	if newVal > 0 {
		parentID := skl.ParentSkl
		parentKey := skill.NameByID(parentID)
		if parent := p.profile.Data(parentKey); parent == nil {
			p.CreateSkill(parentID)
		}
	}
	return nil
}

func (p *Pawn) Learn(ID int) error {
	sklKey := skill.NameByID(ID)
	if sklKey == "UNDEFINED" {
		return fmt.Errorf("cann't update skill")
	}
	skl := p.ReadSkill(ID)
	if skl != nil {
		p.UpdateSkill(ID, skl.ValueInt+1)
		return nil
	}
	p.CreateSkill(ID)
	p.UpdateSkill(ID, 1)
	return nil
}

func (p *Pawn) DeleteSkill(ID int) error {
	sklKey := skill.NameByID(ID)
	if sklKey == "UNDEFINED" {
		return fmt.Errorf("cann't delete skill")
	}
	p.profile.Delete(sklKey)
	return nil
}

func (p *Pawn) IncreaseSkill(id int) error {
	for {
		switch skill.Increase(p.profile, id) {
		case skill.MustChooseErr:
			switch id {
			case skill.One_Art, skill.One_Trade:
				newId, err := ChooseExactSkillID(p.controlType, id)
				if err != nil {
					return fmt.Errorf("%v %v", newId, err.Error())
				}
				if skill.Increase(p.profile, newId) == nil {
					id = newId
					continue
				}
			}
		case skill.KKSruleNotAllow:
			newId, err := ChoseKnowledgeID(p.controlType, id)
			if err != nil {
				return fmt.Errorf("%v %v", newId, err.Error())
			}
			if err = skill.Increase(p.profile, newId); err == nil {
				id = newId
				continue
			}
		default:
			return p.Learn(id)
		}
		return nil
	}
}
