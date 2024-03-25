package character

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galdoba/TravellerTools/hostile/character/career"
	"github.com/Galdoba/TravellerTools/hostile/character/characteristic"
	"github.com/Galdoba/TravellerTools/hostile/character/skill"
	"github.com/Galdoba/TravellerTools/pkg/decidion"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/terminal"
)

const (
	KeyManual                    = "MANUAL"
	ValTrue                      = "true"
	ValFalse                     = "false"
	KeySeed                      = "SEED"
	KeyUPP                       = "UNIVERSAL PERSONALITY PROFILE"
	KeyQualificationRetryAllowed = "QUALIFICATION RETRY ATEMPTS ALLOWED"
	advanceCO                    = "Advance After Commision"
	advanceNCO                   = "Advance NCO Ranks"
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
	Balance           int
	TotalTerms        int
	drafted           bool
	nextEvent         string
}

func NewCharacter() *Character {
	ch := Character{}
	ch.SkillSet = skill.NewSkillSet()
	ch.Age = 18
	ch.Name = "NO NAME"
	ch.nextEvent = EVENT_RollCharacteristics
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
	ch.FlushScreen()
	return ch, nil
}

func (ch *Character) ConsumeBenefits() error {
	ch.Benefits, ch.Balance = confirmBenefits(ch.Benefits)
	for _, benefit := range ch.Benefits {
		if err := ch.gain(benefit); err != nil {
			return fmt.Errorf("gain benefit: %v", err.Error())
		}
	}
	ch.nextEvent = EVENT_EndGeneration
	return nil
}

func confirmBenefits(bnfts []string) ([]string, int) {
	newList := []string{}
	cash := 0
	for _, b := range bnfts {
		switch b {
		default:
			newList = append(newList, b)
		case "$500":
			cash += 500
		case "$1000":
			cash += 1000
		case "$5000":
			cash += 5000
		case "$8000":
			cash += 8000
		case "$10000":
			cash += 10000
		case "$20000":
			cash += 20000
		case "Award":
			newList = append(newList, "Award")
			newList = append(newList, "+1 SOC")
		}
	}
	return newList, cash
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
		keep := false
		for !keep {
			fmt.Println(" STR  DEX  END  INT  EDU  SOC ")
			for i := range ch.CharSet.Chars {
				ch.CharSet.Chars[i].Roll(DICE)
				fmt.Printf("  %v  ", ch.CharSet.Chars[i].Maximum.Code())
				if ch.PC {
					time.Sleep(time.Second)
				}

			}
			fmt.Println("")
			switch ch.PC {
			case true:
				if decidion.Manual_One("Keep characteristics?", "Yes", "No") == "Yes" {
					keep = true
				}
			case false:
				keep = true
			}
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
	case true:
		selected := decidion.Manual_One("Select Homeworld", []string{"Earth", "RANDOM"}...)
		switch selected {
		default:
			ch.Homeworld = selected
		case "RANDOM":
			ch.Homeworld = "Earth"
			switch DICE.Sroll("2d6") {
			case 9, 10, 11, 12:
				ch.Homeworld = "Colony (" + decidion.Random_One(DICE, "Abyss", "Armstrong", "Columbia", "Defiance", "Hamilton") + ")"
			}
		}
	}
	ch.Inform("Homeworld: " + ch.Homeworld)
	ch.nextEvent = EVENT_CHOOSE_BGSKILLS

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
	case true:
		skillsChosen, skillPool = decidion.Manual_Few_Exclude(n, fmt.Sprintf("Select %v background skills", n), skillPool...)
	}
	for i := range skillsChosen {
		if strings.HasPrefix(skillsChosen[i], skill.SkillStr(skill.Vechicle)) {
			_, newSkl := ch.chooseCascadSkill()
			skillsChosen[i] = strings.ReplaceAll(skillsChosen[i], skill.SkillStr(skill.Vechicle), newSkl)
		}
		if err := ch.gain(skillsChosen[i]); err != nil {
			return err
		}
	}
	// fmt.Println(skillsChosen)
	ch.nextEvent = EVENT_CHOOSE_CAREER
	return nil
}

