package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/utils"
)

const (
	DefaultValue             = iota
	SubsectorEmpty           = "Empty"
	KEY_SUBSECTOR_TYPE       = "KEY_SUBSECTOR_TYPE"
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
	KEY_StarSystem_TYPE      = "KEY_StarSystem_TYPE"
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
	KEY_POPULATION           = "KEY_POPULATION"
	PopulationON             = "ForcePopulated"
	PopulationOFF            = "ForceNotPopulated"
	PopulationAuto           = "ForceHabitableOnly"
)

type GenerationState struct {
	Dice          *dice.Dicepool
	SystemName    string
	ConcludedStep int
	NextStep      int
	System        *StarSystem
	vocal         bool
	importedData  []importData
}

type Generator interface {
	Import() error
}

type GeneratorOptions struct {
	key string
	val string
}

func AddOption(key, val string) GeneratorOptions {
	return GeneratorOptions{key, val}
}

func NewGenerator(name string, opt ...GeneratorOptions) (*GenerationState, error) {
	gs := GenerationState{}
	gs.Dice = dice.New().SetSeed(name)
	//gs.Dice.Vocal()
	gs.vocal = true
	gs.SystemName = name
	subsectorType := SubsectorAverage
	systemType := StarSystemRealistic
	populationType := PopulationAuto
	for _, option := range opt {
		switch option.key {
		case KEY_SUBSECTOR_TYPE:
			subsectorType = setSubsectorTypeOption(option.val)
		case KEY_StarSystem_TYPE:
			systemType = setSystemTypeOption(option.val)
		case KEY_POPULATION:
			populationType = setPopulationOption(option.val)
		}
	}
	gs.NextStep = 1
	sts, err := gs.NewStarSystem(subsectorType, systemType, populationType)
	if err != nil {
		return &gs, err
	}
	gs.System = sts
	gs.System.GasGigants = -1
	gs.System.Belts = -1
	gs.System.RockyPlanets = -1
	return &gs, nil
}

func setSubsectorTypeOption(val string) string {
	switch val {
	default:
		return SubsectorAverage
	case SubsectorScattered, SubsectorDispersed, SubsectorAverage, SubsectorCrowded:
		return val
	}
}

func setSystemTypeOption(val string) string {
	switch val {
	default:
		return StarSystemSemiRealistic
	case StarSystemRealistic, StarSystemSemiRealistic, StarSystemFantastic:
		return val
	}
}

func setPopulationOption(val string) string {
	switch val {
	default:
		return PopulationAuto
	case PopulationON, PopulationOFF, PopulationAuto:
		return val
	}
}

func (gs *GenerationState) NewStarSystem(stsType, ssType, pType string) (*StarSystem, error) {
	ss := StarSystem{}
	ss.subsectorType = stsType
	ss.starSystemType = ssType
	ss.populationType = pType
	ss.starPopulation = StarPopulationUNKNOWN
	ss.ObjectType = ObjectUNDEFINED
	return &ss, nil
}

type StarSystem struct {
	subsectorType  string
	starSystemType string
	populationType string
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

func BlackHole() *star {
	bh := star{}
	bh.mass = 999999
	bh.temperature = 1
	bh.luminocity = 0.000000001
	bh.innerLimit = 0.001
	bh.habitableLow = -999
	bh.habitableHigh = -999
	bh.snowLine = 0.002
	bh.outerLimit = 0.2
	bh.class = "Black Hole"
	return &bh
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

type ggiant struct {
	size         int
	descr        string
	spawnedAtAU  float64
	migratedToAU float64
	num          int
	comment      string
	ring         string
	moons        []*rockyPlanet
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
	port          string
	sizeCode      string
	atmoCode      string
	hydrCode      string
	popCode       string
	govCode       string
	lawCode       string
	tlCode        string
	eccentricity  float64
	comment       string
	sizeType      string
	mass          float64
	radius        float64
	potentialAtmo string
	habZone       string
	moons         []*rockyPlanet
	moonOrbit     int
	nativeLife    int
	uwpStr        string
}

func (p *rockyPlanet) getUWP() string {
	return p.uwpStr
}

type detailDataExport struct {
	orbit        float64
	sizeCode     string
	atmoCode     string
	hydrCode     string
	eccentricity float64
	habzone      string
	mw           bool
}

func (dde *detailDataExport) Orbit() float64 {
	return dde.orbit
}
func (dde *detailDataExport) SizeCode() string {
	return dde.sizeCode
}
func (dde *detailDataExport) AtmoCode() string {
	return dde.atmoCode
}
func (dde *detailDataExport) HydrCode() string {
	return dde.hydrCode
}
func (dde *detailDataExport) Eccentricity() float64 {
	return dde.eccentricity
}
func (dde *detailDataExport) Habzone() string {
	return dde.habzone
}
func (dde *detailDataExport) IsMW() bool {
	return dde.mw
}

func NewPlanetExport(orbit float64, sizeCode string, atmoCode string, hydrCode string, eccentricity float64, habzone string, mw bool) *detailDataExport {
	dde := detailDataExport{}
	dde.orbit = orbit
	dde.sizeCode = sizeCode
	dde.atmoCode = atmoCode
	dde.hydrCode = hydrCode
	dde.eccentricity = eccentricity
	dde.habzone = habzone
	dde.mw = mw
	return &dde
}

func (rp *rockyPlanet) ExportDetails() *detailDataExport {
	dde := detailDataExport{}
	dde.orbit = rp.orbit
	dde.sizeCode = rp.sizeCode
	dde.atmoCode = rp.atmoCode
	dde.hydrCode = rp.hydrCode
	dde.eccentricity = rp.eccentricity
	dde.habzone = rp.habZone
	if strings.Contains(rp.comment, "Mainworld") {
		dde.mw = true
	}
	return &dde
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
	port         string
	popCode      string
	govCode      string
	lawCode      string
	tlCode       string
	composition  string
	majorSizeAst int
	width        float64
	lowBorder    float64
	hiBorder     float64
	zone         string
	comment      string
	uwpStr       string
}

func (b *belt) getUWP() string {
	return b.uwpStr
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
	if err := gs.callImport("MW_NAME"); err != nil {
		return nil
	}
	for gs.ConcludedStep < 20 {
		fmt.Printf("Generating Step %v: ", gs.NextStep)
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
			//err = gs.Step12()
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
		case 19:
			err = gs.Step19()
			// case 20:
			// 	err = gs.Step20()
			// 	if err == nil {
			// 		return nil
			// 	}
		}
		fmt.Printf("concluded\r")
		//gs.trackStatus()
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

/*
TESTRUN

*/
