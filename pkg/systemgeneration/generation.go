package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

const (
	DefaultValue             = iota
	SubsectorEmpty           = "Empty"
	SubsectorScattered       = "Scattered"
	SubsectorDispersed       = "Dispersed"
	SubsectorAverage         = "Average"
	SubsectorCrowded         = "Crowded"
	SubsectorDense           = "Dense"
	ObjectUNDEFINED          = "UNDEFINED"
	ObjectNONE               = "NONE"
	ObjectPRESENT            = "PRESENT"
	ObjectStar               = "Star"
	ObjectBrownDwarf         = "BrownDwarf"
	ObjectRoguePlanet        = "RoguePlanet"
	ObjectRogueGasGigant     = "RogueGasGigant"
	ObjectNeutronStar        = "NeutronStar"
	ObjectNebula             = "Nebula"
	ObjectBlackHole          = "BlackHole"
	StarSystemRealistic      = "Realistic"
	StarSystemSemiRealistic  = "Semi-Realistic"
	StarSystemFantastic      = "Fantastic"
	StarPopulationUNKNOWN    = "Unknown"
	StarPopulationSolo       = "Single Star System"
	StarPopulationBinary     = "Binary Star System"
	StarPopulationTrinary    = "Trinary Star System"
	StarPopulationQuatenary  = "Quatenary Star System"
	StarPopulationQuintenary = "Quintenary Star System"
	StarDistancePrimary      = "Primary"
	StarDistanceContact      = "Contact"
	StarDistanceClose        = "Close"
	StarDistanceNear         = "Near"
	StarDistanceFar          = "Far"
	StarDistanceDistant      = "Distant"
)

type GenerationState struct {
	Dice          *dice.Dicepool
	SystemName    string
	ConcludedStep int
	NextStep      int
	System        *StarSystem
	vocal         bool
}

type Generator interface {
}

func NewGenerator(name string) (*GenerationState, error) {
	gs := GenerationState{}
	gs.Dice = dice.New().SetSeed(name)
	gs.Dice.Vocal()
	gs.vocal = true
	gs.SystemName = name
	gs.debug(fmt.Sprintf("SystemName set as %v", name))
	gs.NextStep = 1
	sts, err := gs.NewStarSystem(SubsectorAverage, StarSystemRealistic)
	if err != nil {
		return &gs, err
	}
	gs.System = sts
	return &gs, nil
}

func (gs *GenerationState) NewStarSystem(stsType, ssType string) (*StarSystem, error) {
	ss := StarSystem{}
	ss.subsectorType = stsType
	ss.starSystemType = ssType
	ss.starPopulation = StarPopulationUNKNOWN
	ss.ObjectType = ObjectUNDEFINED
	gs.debug(fmt.Sprintf("ObjectType set as %v\n", ObjectUNDEFINED))
	gs.debug(fmt.Sprintf("starSystemType set as %v\n", ssType))
	return &ss, nil
}

type StarSystem struct {
	subsectorType  string
	starSystemType string
	starPopulation string
	ObjectType     string
	Stars          []star
	GasGigants     int
	Belts          int
}

type star struct {
	class                 string
	num                   int
	size                  string
	generated             bool
	distanceType          string
	distanceFromPrimaryAU float64
	temperature           int
	mass                  float64
	luminocity            float64
	innerLimit            float64
	habitableLow          float64
	habitableHigh         float64
	snowLine              float64
	outerLimit            float64
}

