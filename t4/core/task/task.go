package task

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

const (
	DefaultValue = -1000

	Unresolved         = 110
	SpectacularFailure = 111
	Failure            = 112
	Success            = 113
	SpectacularSuccess = 114
	TaskEasy           = 1000
	TaskAverage        = 2000
	TaskDifficult      = 3000
	TaskFormidable     = 4000
	TaskStaggering     = 5000
	TaskImpossible     = 6000
	TN_Random          = 10000
)

type Task struct {
	resolver     *dice.Dicepool
	description  string
	difficulty   int
	targetNumber int
	modifiers    []mod
	result       int
	seedFixed    bool
	roll         int
}

type mod struct {
	modType      string
	modInfluence int
}

func New(tn int, dificulty int) *Task {
	tsk := Task{}
	tsk.difficulty = dificulty
	tsk.targetNumber = tn
	tsk.result = Unresolved
	return &tsk
}

func (t *Task) SetResolver(dice *dice.Dicepool) {
	t.resolver = dice
}

func (t *Task) AddDescription(descr string) {
	t.description = descr
}

func (t *Task) Result() int {
	return t.result
}

// func New(descr string, diff, tn int, mods ...mod) (*Task, error) {
// 	tsk := Task{}
// 	tsk.difficulty = diff
// 	tsk.targetNumber = tn
// 	tsk.result = Unresolved
// 	tsk.description = descr
// 	for _, m := range mods {
// 		if err := checkMod(m); err != nil {
// 			return &tsk, err
// 		}
// 	}
// 	return &tsk, nil
// }

func checkMod(m mod) error {
	//testChange
	//testChange2
	//testChange3
	//testChange4
	if m.modType == "" {
		return fmt.Errorf("Modifier description is not set")
	}
	if m.modInfluence < 1 {
		return fmt.Errorf("Modifier inpfluence is not set correctly (%v:%v)", m.modType, m.modInfluence)
	}
	return nil
}

func (t *Task) Resolve() int {
	if t.resolver == nil {
		t.resolver = dice.New()
	}
	switch t.difficulty {
	case TaskEasy:
		t.roll = t.resolver.Sroll("1d6") + t.resolver.Sroll("1d3")
		if t.roll == 2 {
			t.result = SpectacularFailure
		}
		if t.roll == 9 {
			t.result = SpectacularSuccess
		}
	case TaskAverage:
		t.roll = t.resolver.Sroll("2d6")
		if t.roll == 2 {
			t.result = SpectacularFailure
		}
		if t.roll == 12 {
			t.result = SpectacularSuccess
		}
	case TaskDifficult:
		t.roll = t.resolver.Sroll("2d6") + t.resolver.Sroll("1d3")
		if t.roll == 3 {
			t.result = SpectacularFailure
		}
		if t.roll == 15 {
			t.result = SpectacularSuccess
		}
	case TaskFormidable:
		t.roll = t.resolver.Sroll("3d6")
		if t.roll == 3 {
			t.result = SpectacularFailure
		}
		if t.roll == 18 {
			t.result = SpectacularSuccess
		}
	case TaskStaggering:
		t.roll = t.resolver.Sroll("3d6") + t.resolver.Sroll("1d3")
		if t.roll == 4 {
			t.result = SpectacularFailure
		}
		if t.roll == 21 {
			t.result = SpectacularSuccess
		}
	case TaskImpossible:
		t.roll = t.resolver.Sroll("4d6")
		if t.roll == 4 {
			t.result = SpectacularFailure
		}
		if t.roll == 24 {
			t.result = SpectacularSuccess
		}
	}
	if t.result != Unresolved {
		return t.result
	}

	t.result = Failure
	if t.roll <= t.targetNumber {
		t.result = Success
	}
	return t.result
}

func (t *Task) String() string {
	str := fmt.Sprintf("Task Prase: %v\n", t.description)
	return str
}
