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
	if err := gs.setPlanetDetails(); err != nil {
		return err
	}
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

func (gs *GenerationState) setPlanetDetails() error {
	for i, star := range gs.System.Stars {
		for _, orbit := range star.orbitDistances {
			if planet, ok := star.orbit[orbit].(*rockyPlanet); ok == true {
				// if err := planet.rollSizeType(gs.Dice); err != nil {
				// 	return err
				// }
				// if err := planet.rollSize(gs.Dice); err != nil {
				// 	return err
				// }
				// if err := planet.rollAtmo(gs.Dice); err != nil {
				// 	return err
				// }
				// if err := planet.rollHydr(gs.Dice); err != nil {
				// 	return err
				// }
				// if planet.orbit > gs.System.Stars[i].habitableHigh {
				// 	planet.comment = " Cold"
				// }
				// if planet.orbit > gs.System.Stars[i].snowLine {
				// 	planet.comment = " Frozen"
				// }
				planet, err := gs.detailPlanet(planet)
				if err != nil {
					return err
				}
				gs.System.Stars[i].orbit[orbit] = planet
			}
		}
	}
	return nil
}

func (gs *GenerationState) detailPlanet(planet *rockyPlanet) (*rockyPlanet, error) {
	if err := planet.rollSizeType(gs.Dice); err != nil {
		return planet, err
	}
	if err := planet.rollSize(gs.Dice); err != nil {
		return planet, err
	}
	if err := planet.rollAtmo(gs.Dice); err != nil {
		return planet, err
	}
	if err := planet.rollHydr(gs.Dice); err != nil {
		return planet, err
	}

	return planet, nil
}

func (p *rockyPlanet) rollSize(dp *dice.Dicepool) error {
	switch p.habZone {
	default:
		return fmt.Errorf("p.habZone = %v", p.habZone)
	case habZoneInner, habZoneHabitable, habZoneOuter:
	}
	switch p.sizeType {
	default:
		return fmt.Errorf("p.sizeType  = %v", p.sizeType)
	case sizeDwarf:
		r := dp.Roll("1d6").Sum()
		switch r {
		case 1, 3, 5:
			p.sizeCode = "0"
		case 2, 4, 6:
			p.sizeCode = "1"
		}
	case sizeMercurian:
		r := dp.Roll("1d6").Sum()
		switch r {
		case 1, 3, 5:
			p.sizeCode = "2"
		case 2, 4, 6:
			p.sizeCode = "3"
		}
	case sizeSubterran:
		r := dp.Roll("1d6").Sum()
		switch r {
		case 1, 2:
			p.sizeCode = "4"
		case 3, 4:
			p.sizeCode = "5"
		case 5, 6:
			p.sizeCode = "6"
		}
	case sizeTerran:
		r := dp.Roll("2d6").Sum()
		switch r {
		case 2, 3, 4, 5:
			p.sizeCode = "7"
		case 6, 7, 8, 9:
			p.sizeCode = "8"
		case 10, 11, 12:
			p.sizeCode = "9"
		}
	case sizeSuperterran:
		r := dp.Roll("2d10").Sum()
		switch r {
		case 2, 3:
			p.sizeCode = "A"
		case 4, 5:
			p.sizeCode = "B"
		case 6, 7:
			p.sizeCode = "C"
		case 8, 9:
			p.sizeCode = "D"
		case 10:
			p.sizeCode = "E"
		case 11:
			p.sizeCode = "F"
		case 12:
			p.sizeCode = "G"
		case 13:
			p.sizeCode = "H"
		case 14:
			p.sizeCode = "J"
		case 15:
			p.sizeCode = "K"
		case 16:
			p.sizeCode = "L"
		case 17:
			p.sizeCode = "M"
		case 18:
			p.sizeCode = "N"
		case 19:
			p.sizeCode = "P"
		case 20:
			p.sizeCode = "Q"
		}
	}

	return nil
}

