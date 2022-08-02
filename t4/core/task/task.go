package task

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
)

const (
	DefaultValue       = -1000
	Unresolved         = 10
	SpectacularFailure = 11
	Failure            = 12
	Success            = 13
	SpectacularSuccess = 14
	TaskEasy           = 1000
	TaskAverage        = 2000
	TaskDifficult      = 3000
	TaskFormidable     = 4000
	TaskStaggering     = 5000
	TaskImpossible     = 6000
)

type Task struct {
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

func New(descr string, diff, tn int, mods ...mod) (*Task, error) {
	tsk := Task{}
	tsk.difficulty = diff
	tsk.targetNumber = tn
	tsk.result = Unresolved
	tsk.description = descr
	for _, m := range mods {
		if err := checkMod(m); err != nil {
			return &tsk, err
		}
	}
	return &tsk, nil
}

func checkMod(m mod) error {
	//testChange
	if m.modType == "" {
		return fmt.Errorf("Modifier description is not set")
	}
	if m.modInfluence < 1 {
		return fmt.Errorf("Modifier inpfluence is not set correctly (%v:%v)", m.modType, m.modInfluence)
	}
	return nil
}

func (t *Task) Resolve() {
	dp := dice.New()
	if t.seedFixed {
		dp = dp.SetSeed(t.description + t.description)
	}

	switch t.difficulty {
	case TaskEasy:
		t.roll = dp.Roll("1d6").Sum() + dp.Roll("1d3").Sum()
		if t.roll == 2 {
			t.result = SpectacularFailure
		}
		if t.roll == 9 {
			t.result = SpectacularSuccess
		}
	case TaskAverage:
		t.roll = dp.Roll("2d6").Sum()
		if t.roll == 2 {
			t.result = SpectacularFailure
		}
		if t.roll == 12 {
			t.result = SpectacularSuccess
		}
	case TaskDifficult:
		t.roll = dp.Roll("2d6").Sum() + dp.Roll("1d3").Sum()
		if t.roll == 3 {
			t.result = SpectacularFailure
		}
		if t.roll == 15 {
			t.result = SpectacularSuccess
		}
	case TaskFormidable:
		t.roll = dp.Roll("3d6").Sum()
		if t.roll == 3 {
			t.result = SpectacularFailure
		}
		if t.roll == 18 {
			t.result = SpectacularSuccess
		}
	case TaskStaggering:
		t.roll = dp.Roll("3d6").Sum() + dp.Roll("1d3").Sum()
		if t.roll == 4 {
			t.result = SpectacularFailure
		}
		if t.roll == 21 {
			t.result = SpectacularSuccess
		}
	case TaskImpossible:
		t.roll = dp.Roll("4d6").Sum()
		if t.roll == 4 {
			t.result = SpectacularFailure
		}
		if t.roll == 24 {
			t.result = SpectacularSuccess
		}
	}
	if t.result != Unresolved {
		return
	}
	modTN := t.targetNumber
	for _, m := range t.modifiers {
		modTN += m.modInfluence
	}
	t.result = Failure
	if t.roll <= modTN {
		t.result = Success
	}
	return
}

func (t *Task) String() string {
	str := fmt.Sprintf("Task Prase: %v\n", t.description)
	return str
}
