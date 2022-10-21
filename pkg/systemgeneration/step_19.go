package systemgeneration

import (
	"fmt"
	"strings"
)

func (gs *GenerationState) Step19() error {
	if gs.NextStep != 19 {
		return fmt.Errorf("not actual step")
	}
	//ЗАРЕЗЕРВИРОВАНО
	if err := gs.placeMoons(); err != nil {
		return err
	}
	gs.ConcludedStep = 19
	gs.NextStep = 20
	switch gs.NextStep {
	case 20:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}
	return nil
}

func (gs *GenerationState) placeMoons() error {
	for _, star := range gs.System.Stars {
		for _, body := range star.orbit {
			if planet, ok := body.(*rockyPlanet); ok {
				moons := -10
				switch planet.sizeCode {
				case "0":
					moons = 0
				case "1", "2", "3":
					moons = gs.Dice.Roll("1d6").DM(-4).Sum()
				case "4", "5", "6":
					moons = gs.Dice.Roll("1d6").DM(-3).Sum()
				default:
					moons = gs.Dice.Roll("1d6").DM(-2).Sum()
				}
				if moons < -9 {
					return fmt.Errorf("moons was not allocated for planet [%v]", planet)
				}
				if moons < 0 {
					moons = 0
				}
				for i := 0; i < moons; i++ {
					moon := &rockyPlanet{}
					moon.habZone = planet.habZone
					moon.star = planet.star
					moonSize := -1
					switch planet.sizeType {
					case "1", "2", "3", "4", "5":
						moonSize = gs.Dice.Roll("1d2").DM(-1).Sum()
					case "6", "7", "8":
						moonSize = gs.Dice.Roll("1d6").DM(-3).Sum()
					case "A", "B", "C", "D", "E", "F":
						moonSize = gs.Dice.Roll("1d6").DM(-2).Sum()
					case "G", "H", "J", "K", "L", "M", "N", "P", "Q":
						moonSize = gs.Dice.Roll("1d6").DM(-1).Sum()
					}
					if moonSize < 0 {
						moonSize = 0
					}
					moon.sizeCode = fmt.Sprintf("%v", moonSize)
					moon.orbit = planet.orbit
					moon, err := gs.detailPlanet(moon)
					if err != nil {
						return err
					}
					moon.moonOrbit = 2 * gs.Dice.Roll("2d10").Sum()
					planet.moons = append(planet.moons, moon)
				}

			}
			if ggigant, ok := body.(*ggiant); ok {
				moons := -10
				switch {
				case strings.Contains(ggigant.descr, GasGigantNeptunian):
					moons = gs.Dice.Roll("1d6").DM(-1).Sum()
				case strings.Contains(ggigant.descr, GasGigantJovian):
					moons = gs.Dice.Roll("1d6").DM(-1).Sum()
				}
				if moons < -9 {
					return fmt.Errorf("moons was not allocated for ggigant [%v]", ggigant)
				}
				if moons < 0 {
					moons = 0
				}
				for i := 0; i < moons; i++ {
					moon := &rockyPlanet{}
					moon.habZone = star.habZone(ggigant.spawnedAtAU)
					moon.star = star.Describe()
					moonSize := -1
					switch {
					case strings.Contains(ggigant.descr, GasGigantNeptunian):
						moonSize = gs.Dice.Roll("1d6").Sum()
					case strings.Contains(ggigant.descr, GasGigantJovian):
						moonSize = gs.Dice.Roll("1d6").DM(1).Sum()
					}
					if moonSize < 0 {
						moonSize = 0
					}
					moon.sizeCode = fmt.Sprintf("%v", moonSize)
					moon, err := gs.detailPlanet(moon)
					if err != nil {
						return err
					}
					moon.moonOrbit = 5*gs.Dice.Roll("1d6").Sum() + gs.Dice.Roll("1d10").DM(-1).Sum()
					ggigant.moons = append(ggigant.moons, moon)
				}
				ring := "No ring"
				if gs.Dice.Roll("1d3").Sum() == 3 {
					dm := 0
					switch {
					case strings.Contains(ggigant.descr, GasGigantJovian):
						dm = 2
					}
					switch gs.Dice.Roll("1d10").DM(dm).Sum() {
					case 1, 2, 3, 4:
						ring = "No ring"
					case 5, 6:
						ring = fmt.Sprintf("A ring system with %v rings with a width of %v km", gs.Dice.Roll("1d10").Sum(), gs.Dice.Roll("1d10").Sum())
					case 7, 8:
						ring = fmt.Sprintf("A ring system with %v rings with a width of %v km", gs.Dice.Roll("1d6").Sum(), gs.Dice.Roll("1d100").Sum())
					case 9, 10:
						ring = fmt.Sprintf("A ring system with %v rings with a width of %v km", gs.Dice.Roll("1d10").Sum(), gs.Dice.Roll("1d100").Sum())
					case 11, 12:
						ring = fmt.Sprintf("A ring system with %v rings with a width of %v km", gs.Dice.Roll("1d10").Sum(), gs.Dice.Roll("1d10000").Sum())
					}
				}
				ggigant.ring = ring

			}
		}
		star.cleanOrbitDistances()
	}
	return nil
}