func (ch *Character) gain(bonus string) error {
	if strings.Contains(bonus, " OR ") {
		opts := strings.Split(bonus, " OR ")
		switch ch.PC {
		case false:
			bonus = decidion.Random_One(DICE, opts...)
		case true:
			bonus = decidion.Manual_One("Select bonus", opts...)
		}
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
	// fmt.Println("cascad skill chosen:", i, str)
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
		case true:
			ch.FlushScreen()
			careerName, careerList = decidion.Manual_One_Exclude("Select career", careerList...)
		}
		crr, err := career.StartCareer(careerName, DICE, ch.CharSet, false)
		if err != nil {
			return fmt.Errorf("can't start career: %v", err.Error())
		}
		if crr.Qualify(DICE, ch.CharSet) {
			ch.Inform("Qualification Passed")
			ch.Career = crr
			ch.nextEvent = EVENT_CAREER_CYCLE
			return nil
		}

		ch.Inform("Qualification Failed")
	}
	// fmt.Println("Draft"
	r := DICE.Sroll("1d6")
	switch r {
	case 1:
		careerName = career.Ranger
	case 2, 3, 4:
		careerName = career.Colonist
	case 5, 6:
		careerName = career.Roughneck
	}
	ch.Inform("Drafted to " + careerName)
	// fmt.Println("drafted to", careerName)
	crr, err := career.StartCareer(careerName, DICE, ch.CharSet, true)
	if err != nil {
		return fmt.Errorf("can't start career by draft: %v", err)
	}
	ch.drafted = true
	ch.Career = crr
	ch.nextEvent = EVENT_CAREER_CYCLE
	return nil
}

func (ch *Character) advancedEducation() bool {
	return ch.CharSet.Chars[characteristic.EDU].Maximum.Value() >= 8

}

func (ch *Character) CareerCycle(options map[string]string) error {
	// ch.Career := ch.Career
	inCareer := true
	term := 1
	if err := ch.gain(ch.Career.RankBonus()); err != nil {
		return err
	}
	for inCareer {
		ch.FlushScreen()
		fmt.Printf("Career: %v term %v\n", ch.Career.Name(), term)
		//background
		if term == 1 {
			ch.Inform("Basic Training:")
			if err := ch.gain(ch.Career.Train(DICE, ch.PC, ch.advancedEducation())); err != nil {
				return err
			}
			if err := ch.gain(ch.Career.Train(DICE, ch.PC, ch.advancedEducation())); err != nil {
				return err
			}
			ch.FlushScreen()
		}
		ch.Inform("Survival Roll:")
		if !ch.Career.Survived(DICE, ch.CharSet) {
			ch.Inform("Failed")
			switch DICE.Sroll("1d6") {
			case 1, 2:
				ch.Inform("Injury Roll:")
				ch.nextEvent = EVENT_INJURY
				return nil
			}
			break
		}
		ch.Inform("Success!!")
		ch.FlushScreen()
		advanceType := advanceCO
		if ch.Career.CanAdvance(true) && (ch.Career.Name() == career.MilitarySpacer || ch.Career.Name() == career.Marine) {
			switch ch.PC {
			case false:
				advanceType = decidion.Random_One(DICE, advanceCO, advanceNCO)
			case true:
				advanceType = decidion.Manual_One("Select advancment type:", advanceCO, advanceNCO)
			}
		}
		//CommisionReceived
		switch advanceType {
		case advanceCO:
			if !ch.CommisionReceived && ch.Career.CommisionReceived(DICE, ch.CharSet) { //roll commision if needed
				ch.Inform("Commision Received:")
				ch.CommisionReceived = true
				if err := ch.gain(ch.Career.Train(DICE, ch.PC, ch.advancedEducation())); err != nil {
					return err
				}
			}
			if ch.Career.AdvancementReceived(DICE, ch.CharSet, false) {
				ch.Inform("Advancement Received")
				// fmt.Printf("advancement RECEIVED on term %v\n", term)
				ch.gain(ch.Career.RankBonus())
				if err := ch.gain(ch.Career.Train(DICE, ch.PC, ch.advancedEducation())); err != nil {
					return err
				}
			}
		case advanceNCO:
			if ch.Career.AdvancementReceived(DICE, ch.CharSet, true) {
				// fmt.Printf("advancement RECEIVED on term %v\n", term)
				ch.gain(ch.Career.RankBonus())
				ch.Inform("Advancement Received (NCO)")
				if err := ch.gain(ch.Career.Train(DICE, ch.PC, ch.advancedEducation())); err != nil {
					return err
				}
			}
		}
		//sturdy
		if term != 1 {
			ch.Inform("Training:")
			if err := ch.gain(ch.Career.Train(DICE, ch.PC, ch.advancedEducation())); err != nil {
				return err
			}
		}
		//age
		ch.Age += 4

		ch.FlushScreen()
		ch.Inform(fmt.Sprintf("Character Age = %v", ch.Age))
		agingBorder := 34
		switch ch.Career.Name() {
		case career.CorporateExec:
			agingBorder = 46
		}
		if ch.Age >= agingBorder {
			ch.Inform("Aging:")
			msg, err := ch.CharSet.AgingRoll(DICE, term, ch.PC)
			if err != nil {
				return err
			}
			ch.Inform(msg)
		}
		//reenlist
		// fmt.Printf("re-enlist after term %v\n", term)
		term++
		ch.TotalTerms++
		if ch.TotalTerms >= 7 { //not realy needed
			break
		}
		if !ch.Career.ReEnlisted(DICE, ch.PC) {
			ch.Inform("Re-enlitment Failed")
			break
		}
	}
	ch.FlushScreen()
	ch.nextEvent = EVENT_MUSTER_OUT
	return nil
}

