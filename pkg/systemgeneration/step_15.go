package systemgeneration

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
)

const (
	sizeDwarf       = "Dwarf"
	sizeMercurian   = "Mercurian"
	sizeSubterran   = "Subterran"
	sizeTerran      = "Terran"
	sizeSuperterran = "Superterran"
)

func (gs *GenerationState) Step15() error {
	fmt.Println("START Step 15")
	if gs.NextStep != 15 {
		return fmt.Errorf("not actual step")
	}
	gs.setPlanetTypes()
	gs.ConcludedStep = 15
	gs.NextStep = 16
	switch gs.NextStep {
	case 16:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	//fmt.Println("END Step 15")
	return nil
}

func (gs *GenerationState) setPlanetTypes() error {
	for i, star := range gs.System.Stars {
		fmt.Println(star)
		lowHab := star.habitableLow
		hiHab := star.habitableHigh
		for _, orbit := range star.orbitDistances {
			fmt.Println(orbit)
			if planet, ok := star.orbit[orbit].(*rockyPlanet); ok == true {
				pos := planet.orbit
				sizeTypeRoll := gs.Dice.Roll("2d6").Sum()
				sizeType := "UNDEFINED"
				p := ""
				switch {
				case pos < lowHab:
					sizeType = innerSizeType(sizeTypeRoll)
					p = "I"
				case pos < hiHab:
					sizeType = habitableSizeType(sizeTypeRoll)
					p = "H"
				default:
					sizeType = outerSizeType(sizeTypeRoll)
					p = "O"
				}
				gs.System.Stars[i].orbit[orbit].(*rockyPlanet).sizeType = sizeType
				size := rollSize(gs.Dice, sizeType, p)
				gs.System.Stars[i].orbit[orbit].(*rockyPlanet).sizeCode = size
				atmo := rollAtmo(gs.Dice, sizeType, p, size)
				gs.System.Stars[i].orbit[orbit].(*rockyPlanet).atmoCode = atmo
				hydr := rollHydr(gs.Dice, sizeType, p, atmo)
				gs.System.Stars[i].orbit[orbit].(*rockyPlanet).hydrCode = hydr

			}
		}
	}
	return nil
}

func rollSize(dp *dice.Dicepool, sType, p string) string {
	switch sType {
	case sizeDwarf:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerDwarfSize(i)
		case "H":
			return habitableDwarfSize(i)
		case "O":
			return outerDwarfSize(i)
		}
	case sizeMercurian:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerMercurianSize(i)
		case "H":
			return habitableMercurianSize(i)
		case "O":
			return outerMercurianSize(i)
		}
	case sizeSubterran:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerSubterranSize(i)
		case "H":
			return habitableSubterranSize(i)
		case "O":
			return outerSubterranSize(i)
		}
	case sizeTerran:
		i := dp.Roll("2d6").Sum()
		switch p {
		case "I":
			return innerTerranSize(i)
		case "H":
			return habitableTerranSize(i)
		case "O":
			return outerTerranSize(i)
		}
	case sizeSuperterran:
		i := dp.Roll("2d10").Sum()
		switch p {
		case "I":
			return innerSuperterranSize(i)
		case "H":
			return habitableSuperterranSize(i)
		case "O":
			return outerSuperterranSize(i)
		}
	}
	return "*"
}

func rollAtmo(dp *dice.Dicepool, sType, p string, size string) string {
	switch sType {
	case sizeDwarf:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerDwarfAtmo(i)
		case "H":
			return habitableDwarfAtmo(i)
		case "O":
			return outerDwarfAtmo(i)
		}
	case sizeMercurian:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerMercurianAtmo(i)
		case "H":
			return habitableMercurianAtmo(i)
		case "O":
			return outerMercurianAtmo(i)
		}
	case sizeSubterran:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerSubterranAtmo(i)
		case "H":
			return habitableSubterranAtmo(i)
		case "O":
			return outerSubterranAtmo(i)
		}
	case sizeTerran:
		i := dp.Roll("2d6").Sum()
		switch p {
		case "I":
			return innerTerranAtmo(i)
		case "H":
			return habitableTerranAtmo(i)
		case "O":
			return outerTerranAtmo(i)
		}
	case sizeSuperterran:
		i := dp.Roll("2d10").Sum()
		switch p {
		case "I":
			dm := 0
			switch size {
			case "A", "B", "C", "D", "E":
			case "F", "G", "H", "J", "K":
				dm = 3
			default:
				dm = 6
			}
			return innerSuperterranAtmo(i + dm)
		case "H":
			return habitableSuperterranAtmo(i)
		case "O":
			return outerSuperterranAtmo(i)
		}
	}
	return "*"
}

