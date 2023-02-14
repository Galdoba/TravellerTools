package upp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
)

var positionCodes = []string{"C1", "C2", "C3", "C4", "C5", "C6", "CS", "CP"}

type upp struct {
	chars      map[string]*characteristic.Frame
	csRevealed bool
	cpRevealed bool
}

func (up *upp) String() string {
	str := ""
	for _, positionCode := range positionCodes {
		if positionCode == "CS" || positionCode == "CP" {
			continue
		}
		str += ehex.New().Set(up.chars[positionCode].Val()).Code()
	}
	return str
}

func newProfile(genetics, genMap string) (*upp, error) {
	//dice := dice.New()

	if len(genetics) != 6 {
		return nil, fmt.Errorf("upp.newProfile(%v, %v): len(%v) != 6", genetics, genMap, genetics)
	}
	if len(genetics) != 6 {
		return nil, fmt.Errorf("upp.newProfile(%v, %v): len(%v) != 6", genetics, genMap, genMap)
	}
	up := upp{}
	//up.genetics = genetics
	up.chars = make(map[string]*characteristic.Frame)
	//genes := strings.Split(up.genetics+"  ", "")
	// for i, code := range positionCodes {
	// 	if chr, err := characteristic.ByGeneticProfile(code, genes[i], genMap); err != nil {
	// 		return nil, fmt.Errorf("characteristic.ByGeneticProfile(%v, %v, %v): %v", code, genes[i], genMap, err.Error())
	// 	} else {
	// 		chr.RollValue(dice)
	// 		up.chars[code] = chr

	// 	}
	// }
	return &up, nil
}

type genedata struct {
	geneProf string
	geneMap  string
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

func isInListStr(elem string, list []string) bool {
	for _, s := range list {
		if s == elem {
			return true
		}
	}
	return false
}

func newGenMap(s string) []int {
	arr := []int{}
	profl := strings.Split(s, "")
	for _, data := range profl {
		i, err := strconv.Atoi(data)
		if err != nil {
			arr = append(arr, 0)
			continue
		}
		arr = append(arr, i)
	}
	return arr
}

func GeneDataManual(genetics, geneMap string) (genedata, error) {
	gd := genedata{genetics, geneMap}
	if !isInListStr(genetics, corectProfiles()) {
		return gd, fmt.Errorf("genetics is invalid '%v'", genetics)
	}
	if !isInListStr(geneMap, corectGenMaps()) {
		return gd, fmt.Errorf("geneMap is invalid '%v'", geneMap)
	}
	return gd, nil
}

func GeneDataRandom() genedata {
	//TODO: добавить на вход модификаторы enviroment
	dice := dice.New()
	genetics := "S"
	genetics += strings.Split("AAAADDDGGGG", "")[dice.Flux()+5]
	genetics += strings.Split("SSSSEEEVVVV", "")[dice.Flux()+5]
	genetics += "I"
	genetics += strings.Split("IIIIEEETTTT", "")[dice.Flux()+5]
	genetics += strings.Split("KKKSSSSCCCC", "")[dice.Flux()+5]
	geneMap := ""
	genetics += strings.Split("11222234567", "")[dice.Flux()+5]
	genetics += strings.Split("11222223333", "")[dice.Flux()+5]
	genetics += strings.Split("11222223333", "")[dice.Flux()+5]
	genetics += strings.Split("11222223333", "")[dice.Flux()+5]
	genetics += strings.Split("11222222233", "")[dice.Flux()+5]
	genetics += strings.Split("11222222222", "")[dice.Flux()+5]
	return genedata{genetics, geneMap}
}

func GeneDataHuman() genedata {
	return genedata{"SDEIES", "222222"}
}

func NewUniversalPersonalityProfile(dice *dice.Dicepool, gd genedata) *upp {
	up := upp{}
	up.chars = make(map[string]*characteristic.Frame)
	gene := strings.Split(gd.geneProf, "")
	genDice := newGenMap(gd.geneMap)
	for i := 0; i < 6; i++ {
		code := "C" + fmt.Sprintf("%v", i+1)
		name := characteristic.GeneticCodeToName(gene[i], code)
		up.chars[code] = characteristic.New(name, genDice[i])
		up.chars[code].RollValue(dice)
	}

	return &up
}
