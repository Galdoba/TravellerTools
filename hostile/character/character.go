package character

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/hostile/character/career"
	"github.com/Galdoba/TravellerTools/hostile/character/characteristic"
	"github.com/Galdoba/TravellerTools/hostile/character/skill"
	"github.com/Galdoba/TravellerTools/pkg/decidion"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	KeyManual                    = "MANUAL"
	ValTrue                      = "true"
	ValFalse                     = "false"
	KeySeed                      = "SEED"
	KeyUPP                       = "UNIVERSAL PERSONALITY PROFILE"
	KeyQualificationRetryAllowed = "QUALIFICATION RETRY ATEMPTS ALLOWED"
)

type Character struct {
	Name              string
	PC                bool
	Homeworld         string
	Age               int
	CommisionReceived bool
	Career            career.CareerState
	CharSet           *characteristic.CharSet
	SkillSet          *skill.SkillSet
	Benefits          []string
	drafted           bool
}

func NewCharacter() *Character {
	ch := Character{}
	ch.SkillSet = skill.NewSkillSet()
	return &ch
}

func (ch *Character) setAsPC() {
	ch.PC = true
}

type generator struct {
	options map[string]string
}

type option struct {
	key string
	val string
}

func Option(key, val string) option {
	return option{key, val}
}

var DICE *dice.Dicepool

func NewGenerator(options ...option) *generator {
	DICE = dice.New()
	g := generator{}
	g.options = make(map[string]string)
	for _, opt := range options {
		g.options[opt.key] = opt.val
		if opt.key == KeySeed {
			DICE.SetSeed(opt.val)
		}
	}
	return &g
}

func (g *generator) Generate() (*Character, error) {
	ch := NewCharacter()
	if _, ok := g.options[KeyManual]; ok {
		ch.setAsPC()
	}
	if err := ch.RollCharacteristics(g.options); err != nil {
		return ch, err
	}
	ch.DetermineHomeworld(g.options)
	ch.ChooseBackgroundSkills(g.options)
	if err := ch.ChooseAndStartCareer(g.options); err != nil {
		return ch, err
	}
	if err := ch.CareerCycle(g.options); err != nil {
		return ch, err
	}
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	// ch.RollCharacteristics(g.dice, g.options)
	fmt.Println(ch.Career.Report())
	return ch, nil
}

func (ch *Character) RollCharacteristics(options map[string]string) error {
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
			ch.CharSet.Chars[i].Roll(DICE)
		}

	}
	return nil
}

func (ch *Character) DetermineHomeworld(options map[string]string) error {
	switch ch.PC {
	case false:
		ch.Homeworld = "Earth"
		switch DICE.Sroll("2d6") {
		case 9, 10, 11, 12:
			ch.Homeworld = "Colony (" + decidion.Random_One(DICE, "Abyss", "Armstrong", "Columbia", "Defiance", "Hamilton") + ")"
		}
	}
	return nil
}

func (ch *Character) ChooseBackgroundSkills(options map[string]string) error {
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
		skillsChosen, skillPool = decidion.Random_Few_Exclude(n, DICE, skillPool...)

	}
	for i := range skillsChosen {
		if strings.HasPrefix(skillsChosen[i], skill.SkillStr(skill.Vechicle)) {
			_, newSkl := ch.chooseCascadSkill()
			skillsChosen[i] = strings.ReplaceAll(skillsChosen[i], skill.SkillStr(skill.Vechicle), newSkl)
		}
		if err := ch.gain(skillsChosen[i]); err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(skillsChosen)
	return nil
}

func (ch *Character) gain(bonus string) error {
	if strings.Contains(bonus, " OR ") {
		return fmt.Errorf("can't gain bonus: %v must be splitted by options", bonus)
	}
	if strings.Contains(bonus, " AND ") {
		return fmt.Errorf("can't gain bonus: %v must be concatenated by options", bonus)
	}
	bonuses := append([]string{}, bonus)
	for _, bonus := range bonuses {
		if err := ch.receiveBonus(bonus); err != nil {
			return err
		}
	}
	return nil
}

func (ch *Character) RunTerm() error {
	return nil
}