func (p *rockyPlanet) rollAtmo(dp *dice.Dicepool) error {
	if p.sizeCode == "" {
		return fmt.Errorf("p.sizeCode is not filled")
	}
	switch p.habZone {
	default:
		return fmt.Errorf("p.habZone = %v", p.habZone)
	case habZoneInner, habZoneHabitable, habZoneOuter:
	}
	switch p.sizeType {
	default:
		return fmt.Errorf("p.sizeType  = %v", p.sizeType)
	case sizeDwarf:
		r := dp.Roll("1d6").Sum()
		switch r {
		case 1, 3, 5:
			p.atmoCode = "0"
		case 2, 4, 6:
			p.atmoCode = "1"
		}
	case sizeMercurian:
		r := dp.Roll("1d6").DM(-3).Sum()
		if r < 3 {
			r = 0
		}
		p.atmoCode = ehex.New().Set(r).Code()
	case sizeSubterran:
		dm := 0
		if p.habZone == habZoneOuter {
			dm = 1
		}
		r := dp.Roll("1d6").DM(dm).Sum()
		p.atmoCode = ehex.New().Set(r).Code()
	case sizeTerran:
		switch p.habZone {
		case habZoneInner:
			r := dp.Roll("2d6").Sum()
			switch r {
			case 2:
				p.atmoCode = "2"
			case 3:
				p.atmoCode = "3"
			case 4, 5:
				p.atmoCode = "4"
			case 6, 7:
				p.atmoCode = "5"
			case 8:
				p.atmoCode = "6"
			case 9:
				p.atmoCode = "7"
			case 10:
				p.atmoCode = "A"
			case 11:
				p.atmoCode = "B"
			case 12:
				p.atmoCode = "C"
			}
		case habZoneHabitable:
			s := ehex.New().Set(p.sizeCode).Value()
			r := dp.Roll("2d6").DM(s - 7).Sum()
			p.atmoCode = ehex.New().Set(r).Code()
		case habZoneOuter:
			r := dp.Roll("2d6").Sum()
			switch r {
			case 2:
				p.atmoCode = "0"
			case 3, 4:
				p.atmoCode = "1"
			case 5, 6:
				p.atmoCode = "A"
			case 7, 8:
				p.atmoCode = "B"
			case 9:
				p.atmoCode = "C"
			case 10:
				p.atmoCode = "D"
			case 11:
				p.atmoCode = "E"
			case 12:
				p.atmoCode = "F"
			}
		}
	case sizeSuperterran:
		dm := -999
		fmt.Println("-----------------", p)
		switch p.habZone {
		case habZoneInner, habZoneHabitable:
			switch p.sizeCode {
			case "A", "B", "C", "D", "E":
				dm = 0
			case "F", "G", "H", "J", "K":
				dm = 3
			case "L", "M", "N", "P", "Q":
				dm = 6
			default:
				return fmt.Errorf("unknown p.sizeCode (%v)", p.sizeCode)
			}
			r := dp.Roll("2d10").DM(dm).Sum()
			switch r {
			default:
				p.atmoCode = "D"
			case 2:
				p.atmoCode = "6"
			case 3:
				p.atmoCode = "7"
			case 4, 5, 6:
				p.atmoCode = "8"
			case 7, 8, 9:
				p.atmoCode = "9"
			case 10, 11, 12:
				p.atmoCode = "E"
			case 13, 14, 15:
				p.atmoCode = "A"
			case 16, 17:
				p.atmoCode = "B"
			case 18, 19:
				p.atmoCode = "C"
			}
		case habZoneOuter:
			switch p.sizeCode {
			case "A", "B", "C", "D", "E":
				dm = 0
			case "F", "G", "H", "J", "K":
				dm = 1
			case "L", "M", "N", "P", "Q":
				dm = 3
			default:
				return fmt.Errorf("unknown p.sizeCode (%v)", p.sizeCode)
			}
			r := dp.Roll("1d6").DM(dm).Sum()
			switch r {
			default:
				p.atmoCode = "D"
			case 1:
				p.atmoCode = "E"
			case 2:
				p.atmoCode = "A"
			case 3:
				p.atmoCode = "B"
			case 4:
				p.atmoCode = "C"

			}
		}
	}
	if p.atmoCode == "" {
		return fmt.Errorf("not asigned ATMO")
	}
	return nil
}

