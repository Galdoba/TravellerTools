package genetics

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func NewTemplate(profile, variations string) *GeneTemplate {
	return &GeneTemplate{profile, variations}
}

type GeneTemplate struct {
	geneProf string
	geneMap  string
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
