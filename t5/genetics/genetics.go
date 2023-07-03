package genetics

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
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
	KEY_GENE_PRF_1     = "GenePrf1"
	KEY_GENE_PRF_2     = "GenePrf2"
	KEY_GENE_PRF_3     = "GenePrf3"
	KEY_GENE_PRF_4     = "GenePrf4"
	KEY_GENE_PRF_5     = "GenePrf5"
	KEY_GENE_PRF_6     = "GenePrf6"
	KEY_GENE_MAP_1     = "GeneMap1"
	KEY_GENE_MAP_2     = "GeneMap2"
	KEY_GENE_MAP_3     = "GeneMap3"
	KEY_GENE_MAP_4     = "GeneMap4"
	KEY_GENE_MAP_5     = "GeneMap5"
	KEY_GENE_MAP_6     = "GeneMap6"
	KEY_LIFE_STG_DUR_1 = "LifeStgDur1"
	KEY_LIFE_STG_DUR_2 = "LifeStgDur2"
	KEY_LIFE_STG_DUR_3 = "LifeStgDur3"
	KEY_LIFE_STG_DUR_4 = "LifeStgDur4"
	KEY_LIFE_STG_DUR_5 = "LifeStgDur5"
	KEY_LIFE_STG_DUR_6 = "LifeStgDur6"
	KEY_LIFE_STG_DUR_7 = "LifeStgDur7"
	KEY_LIFE_STG_DUR_8 = "LifeStgDur8"
	KEY_LIFE_STG_DUR_9 = "LifeStgDur9"
)

// func NewTemplate(profile, variations string) *GeneTemplate {
// 	gp := &GeneTemplate{profile, variations}
// 	return gp
// }

// func EmptyTemplate() *GeneTemplate {
// 	gp := &GeneTemplate{}
// 	return gp
// }

func IsEmpty(gp Genome) bool {
	if gp.Profile()+gp.Variations() == "" {
		return true
	}
	return false
}

// func Check(gp *GeneTemplate) error {
// 	for i, p := range strings.Split(gp.Profile(), "") {
// 		correctValues := []string{}
// 		switch i {
// 		case 0:
// 			correctValues = []string{"S"}
// 		case 1:
// 			correctValues = []string{"D", "A", "G"}
// 		case 2:
// 			correctValues = []string{"E", "S", "V"}
// 		case 3:
// 			correctValues = []string{"I"}
// 		case 4:
// 			correctValues = []string{"E", "T", "I"}
// 		case 6:
// 			correctValues = []string{"S", "C", "K"}
// 		}
// 		if !isInListStr(p, correctValues) {
// 			return fmt.Errorf("gp.Profile(): position %v is incorect", i)
// 		}
// 	}
// 	for i, p := range strings.Split(gp.Variations(), "") {
// 		correctValues := []string{}
// 		switch i {
// 		case 0:
// 			correctValues = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
// 		case 1:
// 			correctValues = []string{"1", "2", "3"}
// 		case 2:
// 			correctValues = []string{"1", "2", "3"}
// 		case 3:
// 			correctValues = []string{"1", "2", "3"}
// 		case 4:
// 			correctValues = []string{"1", "2", "3"}
// 		case 6:
// 			correctValues = []string{"1", "2"}
// 		}
// 		if !isInListStr(p, correctValues) {
// 			return fmt.Errorf("gp.Variations(): position %v is incorect", i)
// 		}
// 	}
// 	return nil
// }

// type GeneTemplate struct {
// 	geneProf string
// 	geneMap  string
// }

type GeneProfile profile.Profile

// func GenomeString(gp GeneProfile) string {
// 	str := ""
// 	val := "?"
// 	for _, key := range []string{
// 		KEY_GENE_PRF_1, KEY_GENE_PRF_2, KEY_GENE_PRF_3, KEY_GENE_PRF_4, KEY_GENE_PRF_5, KEY_GENE_PRF_6,
// 	} {
// 		if gp.Data(key) == nil {
// 			continue
// 		}
// 		switch key {
// 		default:
// 		case KEY_GENE_MAP_1, KEY_GENE_MAP_2, KEY_GENE_MAP_3, KEY_GENE_MAP_4, KEY_GENE_MAP_5, KEY_GENE_MAP_6:
// 			switch gp.Data(key).Code() {
// 			case "1", "2", "3", "4", "5", "6", "7", "8":
// 				val = gp.Data(key).Code()
// 			}

// 		}
// 		str += val
// 		val = "?"
// 	}
// 	return str
// }

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

func Profile(gp GeneProfile) string {
	str := ""
	val := "?"
	for _, key := range []string{
		KEY_GENE_PRF_1, KEY_GENE_PRF_2, KEY_GENE_PRF_3, KEY_GENE_PRF_4, KEY_GENE_PRF_5, KEY_GENE_PRF_6,
	} {
		if gp.Data(key) == nil {
			//str += "?"
			continue
		}
		switch key {
		default:
		case "-":
			val = "-"
		case KEY_GENE_PRF_1:
			switch gp.Data(key).Value() {
			case CHAR_STRENGHT:
				val = "S"
			}
		case KEY_GENE_PRF_2:
			switch gp.Data(key).Value() {
			case CHAR_DEXTERITY:
				val = "D"
			case CHAR_AGILITY:
				val = "A"
			case CHAR_GRACE:
				val = "G"
			}
		case KEY_GENE_PRF_3:
			switch gp.Data(key).Value() {
			case CHAR_ENDURANCE:
				val = "E"
			case CHAR_STAMINA:
				val = "S"
			case CHAR_VIGOR:
				val = "V"
			}
		case KEY_GENE_PRF_4:
			switch gp.Data(key).Value() {
			case CHAR_INTELLIGENCE:
				val = "I"
			}
		case KEY_GENE_PRF_5:
			switch gp.Data(key).Value() {
			case CHAR_EDUCATION:
				val = "E"
			case CHAR_TRAINING:
				val = "T"
			case CHAR_INSTINCT:
				val = "I"
			}
		case KEY_GENE_PRF_6:
			switch gp.Data(key).Value() {
			case CHAR_SOCIAL:
				val = "S"
			case CHAR_CHARISMA:
				val = "C"
			case CHAR_CASTE:
				val = "K"
			}

		}
		str += val
		val = "?"
	}
	return str
}

