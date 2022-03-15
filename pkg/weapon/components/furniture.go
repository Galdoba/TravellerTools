package components

import "fmt"

type Furniture struct {
	stock          int
	bipod          int
	supportMount   bool
	modularization bool
}

func NewFurniture(instructions ...int) (*Furniture, error) {
	f := Furniture{}
	err := fmt.Errorf("Error was not adressed")
	if err = analizeFurniture(instructions); err != nil {
		return &f, err
	}
	f.stock = FRNTR_STOCKLESS
	f.bipod = FRNTR_BIPOD_ABSENT
	for _, instr := range instructions {
		if !isFurniture(instr) {
			continue
		}
		f.addFurniture(instr)
	}
	return &f, err
}

func analizeFurniture(instructions []int) error {
	switch timesCrossed(instructions, []int{FRNTR_STOCKLESS, FRNTR_STOCK_FOLDING, FRNTR_STOCK_FULL}) {
	case 0, 1:
	default:
		return fmt.Errorf("Furniture: Stock instructions confusing")
	}
	switch timesCrossed(instructions, []int{FRNTR_BIPOD_ABSENT, FRNTR_BIPOD_DETACHABLE, FRNTR_BIPOD_FIXED}) {
	case 0, 1:
	default:
		return fmt.Errorf("Furniture: Bipod instructions confusing")
	}
	return nil
}

func isFurniture(i int) bool {
	if i <= FRNTR_STOCKLESS && i >= FRNTR_SUPPORT_MOUNT {
		return true
	}
	return false
}

func (f *Furniture) addFurniture(i int) {
	switch i {
	case FRNTR_STOCKLESS, FRNTR_STOCK_FOLDING, FRNTR_STOCK_FULL:
		f.stock = i
	case FRNTR_BIPOD_ABSENT, FRNTR_BIPOD_FIXED, FRNTR_BIPOD_DETACHABLE:
		f.bipod = i
	case FRNTR_MODULARIZATION, FRNTR_SUPPORT_MOUNT:
		f.modularization = true
	}
}