func rollHydr(dp *dice.Dicepool, sType, p, atmo string) string {
	i := 0
	switch sType {
	case sizeDwarf:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerDwarfHydr(i)
		case "H":
			return habitableDwarfHydr(i)
		case "O":
			return outerDwarfHydr(i)
		}
	case sizeMercurian:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerMercurianHydr(i)
		case "H":
			return habitableMercurianHydr(i)
		case "O":
			return outerMercurianHydr(i)
		}
	case sizeSubterran:
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			return innerSubterranHydr(i)
		case "H":
			return habitableSubterranHydr(i)
		case "O":
			return outerSubterranHydr(i)
		}
	case sizeTerran:
		a := ehex.New().Set(atmo).Value()
		i := dp.Roll("1d6").Sum()
		switch p {
		case "I":
			switch atmo {
			case "0", "1", "2", "3":
				return "0"
			case "5", "6", "8":
				i = dp.Roll("1d6").DM(-10 + a).Sum()
			case "4", "7", "9":
				i = dp.Roll("1d6").DM(-11 + a).Sum()
			case "A", "B", "C", "D", "E", "F":
				return "0"
			}
			if i < 0 {
				return "0"
			}
			return fmt.Sprintf("%v", i)
		case "H":
			return habitableTerranHydr(i)
		case "O":
			return outerTerranHydr(i)
		}
	case sizeSuperterran:
		switch p {
		case "I":

			a := ehex.New().Set(atmo).Value()
			switch atmo {
			case "6", "8":
				i = dp.Roll("1d6").DM(-10 + a).Sum()
			case "7", "9":
				i = dp.Roll("1d6").DM(-11 + a).Sum()
			default:
				return "0"
			}
			if i < 0 {
				i = 0
			}
			return fmt.Sprintf("%v", i)
			//return innerSuperterranHydr(i)
		case "H":
			return habitableSuperterranHydr(i)
		case "O":
			return outerSuperterranHydr(i)
		}
	}
	return "*"
}

func innerSizeType(roll int) string {
	switch roll {
	default:
		return "Error"
	case 2, 3, 4:
		return sizeDwarf
	case 5, 6, 7:
		return sizeMercurian
	case 8, 9:
		return sizeSubterran
	case 10, 11:
		return sizeTerran
	case 12:
		return sizeSuperterran
	}
}

func habitableSizeType(roll int) string {
	switch roll {
	default:
		return "Error"
	case 2:
		return sizeDwarf
	case 3:
		return sizeMercurian
	case 4, 5, 6:
		return sizeSubterran
	case 7, 8, 9, 10, 11:
		return sizeTerran
	case 12:
		return sizeSuperterran
	}
}

