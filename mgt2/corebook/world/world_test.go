package world

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	data := [][]string{}
	for i := 100; i < 125; i++ {
		c := NewConstructor(
			Instruction(KEY_SEED, fmt.Sprintf("Test %v", i+25)),
			Instruction(KEY_NAME, ""),
		)
		w, err := c.Create()
		data = append(data, w.ShortData())
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	PrintAsTable(data...)
}

func PrintAsTable(flds ...[]string) {
	lMap := make(map[int]int)
	for _, fld := range flds {
		for i, f := range fld {
			if lMap[i] < len(f) {
				lMap[i] = len(f)
			}
		}

	}
	for _, fld := range flds {
		for i, f := range fld {
			for len(f) < lMap[i] {
				f += " "
			}
			fmt.Print(f + "  ")
		}
		fmt.Print("\n")
	}
}
