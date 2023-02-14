package characteristic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	GP_Human = "SDEIES"
)

type Frame struct {
	positionCode   string
	abb            string
	name           string
	human          bool //human or analog
	description    string
	geneticProfile string
	charType       string
	generationDice int //genes
	value          int
	geneticValue   int
	mods           []mod
	//generationDM   int
}

type mod struct {
	name      string
	descr     string
	influence int
}

func (c *Frame) Code() string {
	return c.positionCode
}

func (c *Frame) Name() string {
	return c.name
}

func (c *Frame) Abb() string {
	return c.abb
}

func (c *Frame) GP() string {
	return c.geneticProfile
}

func (c *Frame) Genes() int {
	return c.generationDice
}

func (c *Frame) Val() int {
	return c.value
}

func (c *Frame) Actual() int {
	sum := c.value
	for _, v := range c.mods {
		sum += v.influence
	}
	return sum
}

func (c *Frame) SetValue(v int) {
	c.value = v
}

func (c *Frame) SetGene(v int) {
	c.generationDice = v
}

func (c *Frame) RollValue(dice *dice.Dicepool) error {
	r := c.generationDice
	c.geneticValue = dice.Sroll("1d6")
	switch r {
	default:
		return fmt.Errorf("generation dice invalid")
	case 1:
		c.value = c.geneticValue
	case 2:
		c.value = c.geneticValue + dice.Sroll("1d6")
	case 3:
		c.value = c.geneticValue + dice.Sroll("2d6")
	case 4:
		c.value = c.geneticValue + dice.Sroll("1d6") + 12
	case 5:
		c.value = c.geneticValue + dice.Sroll("2d6") + 12
	case 6:
		c.value = c.geneticValue + dice.Sroll("3d6") + 12
	case 7:
		c.value = c.geneticValue + dice.Sroll("4d6") + 12
	case 8:
		c.value = c.geneticValue + dice.Sroll("5d6") + 12
	}
	return nil
}

const (
	Strength       = "Strength"
	Dexterity      = "Dexterity"
	Agility        = "Agility"
	Grace          = "Grace"
	Endurance      = "Endurance"
	Stamina        = "Stamina"
	Vigor          = "Vigor"
	Intelligence   = "Intelligence"
	Education      = "Education"
	Training       = "Training"
	Instinct       = "Instinct"
	SocialStanding = "Social Standing"
	Charisma       = "Charisma"
	Caste          = "Caste"
	Sanity         = "Sanity"
	Psionics       = "Psionics"
	PseudoCHR      = "Pseudo Characteristic"
	TYPE_PHYSICAL  = "Physical"
	TYPE_MENTAL    = "Mental"
	TYPE_SOCICAL   = "Social"
	TYPE_OBSCURE   = "Obscure"
)

func New(name string, genDice int) *Frame {
	chr := Frame{}
	chr.name = name
	chr.generationDice = genDice
	switch name {
	default:
		return nil
	case Strength:
		chr.positionCode = "C1"
		chr.abb = "Str"
		chr.human = true
		chr.description = "physical power"
		chr.geneticProfile = "S"
		chr.charType = TYPE_PHYSICAL
	case Dexterity:
		chr.positionCode = "C2"
		chr.abb = "Dex"
		chr.human = true
		chr.description = "hand-eye co-ordination"
		chr.geneticProfile = "D"
		chr.charType = TYPE_PHYSICAL
	case Agility:
		chr.positionCode = "C2"
		chr.abb = "Agi"
		chr.human = false
		chr.description = "body co-ordination"
		chr.geneticProfile = "A"
		chr.charType = TYPE_PHYSICAL
	case Grace:
		chr.positionCode = "C2"
		chr.abb = "Gra"
		chr.human = false
		chr.description = "body-limb co-ordination"
		chr.geneticProfile = "G"
		chr.charType = TYPE_PHYSICAL
	case Endurance:
		chr.positionCode = "C3"
		chr.abb = "End"
		chr.human = true
		chr.description = "resistance to fatigue"
		chr.geneticProfile = "E"
		chr.charType = TYPE_PHYSICAL
	case Stamina:
		chr.positionCode = "C3"
		chr.abb = "Sta"
		chr.human = false
		chr.description = "long-term task persistence"
		chr.geneticProfile = "S"
		chr.charType = TYPE_PHYSICAL
	case Vigor:
		chr.positionCode = "C3"
		chr.abb = "Vig"
		chr.human = false
		chr.description = "short-term fatigue resistance"
		chr.geneticProfile = "V"
		chr.charType = TYPE_PHYSICAL
	case Intelligence:
		chr.positionCode = "C4"
		chr.abb = "Int"
		chr.human = true
		chr.description = "ability to think and reason"
		chr.geneticProfile = "I"
		chr.charType = TYPE_MENTAL
	case Education:
		chr.positionCode = "C5"
		chr.abb = "Edu"
		chr.human = true
		chr.description = "achievement level in school"
		chr.geneticProfile = "E"
		chr.charType = TYPE_MENTAL
	case Training:
		chr.positionCode = "C5"
		chr.abb = "Tra"
		chr.human = false
		chr.description = "based on cultural heritage"
		chr.geneticProfile = "T"
		chr.charType = TYPE_MENTAL
	case Instinct:
		chr.positionCode = "C5"
		chr.abb = "Ins"
		chr.human = false
		chr.description = "based on genetic heritage"
		chr.geneticProfile = "I"
		chr.charType = TYPE_MENTAL
	case SocialStanding:
		chr.positionCode = "C6"
		chr.abb = "Soc"
		chr.human = true
		chr.description = "large group hierarchy"
		chr.geneticProfile = "S"
		chr.charType = TYPE_SOCICAL
	case Charisma:
		chr.positionCode = "C6"
		chr.abb = "Cha"
		chr.human = false
		chr.description = "small group hierarchy"
		chr.geneticProfile = "C"
		chr.charType = TYPE_SOCICAL
	case Caste:
		chr.positionCode = "C6"
		chr.abb = "Cas"
		chr.human = false
		chr.description = "genetic group hierarchy"
		chr.geneticProfile = "K"
		chr.charType = SocialStanding
	case Sanity:
		chr.positionCode = "CS"
		chr.abb = "San"
		chr.human = true
		chr.description = "mental health and stability"
		chr.geneticProfile = "S"
		chr.charType = TYPE_OBSCURE
	case Psionics:
		chr.positionCode = "SP"
		chr.abb = "Psi"
		chr.human = true
		chr.description = "extra-sensory mental power"
		chr.geneticProfile = "P"
		chr.charType = TYPE_OBSCURE
	case PseudoCHR:
		chr.positionCode = ""
		chr.abb = ""
		chr.human = true
		chr.description = "Pseudo characteristic"
		chr.geneticProfile = "?"
		chr.charType = TYPE_OBSCURE
		chr.generationDice = 0
		chr.value = 7
	}
	return &chr
}

