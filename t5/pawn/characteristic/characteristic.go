package characteristic

import "fmt"

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
	TYPE_PHYSICAL  = "Physical"
	TYPE_MENTAL    = "Mental"
	TYPE_SOCICAL   = "Social"
	TYPE_OBSCURE   = "Obscure"
)

func New(name string) *Frame {
	chr := Frame{}
	switch name {
	default:
		return nil
	case Strength:
		chr = Frame{
			positionCode:   "C1",
			abb:            "Str",
			name:           name,
			human:          true,
			description:    "physical power",
			geneticProfile: "S",
			charType:       TYPE_PHYSICAL,
		}
	case Dexterity:
		chr = Frame{
			positionCode:   "C2",
			abb:            "Dex",
			name:           name,
			human:          true,
			description:    "hand-eye co-ordination",
			geneticProfile: "D",
			charType:       TYPE_PHYSICAL,
		}
	case Agility:
		chr = Frame{
			positionCode:   "C2",
			abb:            "Agi",
			name:           name,
			human:          false,
			description:    "body co-ordination",
			geneticProfile: "A",
			charType:       TYPE_PHYSICAL,
		}
	case Grace:
		chr = Frame{
			positionCode:   "C2",
			abb:            "Gra",
			name:           name,
			human:          false,
			description:    "body-limb co-ordination",
			geneticProfile: "G",
			charType:       TYPE_PHYSICAL,
		}
	case Endurance:
		chr = Frame{
			positionCode:   "C3",
			abb:            "End",
			name:           name,
			human:          true,
			description:    "resistance to fatigue",
			geneticProfile: "E",
			charType:       TYPE_PHYSICAL,
		}
	case Stamina:
		chr = Frame{
			positionCode:   "C3",
			abb:            "Sta",
			name:           name,
			human:          false,
			description:    "long-term task persistence",
			geneticProfile: "S",
			charType:       TYPE_PHYSICAL,
		}
	case Vigor:
		chr = Frame{
			positionCode:   "C3",
			abb:            "Vig",
			name:           name,
			human:          false,
			description:    "short-term fatigue resistance",
			geneticProfile: "V",
			charType:       TYPE_PHYSICAL,
		}
	case Intelligence:
		chr = Frame{
			positionCode:   "C4",
			abb:            "Int",
			name:           name,
			human:          true,
			description:    "ability to think and reason",
			geneticProfile: "I",
			charType:       TYPE_MENTAL,
		}
	case Education:
		chr = Frame{
			positionCode:   "C5",
			abb:            "Edu",
			name:           name,
			human:          true,
			description:    "achievement level in school",
			geneticProfile: "E",
			charType:       TYPE_MENTAL,
		}
	case Training:
		chr = Frame{
			positionCode:   "C5",
			abb:            "Tra",
			name:           name,
			human:          false,
			description:    "based on cultural heritage",
			geneticProfile: "T",
			charType:       TYPE_MENTAL,
		}
	case Instinct:
		chr = Frame{
			positionCode:   "C5",
			abb:            "Ins",
			name:           name,
			human:          false,
			description:    "based on genetic heritage",
			geneticProfile: "I",
			charType:       TYPE_MENTAL,
		}
	case SocialStanding:
		chr = Frame{
			positionCode:   "C6",
			abb:            "Soc",
			name:           name,
			human:          true,
			description:    "large group hierarchy",
			geneticProfile: "S",
			charType:       TYPE_SOCICAL,
		}
	case Charisma:
		chr = Frame{
			positionCode:   "C6",
			abb:            "Cha",
			name:           name,
			human:          false,
			description:    "small group hierarchy",
			geneticProfile: "C",
			charType:       TYPE_SOCICAL,
		}
	case Caste:
		chr = Frame{
			positionCode:   "C6",
			abb:            "Cas",
			name:           name,
			human:          false,
			description:    "genetic group hierarchy",
			geneticProfile: "K",
			charType:       SocialStanding,
		}
	case Sanity:
		chr = Frame{
			positionCode:   "CS",
			abb:            "San",
			name:           name,
			human:          true,
			description:    "mental health and stability",
			geneticProfile: "S",
			charType:       TYPE_OBSCURE,
		}
	case Psionics:
		chr = Frame{
			positionCode:   "SP",
			abb:            "Psi",
			name:           name,
			human:          true,
			description:    "extra-sensory mental power",
			geneticProfile: "P",
			charType:       TYPE_OBSCURE,
		}
	}
	return &chr
}

func ByGeneticProfile(code, geneticProfile string) (*Frame, error) {
	switch code {
	default:
		return nil, fmt.Errorf("unexpected code = '%v'", code)
	case "C1":
		if geneticProfile == "S" {
			return New(Strength), nil
		}
	case "C2":
		if geneticProfile == "D" {
			return New(Dexterity), nil
		}
		if geneticProfile == "A" {
			return New(Agility), nil
		}
		if geneticProfile == "G" {
			return New(Grace), nil
		}
	case "C3":
		if geneticProfile == "E" {
			return New(Endurance), nil
		}
		if geneticProfile == "S" {
			return New(Stamina), nil
		}
		if geneticProfile == "V" {
			return New(Vigor), nil
		}
	case "C4":
		if geneticProfile == "I" {
			return New(Intelligence), nil
		}
	case "C5":
		if geneticProfile == "E" {
			return New(Education), nil
		}
		if geneticProfile == "T" {
			return New(Training), nil
		}
		if geneticProfile == "I" {
			return New(Instinct), nil
		}
	case "C6":
		if geneticProfile == "S" {
			return New(SocialStanding), nil
		}
		if geneticProfile == "C" {
			return New(Charisma), nil
		}
		if geneticProfile == "K" {
			return New(Caste), nil
		}
	case "CS":
		return New(Sanity), nil
	case "CP":
		return New(Psionics), nil

	}
	return nil, fmt.Errorf("unexpected combination ('%v' , '%v')", code, geneticProfile)
}