func Variations(gp GeneProfile) string {
	str := ""
	val := "?"
	for _, key := range []string{
		KEY_GENE_MAP_1, KEY_GENE_MAP_2, KEY_GENE_MAP_3, KEY_GENE_MAP_4, KEY_GENE_MAP_5, KEY_GENE_MAP_6,
	} {
		if gp.Data(key) == nil {
			//str += "?"
			continue
		}
		switch key {
		case KEY_GENE_MAP_1, KEY_GENE_MAP_2, KEY_GENE_MAP_3, KEY_GENE_MAP_4, KEY_GENE_MAP_5, KEY_GENE_MAP_6:
			switch gp.Data(key).Code() {
			case "1", "2", "3", "4", "5", "6", "7", "8":
				val = gp.Data(key).Code()
			}

		}
		ehex.New().Set("X").Encode("Unknown")
		str += val
		val = "?"
	}
	return str
}

func GeneTemplateHuman() GeneProfile {
	gp := NewGeneData("SDEIES", "222222")
	return gp
}

func RollGenome(dice *dice.Dicepool) GeneProfile {
	geneProf := rollGeneProfile(dice)
	geneMap := rollGenemap(geneProf, dice)
	gp := NewGeneData(geneProf, geneMap)
	return gp
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

func GenomeCompatability(genome1, genome2 GeneProfile) int {
	gp1 := strings.Split(Profile(genome1)+Variations(genome1), "")
	gp2 := strings.Split(Profile(genome2)+Variations(genome2), "")
	match := -2
	for i := range gp1 {
		if gp1[i] == gp2[i] {
			match++
		}
	}
	return match
}

func FromProfile(prf profile.Profile) GeneProfile {
	gen := profile.New()
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
		data := prf.Data(key)
		if data == nil {
			return nil
		}
		gen.Inject(key, data)
	}
	return gen
}

func BlankGenome() GeneProfile {
	gemoneMap := profile.New()
	for _, c := range []string{"1", "2", "3", "4", "5", "6", "S", "P"} {
		gemoneMap.Inject("PRF_C"+fmt.Sprintf("%v", c), ehex.New().Set("X"))
		gemoneMap.Inject("MAP_C"+fmt.Sprintf("%v", c), ehex.New().Set("X"))
		gemoneMap.Inject("GENE_C"+fmt.Sprintf("%v", c), ehex.New().Set("X"))
		gemoneMap.Inject("MUT_C"+fmt.Sprintf("%v", c), ehex.New().Set("X"))
		gemoneMap.Inject("C"+fmt.Sprintf("%v", c), ehex.New().Set(0))
	}
	return gemoneMap
}

func HumanGeneString() string {
	return "01478BEF 22222222 0000XXXX BBBBXXXX 00000000"
	//      01478B 222222 4456XX BGBRBB
	//      1100100110010000 222222 4456XX BGBRBB" - speceis:individual
	//      SDEIES 222222 4566XX BGBRBB
}

//func FillGenomeProfile(gen GeneProfile, geneString string) GeneProfile {
// 	data := strings.Split(geneString, "")
// 	codes := charCodes()
// 	for i, val := range data {
// 		v := ehex.New().Set(val)
// 		switch i {
// 		case 0:
// //			gen.Inject()
// 		}
// 	}
//}

func charTypes() []string {
	keys := []string{}
	for i := 0; i <= 15; i++ {
		keys = append(keys, fmt.Sprintf("CHRTYPE%v", i))
	}
	return keys
}

func charCodes() []string {
	keys := []string{"C1", "C2", "C3", "C4", "C5", "C6", "CS", "CP"}
	return keys
}

func characteristics() []string {
	keys := []string{
		"Strength",
		"Dexterity",
		"Agility",
		"Grace",
		"Endurance",
		"Stamina",
		"Vigor",
		"Intelligence",
		"Education",
		"Training",
		"Instinct",
		"Social",
		"Charisma",
		"Caste",
		"Sanity",
		"Psionics"}
	return keys
}

type environmentDM struct {
	size    ehex.Ehex
	atmo    ehex.Ehex
	hydr    ehex.Ehex
	terrain string
}

func EnvironmentDM(size, atmo, hydr ehex.Ehex, terrain string) int {
	dm := 0
	if atmo.Value() >= 8 {
		dm = dm - 2
	}
	if size.Value() <= 5 {
		dm = dm - 1
	}
	if hydr.Value() >= 6 {
		dm = dm + 1
	}
	if hydr.Value() >= 9 {
		dm = dm + 1
	}
	return dm
}