func (ch *Character) MusterOut(options map[string]string) error {
	if len(ch.Benefits) != 0 {
		return fmt.Errorf("Benefits are not NIL")
	}
	ch.Benefits = ch.Career.MusterOut(DICE, ch.SkillSet.SkillVal(skill.Gambling) >= 0, ch.PC)
	ch.nextEvent = EVENT_BENEFITS
	return nil
}

func (ch *Character) receiveBonus(bonus string) error {
	if bonus == "" {
		return nil
	}
	id, val := skill.FromText(bonus)
	if id == skill.Vechicle {
		id, _ = ch.chooseCascadSkill()
	}
	if id != -1 {
		switch val {
		case 0, 1:
			err := ch.SkillSet.Gain(bonus)
			if err != nil {
				return err
			}
		case 999:
			err := ch.SkillSet.Increase(id)
			if err != nil {
				return nil
			}
		default:
			fmt.Println(val)
			panic("+++++++" + bonus)
		}
	}
	if id != -1 && (val == 0 || val == 1) {
		err := ch.SkillSet.Gain(fmt.Sprintf("%v %v", skill.SkillStr(id), val))
		if err != nil {
			return err
		}
	}
	id, val = characteristic.FromText(bonus)
	switch id {
	case characteristic.STR, characteristic.DEX, characteristic.END, characteristic.INT, characteristic.EDU, characteristic.SOC, characteristic.INST:
		ch.CharSet.Chars[id].ChangeMaximumBy(val)
	}
	// ch.FlushScreen()
	ch.Inform(fmt.Sprintf("Received: %v", bonus))
	// ch.Benefits = append(ch.Benefits, bonus)
	return nil
}

func (ch *Character) Form() string {
	form := ""
	form += fmt.Sprintf("|SUBJECT NAME        : %v\n", ch.Name)
	form += fmt.Sprintf("|UNIVERSAL PROFILE   : %v\n", ch.CharSet.String())
	form += fmt.Sprintf("|HOMEWORLD           : %v\n", ch.Homeworld)
	jobTitle := "NONE"
	careerName := "NONE"
	if ch.Career != nil {
		jobTitle = ch.Career.JobTitle()
		careerName = ch.Career.Name() + fmt.Sprintf(" %v", ch.TotalTerms)
	}
	form += fmt.Sprintf("|JOB TITLE           : %v (%v)\n", jobTitle, careerName)
	form += fmt.Sprintf("|AGE           : %v\n", ch.Age)
	list := ch.SkillSet.List()
	if len(list) > 0 {
		form += fmt.Sprintf("|TRAINING AND SKILLS :\n")
		for _, l := range list {
			form += fmt.Sprintf("|    %v\n", l)
		}
	}
	if len(ch.Benefits) > 0 {

		form += fmt.Sprintf("|LIST OF BENEFITS    :\n")
		for i := range ch.Benefits {
			form += fmt.Sprintf("|    %v\n", ch.Benefits[i])
		}
	}
	if ch.Balance != 0 {
		form += fmt.Sprintf("|BALANCE             : $%v\n", ch.Balance)
	}
	return form
}

func (ch *Character) FlushScreen() {
	terminal.Clear()
	fmt.Println(ch.Form())
}

func (ch *Character) Inform(msg string) {
	fmt.Println(msg)
	time.Sleep(time.Second)
}

func (ch *Character) Injury() error {
	err := fmt.Errorf("injury not rolled")
	r := DICE.Sroll("1d6")
	msg := ""
	switch r {
	default:
		return fmt.Errorf("func injury%v() not implemented", r)
	case 1:
		msg, err = ch.injury1()
	case 2:
		msg, err = ch.injury2()
	case 3:
		msg, err = ch.injury3()
	case 4:
		msg, err = ch.injury4()
	case 5:
		msg, err = ch.injury5()
	case 6:
		msg, err = ch.injury6()
	}
	ch.Inform(msg)
	ch.nextEvent = EVENT_MUSTER_OUT
	return err
}