func (gs *GenerationState) GenerateData() error {
	err := fmt.Errorf("initial error")
	err = nil
	for gs.ConcludedStep < 20 {
		switch gs.NextStep {
		default:
			err = fmt.Errorf("gs.NextStep = %v unimplemented", gs.NextStep)
		case 1:
			err = gs.Step01()
		case 2:
			err = gs.Step02()
		case 3:
			err = gs.Step03()
		case 4:
			err = gs.Step04()
		case 5:
			err = gs.Step05()
		case 6:
			err = gs.Step06()
		case 7:
			err = gs.Step07()
		case 8:
			err = gs.Step08()
		case 9:
			err = gs.Step09()
		case 10:
			err = gs.Step10()
		case 11:
			err = gs.Step11()
		case 20:
			err = gs.Step20()
			if err == nil {
				return nil
			}
		}
		gs.trackStatus()
		if err != nil {
			return fmt.Errorf("GenerateData returned err=%v\n concluded Step = %v\n next step = %v", err.Error(), gs.ConcludedStep, gs.NextStep)
		}
	}
	return fmt.Errorf("unresolved generation\n concluded Step = %v\n next step = %v", gs.ConcludedStep, gs.NextStep)
}

func (gs *GenerationState) Step01() error {
	fmt.Println("START Step 01")
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
	gs.ConcludedStep = 1
	fmt.Println("END Step 01")
	return nil
}