type genDice struct {
	dice []int
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

func GeneticCodeToName(gen, code string) string {
	switch code {
	case "C1":
		return Strength
	case "C2":
		switch gen {
		case "A":
			return Agility
		case "D":
			return Dexterity
		case "G":
			return Grace
		}
	case "C3":
		switch gen {
		case "S":
			return Stamina
		case "E":
			return Endurance
		case "V":
			return Vigor
		}
	case "C4":
		return Intelligence
	case "C5":
		switch gen {
		case "I":
			return Instinct
		case "E":
			return Education
		case "T":
			return Training
		}
	case "C6":
		switch gen {
		case "K":
			return Caste
		case "S":
			return SocialStanding
		case "C":
			return Charisma
		}
	}
	return "ERROR: bad data (" + gen + " , " + code + ")"
}

func ByGeneticProfile(code, geneticProfile, genMap string) (*Frame, error) {
	genDiceArr := newGenMap(genMap)
	pos, _ := strconv.Atoi(strings.TrimPrefix(code, "C"))
	switch code {
	default:
		return nil, fmt.Errorf("unexpected code = '%v'", code)
	case "C1":
		if geneticProfile == "S" {
			return New(Strength, genDiceArr[pos]), nil
		}
	case "C2":
		if geneticProfile == "D" {
			return New(Dexterity, genDiceArr[pos]), nil
		}
		if geneticProfile == "A" {
			return New(Agility, genDiceArr[pos]), nil
		}
		if geneticProfile == "G" {
			return New(Grace, genDiceArr[pos]), nil
		}
	case "C3":
		if geneticProfile == "E" {
			return New(Endurance, genDiceArr[pos]), nil
		}
		if geneticProfile == "S" {
			return New(Stamina, genDiceArr[pos]), nil
		}
		if geneticProfile == "V" {
			return New(Vigor, genDiceArr[pos]), nil
		}
	case "C4":
		if geneticProfile == "I" {
			return New(Intelligence, genDiceArr[pos]), nil
		}
	case "C5":
		if geneticProfile == "E" {
			return New(Education, genDiceArr[pos]), nil
		}
		if geneticProfile == "T" {
			return New(Training, genDiceArr[pos]), nil
		}
		if geneticProfile == "I" {
			return New(Instinct, genDiceArr[pos]), nil
		}
	case "C6":
		if geneticProfile == "S" {
			return New(SocialStanding, genDiceArr[pos]), nil
		}
		if geneticProfile == "C" {
			return New(Charisma, genDiceArr[pos]), nil
		}
		if geneticProfile == "K" {
			return New(Caste, genDiceArr[pos]), nil
		}
	case "CS":
		return New(Sanity, genDiceArr[pos]), nil
	case "CP":
		return New(Psionics, genDiceArr[pos]), nil

	}
	return nil, fmt.Errorf("unexpected combination ('%v' , '%v')", code, geneticProfile)
}

func (chr *Frame) CheckAverage(dice *dice.Dicepool) {

}
