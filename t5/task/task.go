package task

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/t5/pawn/skill"
)

type t5task struct {
	typeOfTask      string
	charCode        string
	skillID         int
	mods            []mod
	duration        taskDurationData
	specialTypeTask string
	dieRoll         int
	dieUncertain    int
	prob            float64
	rr              rollResult
	////////
	phrase string
}

type Asset interface {
	Name() string
	Actual() int
	SetValue(v int) //for testing
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
			if tsk.charCode != "" {
				return &tsk, fmt.Errorf("instruction repeated: Characteristic: %v", in)
			}
			tsk.charCode = in.valS
		case in.key == instKey_Skill:
			if tsk.skillID != 0 {
				return &tsk, fmt.Errorf("instruction repeated: Skill: %v", in)
			}
			tsk.skillID = in.valI
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
	if tsk.charCode != "" {
		s += tsk.charCode + " + "
	}
	if tsk.skillID != 0 {
		s += skill.NameByID(tsk.skillID) + " + "
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

func TargetNumber(as ...Asset) int {
	tn := 0
	for _, v := range as {
		tn += v.Actual()
	}
	return tn
}

func (tsk *t5task) Resolve(dice *dice.Dicepool, tn int) error {
	code := fmt.Sprintf("%vd6", tsk.dieRoll)
	res := dice.Roll(code).Result()

	tsk.rr = newResult(res, tn)
	return nil
}
