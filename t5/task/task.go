package task

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type t5task struct {
	typeOfTask      string
	char            string
	skill           string
	mods            []mod
	duration        taskDurationData
	specialTypeTask string
	dieRoll         int
	dieUncertain    int
	prob            float64
	////////
	phrase string
}

type Asset interface {
	Name() string
	Actual() int
	SetValue(v int)
}

//task.State()

type instruction struct {
	key      string
	valS     string
	valI     int
	valInt64 int64
	valB     bool
}

/*
Phrase - string
Difficulty - int
Assets - int
comments - string
*/
const (
	ignoreInstruction      = "ignore_instruction"
	instKey_TaskPhrase     = "Task Phrase"
	instKey_Duration       = "Duration"
	instKey_BaseDifficulty = "Difficulty"
	instKey_Chararteristic = "Characteristic"
	instKey_Skill          = "Skill"
	instKey_Mod            = "Mod"
)

func New(inst ...instruction) (*t5task, error) {
	tsk := t5task{}
	for _, in := range inst {
		switch {
		default:
			return nil, fmt.Errorf("unknown instruction type '%v'", in.key)
		case in.key == ignoreInstruction:
		case in.key == instKey_TaskPhrase:
			if tsk.phrase != "" {
				return &tsk, fmt.Errorf("instruction repeated: Task Phrase: %v", in)
			}
			tsk.phrase = in.valS
		case in.key == instKey_Duration:
			if tsk.duration.durationType != Ignored {
				return &tsk, fmt.Errorf("instruction repeated: Duration: %v", in)
			}
			tsk.duration.durationType = in.valI
			tsk.duration.units = in.valInt64
		case in.key == instKey_BaseDifficulty:
			if tsk.dieRoll != 0 {
				return &tsk, fmt.Errorf("instruction repeated: Difficulty: %v", in)
			}
			tsk.dieRoll = in.valI
		case in.key == instKey_Chararteristic:
			if tsk.char != "" {
				return &tsk, fmt.Errorf("instruction repeated: Characteristic: %v", in)
			}
			tsk.char = in.valS
		case in.key == instKey_Skill:
			if tsk.skill != "" {
				return &tsk, fmt.Errorf("instruction repeated: Skill: %v", in)
			}
			tsk.skill = in.valS
		case in.key == instKey_Mod:
			for _, m := range tsk.mods {
				if strings.Contains(m.text, in.valS) {
					return &tsk, fmt.Errorf("instruction repeated: Modifier: %v", in)
				}
			}
			tsk.mods = append(tsk.mods, mod{
				text:     in.valS,
				val:      in.valI,
				required: in.valB,
			})
		}

	}

	return &tsk, nil
}

func (tsk *t5task) toString() string {
	s := tsk.phrase
	if tsk.duration.durationType != Ignored {
		s += fmt.Sprintf(" %v", tsk.duration.describe())
	}
	s += fmt.Sprintf("\n  %vD <= (", tsk.dieRoll)
	if tsk.char != "" {
		s += tsk.char + " + "
	}
	if tsk.skill != "" {
		s += tsk.skill + " + "
	}

	s = strings.TrimSuffix(s, " + ")
	s += ")"
	fmt.Println(tsk.mods)
	for _, m := range tsk.mods {
		s += "\n "
		switch m.val >= 0 {
		case true:
			s += " +" + fmt.Sprintf("%v for %v", math.Abs(float64(m.val)), m.text)
		case false:
			s += " -" + fmt.Sprintf("%v for %v", math.Abs(float64(m.val)), m.text)
		}
		if m.required {
			s += " (required)"
		}
	}

	return s
}

func (tsk *t5task) fill(dataMap map[string]int) {

}

func (tsk *t5task) Resolve(dice *dice.Dicepool) error {
	return nil
}
