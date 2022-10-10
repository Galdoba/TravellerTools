package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/utils"
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
	GasGigantNeptunian       = "Neptunian"
	GasGigantJovian          = "Jovian"
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
	//gs.Dice.Vocal()
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
	Stars          []*star
	GasGigants     int
	GG             []*ggiant
	Belts          int
	BeltData       []*belt
	body           []StellarBody
	MW_UWP         string
	RockyPlanets   int
}

type star struct {
	class                 string
	num                   int
	size                  string
	rank                  string
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
	orbit                 map[float64]StellarBody
	orbitDistances        []float64
}

type bodyHolder struct {
	comment string
}

func (bh *bodyHolder) Describe() string {
	return fmt.Sprintf("%v", bh.comment)
}

func (bh *bodyHolder) setComment(s string) {
	bh.comment = s
}

func (s *star) Describe() string {
	return fmt.Sprintf("%v: %v", s.rank, s.Code())
}

type ggiant struct {
	size         int
	descr        string
	spawnedAtAU  float64
	migratedToAU float64
	num          int
	comment      string
}

func (g *ggiant) Describe() string {
	orbit := g.spawnedAtAU
	if g.migratedToAU != 0 {
		orbit = g.migratedToAU
	}
	return fmt.Sprintf("%v AU	Gas Gigant %v	         	%v	%v", orbit, g.num, g.descr, g.comment)
}

type rockyPlanet struct {
	//stellarBody
	num           int
	star          string
	orbit         float64
	sizeCode      string
	atmoCode      string
	hydrCode      string
	eccentricity  float64
	comment       string
	sizeType      string
	mass          float64
	radius        float64
	potentialAtmo string
	habZone       string
}

func (rp *rockyPlanet) Describe() string {
	return fmt.Sprintf("%v AU	Planet %v	_%v%v%v___-_	%v	%v	%v", rp.orbit, rp.num, rp.sizeCode, rp.atmoCode, rp.hydrCode, rp.sizeType, rp.eccentricity, rp.comment)
}

type jumpZoneBorder struct {
	zone  string
	orbit float64
}

func (jzb *jumpZoneBorder) Describe() string {
	return fmt.Sprintf("Jump Zone Border: %v - %v", jzb.zone, jzb.orbit)
}

type belt struct {
	//stellarBody
	num          int
	star         string
	orbit        float64
	sizeCode     string
	atmoCode     string
	hydrCode     string
	composition  string
	majorSizeAst int
	width        float64
	lowBorder    float64
	hiBorder     float64
	zone         string
	comment      string
}

func (b *belt) Describe() string {
	return fmt.Sprintf("%v AU	Belt %v	%v%v%v			%v", b.orbit, b.num, b.sizeCode, b.atmoCode, b.hydrCode, b.composition)
}

func (b *belt) Width() float64 {
	return b.width
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
		case 12:
			err = gs.Step12()
		case 13:
			err = gs.Step13()
		case 14:
			err = gs.Step14()
		case 15:
			err = gs.Step15()
		case 16:
			err = gs.Step16()
		case 17:
			err = gs.Step17()
		case 18:
			err = gs.Step18()
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

type StellarBody interface {
	Describe() string
}

func roundFloat(f float64, unit float64) float64 {

	return utils.RoundFloat64(f, int(unit))
	//
	// return math.Trunc(f/unit) * unit
}

func sizeOfGG(dp *dice.Dicepool, descr string) int {
	switch descr {
	case GasGigantNeptunian:
		switch dp.Roll("2d6").Sum() {
		case 2:
			return 30
		case 3:
			return 32
		case 4:
			return 35
		case 5:
			return 37
		case 6:
			return 40
		case 7:
			return 42
		case 8:
			return 45
		case 9:
			return 47
		case 10:
			return 50
		case 11:
			return 55
		case 12:
			return 57
		}
	case GasGigantJovian:
		switch dp.Roll("2d6").Sum() {
		case 2:
			return 60
		case 3:
			return 70
		case 4:
			return 80
		case 5:
			return 90
		case 6:
			return 100
		case 7:
			return 110
		case 8:
			return 120
		case 9:
			return 130
		case 10:
			return 140
		case 11:
			return 150
		case 12:
			return 160
		case 13:
			return 170
		case 14:
			return 180
		case 15:
			return 190
		case 16:
			return 200
		case 17:
			return 210
		case 18:
			return 220
		case 19:
			return 230
		case 20:
			return 240
		}
	}
	return -99
}

func starDistanceToClosest(stars []*star, i int) float64 {
	l := len(stars)
	dist1 := 0.0
	dist2 := 0.0
	if i > 0 {
		dist1 = stars[i].distanceFromPrimaryAU - stars[i-1].distanceFromPrimaryAU
	}
	if i+1 < l {
		dist2 = stars[i+1].distanceFromPrimaryAU - stars[i].distanceFromPrimaryAU
	}
	d := dist1 - dist2
	if d < 0 {
		d = d * -1
	}
	return d
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
	if gs.ConcludedStep < 17 {
		return
	}
	fmt.Println("SYSTEM SO FAR:")

	gs.System.printSystemSheet()
	fmt.Println(" ")
}

func (sys *StarSystem) printSystemSheet() {
	for i, bod := range sys.body {
		fmt.Println("--", i, "--", bod.Describe())
	}
	for _, star := range sys.Stars {
		fmt.Println("star.orbit")
		for i := 0; i < 640001; i++ {
			fl := float64(i) / 100
			if v, ok := star.orbit[fl]; ok == true {
				fmt.Printf("Orbit %v = %v\n", fl, v)
			}
		}
	}
}

func (s *star) Code() string {
	code := fmt.Sprintf("%v%v %v", s.class, s.num, s.size)
	code = strings.TrimSpace(code)
	return code
}

/*
TESTRUN

*/
