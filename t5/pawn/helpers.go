package pawn

import (
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/t5/genetics"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

const (
	CheckEasy    = -1
	CheckAverage = 0
	CheckHard    = 1
)

func (p *Pawn) Genome() (genetics.GeneProfile, error) {
	gp := profile.New()
	nodata := 0
	for _, key := range []string{
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
	} {
		data := p.profile.Data(key)
		if data == nil {
			nodata++
			continue
		}
		gp.Inject(key, data.Code())
	}
	if nodata == 12 {
		return genetics.NewGeneData(genetics.HumanGeneData()), nil
	}
	return gp, nil
}

func (p *Pawn) Characteristic(key int) int {
	switch key {
	case C1, CHAR_STRENGHT:
		return p.profile.Data(KEY_VAL_C1).Value()
	case C2:
		return p.profile.Data(KEY_VAL_C2).Value()
	case C3:
		return p.profile.Data(KEY_VAL_C3).Value()
	case C4, CHAR_INTELLIGENCE:
		return p.profile.Data(KEY_VAL_C4).Value()
	case C5:
		return p.profile.Data(KEY_VAL_C5).Value()
	case C6:
		return p.profile.Data(KEY_VAL_C6).Value()
	case CHAR_DEXTERITY:
		switch p.profile.Data(KEY_GENE_PRF_2).Value() {
		default:
			return (p.profile.Data(KEY_VAL_C2).Value() + 1) / 2
		case CHAR_DEXTERITY:
			return p.profile.Data(KEY_VAL_C2).Value()
		}
	case CHAR_AGILITY:
		switch p.profile.Data(KEY_GENE_PRF_2).Value() {
		default:
			return (p.profile.Data(KEY_VAL_C2).Value() + 1) / 2
		case CHAR_AGILITY:
			return p.profile.Data(KEY_VAL_C2).Value()
		}
	case CHAR_GRACE:
		switch p.profile.Data(KEY_GENE_PRF_2).Value() {
		default:
			return (p.profile.Data(KEY_VAL_C2).Value() + 1) / 2
		case CHAR_GRACE:
			return p.profile.Data(KEY_VAL_C2).Value()
		}
	case CHAR_ENDURANCE:
		switch p.profile.Data(KEY_GENE_PRF_3).Value() {
		default:
			return (p.profile.Data(KEY_VAL_C3).Value() + 1) / 2
		case CHAR_ENDURANCE:
			return p.profile.Data(KEY_VAL_C3).Value()
		}
	case CHAR_STAMINA:
		switch p.profile.Data(KEY_GENE_PRF_3).Value() {
		default:
			return (p.profile.Data(KEY_VAL_C3).Value() + 1) / 2
		case CHAR_STAMINA:
			return p.profile.Data(KEY_VAL_C3).Value()
		}
	case CHAR_VIGOR:
		switch p.profile.Data(KEY_GENE_PRF_3).Value() {
		default:
			return (p.profile.Data(KEY_VAL_C3).Value() + 1) / 2
		case CHAR_VIGOR:
			return p.profile.Data(KEY_VAL_C3).Value()
		}

	case CHAR_EDUCATION:
		switch p.profile.Data(KEY_GENE_PRF_5).Value() {
		case CHAR_TRAINING:
			return (p.profile.Data(KEY_VAL_C5).Value() + 1) / 2
		case CHAR_EDUCATION:
			return p.profile.Data(KEY_VAL_C5).Value()
		case CHAR_INSTINCT:
			return 4
		}
	case CHAR_TRAINING:
		switch p.profile.Data(KEY_GENE_PRF_5).Value() {
		case CHAR_TRAINING:
			return p.profile.Data(KEY_VAL_C5).Value()
		case CHAR_EDUCATION:
			return (p.profile.Data(KEY_VAL_C5).Value() + 1) / 2
		case CHAR_INSTINCT:
			return 4
		}
	case CHAR_INSTINCT:
		switch p.profile.Data(KEY_GENE_PRF_5).Value() {
		default:
			return 4
		case CHAR_INSTINCT:
			return p.profile.Data(KEY_VAL_C5).Value()
		}

	case CHAR_SOCIAL:
		switch p.profile.Data(KEY_GENE_PRF_6).Value() {
		case CHAR_SOCIAL:
			return p.profile.Data(KEY_VAL_C6).Value()
		case CHAR_CHARISMA:
			return (p.profile.Data(KEY_VAL_C6).Value() + 1) / 2
		case CHAR_CASTE:
			return 4
		}
	case CHAR_CHARISMA:
		switch p.profile.Data(KEY_GENE_PRF_6).Value() {
		case CHAR_SOCIAL:
			return (p.profile.Data(KEY_VAL_C6).Value() + 1) / 2
		case CHAR_CHARISMA:
			return p.profile.Data(KEY_VAL_C6).Value()
		case CHAR_CASTE:
			return 4
		}
	case CHAR_CASTE:
		switch p.profile.Data(KEY_GENE_PRF_6).Value() {
		default:
			return 4
		case CHAR_CASTE:
			return p.profile.Data(KEY_VAL_C6).Value()
		}

	}
	return 4
}

func (p *Pawn) Skill(key int) int {
	keyStr := skill.NameByID(key)
	if skl := p.profile.Data(keyStr); skl != nil {
		return skl.Value()
	}
	return -1
}

func charID2Map(id int) string {
	mp := ""
	switch id {
	case CHAR_STRENGHT, C1:
		mp = KEY_GENE_MAP_1
	case CHAR_DEXTERITY, CHAR_AGILITY, CHAR_GRACE, C2:
		mp = KEY_GENE_MAP_2
	case CHAR_ENDURANCE, CHAR_STAMINA, CHAR_VIGOR, C3:
		mp = KEY_GENE_MAP_3
	case CHAR_INTELLIGENCE, C4:
		mp = KEY_GENE_MAP_4
	case CHAR_EDUCATION, CHAR_TRAINING, CHAR_INSTINCT, C5:
		mp = KEY_GENE_MAP_5
	case CHAR_SOCIAL, CHAR_CHARISMA, CHAR_CASTE, C6:
		mp = KEY_GENE_MAP_6
	}
	return mp
}

func (p *Pawn) CheckCharacteristic(diff, asset int) bool {
	chr := characteristic.FromProfile(p.profile, asset)
	switch p.controlType {
	case control_Random:
		return chr.Check(diff, dice.New())
	}
	// chr := p.Characteristic(asset)
	// df := p.profile.Data(charID2Map(asset)).Value() + diff
	// if df < 1 {
	// 	return true
	// }
	// switch p.controlType {
	// default:
	// 	panic("control type not implemented")
	// case control_Random:
	// 	dice := dice.New()
	// 	dice.Vocal()
	// 	if dice.Sroll(fmt.Sprintf("%vd6", df)) <= chr {
	// 		return true
	// 	}
	// }
	panic(0)
	return false
}
