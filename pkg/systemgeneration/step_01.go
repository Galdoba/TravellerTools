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
	for _, imported := range gs.importedData {
		fmt.Println(imported)
		switch imported.dataKey {
		case "Stellar":
			if err := gs.injectStellar(imported.data); err != nil {
				return err
			}
		}
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
		starData := getTableValues(d)
		if starData.innerLimit != 0 {
			stars = append(stars, d)
		}
		if i != 0 {
			starData := getTableValues(dt[i-1] + " " + dt[i])
			if starData.innerLimit != 0 {
				stars = append(stars, dt[i-1]+" "+dt[i])
			}
		}
	}
	check := ""
	for _, st := range stars {
		fmt.Println("detected", st)
		check = st + " "
	}
	check = strings.TrimSuffix(check, " ")
	if check != stellar {
		return []string{}, fmt.Errorf("'%v' != '%v'", stellar, check)
	}

	return stars, nil
}
