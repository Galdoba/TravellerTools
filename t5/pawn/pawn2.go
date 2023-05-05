package pawn

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/devtools/errmaker"
)

const (
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
	CHAR_STRENGHT  = iota
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
)

type pawn2 struct {
	controlType int
	name        string
	profile     profile.Profile
}

func New2() *pawn2 {
	p := pawn2{}
	p.profile = profile.New()
	return &p
}

func (p *pawn2) String() string {
	str := "UPP: "
	keys := []string{KEY_VAL_C1, KEY_VAL_C2, KEY_VAL_C3, KEY_VAL_C4, KEY_VAL_C5, KEY_VAL_C6}
	for _, k := range keys {
		str += p.profile.Data(k).Code()
	}
	return str
}

func (p *pawn2) RollCharacteristics(dice *dice.Dicepool) error {
	genome, err := p.GenomeTemplate()
	if err != nil {
		return errmaker.ErrorFrom(err)
	}

	p.InjectGenetics(genome)
	keys := []string{KEY_VAL_C1, KEY_VAL_C2, KEY_VAL_C3, KEY_VAL_C4, KEY_VAL_C5, KEY_VAL_C6}
	//fmt.Println(genome.Variations())
	mapKeys := []string{KEY_GENE_MAP_1, KEY_GENE_MAP_2, KEY_GENE_MAP_3, KEY_GENE_MAP_4, KEY_GENE_MAP_5, KEY_GENE_MAP_6}
	for i, mKey := range mapKeys {
		diceNbr := p.profile.Data(mKey).Code()
		set := dice.Sroll(fmt.Sprintf("%vd6", diceNbr))
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

func (p *pawn2) InjectGenetics(gp profile.Profile) error {

	//////////
	keys := CharacteristicProfileKeys()
	for _, key := range keys {
		p.profile.Inject(key, gp.Data(key).Code())
	}
	//////////
	return nil
}
