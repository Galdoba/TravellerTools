package character

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/hostile/character/characteristic"
	"github.com/Galdoba/TravellerTools/hostile/character/skill"
	"github.com/Galdoba/TravellerTools/pkg/decidion"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	KeyManual = "MANUAL"
	ValTrue   = "true"
	ValFalse  = "false"
	KeySeed   = "SEED"
	KeyUPP    = "UNIVERSAL PERSONALITY PROFILE"
)

type Character struct {
	Name      string
	PC        bool
	Homeworld string
	Age       int
	// Career         career.CareerPath
	CharSet  *characteristic.CharSet
	SkillSet skill.SkillSet
}

func NewCharacter() *Character {
	ch := Character{}

	return &ch
}

func (ch *Character) setAsPC() {
	ch.PC = true
}

type generator struct {
	dice    *dice.Dicepool
	options map[string]string
}

type option struct {
	key string
	val string
}

func Option(key, val string) option {
	return option{key, val}
}

func NewGenerator(options ...option) *generator {
	g := generator{}
	g.dice = dice.New()
	g.options = make(map[string]string)
	for _, opt := range options {
		g.options[opt.key] = opt.val
		if opt.key == KeySeed {
			g.dice.SetSeed(opt.val)
		}
	}
	return &g
}

func (g *generator) Generate() (*Character, error) {
	ch := NewCharacter()
	if _, ok := g.options[KeyManual]; ok {
		ch.setAsPC()
	}
	ch.RollCharacteristics(g.dice, g.options)
	ch.DetermineHomeworld(g.dice, g.options)
	ch.ChooseBackgroundSkills(g.dice, g.options)
	// ch.CareerCycle(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	return ch, nil
}

func (ch *Character) RollCharacteristics(dice *dice.Dicepool, options map[string]string) error {
	chrSet, err := characteristic.NewCharSet(characteristic.Human()...)

	if err != nil {
		return err
	}
	ch.CharSet = chrSet
	if val, ok := options[KeyUPP]; ok {
		if val == "" {
			return fmt.Errorf("can't roll characteristics: option %v met but not set", KeyUPP)
		}
		if len(val) != 6 {
			return fmt.Errorf("can't roll characteristics: option %v is invalid '%v'", KeyUPP, val)
		}
		for i, v := range strings.Split(val, "") {
			chr, err := characteristic.New(i)
			if err != nil {
				return err
			}
			chr.Maximum = ehex.New().Set(v)
			ch.CharSet.Chars[i] = chr
		}
	} else {
		for i := range ch.CharSet.Chars {
			ch.CharSet.Chars[i].Roll(dice)
		}

	}
	return nil
}

func (ch *Character) DetermineHomeworld(dice *dice.Dicepool, options map[string]string) error {
	switch ch.PC {
	case false:
		ch.Homeworld = "Earth"
		switch dice.Sroll("2d6") {
		case 9, 10, 11, 12:
			ch.Homeworld = "Colony (" + decidion.Random_One(dice, "Abyss", "Armstrong", "Columbia", "Defiance", "Hamilton") + ")"
		}
	}
	return nil
}

func (ch *Character) ChooseBackgroundSkills(dice *dice.Dicepool, options map[string]string) error {
	skillPool := []string{
		skill.SkillStr(skill.Administration) + " 0",
		skill.SkillStr(skill.Agriculture) + " 0",
		skill.SkillStr(skill.Comms) + " 0",
		skill.SkillStr(skill.Computer) + " 0",
		skill.SkillStr(skill.Electronics) + " 0",
		skill.SkillStr(skill.Engineering) + " 0",
		skill.SkillStr(skill.Gambling) + " 0",
		skill.SkillStr(skill.Investigate) + " 0",
		skill.SkillStr(skill.Liason) + " 0",
		skill.SkillStr(skill.Mechanical) + " 0",
		skill.SkillStr(skill.Medical) + " 0",
		skill.SkillStr(skill.Steward) + " 0",
		skill.SkillStr(skill.Survival) + " 0",
	}
	switch ch.Homeworld {
	case "Earth":
		skillPool = append(skillPool, skill.SkillStr(skill.Ground_Vechicle)+" 0")
		skillPool = append(skillPool, skill.SkillStr(skill.Brawling)+" 0")
		skillPool = append(skillPool, skill.SkillStr(skill.Streetwise)+" 0")
		skillPool = append(skillPool, skill.SkillStr(skill.Carousing)+" 0")
	default:
		skillPool = append(skillPool, skill.SkillStr(skill.Vacc_Suit)+" 0")
		skillPool = append(skillPool, skill.SkillStr(skill.Survival)+" 0")
		skillPool = append(skillPool, skill.SkillStr(skill.Brawling)+" 0")
		skillPool = append(skillPool, skill.SkillStr(skill.Vechicle)+" 0")
	}
	n := ch.CharSet.Chars[characteristic.EDU].Mod() + 3
	skillsChosen := []string{}
	switch ch.PC {
	case false:
		skillsChosen, skillPool = decidion.Random_Few_Exclude(n, dice, skillPool...)

	}
	fmt.Println(skillsChosen)
	return nil
}

func (ch *Character) gain(bonus string) error {
	return nil
}

func (ch *Character) RunTerm() error {
	return nil
}