func innerDwarfSize(i int) string {
	switch i {
	default:
		return "X"
	case 1, 3, 5:
		return "0"
	case 2, 4, 6:
		return "1"
	}
}
func innerDwarfAtmo(i int) string {
	switch i {
	default:
		return "X"
	case 1, 3, 5:
		return "0"
	case 2, 4, 6:
		return "1"
	}
}
func innerDwarfHydr(roll int) string {
	return "0"
}
func innerMercurianSize(i int) string {
	switch i {
	default:
		return "X"
	case 1, 3, 5:
		return "2"
	case 2, 4, 6:
		return "3"
	}
}
func innerMercurianAtmo(i int) string {
	switch {
	default:
		return "0"
	case i-3 > 0:
		return fmt.Sprintf("%v", i-3)
	}
}
func innerMercurianHydr(roll int) string {
	return "0"
}
func innerSubterranSize(roll int) string {
	switch roll {
	case 1, 2:
		return "4"
	case 3, 4:
		return "5"
	case 5, 6:
		return "6"
	}
	return "X"
}
func innerSubterranAtmo(roll int) string {
	return fmt.Sprintf("%v", roll)
}
func innerSubterranHydr(roll int) string {
	return "0"
}
func innerTerranSize(roll int) string {
	switch roll {
	case 2, 3, 4, 5:
		return "7"
	case 6, 7, 8, 9:
		return "8"
	case 10, 11, 12:
		return "9"
	}
	return "X"
}
func innerTerranAtmo(roll int) string {
	switch roll {
	case 2:
		return "2"
	case 3:
		return "3"
	case 4, 5:
		return "4"
	case 6, 7:
		return "5"
	case 8:
		return "6"
	case 9:
		return "7"
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	}
	return "X"
}
func innerTerranHydr(roll int) string {
	return "X"
}
func innerSuperterranSize(roll int) string {
	switch roll {
	case 2, 3:
		return "A"
	case 4, 5:
		return "B"
	case 6, 7:
		return "C"
	case 8, 9:
		return "D"
	case 10:
		return "E"
	case 11:
		return "F"
	case 12:
		return "G"
	case 13:
		return "H"
	case 14:
		return "J"
	case 15:
		return "K"
	case 16:
		return "L"
	case 17:
		return "M"
	case 18:
		return "N"
	case 19:
		return "P"
	case 20:
		return "Q"
	}
	return "X"
}
func innerSuperterranAtmo(roll int) string {
	switch roll {
	case 2:
		return "6"
	case 3:
		return "7"
	case 4, 5, 6:
		return "8"
	case 7, 8, 9:
		return "9"
	case 10, 11, 12:
		return "E"
	case 13, 14, 15:
		return "A"
	case 16, 17:
		return "B"
	case 18, 19:
		return "C"
	default:
		if roll >= 20 {
			return "D"
		}
	}
	return "X"
}
func innerSuperterranHydr(roll int) string {
	return "X"
}
func habitableDwarfSize(roll int) string {
	return "X"
}
func habitableDwarfAtmo(roll int) string {
	return "X"
}
func habitableDwarfHydr(roll int) string {
	return "X"
}
func habitableMercurianSize(roll int) string {
	return "X"
}
func habitableMercurianAtmo(roll int) string {
	return "X"
}
func habitableMercurianHydr(roll int) string {
	return "X"
}
func habitableSubterranSize(roll int) string {
	switch roll {
	case 1, 2:
		return "4"
	case 3, 4:
		return "5"
	case 5, 6:
		return "6"
	}
	return "X"
}
func habitableSubterranAtmo(roll int) string {
	return "X"
}
func habitableSubterranHydr(roll int) string {
	return "X"
}
func habitableTerranSize(roll int) string {
	return "X"
}
func habitableTerranAtmo(roll int) string {
	return "X"
}
func habitableTerranHydr(roll int) string {
	return "X"
}
func habitableSuperterranSize(roll int) string {
	return "X"
}
func habitableSuperterranAtmo(roll int) string {
	return "X"
}
func habitableSuperterranHydr(roll int) string {
	return "X"
}
func outerDwarfSize(roll int) string {
	return "X"
}
func outerDwarfAtmo(roll int) string {
	return "X"
}
func outerDwarfHydr(roll int) string {
	return "X"
}
func outerMercurianSize(roll int) string {
	return "X"
}
func outerMercurianAtmo(roll int) string {
	return "X"
}
func outerMercurianHydr(roll int) string {
	return "X"
}
func outerSubterranSize(roll int) string {
	switch roll {
	case 1, 2:
		return "4"
	case 3, 4:
		return "5"
	case 5, 6:
		return "6"
	}
	return "X"
}
func outerSubterranAtmo(roll int) string {
	return "X"
}
func outerSubterranHydr(roll int) string {
	return "X"
}
func outerTerranSize(roll int) string {
	return "X"
}
func outerTerranAtmo(roll int) string {
	return "X"
}
func outerTerranHydr(roll int) string {
	return "X"
}
func outerSuperterranSize(roll int) string {
	return "X"
}
func outerSuperterranAtmo(roll int) string {
	return "X"
}
func outerSuperterranHydr(roll int) string {
	return "X"
}

func outerSizeType(roll int) string {
	switch roll {
	default:
		return "Error"
	case 2, 3, 4, 5, 6:
		return sizeDwarf
	case 7, 8, 9:
		return sizeMercurian
	case 10:
		return sizeSubterran
	case 11:
		return sizeTerran
	case 12:
		return sizeSuperterran
	}
}
