package genetics

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
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
)

func NewTemplate(profile, variations string) *GeneTemplate {
	gp := &GeneTemplate{profile, variations}
	return gp
}

func EmptyTemplate() *GeneTemplate {
	gp := &GeneTemplate{}
	return gp
}

func IsEmpty(gp Genome) bool {
	if gp.Profile()+gp.Variations() == "" {
		return true
	}
	return false
}

func Check(gp *GeneTemplate) error {
	for i, p := range strings.Split(gp.Profile(), "") {
		correctValues := []string{}
		switch i {
		case 0:
			correctValues = []string{"S"}
		case 1:
			correctValues = []string{"D", "A", "G"}
		case 2:
			correctValues = []string{"E", "S", "V"}
		case 3:
			correctValues = []string{"I"}
		case 4:
			correctValues = []string{"E", "T", "I"}
		case 6:
			correctValues = []string{"S", "C", "K"}
		}
		if !isInListStr(p, correctValues) {
			return fmt.Errorf("gp.Profile(): position %v is incorect", i)
		}
	}
	for i, p := range strings.Split(gp.Variations(), "") {
		correctValues := []string{}
		switch i {
		case 0:
			correctValues = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
		case 1:
			correctValues = []string{"1", "2", "3"}
		case 2:
			correctValues = []string{"1", "2", "3"}
		case 3:
			correctValues = []string{"1", "2", "3"}
		case 4:
			correctValues = []string{"1", "2", "3"}
		case 6:
			correctValues = []string{"1", "2"}
		}
		if !isInListStr(p, correctValues) {
			return fmt.Errorf("gp.Variations(): position %v is incorect", i)
		}
	}
	return nil
}

type GeneTemplate struct {
	geneProf string
	geneMap  string
}

type GeneProfile profile.Profile

func HumanGeneData() (string, string) {
	return "SDEIES", "222222"
}

func NewGeneData(prfl, variations string) GeneProfile {
	gd := profile.New()
	for i, p := range strings.Split(prfl, "") {
		switch i {
		case 0:
			gd.Inject(KEY_GENE_PRF_1, CHAR_STRENGHT)
		case 1:
			switch p {
			case "D":
				gd.Inject(KEY_GENE_PRF_2, CHAR_DEXTERITY)
			case "A":
				gd.Inject(KEY_GENE_PRF_2, CHAR_AGILITY)
			case "G":
				gd.Inject(KEY_GENE_PRF_2, CHAR_GRACE)
			}
		case 2:
			switch p {
			case "E":
				gd.Inject(KEY_GENE_PRF_3, CHAR_ENDURANCE)
			case "S":
				gd.Inject(KEY_GENE_PRF_3, CHAR_STAMINA)
			case "V":
				gd.Inject(KEY_GENE_PRF_3, CHAR_VIGOR)
			}
		case 3:
			gd.Inject(KEY_GENE_PRF_4, CHAR_INTELLIGENCE)
		case 4:
			switch p {
			case "E":
				gd.Inject(KEY_GENE_PRF_5, CHAR_EDUCATION)
			case "T":
				gd.Inject(KEY_GENE_PRF_5, CHAR_TRAINING)
			case "I":
				gd.Inject(KEY_GENE_PRF_5, CHAR_INSTINCT)
			}
		case 5:
			switch p {
			case "S":
				gd.Inject(KEY_GENE_PRF_6, CHAR_SOCIAL)
			case "C":
				gd.Inject(KEY_GENE_PRF_6, CHAR_CHARISMA)
			case "K":
				gd.Inject(KEY_GENE_PRF_6, CHAR_CASTE)
			}

		}
	}
	for i, v := range strings.Split(variations, "") {
		switch i {
		case 0:
			gd.Inject(KEY_GENE_MAP_1, v)
		case 1:
			gd.Inject(KEY_GENE_MAP_2, v)
		case 2:
			gd.Inject(KEY_GENE_MAP_3, v)
		case 3:
			gd.Inject(KEY_GENE_MAP_4, v)
		case 4:
			gd.Inject(KEY_GENE_MAP_5, v)
		case 5:
			gd.Inject(KEY_GENE_MAP_6, v)
		}
	}
	return gd
}

type Genome interface {
	Profile() string
	Variations() string
}

func (gt *GeneTemplate) Profile() string {
	return gt.geneProf
}

func (gt *GeneTemplate) Variations() string {
	return gt.geneMap
}

func GeneTemplateHuman() *GeneTemplate {
	return &GeneTemplate{"SDEIES", "222222"}
}

