package landing

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"gopkg.in/AlecAivazis/survey.v1"
)

const (
	speedStationary = iota
	speedCreep
	speedSlowPort
	speedPort
	speedSlowTransfer
	speedOrbitalTransfer
	speedSlowTransit
	speedTransit
	speedFastTransit
	speedExtreme
	worldStarport
	worldSize
	worldAtmo
	weather
	shiptype
)

type speed struct {
	val             string
	descr           string
	cumulativeTrust int
	deverityCode    string
	err             error
}

func setSpeed(val int) speed {
	sp := speed{}
	switch val {
	default:
		sp.err = fmt.Errorf("unknown val (%v)", val)
		return sp
	case speedStationary:
		sp.val = "Stationary"
		sp.descr = "Ship is not moving."
		sp.cumulativeTrust = 0
		sp.deverityCode = "0d6"
	case speedCreep:
		sp.val = "Creep"
		sp.descr = "Extremely slow movement such as in final docking or when making small adjustments to position."
		sp.cumulativeTrust = 1
		sp.deverityCode = "1d3"
	case speedSlowPort:
		sp.val = "Slow Port"
		sp.descr = "Careful movement within the confines of a port, typically used by large ships and in crowded areas."
		sp.cumulativeTrust = 2
		sp.deverityCode = "1d6"
	case speedPort:
		sp.val = "Port"
		sp.descr = "Movement within the confines of a port is normally carried out at this speed, balancing brisk and efficient against safety."
		sp.cumulativeTrust = 4
		sp.deverityCode = "2d6"
	case speedSlowTransfer:
		sp.val = "Slow Transfer"
		sp.descr = "Cautious manoeuvres of a sort normally used in crowded orbital space."
		sp.cumulativeTrust = 8
		sp.deverityCode = "3d6"
	case speedOrbitalTransfer:
		sp.val = "Orbital Transfer"
		sp.descr = "Standard speed at which ships enter and leave orbit, or manoeuvre in orbital space."
		sp.cumulativeTrust = 16
		sp.deverityCode = "4d6"
	case speedSlowTransit:
		sp.val = "Slow Transit"
		sp.descr = "Slower vessels or ones moving cautiously will typically be manoeuvring at this rate. Traffic control will normally object to ships moving faster than this pace within its controlled area."
		sp.cumulativeTrust = 32
		sp.deverityCode = "5d6"
	case speedTransit:
		sp.val = "Transit"
		sp.descr = "A vessel moving in or out between jump point and orbital space will normally be moving at 'transit' speed."
		sp.cumulativeTrust = 64
		sp.deverityCode = "6d6"
	case speedFastTransit:
		sp.val = "Fast Transit"
		sp.descr = "The vessel has been accelerating hard for some time, or is arriving from a fast interplanetary transit without slowing down. A system defence craft boosting to intercept a contact, or a ship making an urgent run (such as a smuggler or fast courier) might be travelling at this rate."
		sp.cumulativeTrust = 128
		sp.deverityCode = "7d6"
	case speedExtreme:
		sp.val = "Extreme"
		sp.descr = "The vessel is going far too fast to enter orbit without a hard and sustained braking manoeuvre. Any ship going this fast in orbital space will attract both attention and alarm."
		sp.cumulativeTrust = 256
		sp.deverityCode = "8d6"
	}
	return sp
}

type Landing struct {
	uwp        string
	speed      speed
	pilotDM    map[int]int
	severityDM int
	difficulty int
	descr      string
}

func (l *Landing) String() string {
	dm := 0
	for _, v := range l.pilotDM {
		dm += v
	}

	return fmt.Sprintf("\n%v, DM[%v]", newTask(l.descr, l.difficulty).String(), dm)
}

type Port interface {
	MW_Name() string
	MW_UWP() string
}

func Preapare(port Port) (*Landing, error) {
	l := Landing{}
	l.uwp = port.MW_UWP()
	l.pilotDM = make(map[int]int)
	l.evaluatePlanetaryConditions()
	l.pilotDM[shiptype] = shipTypeDM()

	l.setDifficulty()

	return &l, nil
}

