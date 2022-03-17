package components

import (
	"fmt"
	"strconv"
	"testing"

	combinations "github.com/mxschmitt/golang-combinations"
)

func input() [][]string {
	dt := []int{BRL_len_MINIMAL, BRL_len_SHORT, BRL_len_HANDGUN, BRL_len_ASSAULT, BRL_len_CARBINE, BRL_len_RIFLE, BRL_len_LONG, BRL_len_VERY_LONG, BRL_weight_HEAVY, BRL_weight_STANDARD, WRONG_INSTRUCTION}
	dtStr := []string{}
	for _, s := range dt {
		dtStr = append(dtStr, strconv.Itoa(s))
	}
	return combinations.Combinations(dtStr, 2)
}

func TestBarrel(t *testing.T) {
	input := input()
	for _, try := range input {
		instructions := []int{}
		for _, t := range try {
			in, _ := strconv.Atoi(t)
			instructions = append(instructions, in)
		}
		if !(isBarrelLentgh(instructions[0]) && isBarrelWeight(instructions[1])) {
			continue
		}
		fmt.Println("Test", Verbal(instructions[0]), Verbal(instructions[1]))
		brl, err := NewBarrel(instructions...)
		if err != nil {
			t.Errorf("error: %v", err.Error())
		}
		fmt.Println("barrel struct:", brl)
		if brl.lenght == _UNDEFINED_ {
			t.Errorf("barrel lentgh is undefined")
		}
		if !isBarrelLentgh(brl.lenght) {
			t.Errorf("barrel lentgh value incorect (%v)", brl.lenght)
		}
	}
}
