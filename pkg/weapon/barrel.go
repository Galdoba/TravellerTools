package weapon

import "fmt"

type barrel struct {
	lenght int
	heavy  bool
}

func newBarrel(instructions ...int) (*barrel, error) {
	brl := barrel{}
	err := fmt.Errorf("err value not addressed")
	err = inputInstructionsCheck(instructions)
	if err != nil {
		return &brl, err
	}
	for _, i := range instructions {
		switch isBarrelLentgh(i) {
		case true:
			brl.lenght = i
		case false:
			if i == brl_weight_HEAVY {
				brl.heavy = true
			}
		}
	}
	return &brl, err
}

func isBarrelLentgh(i int) bool {
	if i >= brl_len_MINIMAL && i <= brl_len_VERY_LONG {
		return true
	}
	return false
}

func isBarrelWeight(i int) bool {
	if i >= brl_weight_STANDARD && i <= brl_weight_HEAVY {
		return true
	}
	return false
}

func inputInstructionsCheck(inst []int) error {
	instrMap := make(map[int]int)
	if len(inst) != 2 {
		return fmt.Errorf("instructions confusing: expect 2 but have (%v)", len(inst))
	}
	for _, i := range inst {
		instrMap[i]++
	}
	if instrMap[brl_weight_STANDARD]+instrMap[brl_weight_HEAVY] == 0 {
		return fmt.Errorf("barrel weight not assignned")
	}
	if instrMap[brl_weight_STANDARD]+instrMap[brl_weight_HEAVY] == 0 {
		return fmt.Errorf("barrel weight instructions is confusing (heavy=%v std=%v)", instrMap[brl_weight_STANDARD], instrMap[brl_weight_HEAVY])
	}
	return nil
}