func (l *Landing) setDifficulty() {
	answer := ""
	prompt := &survey.Select{
		Message: "This is",
		Options: []string{"Emergency crash-landing", "Heavy landing", "Standard landing", "Smooth landing", "Perfect landing"},
	}
	valid := survey.ComposeValidators()
	survey.AskOne(prompt, &answer, valid)
	l.descr = answer
	switch answer {
	case "Emergency crash-landing":
		l.difficulty = 4
	case "Heavy landing":
		l.difficulty = 6
	case "Standard landing":
		l.difficulty = 8
	case "Smooth landing":
		l.difficulty = 10
	case "Perfect landing":
		l.difficulty = 12
	}
}

func (l *Landing) evaluatePlanetaryConditions() {
	uwp, _ := uwp.FromString0(l.uwp)
	st := uwp.Data(profile.KEY_PORT).Code()
	sz := uwp.Data(profile.KEY_SIZE).Value()
	at := uwp.Data(profile.KEY_ATMO).Value()
	hd := uwp.Data(profile.KEY_HYDR).Value()
	stDM := 0
	switch st {
	case "A", "B":
		stDM = 2
		fmt.Println("Starport Guidence System ACTIVE (+2)")
	case "C", "D":
		fmt.Println("Starport Guidence System PASSIVE (+0)")
	case "E", "X":
		stDM = -2
		fmt.Println("Starport Guidence System ABSENT (-2)")
	}
	l.pilotDM[worldStarport] = stDM
	sizeDM := (-1 * sz) + 9
	if sz >= 0 {
		sizeDM = 0
	}
	fmt.Printf("World Size (%v)\n", sizeDM)
	l.pilotDM[worldSize] = sizeDM
	atmDM := 0
	switch at {
	case 6, 7:
		atmDM = -1
		fmt.Println("Standard Atmosphere (-1)")
	case 8, 9:
		atmDM = -2
		fmt.Println("Dense Atmosphere (-2)")
	case 13:
		atmDM = -3
		fmt.Println("Very Dense Atmosphere (-3)")
	}
	l.pilotDM[worldAtmo] = atmDM
	wx := 0
	if at > 0 {
		wx = (at * hd) / sz
	}
	if wx > 2 {
		dp := dice.New()
		r := dp.Roll("2d6").DM(-1 * wx).Sum()
		if r < 0 {
			l.pilotDM[weather] = r / 2
		}
		if l.pilotDM[weather] > 0 {
			l.pilotDM[weather] = 0
		}
	}
	switch l.pilotDM[weather] {
	case 0:
		fmt.Println("Weather is good or not affecting the landing...")
	case -1:
		fmt.Println("Some Minor Weather Conditions... (-1)")
	case -2:
		fmt.Println("Some Serious Weather Conditions... (-2)")
	case -3:
		fmt.Println("Some Violent Weather Conditions... (-3)")
	case -4:
		fmt.Println("Some Extereme Weather Conditions... (-4)")
	default:
		fmt.Printf("Some DEADLY Weather Conditions... (%v)", l.pilotDM[weather])
	}
}

func shipTypeDM() int {
	answer := ""
	prompt := &survey.Select{
		Message: "Starship Streamlined?",
		Options: []string{"Yes      (DM:  0)", "Partialy (DM: -2)", "No       (DM: -4)"},
	}
	valid := survey.ComposeValidators()
	survey.AskOne(prompt, &answer, valid)
	dm := 0
	switch answer {
	case "Partialy (DM: -2)":
		dm = -2
	case "No       (DM: -4)":
		dm = -4
	}
	return dm
}

type task struct {
	name  string
	skill string
	atr   string
	dif   int
}

func newTask(name string, dif int) task {
	t := task{}
	t.name = name
	t.dif = dif
	return t
}

func (t task) String() string {
	difDsc := ""
	switch t.dif {
	case 4:
		difDsc = "Easy (4+)"
	case 6:
		difDsc = "Routine (6+)"
	case 8:
		difDsc = "Average (8+)"
	case 10:
		difDsc = "Difficult (10+)"
	case 12:
		difDsc = "Very Difficult (12+)"

	}
	return fmt.Sprintf("%v: %v Pilot check (1D minutes, DEX)", t.name, difDsc)
}
