package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

func (gs *GenerationState) Step16() error {
	if gs.NextStep != 16 {
		return fmt.Errorf("not actual step")
	}
	if err := gs.setBeltDetails(); err != nil {
		return err
	}
	if err := gs.adjustBeltPositions(); err != nil {
		return err
	}
	gs.ConcludedStep = 16
	gs.NextStep = 17
	switch gs.NextStep {
	case 17:
	default:
		return fmt.Errorf("gs.NextStep imposible")
	}

	return nil
}

func (gs *GenerationState) placeBelts() error {
	beltNum := gs.System.Belts
	beltMarkers := gs.canPlaceBodies()
	if beltNum > freeSlots(beltMarkers) {
		beltNum = freeSlots(beltMarkers)
	}
	for i := 0; i < beltNum; i++ {
		st := gs.Dice.Roll(fmt.Sprintf("1d%v", len(gs.System.Stars))).DM(-1).Sum()
		star := gs.System.Stars[st]
		if len(star.orbitDistances) < 1 {
			i--
			continue
		}
		try := 0
		for {
			orb := gs.Dice.Roll(fmt.Sprintf("1d%v", len(star.orbitDistances))).DM(-1).Sum()
			if orb-try < 0 {
				i--
				break
			}
			dist := star.orbitDistances[orb-try]
			if v, ok := star.orbit[dist]; ok == true {
				if !strings.Contains(v.Describe(), "empty orbit") {
					try++
					continue
				}
				belt := &belt{num: i + 1, star: star.Describe(), orbit: dist, zone: habZoneInner, comment: "Belt"}
				if belt.orbit >= star.snowLine {
					belt.zone = habZoneOuter
				}
				gs.System.Stars[st].orbit[dist] = belt
				break
			}
		}

	}

	return nil
}

func (gs *GenerationState) adjustBeltPositions() error {
	for _, star := range gs.System.Stars {
		for _, orbit := range star.orbitDistances {
			if belt, ok := star.orbit[orbit].(*belt); ok == true {
				ajusted := false
				for o := belt.orbit; o > belt.lowBorder; o = o - 0.01 {
					o = roundFloat(o, 2)
					if o == belt.orbit {
						continue
					}
					if _, ok := star.orbit[o]; ok == true || o < star.innerLimit {
						belt.lowBorder = roundFloat(o+0.1, 2)
						ajusted = true
						break
					}
				}
				for o := belt.orbit; o < belt.hiBorder; o = o + 0.01 {
					o = roundFloat(o, 2)
					if o == belt.orbit {
						continue
					}
					if _, ok := star.orbit[o]; ok || o > star.outerLimit {
						belt.hiBorder = roundFloat(o-0.1, 2)
						ajusted = true
						break
					}
				}
				if ajusted {
					belt.width = roundFloat((belt.lowBorder+belt.hiBorder)/2, 2)
					belt.orbit = roundFloat(belt.lowBorder+belt.width/2, 2)
					delete(star.orbit, orbit)
					star.orbit[belt.orbit] = belt
				}
			}

		}
	}
	return nil
}

func canStayInBelt(body StellarBody) bool {
	if planet, ok := body.(*rockyPlanet); ok == true {
		switch planet.sizeCode {
		case "0", "1", "2", "3":
			return true
		}
	}
	return false
}

func (gs *GenerationState) setBeltDetails() error {
	if err := gs.placeBelts(); err != nil {
		return err
	}
	for _, star := range gs.System.Stars {
		for _, orbit := range star.orbitDistances {
			if belt, ok := star.orbit[orbit].(*belt); ok == true {
				if err := belt.rollComposition(gs.Dice); err != nil {
					return err
				}
				if err := belt.rollAsteroidsMajorSize(gs.Dice); err != nil {
					return err
				}
				if err := belt.rollWidth(gs.Dice); err != nil {
					return err
				}

			}
		}
	}
	return nil
}

func (b *belt) rollWidth(dp *dice.Dicepool) error {
	r := dp.Roll("2d6").Sum()
	switch b.zone {
	default:
		return fmt.Errorf("unknown belt zone: %v", b)
	case habZoneInner:
		b.width = innerWidth(r)
	case habZoneOuter:
		b.width = outerWidth(r)
	}
	if b.width == 0.0 {
		return fmt.Errorf("belt.width = UNDEFINED")
	}
	b.lowBorder = roundFloat(b.orbit-(b.width/2), 3)
	b.hiBorder = roundFloat(b.orbit+(b.width/2), 3)
	return nil
}

func innerWidth(i int) float64 {
	switch i {
	case 2:
		return 0.001
	case 3:
		return 0.005
	case 4:
		return 0.010
	case 5:
		return 0.025
	case 6:
		return 0.050
	case 7:
		return 0.075
	case 8:
		return 0.100
	case 9:
		return 0.125
	case 10:
		return 0.150
	case 11:
		return 0.175
	case 12:
		return 0.200
	}
	return 0
}

func outerWidth(i int) float64 {
	switch i {
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 7
	case 5:
		return 10
	case 6:
		return 12
	case 7:
		return 15
	case 8:
		return 20
	case 9:
		return 25
	case 10:
		return 30
	case 11:
		return 35
	case 12:
		return 40
	}
	return 0
}

func (b *belt) rollComposition(dp *dice.Dicepool) error {
	r := dp.Roll("2d6").Sum()
	switch b.zone {
	default:
		return fmt.Errorf("unknown belt zone: %v", b)
	case habZoneInner:
		b.composition = innerComposition(r)
	case habZoneOuter:
		b.composition = outerComposition(r)
	}
	if b.composition == "UNDEFINED" {
		return fmt.Errorf("belt.composition = UNDEFINED")
	}
	return nil
}

func innerComposition(i int) string {
	switch i {
	case 2:
		return "85% carbonacerous, 13% silicate, 2% metal"
	case 3, 4:
		return "80% carbonacerous, 15% silicate, 5% metal"
	case 5, 6, 7:
		return "75% carbonacerous, 17% silicate, 8% metal"
	case 8, 9:
		return "75% carbonacerous, 15% silicate, 10% metal"
	case 10, 11:
		return "75% carbonacerous, 13% silicate, 12% metal"
	case 12:
		return "78% carbonacerous, 10% silicate, 12% metal"
	}
	return "UNDEFINED"
}
func outerComposition(i int) string {
	switch i {
	case 2:
		return "70% rock, 20% hydrocarbons, 10% water ice"
	case 3, 4:
		return "50% rock, 40% hydrocarbons, 10% water ice"
	case 5, 6, 7:
		return "40% rock, 50% hydrocarbons, 10% water ice"
	case 8, 9:
		return "30% rock, 60% hydrocarbons, 20% water ice"
	case 10, 11:
		return "10% rock, 70% hydrocarbons, 20% water ice"
	case 12:
		return "70% hydrocarbons, 30% water ice"
	}
	return "UNDEFINED"
}

func (b *belt) rollAsteroidsMajorSize(dp *dice.Dicepool) error {
	r := dp.Roll("2d6").Sum()
	b.majorSizeAst = asteroidsMajorSize(r)
	if b.majorSizeAst == 0 {
		return fmt.Errorf("belt asteroids major size UNDEFINED")
	}
	return nil
}

func asteroidsMajorSize(i int) int {
	switch i {
	case 2:
		return 1
	case 3:
		return 10
	case 4:
		return 50
	case 5:
		return 100
	case 6:
		return 250
	case 7:
		return 500
	case 8:
		return 1000
	case 9:
		return 10000
	case 10:
		return 25000
	case 11:
		return 50000
	case 12:
		return 100000
	}
	return 0
}
