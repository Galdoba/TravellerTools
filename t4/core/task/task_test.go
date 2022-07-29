package task

import (
	"fmt"
	"testing"
)

type taskExample struct {
	descr string
	diff  int
	tn    int
	mods  []mod
}

func taskInputExamples() []taskExample {
	return []taskExample{
		{"[test task]", TaskAverage, 7, nil},
		{"[test task with mods]", TaskAverage, 7, []mod{{"test mod", 1}, {"test mod 2", 3}}},
	}
}

func TestTask(t *testing.T) {
	for i, input := range taskInputExamples() {
		fmt.Printf("Test %v: %v\n", i, input)
		tsk, err := New(input.descr, input.diff, input.tn, input.mods...)
		if tsk == nil {
			t.Errorf("func returned no object")
		}
		if err != nil {
			t.Errorf("func returned error: %v", err.Error())
		}
		if tsk.description == "" {
			t.Errorf("task description is not set (expect %v)", input.descr)
		}
		switch tsk.difficulty {
		default:
			t.Errorf("task difficulty is %v (unknown value)", tsk.difficulty)
		case TaskEasy, TaskAverage, TaskDifficult, TaskFormidable, TaskStaggering, TaskImpossible:
			if tsk.difficulty != input.diff {
				t.Errorf("task difficulty is %v (expect %v)", tsk.difficulty, input.diff)
			}
		}
		if tsk.targetNumber == DefaultValue {
			t.Errorf("task target number is not set")
		}
		if tsk.targetNumber != input.tn {
			t.Errorf("task target number is %v (expect %v)", tsk.targetNumber, input.tn)
		}
		tsk.Resolve()
		fmt.Printf("Resolve: %v\n", tsk.result)
	}
}
