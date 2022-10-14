package systemgeneration

import (
	"fmt"
	"strconv"
	"strings"
)

func (gs *GenerationState) Step01() error {
	if gs.NextStep != 1 {
		fmt.Errorf("not actual step")
	}
	tn := 0
	switch gs.System.subsectorType {
	case SubsectorEmpty:
		tn = 5
	case SubsectorScattered:
		tn = 20
	case SubsectorDispersed:
		tn = 35
	case SubsectorAverage:
		tn = 50
	case SubsectorCrowded:
		tn = 60
	case SubsectorDense:
		tn = 75
	}
	gs.System.ObjectType = ObjectNONE
	gs.debug("ObjectType set as NONE")
	presenceRoll := gs.Dice.Roll("1d100").Sum()
	if presenceRoll <= tn {
		gs.System.ObjectType = ObjectPRESENT
		gs.debug("ObjectType set as PRESENT")
	}
	switch gs.System.ObjectType {
	default:
		return fmt.Errorf("system ObjectType is invalid")
	case ObjectNONE:
		gs.debug("ObjectType Is not in the hex: END GENERATION")
		gs.NextStep = 20
	case ObjectPRESENT:
		gs.NextStep = 2
	}
	fmt.Println("IMPORTING:")
	if err := gs.callImport("Stellar"); err != nil {
		return nil
	}
	fmt.Println("IMPORTING DONE:")

	gs.ConcludedStep = 1
	return nil
}

func (gs *GenerationState) injectStellar(stellar string) error {
	stars, err := decodeStellar(stellar)
	if err != nil {
		return err
	}
	switch len(stars) {
	case 1:
		gs.System.starPopulation = StarPopulationSolo
	case 2:
		gs.System.starPopulation = StarPopulationBinary
	case 3:
		gs.System.starPopulation = StarPopulationTrinary
	case 4:
		gs.System.starPopulation = StarPopulationQuatenary
	case 5:
		gs.System.starPopulation = StarPopulationQuintenary
	}
	gs.System.ObjectType = ObjectStar
	for _, starCode := range stars {
		class, num, size := decodeStar(starCode)
		if num == -1 {
			num = gs.Dice.Roll("1d10").DM(-1).Sum()
		}
		if class == "BD" {
			dwarfTypeRoll := gs.Dice.Roll("1d100").Sum()
			switch {
			case dwarfTypeRoll <= 50:
				class = "L"
			case dwarfTypeRoll <= 75:
				class = "T"
			case dwarfTypeRoll <= 100:
				class = "Y"
			}
		}
		str := &star{class: class, num: num, size: size}
		str.LoadValues()
		gs.System.Stars = append(gs.System.Stars, str)
	}
	gs.NextStep = 7

	return nil
}

func decodeStar(star string) (string, int, string) {
	class := ""
	num := -1
	size := ""
	if star == "BD" {
		return "BD", -1, ""
	}
	if star == "D" {
		return "D", -1, ""
	}
	for _, cl := range []string{"O", "B", "A", "F", "G", "K", "M", "L", "T", "Y"} {
		switch {
		case strings.Contains(star, cl):
			class = cl
		}
	}
	for _, n := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		switch {
		case strings.Contains(star, n):
			nu, _ := strconv.Atoi(n)
			num = nu
		}
	}
	s := strings.Split(star, " ")
	if len(s) > 1 {
		size = s[1]
	}
	return class, num, size
}

func decodeStellar(stellar string) ([]string, error) {
	dt := strings.Split(stellar, " ")
	stars := []string{}
	for i, d := range dt {
		switch d {
		case "BD":
			stars = append(stars, "BD")
		case "D":
			stars = append(stars, "D")
		case "Ia", "Ib", "II", "III", "IV", "V", "VI":
			stars = append(stars, dt[i-1]+" "+d)
		default:
		}
	}
	for _, str := range stars {
		switch str {
		case "BD":
			continue
		case "D":
			continue
		}
		checked := false
		try := 0
		for !checked {
			data := getTableValues(str)
			if data.star != "" {
				break
			}
			if data.star == "" && try == 0 {
				str = strings.ReplaceAll(str, "VI", "V")
				try++
				continue
			}
			if data.star == "" && try == 1 {
				str = strings.ReplaceAll(str, "IV", "V")
				try++
				continue
			}
			try++
			if try > 5 {
				return stars, fmt.Errorf("not matched %v %v %v", str, stars, stellar)
			}
		}

	}

	return stars, nil
}
