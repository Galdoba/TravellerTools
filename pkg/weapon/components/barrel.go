package components

import "fmt"

type Barrel struct {
	lenght int
	heavy  bool
}

func NewBarrel(instructions ...int) (*Barrel, error) {
	brl := Barrel{}
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
			if i == BRL_weight_HEAVY {
				brl.heavy = true
			}
		}
	}
	return &brl, err
}

func isBarrelLentgh(i int) bool {
	if i >= BRL_len_MINIMAL && i <= BRL_len_VERY_LONG {
		return true
	}
	return false
}

func isBarrelWeight(i int) bool {
	if i >= BRL_weight_STANDARD && i <= BRL_weight_HEAVY {
		return true
	}
	return false
}

func inputInstructionsCheck(inst []int) error {
	instrMap := make(map[int]int)
	if len(inst) != 2 {
		//return fmt.Errorf("instructions confusing: expect 2 but have (%v)", len(inst))
	}
	for _, i := range inst {
		instrMap[i]++
	}
	if instrMap[BRL_weight_STANDARD]+instrMap[BRL_weight_HEAVY] == 0 {
		//return fmt.Errorf("Barrel weight not assignned")
	}
	if instrMap[BRL_weight_STANDARD]+instrMap[BRL_weight_HEAVY] > 1 {
		return fmt.Errorf("Barrel weight instructions is confusing (heavy=%v std=%v)", instrMap[BRL_weight_STANDARD], instrMap[BRL_weight_HEAVY])
	}
	return nil
}
