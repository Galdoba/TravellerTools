package components

import "fmt"

type Accessoire struct {
	ammoFeed int
	other    []int
}

func NewAccessoires(instructions ...int) (*Accessoire, error) {
	err := fmt.Errorf("error was not addressed")
	a := Accessoire{}
	if err = analizeAccessoires(instructions); err != nil {
		return &a, err
	}
	a.addAccessoire(instructions...)
	return &a, err
}

func (a *Accessoire) addAccessoire(instructions ...int) {
	for _, inst := range instructions {
		if !isAccessoire(inst) {
			continue
		}
		switch inst {
		case ACSR_AFD_MAGAZINE_FIXED, ACSR_AFD_MAGAZINE_STANDARD, ACSR_AFD_MAGAZINE_EXTENDED, ACSR_AFD_MAGAZINE_DRUM, ACSR_AFD_BELT, ACSR_AFD_CLIPS:
			a.ammoFeed = inst
		default:
			a.other = append(a.other, inst)
		}
	}
	if len(a.other) == 0 {
		a.other = append(a.other, ACSR_ABSENT)
	}
}

func analizeAccessoires(instructions []int) error {
	switch {
	default:
		//return fmt.Errorf("Accesoires: not Implemented: ammoFeed/sighting/other")
	case timesCrossed(instructions, []int{ACSR_AFD_MAGAZINE_FIXED, ACSR_AFD_MAGAZINE_STANDARD, ACSR_AFD_MAGAZINE_EXTENDED, ACSR_AFD_MAGAZINE_DRUM,
		ACSR_AFD_BELT, ACSR_AFD_CLIPS}) > 1:
		return fmt.Errorf("Accessoires: multiple Ammo Feeder Device instructions")
	case timesCrossed(instructions, []int{ACSR_SCOPE_BASIC, ACSR_SCOPE_LONG_RANGE, ACSR_SCOPE_LOW_LIGHT,
		ACSR_SCOPE_THERMAL, ACSR_SCOPE_COMBINATION, ACSR_SCOPE_MULTISPECTRAL}) > 1:
		return fmt.Errorf("Accessoires: multiple Scope instructions")
	case timesCrossed(instructions, []int{ACSR_SUPPRESSOR_BASIC, ACSR_SUPPRESSOR_STANDARD, ACSR_SUPPRESSOR_EXTREME}) > 1:
		return fmt.Errorf("Accessoires: multiple Suppressor instructions")
	}
	return nil
}

func isAccessoire(i int) bool {
	if i >= ACSR_SUPPRESSOR_BASIC && i <= ACSR_OTHER_STABILISATION {
		return true
	}
	return false
}