func GeneTemplateManual(genetics, geneMap, seed string) (GeneTemplate, error) {
	gd := GeneTemplate{genetics, geneMap}
	if genetics == "" {
		gd.geneProf = randomGeneProfile(seed)

	}
	if !isInListStr(gd.geneProf, corectProfiles()) {
		return gd, fmt.Errorf("genetics is invalid '%v'", genetics)
	}
	if geneMap == "" {
		gd.geneMap = randomGenemap(gd.geneProf, seed+seed)
	}
	if !isInListStr(gd.geneMap, corectGenMaps()) {
		return gd, fmt.Errorf("geneMap is invalid '%v'", geneMap)
	}
	fmt.Println(gd)
	return gd, nil
}

func RollGenome(dice *dice.Dicepool) *GeneTemplate {
	gt := GeneTemplate{}
	gt.geneProf = rollGeneProfile(dice)
	gt.geneMap = rollGenemap(gt.geneProf, dice)
	return &gt
}

func isInListStr(elem string, list []string) bool {
	for _, s := range list {
		if s == elem {
			return true
		}
	}
	return false
}

func randomGeneProfile(seed string) string {
	dice := dice.New().SetSeed(seed)
	genetics := "S"
	genetics += strings.Split("AAAADDDGGGG", "")[dice.Flux()+5]
	genetics += strings.Split("SSSSEEEVVVV", "")[dice.Flux()+5]
	genetics += "I"
	genetics += strings.Split("IIIIEEETTTT", "")[dice.Flux()+5]
	genetics += strings.Split("KKKSSSSCCCC", "")[dice.Flux()+5]
	return genetics
}

func rollGeneProfile(dice *dice.Dicepool) string {
	genetics := "S"
	genetics += strings.Split("AAAADDDGGGG", "")[dice.Flux()+5]
	genetics += strings.Split("SSSSEEEVVVV", "")[dice.Flux()+5]
	genetics += "I"
	genetics += strings.Split("IIIIEEETTTT", "")[dice.Flux()+5]
	genetics += strings.Split("KKKSSSSCCCC", "")[dice.Flux()+5]
	return genetics
}

func newGeneMap(geneprof, genemap, seed string) string {
	if genemap == "" {
		genemap = randomGenemap(geneprof, seed)
	}
	return genemap
}

func randomGenemap(geneprof, seed string) string {
	//без учета экологических факторов
	dice := *dice.New().SetSeed(seed)
	genemmap := ""
	genemmap += strings.Split("11222234567", "")[dice.Flux()+5]
	genemmap += strings.Split("11222223333", "")[dice.Flux()+5]
	genemmap += strings.Split("11222223333", "")[dice.Flux()+5]
	genemmap += strings.Split("11222223333", "")[dice.Flux()+5]
	switch strings.Split(geneprof, "")[4] {
	default:
		genemmap += strings.Split("11222222233", "")[dice.Flux()+5]
	case "I":
		genemmap += "2"
	}
	genemmap += strings.Split("11222222222", "")[dice.Flux()+5]
	return genemmap
}

func rollGenemap(geneprof string, dice *dice.Dicepool) string {
	//без учета экологических факторов
	genemmap := ""
	genemmap += strings.Split("11222234567", "")[dice.Flux()+5]
	genemmap += strings.Split("11222223333", "")[dice.Flux()+5]
	genemmap += strings.Split("11222223333", "")[dice.Flux()+5]
	genemmap += strings.Split("11222223333", "")[dice.Flux()+5]
	switch strings.Split(geneprof, "")[4] {
	default:
		genemmap += strings.Split("11222222233", "")[dice.Flux()+5]
	case "I":
		genemmap += "2"
	}
	genemmap += strings.Split("11222222222", "")[dice.Flux()+5]
	return genemmap
}

func corectProfiles() []string {
	gp := []string{}
	for _, c1 := range []string{"S"} {
		for _, c2 := range []string{"D", "A", "G"} {
			for _, c3 := range []string{"E", "S", "V"} {
				for _, c4 := range []string{"I"} {
					for _, c5 := range []string{"E", "T", "I"} {
						for _, c6 := range []string{"S", "C", "K"} {
							gp = append(gp, c1+c2+c3+c4+c5+c6)
						}
					}
				}
			}
		}
	}
	return gp
}

func corectGenMaps() []string {
	gp := []string{}
	for _, c1 := range []string{"1", "2", "3", "4", "5", "6", "7", "8"} {
		for _, c2 := range []string{"1", "2", "3"} {
			for _, c3 := range []string{"1", "2", "3"} {
				for _, c4 := range []string{"1", "2", "3"} {
					for _, c5 := range []string{"1", "2", "3"} {
						for _, c6 := range []string{"1", "2"} {
							gp = append(gp, c1+c2+c3+c4+c5+c6)
						}
					}
				}
			}
		}
	}
	return gp
}

func GenomeCompatability(genome1, genome2 Genome) int {
	gp1 := strings.Split(genome1.Profile()+genome1.Variations(), "")
	gp2 := strings.Split(genome2.Profile()+genome2.Variations(), "")
	match := -2
	for i := range genome1.Profile() {
		if gp1[i] == gp2[i] {
			match++
		}
	}
	return match
}
