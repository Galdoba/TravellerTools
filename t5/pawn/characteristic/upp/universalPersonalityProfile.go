package upp

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/t5/pawn/characteristic"
)

var positionCodes = []string{"C1", "C2", "C3", "C4", "C5", "C6", "CS", "CP"}

type upp struct {
	genetics   string
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

func newProfile(genetics string) (*upp, error) {
	dice := dice.New()

	if len(genetics) != 6 {
		return nil, fmt.Errorf("upp.newProfile(%v): len(%v) != 6", genetics, genetics)
	}
	up := upp{}
	up.genetics = genetics
	up.chars = make(map[string]*characteristic.Frame)
	genes := strings.Split(up.genetics+"  ", "")
	for i, code := range positionCodes {
		if chr, err := characteristic.ByGeneticProfile(code, genes[i]); err != nil {
			return nil, fmt.Errorf("characteristic.ByGeneticProfile(%v, %v): %v", code, genes[i], err.Error())
		} else {
			chr.SetValue(dice.Sroll("2d6"))
			up.chars[code] = chr

		}
	}
	return &up, nil
}