func (gs *GenerationState) Step02() error {
	fmt.Println("START Step 02")
	if gs.NextStep != 2 {
		return fmt.Errorf("not actual step")
	}
	typeRoll := gs.Dice.Roll("1d100").Sum()
	switch {
	case typeRoll <= 80:
		gs.System.ObjectType = ObjectStar
		gs.NextStep = 4
	case typeRoll <= 88:
		gs.System.ObjectType = ObjectBrownDwarf
		gs.NextStep = 3
	case typeRoll <= 94:
		gs.System.ObjectType = ObjectRoguePlanet
		gs.NextStep = 15
	case typeRoll <= 97:
		gs.System.ObjectType = ObjectRogueGasGigant
		gs.NextStep = 13
	case typeRoll <= 98:
		gs.System.ObjectType = ObjectNeutronStar
		gs.NextStep = 18
	case typeRoll <= 99:
		gs.System.ObjectType = ObjectNebula
		gs.NextStep = 18
	case typeRoll <= 100:
		gs.System.ObjectType = ObjectBlackHole
		gs.NextStep = 20
	}
	gs.debug(fmt.Sprintf("gs.System.ObjectType set as %v", gs.System.ObjectType))
	gs.ConcludedStep = 2
	switch gs.NextStep {
	case 4, 3, 15, 13, 18, 20:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 02")
	return nil
}

func (gs *GenerationState) Step03() error {
	fmt.Println("START Step 03")
	if gs.NextStep != 3 {
		return fmt.Errorf("not actual step")
	}
	starType := ""
	dwarfTypeRoll := gs.Dice.Roll("1d100").Sum()
	switch {
	case dwarfTypeRoll <= 50:
		starType = "L"
	case dwarfTypeRoll <= 75:
		starType = "T"
	case dwarfTypeRoll <= 100:
		starType = "Y"
	}
	str := star{class: starType, num: -1, size: ""}
	gs.System.Stars = append(gs.System.Stars, str)
	gs.debug("starType is " + starType)
	gs.NextStep = 5
	gs.debug(fmt.Sprintf("gs.System.ObjectType set as %v", gs.System.ObjectType))
	gs.ConcludedStep = 3
	switch gs.NextStep {
	case 5:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 03")
	return nil
}

func (gs *GenerationState) Step04() error {
	fmt.Println("START Step 04")
	if gs.NextStep != 4 {
		return fmt.Errorf("not actual step")
	}
	starType := ""
	starTypeRoll := gs.Dice.Roll("1d100").Sum()
	tn := []int{}
	switch gs.System.starSystemType {
	default:
		return fmt.Errorf("unknown gs.System.starSystemType (%v)", gs.System.starSystemType)
	case StarSystemRealistic:
		tn = []int{80, 88, 94, 97, 98, 99, 100}
	case StarSystemSemiRealistic:
		tn = []int{50, 77, 90, 97, 98, 99, 100}
	case StarSystemFantastic:
		tn = []int{25, 50, 75, 97, 98, 99, 100}
	}
	switch {
	case starTypeRoll <= tn[0]:
		starType = "M"
	case starTypeRoll <= tn[1]:
		starType = "K"
	case starTypeRoll <= tn[2]:
		starType = "G"
	case starTypeRoll <= tn[3]:
		starType = "F"
	case starTypeRoll <= tn[4]:
		starType = "A"
	case starTypeRoll <= tn[5]:
		starType = "B"
	case starTypeRoll <= tn[6]:
		starType = "O"
	}
	str := star{class: starType, num: -1, size: ""}
	gs.System.Stars = append(gs.System.Stars, str)
	gs.debug("starType is " + starType)
	gs.NextStep = 5
	gs.debug(fmt.Sprintf("gs.System.ObjectType set as %v", gs.System.ObjectType))
	gs.ConcludedStep = 4
	switch gs.NextStep {
	case 5:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 04")
	return nil
}

func (gs *GenerationState) Step05() error {
	fmt.Println("START Step 05")
	if gs.NextStep != 5 {
		return fmt.Errorf("not actual step")
	}
	numRoll := gs.Dice.Roll("1d10").Sum()
	for i, star := range gs.System.Stars {
		if star.num != -1 {
			continue
		}
		gs.System.Stars[i].num += numRoll
		gs.debug(fmt.Sprintf("gs.System.Stars[%v].num set as %v", i, gs.System.Stars[i].num))
		gs.ConcludedStep = 5
		switch gs.System.Stars[i].class {
		case "O", "B", "A", "F", "G", "K", "M":
			gs.NextStep = 6
		case "L", "T", "Y":
			gs.NextStep = 7
		}
	}
	switch gs.NextStep {
	case 6, 7:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 05")
	return nil
}

func (gs *GenerationState) Step06() error {
	fmt.Println("START Step 06")
	if gs.NextStep != 6 {
		return fmt.Errorf("not actual step")
	}
	lumRoll1 := gs.Dice.Roll("1d100").Sum()
	lumRoll2 := gs.Dice.Roll("1d10").Sum()
	for i, star := range gs.System.Stars {
		switch star.class {
		case "L", "T", "Y":
			continue
		}
		if star.size != "" {
			continue
		}
		switch {
		case lumRoll1 <= 90:
			gs.System.Stars[i].size = "V"
		case lumRoll1 <= 94:
			gs.System.Stars[i].size = "IV"
		case lumRoll1 <= 96:
			gs.System.Stars[i].size = "D"
		case lumRoll1 <= 99:
			gs.System.Stars[i].size = "III"
		case lumRoll1 <= 100:
			switch {
			case lumRoll2 <= 4:
				gs.System.Stars[i].size = "II"
			case lumRoll2 <= 6:
				gs.System.Stars[i].size = "VI"
			case lumRoll2 <= 8:
				gs.System.Stars[i].size = "Ia"
			case lumRoll2 <= 10:
				gs.System.Stars[i].size = "Ib"
			}
		}
		if gs.System.Stars[i].size == "D" {
			gs.System.Stars[i].class = "D"
			gs.System.Stars[i].size = ""
		}
		// if gs.System.Stars[i].size != "" {
		// 	continue
		// }

		switch gs.System.Stars[i].size {
		case "O", "B", "A", "F":
			if gs.System.Stars[i].size == "VI" {
				gs.System.Stars[i].size = "V"
			}
		}
	}
	gs.ConcludedStep = 6
	gs.NextStep = 7
	switch gs.NextStep {
	case 7:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 06")
	return nil
}

func (gs *GenerationState) Step07() error {
	fmt.Println("START Step 07")
	if gs.NextStep != 7 {
		return fmt.Errorf("not actual step")
	}
	switch gs.System.starPopulation {
	default:
		return fmt.Errorf("star population unexpected")
	case StarPopulationSolo, StarPopulationBinary, StarPopulationTrinary, StarPopulationQuatenary, StarPopulationQuintenary:
		fmt.Printf("System: %v, have %v, want %v\n", gs.System.starPopulation, len(gs.System.Stars), strSystToNum(gs.System.starPopulation))
		if len(gs.System.Stars) < strSystToNum(gs.System.starPopulation) {
			gs.NextStep = 4
			return nil
		}
		gs.ConcludedStep = 7
		gs.NextStep = 8
		gs.System.Stars = sortStars(gs.System.Stars)
	case StarPopulationUNKNOWN:
		tn := []int{}
		switch gs.System.Stars[0].class {
		case "O", "B", "A":
			tn = []int{10, 90, 98, 99, 100}
		case "F", "G", "K":
			tn = []int{45, 99, 100, 599, 999}
		case "M", "L", "T", "Y", "":
			tn = []int{69, 98, 100, 200, 300}
		}
		strComposRoll := gs.Dice.Roll("1d100").Sum()
		switch {
		case strComposRoll <= tn[0]:
			gs.System.starPopulation = StarPopulationSolo
		case strComposRoll <= tn[1]:
			gs.System.starPopulation = StarPopulationBinary
		case strComposRoll <= tn[2]:
			gs.System.starPopulation = StarPopulationTrinary
		case strComposRoll <= tn[3]:
			gs.System.starPopulation = StarPopulationQuatenary
		case strComposRoll <= tn[4]:
			gs.System.starPopulation = StarPopulationQuintenary
		}
		return nil
	}
	switch gs.NextStep {
	case 4, 8:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 07")
	return nil
}

func (gs *GenerationState) Step08() error {
	fmt.Println("START Step 08")
	if gs.NextStep != 8 {
		return fmt.Errorf("not actual step")
	}
	switch gs.System.starPopulation {
	default:
		return fmt.Errorf("imposible population at step 08")

	case StarPopulationSolo, StarPopulationBinary, StarPopulationTrinary, StarPopulationQuatenary, StarPopulationQuintenary:
		for i, _ := range gs.System.Stars {
			if i == 0 {
				gs.System.Stars[0].distanceType = "Primary"
				gs.System.Stars[0].distanceFromPrimaryAU = 0.0
				fmt.Println(gs.System.Stars[i])
				continue
			}
			for gs.System.Stars[i].distanceFromPrimaryAU <= gs.System.Stars[i-1].distanceFromPrimaryAU {
				dist, au := rollDistance(gs.Dice)
				gs.System.Stars[i].distanceType = dist
				gs.System.Stars[i].distanceFromPrimaryAU = au
				if dist == StarDistanceContact {
					gs.System.Stars[i].distanceFromPrimaryAU = gs.System.Stars[i-1].distanceFromPrimaryAU
				}

			}
			fmt.Println(gs.System.Stars[i])
		}
		gs.ConcludedStep = 8
		gs.NextStep = 9
	}
	switch gs.NextStep {
	case 9:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 08")
	return nil
}

func (gs *GenerationState) Step09() error {
	fmt.Println("START Step 09")
	if gs.NextStep != 9 {
		return fmt.Errorf("not actual step")
	}
	switch gs.System.Stars[0].class {
	case "O", "B", "A", "F", "G", "K", "M":
		gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-2).Sum()
	case "L":
		gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-4).Sum()
	case "T":
		gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-5).Sum()
	case "Y":
		gs.System.GasGigants = gs.Dice.Roll("1d6").DM(-6).Sum()
	}
	if gs.System.GasGigants < 0 {
		gs.System.GasGigants = 0
	}
	fmt.Println("Gas Gigants:", gs.System.GasGigants)
	gs.ConcludedStep = 9
	gs.NextStep = 10
	switch gs.NextStep {
	case 10:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 09")
	return nil
}

func (gs *GenerationState) Step10() error {
	fmt.Println("START Step 10")
	if gs.NextStep != 10 {
		return fmt.Errorf("not actual step")
	}
	gs.System.Belts = gs.Dice.Roll("1d6").DM(-3).Sum()
	if gs.System.Belts < 0 {
		gs.System.Belts = 0
	}
	fmt.Println("Belts:", gs.System.Belts)
	gs.ConcludedStep = 10
	gs.NextStep = 11
	switch gs.NextStep {
	case 11:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 10")
	return nil
}

func (gs *GenerationState) Step11() error {
	fmt.Println("START Step 11")
	if gs.NextStep != 11 {
		return fmt.Errorf("not actual step")
	}
	for i := range gs.System.Stars {
		gs.System.Stars[i].LoadValues()
	}

	gs.ConcludedStep = 11
	gs.NextStep = 12
	switch gs.NextStep {
	case 12:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	fmt.Println("END Step 10")
	return nil
}

func rollDistance(dp *dice.Dicepool) (string, float64) {
	r1 := dp.Roll("1d100").Sum()
	r2 := dp.Roll("1d100").Sum()
	dist := distChart(r1)
	au := auChart(r2, dist)
	return dist, au
}

func distChart(i int) string {
	distChart := []int{10, 30, 50, 80, 100}
	switch {
	case i <= distChart[0]:
		return StarDistanceContact
	case i <= distChart[1]:
		return StarDistanceClose
	case i <= distChart[2]:
		return StarDistanceNear
	case i <= distChart[3]:
		return StarDistanceFar
	case i <= distChart[4]:
		return StarDistanceDistant
	}
	return "DISTANCE UNDEFINED"
}

func auChart(i int, dType string) float64 {
	distChart := []float64{}
	switch dType {
	case StarDistanceContact:
		return -1
	case StarDistanceClose:
		distChart = []float64{0.5, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0, 5.5}
	case StarDistanceNear:
		distChart = []float64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	case StarDistanceFar:
		distChart = []float64{100, 150, 200, 250, 300, 350, 400, 450, 500, 550}
	case StarDistanceDistant:
		distChart = []float64{600, 750, 1000, 1500, 2000, 2500, 3000, 4000, 5000, 6000}
	}
	fmt.Println(i, "*---", dType)
	switch {
	case i <= 9:
		return distChart[0]
	case i <= 19:
		return distChart[1]
	case i <= 29:
		return distChart[2]
	case i <= 39:
		return distChart[3]
	case i <= 49:
		return distChart[4]
	case i <= 59:
		return distChart[5]
	case i <= 69:
		return distChart[6]
	case i <= 79:
		return distChart[7]
	case i <= 89:
		return distChart[8]
	case i <= 100:
		return distChart[9]
	}
	return 0.0
}

func strSystToNum(str string) int {
	switch str {
	case StarPopulationSolo:
		return 1
	case StarPopulationBinary:
		return 2
	case StarPopulationTrinary:
		return 3
	case StarPopulationQuatenary:
		return 4
	case StarPopulationQuintenary:
		return 5
	default:
		return -1
	}
}

func sortStars(stars []star) []star {
	strSizes := []int{}
	for _, str := range stars {
		strSizes = append(strSizes, setSize(str))
	}
	newOrder := []star{}
	for i := 1000; i > -10; i-- {
		for v, num := range strSizes {
			if i != num {
				continue
			}
			newOrder = append(newOrder, stars[v])
		}
	}
	return newOrder
}

func setSize(s star) int {
	ss := 0
	for _, scl := range []string{"L", "T", "Y", "M", "K", "G", "F", "A", "B", "O"} {
		if s.class != scl {
			ss += 10
			continue
		}
		ss -= s.num
		break
	}
	for _, scl := range []string{"", "VI", "V", "IV", "III", "II", "Ib", "Ia"} {
		if s.class != scl {
			ss += 100
			// вса
			continue
		}
		break
	}
	return ss
}

func (gs *GenerationState) Step20() error {
	fmt.Println("START Step 20")
	if gs.NextStep != 20 {
		return fmt.Errorf("not actual step")
	}
	gs.NextStep = 99
	fmt.Println("END Step 20")
	return nil
}

func (gs *GenerationState) debug(str string) {
	if gs.vocal {
		fmt.Println(str)
	}
}

func (gs *GenerationState) trackStatus() {
	fmt.Printf("generation steps %v/%v\n", gs.ConcludedStep, gs.NextStep)
}

func (s *star) Code() string {
	code := fmt.Sprintf("%v%v %v", s.class, s.num, s.size)
	code = strings.TrimSpace(code)
	return code
}

/*
TESTRUN

*/
