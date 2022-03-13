package weapon

import (
	"fmt"
	"strconv"
	"testing"

	combinations "github.com/mxschmitt/golang-combinations"
)

func input() [][]string {
	dt := []int{brl_len_MINIMAL, brl_len_SHORT, brl_len_HANDGUN, brl_len_ASSAULT, brl_len_CARBINE, brl_len_RIFLE, brl_len_LONG, brl_len_VERY_LONG, brl_weight_HEAVY, brl_weight_STANDARD, WRONG_INSTRUCTION}
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
		fmt.Println("Test", verbal(instructions[0]), verbal(instructions[1]))
		brl, err := newBarrel(instructions...)
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