func (ch *Character) chooseCascadSkill() (int, string) {
	str := ""
	i := 0
	switch ch.PC {
	case false:
		str = decidion.Random_One(DICE, skill.SkillStr(skill.Aircraft), skill.SkillStr(skill.Ground_Vechicle), skill.SkillStr(skill.Watercraft))
	}
	switch str {
	case skill.SkillStr(skill.Aircraft):
		i = skill.Aircraft
	case skill.SkillStr(skill.Ground_Vechicle):
		i = skill.Ground_Vechicle
	case skill.SkillStr(skill.Watercraft):
		i = skill.Watercraft
	}
	fmt.Println("cascad skill chosen:", i, str)
	return i, str
}

func (ch *Character) ChooseAndStartCareer(options map[string]string) error {
	careerName := ""
	careerList := []string{
		career.Colonist,
		career.CorporateAgent,
		career.CorporateExec,
		career.CommersialSpacer,
		career.Marine,
		career.Marshal,
		career.MilitarySpacer,
		career.Physician,
		career.Ranger,
		career.Rogue,
		career.Roughneck,
		career.Scientist,
		career.SurveyScout,
		career.Technician,
	}
	qra, _ := strconv.Atoi(options[KeyQualificationRetryAllowed])
	atemptsAllowed := 1 + qra
	for i := 0; i < atemptsAllowed; i++ {
		switch ch.PC {
		case false:
			careerName, careerList = decidion.Random_One_Exclude(DICE, careerList...)
		}
		crr, err := career.StartCareer(careerName, DICE, ch.CharSet, false)
		if err != nil {
			return fmt.Errorf("can't start career: %v", err.Error())
		}
		if crr.Qualify(DICE, ch.CharSet) {
			ch.Career = crr
			return nil
		}
	}
	fmt.Println("Draft")
	r := DICE.Sroll("1d6")
	switch r {
	case 1:
		careerName = career.Ranger
	case 2, 3, 4:
		careerName = career.Colonist
	case 5, 6:
		careerName = career.Roughneck
	}
	fmt.Println("drafted to", careerName)
	crr, err := career.StartCareer(careerName, DICE, ch.CharSet, true)
	if err != nil {
		return fmt.Errorf("can't start career by draft: %v", err)
	}
	ch.drafted = true
	ch.Career = crr
	return nil
}

func (ch *Character) CareerCycle(options map[string]string) error {
	career := ch.Career
	fmt.Println(career.Report())
	inCareer := true
	term := 1
	for inCareer {
		fmt.Printf("start term %v\n", term)
		//background
		if term == 1 {
			// fmt.Printf("background skill benefits on term %v\n", term)
			if err := ch.gain(career.Train(DICE, ch.PC)); err != nil {
				return err
			}

			if err := ch.gain(career.Train(DICE, ch.PC)); err != nil {
				return err
			}
		}

		//survival
		// fmt.Printf("survival on term %v\n", term)
		if !career.Survived(DICE, ch.CharSet) {
			return fmt.Errorf("not survived on term %v", term)
		}
		//sturdy
		if term != 1 {
			fmt.Printf("study on term %v\n", term)

			if err := ch.gain(career.Train(DICE, ch.PC)); err != nil {
				return err
			}
		}
		//commision
		if !ch.CommisionReceived && career.CommisionReceived(DICE, ch.CharSet) {

			fmt.Printf("commision RECEIVED on term %v\n", term)
			ch.CommisionReceived = true
			if err := ch.gain(career.Train(DICE, ch.PC)); err != nil {
				return err
			}
		}
		//advance
		fmt.Printf("advancement on term %v\n", term)

		//reenlist
		fmt.Printf("re-enlist after term %v\n", term)
		if term == 5 {
			break
		}
		term++
	}
	return nil
}

func (ch *Character) receiveBonus(bonus string) error {
	id, val := skill.FromText(bonus)
	if id == skill.Vechicle {
		id, _ = ch.chooseCascadSkill()
	}
	if id != -1 && val != -1 {
		return ch.SkillSet.Increase(id)
	}
	if id != -1 && (val == 0 || val == 1) {
		return ch.SkillSet.Gain(fmt.Sprintf("%v %v", skill.SkillStr(id), val))
	}
	id, val = characteristic.FromText(bonus)
	switch id {
	case characteristic.STR, characteristic.DEX, characteristic.END, characteristic.INT, characteristic.EDU, characteristic.SOC, characteristic.INST:
		ch.CharSet.Chars[id].ChangeMaximumBy(val)
		return nil
	}
	switch bonus {
	default:
		panic("Benefit " + bonus)

	}
	ch.Benefits = append(ch.Benefits, bonus)
	return nil
}
