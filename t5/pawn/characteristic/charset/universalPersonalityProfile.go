package charset

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/Galdoba/TravellerTools/pkg/dice"
// 	"github.com/Galdoba/TravellerTools/pkg/ehex"
// 	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
// )

// var positionCodes = []string{"C1", "C2", "C3", "C4", "C5", "C6", "CS", "CP"}

// type CharSet struct {
// 	chars map[string]*characteristic.Frame
// }

// //func (up *CharSet) ProfileKeys() []

// func (up *CharSet) String() string {
// 	str := ""
// 	for _, positionCode := range positionCodes {
// 		if positionCode == "CS" || positionCode == "CP" {
// 			continue
// 		}
// 		str += ehex.New().Set(up.chars[positionCode].Val()).Code()
// 	}
// 	return str
// }

// func newGenMap(s string) []int {
// 	arr := []int{}
// 	profl := strings.Split(s, "")
// 	for _, data := range profl {
// 		i, err := strconv.Atoi(data)
// 		if err != nil {
// 			arr = append(arr, 0)
// 			continue
// 		}
// 		arr = append(arr, i)
// 	}
// 	return arr
// }

// // func GeneDataRandom() genedata {
// // 	//TODO: добавить на вход модификаторы enviroment
// // 	dice := dice.New()
// // 	genetics := "S"
// // 	genetics += strings.Split("AAAADDDGGGG", "")[dice.Flux()+5]
// // 	genetics += strings.Split("SSSSEEEVVVV", "")[dice.Flux()+5]
// // 	genetics += "I"
// // 	genetics += strings.Split("IIIIEEETTTT", "")[dice.Flux()+5]
// // 	genetics += strings.Split("KKKSSSSCCCC", "")[dice.Flux()+5]
// // 	geneMap := ""
// // 	geneMap = randomGenemap(genetics)
// // 	return genedata{genetics, geneMap}
// // }

// // func GeneDataHuman() genedata {
// // 	return genedata{"SDEIES", "222222"}
// // }

// func NewCharSet(dice *dice.Dicepool, geneticProfile, geneticVariations string) *CharSet {
// 	up := CharSet{}
// 	up.chars = make(map[string]*characteristic.Frame)
// 	gene := strings.Split(geneticProfile, "")
// 	genDice := newGenMap(geneticVariations)
// 	for i := 0; i < 6; i++ {
// 		code := "C" + fmt.Sprintf("%v", i+1)
// 		name := characteristic.GeneticCodeToName(gene[i], code)
// 		up.chars[code] = characteristic.New(name, genDice[i])
// 		up.chars[code].RollValue(dice)
// 	}

// 	return &up
// }

// func (up *CharSet) ValueOf(code string) int {
// 	for _, c := range positionCodes {
// 		if c == code {
// 			return up.chars[code].Actual()
// 		}
// 	}
// 	for _, chr := range up.chars {
// 		try := chr.ValueAs(code)
// 		if try > -99 {
// 			return try
// 		}
// 	}
// 	return 0
// }
