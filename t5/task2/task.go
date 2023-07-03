package task2

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile"
)

/*
	Resolver - тот кто предоставляет Аssets
			 		   сравнивается реквезитами
*/

const (
	Unknown = iota
	DIF_Easy
	DIF_Average
	DIF_Difficult
	DIF_Formidable
	DIF_Staggering
	DIF_Hopeless
	DIF_Impossible
	DIF_BeyondImpossible
	VARIABLE
	AUTOMATIC
	DUR_COMBAT_ROUND
	DUR_Minutes
	DUR_SPACE_ROUND
	DUR_Hour
	DUR_Day
	DUR_Week
	DUR_Month
	DUR_Quarter
	DUR_Year
	ABSOLUTE
	RANDOM
	WithoutSkill
	SkillOnly
	SkillOptional
	Mod_Required
	Mod_Optional
	Mod_Automatic
)

type Task struct {
	statement     string
	Difficulty    int
	DurationScale int
	DurationType  int
	startTick     int64
	resolveTick   int64
	SkillUse      int
	CharKey       string
	SkillKey      string
	Mods          []Modifier
	Comments      []string
	err           error
}

type Actor interface {
	Profile() profile.Profile
}

/*
attempt := task.New("Fix engine",3).Assets("C2","Mechanic", Mod("Edu 10+", 2, Mod_Optional))
verdict := atempt.Resolve(actor,dice)



task.New("Begin Noble career",Dif_Automatic).Assets("Soc","", Mod("Soc 11+", 0, Mod_Required))
*/

func New(statement string, difficulty int) *Task {
	ts := Task{}
	ts.statement = statement
	ts.Difficulty = difficulty
	return &ts
}

func (ts *Task) AssetsRequired(charKey, skillKey string, mods ...Modifier) *Task {
	ts.CharKey = charKey
	ts.SkillKey = skillKey
	ts.Mods = append(ts.Mods, mods...)
	return ts
}

func (ts *Task) SkillUseMethod(skillUse int) *Task {
	ts.SkillUse = skillUse
	return ts
}

func (ts *Task) Duration(duration, variance int, startTick int64) *Task {
	ts.DurationScale = duration
	ts.DurationType = variance
	ts.startTick = startTick
	ts.resolveTick = -1
	return ts
}

type Modifier struct {
	Description string
	Value       int
	Use         int
}

type Verdict struct {
	Success      bool
	Spectaculars []int
	err          error
}

func (ts *Task) Resolve(actor Actor, dice *dice.Dicepool) Verdict {
	vrd := Verdict{}
	chr := actor.Profile().Data(ts.CharKey)
	if chr == nil {
		switch ts.SkillUse {
		default:
			vrd.err = fmt.Errorf("cann't resolve: actor has no charkey '%v'", ts.CharKey)
			return vrd
		case SkillOnly:
			chr = ehex.New().Set(7)
		}
	}
	skl := actor.Profile().Data(ts.SkillKey)
	if skl == nil {
		switch ts.SkillUse {
		default:
			vrd.err = fmt.Errorf("cann't resolve: actor has no skillkey '%v'", ts.SkillKey)
			return vrd
		case SkillOptional:
			skl = ehex.New().Set(0)
		case WithoutSkill:
			skl = ehex.New().Set(3)
		}
	}
	mods := []ehex.Ehex{}
	for _, m := range ts.Mods {
		switch m.Use {
		case Mod_Optional:
		}
		mods = append(mods, ehex.New().Set(m.Value))
	}
	return vrd
}