func (ch *Character) injury1() (string, error) {
	charNames := []string{"STR", "DEX", "END"}
	chrName := ""
	vals := []int{DICE.Sroll("1d6"), 2, 2}
	other := []string{}
	msg := "character injured:"
	switch ch.PC {
	case false:
		for i := 0; i < 3; i++ {
			switch i {
			case 0:
				chrName, other = decidion.Random_One_Exclude(DICE, charNames...)
			case 1, 2:
				chrName = decidion.Random_One(DICE, other...)
			}
			id, _ := characteristic.FromText(chrName)
			ch.reduce_characteristic_limited(id, vals[i]*-1)
			msg += fmt.Sprintf(" %v rediced by %v,", chrName, vals[i])
		}
	case true:
		for i := 0; i < 3; i++ {
			switch i {
			case 0:
				chrName, other = decidion.Manual_One_Exclude(fmt.Sprintf("%v of 3: Choose characteristic to reduce by %v", i+1, vals[0]), charNames...)
			case 1, 2:
				chrName = decidion.Manual_One(fmt.Sprintf("%v of 3: Choose characteristic to reduce by %v", i+1, vals[i]), other...)
			}
			id, _ := characteristic.FromText(chrName)
			ch.reduce_characteristic_limited(id, vals[i])
			msg += fmt.Sprintf(" %v rediced by %v,", chrName, vals[i])
		}
	}
	msg = strings.Trim(msg, ",") + "."
	return msg, nil
}

func (ch *Character) injury2() (string, error) {
	charNames := []string{"STR", "DEX", "END"}
	chrName := ""
	val := DICE.Sroll("1d6")
	msg := "character injured:"
	switch ch.PC {
	case false:
		chrName = decidion.Random_One(DICE, charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val*-1)
		msg += fmt.Sprintf(" %v rediced by %v,", chrName, val)
	case true:
		chrName = decidion.Manual_One(fmt.Sprintf("Choose characteristic to reduce by %v", val), charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val)
		msg += fmt.Sprintf(" %v rediced by %v,", chrName, val)
	}
	msg = strings.Trim(msg, ",") + "."
	return msg, nil
}

func (ch *Character) injury3() (string, error) {
	charNames := []string{"STR", "DEX"}
	chrName := ""
	val := 2
	msg := make(map[string]string)
	msg["STR"] = "Missing limb. STR reduced by 2."
	msg["DEX"] = "Missing eye. DEX reduced by 2."
	switch ch.PC {
	case false:
		chrName = decidion.Random_One(DICE, charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val*-1)
	case true:
		chrName = decidion.Manual_One(fmt.Sprintf("Choose characteristic to reduce by %v", val), charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val)
	}
	return msg[chrName], nil
}

func (ch *Character) injury4() (string, error) {
	charNames := []string{"STR", "DEX", "END"}
	chrName := ""
	val := 2
	msg := "Scarred and injured:"
	switch ch.PC {
	case false:
		chrName = decidion.Random_One(DICE, charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val*-1)
		msg += fmt.Sprintf(" %v rediced by %v,", chrName, val)
	case true:
		chrName = decidion.Manual_One(fmt.Sprintf("Choose characteristic to reduce by %v", val), charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val)
		msg += fmt.Sprintf(" %v rediced by %v,", chrName, val)
	}
	msg = strings.Trim(msg, ",") + "."
	return msg, nil
}

func (ch *Character) injury5() (string, error) {
	charNames := []string{"STR", "DEX", "END"}
	chrName := ""
	val := 1
	msg := "Injured:"
	switch ch.PC {
	case false:
		chrName = decidion.Random_One(DICE, charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val*-1)
		msg += fmt.Sprintf(" %v rediced by %v,", chrName, val)
	case true:
		chrName = decidion.Manual_One(fmt.Sprintf("Choose characteristic to reduce by %v", val), charNames...)
		id, _ := characteristic.FromText(chrName)
		ch.reduce_characteristic_limited(id, val)
		msg += fmt.Sprintf(" %v rediced by %v,", chrName, val)
	}
	msg = strings.Trim(msg, ",") + "."
	return msg, nil
}

func (ch *Character) injury6() (string, error) {
	msg := "Lightly injured: No permanent effect"
	return msg, nil
}

func (ch *Character) reduce_characteristic_limited(id, n int) {
	ch.CharSet.Chars[id].ChangeMaximumBy(-1 * n)
	if ch.CharSet.Chars[id].Maximum.Value() == 0 {
		ch.CharSet.Chars[id].ChangeMaximumBy(1)
	}
}
