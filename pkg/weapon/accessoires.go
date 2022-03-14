package weapon

import "fmt"

type accessoire struct {
	suppressor int
	ammoFeed   int
	sighting   int
	other      []int
}

func newAccessoires(instructions ...int) (*accessoire, error) {
	err := fmt.Errorf("error was not addressed")
	a := accessoire{}
	if err = analizeAccessoires(instructions); err != nil {
		return &a, err
	}
	a.addAccessoire(instructions...)
	return &a, err
}

func (a *accessoire) addAccessoire(instructions ...int) {
	for _, inst := range instructions {
		if !isAccessoire(inst) {
			continue
		}
		switch inst {
		case accessoire_SUPPRESSOR_ABSENT, accessoire_SUPPRESSOR_BASIC, accessoire_SUPPRESSOR_STANDARD, accessoire_SUPPRESSOR_EXTREME:
			a.suppressor = inst
		case accessoire_AFD_MAGAZINE_FIXED, accessoire_AFD_MAGAZINE_STANDARD, accessoire_AFD_MAGAZINE_EXTENDED, accessoire_AFD_MAGAZINE_DRUM, accessoire_AFD_BELT, accessoire_AFD_CLIPS:
			a.ammoFeed = inst
		case accessoire_SCOPE_ABSENT, accessoire_SCOPE_BASIC, accessoire_SCOPE_LONG_RANGE, accessoire_SCOPE_LOW_LIGHT, accessoire_SCOPE_THERMAL, accessoire_SCOPE_COMBINATION, accessoire_SCOPE_MULTISPECTRAL, accessoire_SCOPE_LASER_POINTER, accessoire_SCOPE_ISS, accessoire_SCOPE_HOLOGRAPHIC:
			a.sighting = inst
		case accessoire_OTHER_ABSENT, accessoire_OTHER_BAYONET_LUG, accessoire_OTHER_BLING, accessoire_OTHER_FLASHLIGHT, accessoire_OTHER_GRAVITIC_SUPPORT, accessoire_OTHER_GUN_CAMERA, accessoire_OTHER_WEAPON_INTELLIGENT, accessoire_OTHER_WEAPON_SECURE, accessoire_OTHER_STABILISATION:
			a.other = append(a.other, inst)
		}
	}
	if len(a.other) == 0 {
		a.other = append(a.other, accessoire_OTHER_ABSENT)
	}
}

func analizeAccessoires(instructions []int) error {
	switch {
	default:
		//return fmt.Errorf("Accesoires: not Implemented: ammoFeed/sighting/other")

	case timesCrossed(instructions, []int{accessoire_SUPPRESSOR_ABSENT, accessoire_SUPPRESSOR_BASIC, accessoire_SUPPRESSOR_STANDARD, accessoire_SUPPRESSOR_EXTREME}) > 1:
		return fmt.Errorf("Accessoires: multiple Suppressor instructions")
	case timesCrossed(instructions, []int{accessoire_AFD_MAGAZINE_FIXED, accessoire_AFD_MAGAZINE_STANDARD, accessoire_AFD_MAGAZINE_EXTENDED, accessoire_AFD_MAGAZINE_DRUM,
		accessoire_AFD_BELT, accessoire_AFD_CLIPS}) > 1:
		return fmt.Errorf("Accessoires: multiple Ammo Feeder Device instructions")
	case timesCrossed(instructions, []int{accessoire_SCOPE_ABSENT, accessoire_SCOPE_BASIC, accessoire_SCOPE_LONG_RANGE, accessoire_SCOPE_LOW_LIGHT,
		accessoire_SCOPE_THERMAL, accessoire_SCOPE_COMBINATION, accessoire_SCOPE_MULTISPECTRAL, accessoire_SCOPE_LASER_POINTER, accessoire_SCOPE_ISS,
		accessoire_SCOPE_HOLOGRAPHIC}) > 1:
		return fmt.Errorf("Accessoires: multiple Sighting Device instructions")
	case timesCrossed(instructions, []int{accessoire_OTHER_ABSENT, accessoire_OTHER_BAYONET_LUG, accessoire_OTHER_BLING, accessoire_OTHER_FLASHLIGHT,
		accessoire_OTHER_GRAVITIC_SUPPORT, accessoire_OTHER_GUN_CAMERA, accessoire_OTHER_WEAPON_INTELLIGENT, accessoire_OTHER_WEAPON_SECURE,
		accessoire_OTHER_STABILISATION}) < 1:
	}
	return nil
}

func isAccessoire(i int) bool {
	if i >= accessoire_SUPPRESSOR_ABSENT && i <= accessoire_OTHER_STABILISATION {
		return true
	}
	return false
}
