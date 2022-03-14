package weapon

import "fmt"

type furniture struct {
	stock          int
	bipod          int
	supportMount   bool
	modularization bool
}

func newFurniture(instructions ...int) (*furniture, error) {
	f := furniture{}
	err := fmt.Errorf("Error was not adressed")
	if err = analizeFurniture(instructions); err != nil {
		return &f, err
	}
	f.stock = furniture_STOCKLESS
	f.bipod = furniture_BIPOD_ABSENT
	for _, instr := range instructions {
		if !isFurniture(instr) {
			continue
		}
		f.addFurniture(instr)
	}
	return &f, err
}

func analizeFurniture(instructions []int) error {
	switch timesCrossed(instructions, []int{furniture_STOCKLESS, furniture_STOCK_FOLDING, furniture_STOCK_FULL}) {
	case 0, 1:
	default:
		return fmt.Errorf("Furniture: Stock instructions confusing")
	}
	switch timesCrossed(instructions, []int{furniture_BIPOD_ABSENT, furniture_BIPOD_DETACHABLE, furniture_BIPOD_FIXED}) {
	case 0, 1:
	default:
		return fmt.Errorf("Furniture: Bipod instructions confusing")
	}
	return nil
}

func isFurniture(i int) bool {
	if i <= furniture_STOCKLESS && i >= furniture_SUPPORT_MOUNT {
		return true
	}
	return false
}

func (f *furniture) addFurniture(i int) {
	switch i {
	case furniture_STOCKLESS, furniture_STOCK_FOLDING, furniture_STOCK_FULL:
		f.stock = i
	case furniture_BIPOD_ABSENT, furniture_BIPOD_FIXED, furniture_BIPOD_DETACHABLE:
		f.bipod = i
	case furniture_MODULARIZATION, furniture_SUPPORT_MOUNT:
		f.modularization = true
	}
}
