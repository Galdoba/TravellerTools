package task

import (
	"fmt"
	"testing"
)

type taskinput struct {
	descr  string
	dif    int
	assets []taskAsset
}

func inputTask() []taskinput {
	return []taskinput{
		{"fix car", DifficultyAverage, []taskAsset{Asset("Dex", 7), Asset("Mechanic", 2)}},
		{"establish comm contact with pinnacle Crew", DifficultyDifficult, []taskAsset{Asset("Edu", 7), Asset("Communication", 2), Asset("Good Radio", 1)}},
		{"establish comm contact with pinnacle Crew", DifficultyDifficult, []taskAsset{Asset("Edu", 7), Asset("Communication", 2), Asset("Good Radio", 1), Asset("Good Radio", 5)}},
		{"leap a 6.0 meter gap", DifficultyBeyondImpossible, []taskAsset{Asset("Str", 5), Asset("Athletics", 1)}},
	}
}

func TestTask(t *testing.T) {
	for i, input := range inputTask() {
		fmt.Printf("Test %v - %v\n", i, input)
		task, err := New(input.descr, input.dif, input.assets...)
		if task == nil {
			t.Errorf("func returned no object")
		}
		if err != nil {
			t.Errorf("func returned error: %v", err.Error())
			continue
		}
		fmt.Println(task.Phrase())
		fmt.Println(task.Statement())
		task.Resolve()
		fmt.Println(task.Result())
	}

}
