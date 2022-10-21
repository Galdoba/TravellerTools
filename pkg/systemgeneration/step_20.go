package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

func (gs *GenerationState) Step20() error {
	if gs.NextStep != 20 {
		return fmt.Errorf("not actual step")
	}
	printSystem(gs)
	switch gs.System.populationType {
	default:
		fmt.Println(gs.System.populationType)
		return fmt.Errorf("gs.System.populationType = %v", gs.System.populationType)
	case PopulationAuto, PopulationON, PopulationOFF:
		if err := gs.PopulateBodies(); err != nil {
			return err
		}
	}
	printSystem(gs)
	gs.NextStep = 99
	return nil
}

func (gs *GenerationState) PopulateBodies() error {
	fmt.Println(gs.System.populationType)
	mwuwp, err := uwp.FromString(gs.System.MW_UWP)
	if err != nil {
		return err
	}
	for s, star := range gs.System.Stars {
		fmt.Println("Populating star", s)
		for p, orbit := range star.orbitDistances {
			fmt.Println("Populating orbit", p)
			if planet, ok := star.orbit[orbit].(*rockyPlanet); ok == true {
				fmt.Println("planet", p)

				planet.populate(gs.Dice, mwuwp, gs.System.populationType)
				for m, _ := range planet.moons {
					fmt.Println("moon", m)
				}
			}
			if gg, ok := star.orbit[orbit].(*ggiant); ok == true {
				fmt.Println("gg", p)
				for m, _ := range gg.moons {
					fmt.Println("moon", m)
				}
			}
			if _, ok := star.orbit[orbit].(*belt); ok == true {
				fmt.Println("belt", p)
			}
		}
	}
	return nil
}

func (p *rockyPlanet) populate(dp *dice.Dicepool, mwuwp uwp.UWP, popType string) error {
	mwTL := mwuwp.TL()

	switch popType {
	default:
		return fmt.Errorf("????")
	case PopulationOFF:
		p.uwpStr = fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode)
	case PopulationON:
		fmt.Println(p.comment, "COMMENT")
		if strings.Contains(p.comment, "Mainworld") {
			p.injectUwpCodes(mwuwp)
			return nil
		}
		if mwTL < 8 {
			uwpW, _ := uwp.FromString(fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode))
			p.injectUwpCodes(uwpW)
			return nil
		}
		switch p.habZone {
		case habZoneInner:
			switch p.sizeType {
			case sizeDwarf:
				pop := dp.Roll("1d6").DM(-3).Sum()
				if pop < 0 {
					uwpW, _ := uwp.FromString(fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode))
					p.injectUwpCodes(uwpW)
					return nil
				}
				if pop == 0 {
					p.comment = "Under Construction "
				}
				p.popCode = fmt.Sprintf("%v", pop)
				switch dp.Roll("2d6").Sum() {
				case 2:
					uwpW, _ := uwp.FromString(fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode))
					p.injectUwpCodes(uwpW)
				case 3, 4, 5:
					p.govCode = "1"
					p.lawCode = ehex.CodeOf(dp.RollSum("1d6+4"))
					p.comment += "Corporative Compound"
				case 6, 7:
					p.govCode = "6"
					p.lawCode = ehex.New().Set(dp.Roll("1d6+4").Sum()).Code()
					p.comment += "Research Station"
				case 8, 9:
					p.govCode = "6"
					p.lawCode = ehex.New().Set(dp.Roll("1d6+5").Sum()).Code()
					p.comment += "Military Base"
				case 10, 11:
					p.govCode = "6"
					p.lawCode = "B"
					p.comment += "Prison"
				case 12:
					//todo: нужен метод добрасывания отдельных участков UWP
					g := ehex.New().Set(p.popCode).Value()
					p.govCode = ehex.New().Set(dp.Roll("2d6").DM(-7 + g)).Code()
					p.lawCode = ehex.New().Set(dp.Roll("2d6").DM(-7 + g)).Code()
				}
			}
		}
	}
	return nil
}

func (p *rockyPlanet) injectUwpCodes(uwpS uwp.UWP) {
	p.port = uwpS.Starport()
	p.sizeCode = ehex.New().Set(uwpS.Size()).Code()
	p.atmoCode = ehex.New().Set(uwpS.Atmo()).Code()
	p.hydrCode = ehex.New().Set(uwpS.Hydr()).Code()
	p.popCode = ehex.New().Set(uwpS.Pops()).Code()
	p.govCode = ehex.New().Set(uwpS.Govr()).Code()
	p.lawCode = ehex.New().Set(uwpS.Laws()).Code()
	p.tlCode = ehex.New().Set(uwpS.TL()).Code()
}