func (p *rockyPlanet) rollHydr(dp *dice.Dicepool) error {
	if p.atmoCode == "" {
		fmt.Errorf("p.atmoCode is not filled")
	}
	switch p.habZone {
	default:
		return fmt.Errorf("p.habZone = %v", p.habZone)
	case habZoneInner, habZoneHabitable, habZoneOuter:
	}
	switch p.sizeType {
	default:
		return fmt.Errorf("p.sizeType  = %v", p.sizeType)
	case sizeDwarf, sizeMercurian:
		p.hydrCode = "0"
	case sizeSubterran:
		switch p.habZone {
		case habZoneInner:
			p.hydrCode = "0"
		case habZoneHabitable:
			switch p.atmoCode {
			default:
				a := ehex.New().Set(p.atmoCode).Value()
				r := dp.Roll("2d6").DM(-9 + a).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			case "0", "1", "2", "3":
				p.hydrCode = "0"
			}
		case habZoneOuter:
			dm := 0
			switch p.atmoCode {
			default:
				switch p.atmoCode {
				case "4", "5":
					dm = -1
				case "6", "7":
					dm = 1
				}
				r := dp.Roll("1d6").DM(dm).Sum()
				switch r {
				case 0, 1:
					p.hydrCode = "0"
				case 2:
					p.hydrCode = "2"
				case 3:
					p.hydrCode = "4"
				case 4:
					p.hydrCode = "6"
				case 5:
					p.hydrCode = "8"
				case 6, 7:
					p.hydrCode = "A"
				}
				p.hydrCode = ehex.New().Set(r).Code()
			case "0", "1", "2", "3":
				p.hydrCode = "0"
			}
		}
	case sizeTerran:
		switch p.habZone {
		case habZoneInner:
			switch p.atmoCode {
			case "0", "1", "2", "3", "A", "B", "C", "D", "E", "F":
				p.hydrCode = "0"
			case "5", "6", "8":
				a := ehex.New().Set(p.atmoCode).Value()
				r := dp.Roll("1d6").DM(-10 + a).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			case "4", "7", "9":
				a := ehex.New().Set(p.atmoCode).Value()
				r := dp.Roll("1d6").DM(-11 + a).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			}
		case habZoneHabitable:
			switch p.atmoCode {
			case "0", "1", "2", "3", "A", "B", "C", "D", "E", "F":
				p.hydrCode = "0"
			default:
				a := ehex.New().Set(p.atmoCode).Value()
				r := dp.Roll("2d6").DM(-7 + a).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			}
		case habZoneOuter:
			switch p.atmoCode {
			default:
				a := ehex.New().Set(p.atmoCode).Value()
				r := dp.Roll("2d6").DM(-7 + a).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			case "0", "1":
				p.hydrCode = "0"
			case "A", "B", "C", "D", "E", "F":
				r := dp.Roll("1d6").Sum()
				switch r {
				case 1:
					p.hydrCode = "0"
				case 2:
					p.hydrCode = "2"
				case 3:
					p.hydrCode = "4"
				case 4:
					p.hydrCode = "6"
				case 5:
					p.hydrCode = "8"
				case 6:
					p.hydrCode = "A"
				}
			}
		}
	case sizeSuperterran:
		a := ehex.New().Set(p.atmoCode).Value()
		switch p.habZone {
		case habZoneInner:
			switch p.atmoCode {
			case "6", "8":
				r := dp.Roll("1d6").DM(a - 10).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			case "7", "9":
				r := dp.Roll("1d6").DM(a - 11).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			default:
				p.hydrCode = "0"
			}
		case habZoneHabitable, habZoneOuter:
			if p.sizeType == sizeSuperterran {
				//return fmt.Errorf("STOP")
			}
			switch p.atmoCode {
			case "A", "B", "C", "D", "E", "F":
				p.hydrCode = "0"
			default:
				r := dp.Roll("2d6").DM(a - 7).Sum()
				if r < 0 {
					r = 0
				}
				p.hydrCode = ehex.New().Set(r).Code()
			}
			//case habZoneOuter: в книге количество воды почему-то зависит от размера супертерана что не логично
		}
	}
	return nil
}

func (p *rockyPlanet) rollSizeType(dp *dice.Dicepool) error {
	switch p.sizeCode {
	case "":
	case "0", "1":
		p.sizeType = sizeDwarf
	case "2", "3":
		p.sizeType = sizeMercurian
	case "4", "5", "6":
		p.sizeType = sizeSubterran
	case "7", "8", "9":
		p.sizeType = sizeTerran
	default:
		p.sizeType = sizeSubterran
	}
	if p.sizeType != "" {
		return nil
	}
	sizeTypeRoll := dp.Roll("2d6").Sum()
	switch p.habZone {
	case habZoneInner:
		p.sizeType = innerSizeType(sizeTypeRoll)
	case habZoneHabitable:
		p.sizeType = habitableSizeType(sizeTypeRoll)
	case habZoneOuter:
		p.sizeType = outerSizeType(sizeTypeRoll)
	default:
		return fmt.Errorf("habitable zone is '%v', %v", p.habZone, p)
	}
	return nil
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
