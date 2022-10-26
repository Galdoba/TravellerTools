package systemgeneration

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/utils"
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

				if err := planet.populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
					return nil
				}
				for m, _ := range planet.moons {
					fmt.Println("moon", m)
					if err := planet.moons[m].populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
						return err
					}

				}
			}
			if gg, ok := star.orbit[orbit].(*ggiant); ok == true {
				fmt.Println("gg", p)
				for m, _ := range gg.moons {
					fmt.Println("moon", m)
					if err := gg.moons[m].populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
						return err
					}
				}
			}
			if _, ok := star.orbit[orbit].(*belt); ok == true {
				fmt.Println("belt", p)
			}
		}
	}
	return nil
}

func notSettledUWP(p *rockyPlanet) error {
	fmt.Println("NOT SETTLED ---")
	notSettled := fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode)
	p.comment = ""
	p.injectUwpCodes(uwp.Inject(notSettled))
	return nil
}

func (p *rockyPlanet) defineAsCorporative(dice *dice.Dicepool, pop int) {
	gov := 1
	fmt.Println("Corporative")
	pop = utils.Min(pop, 5)
	p.comment += fmt.Sprintf("%v Corporatin %v Facility", dice.PickStrOnly([]string{"Local", "Major"}), dice.PickStrOnly([]string{"Mining", "Science", "Storage"}))
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsMining(dice *dice.Dicepool, pop int, local bool) {
	gov := 1
	fmt.Println("Mining")
	pop = utils.Min(pop, 5)
	p.comment += fmt.Sprintf("%v Corporatin %v Facility", dice.PickStrOnly([]string{"Local", "Major"}), dice.PickStrOnly([]string{"Mining", "Science", "Storage"}))
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsResearch(dice *dice.Dicepool, pop int) {
	fmt.Println("Research")
	gov := 6
	pop = utils.Min(pop, 3)
	p.comment += fmt.Sprintf("Research Base")
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsMilitary(dice *dice.Dicepool, pop int) {
	fmt.Println("Military")
	gov := 6
	pop = utils.Min(pop, 4)
	p.comment += fmt.Sprintf("Military Base")
	law := dice.Sroll("1d6+5")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsPrison(dice *dice.Dicepool, pop int) {
	fmt.Println("Prison")
	gov := 6
	pop = utils.Min(pop, 4)
	p.comment += fmt.Sprintf("Prison Instalation")
	law := dice.Sroll("1d11")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsIndependent(dice *dice.Dicepool, pop int) {
	gov := utils.Max(dice.Sroll("2d6-7")+pop, 0)
	pop = utils.Min(pop, 8)
	p.comment += fmt.Sprintf("Independent World")
	law := utils.Max(dice.Sroll("2d6-7")+gov, 0)
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsColony(dice *dice.Dicepool, pop int) {
	gov := 6
	pop = utils.Min(pop, 8)
	p.comment += fmt.Sprintf("Colony")
	law := dice.Sroll("2d6-1")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsFuelBase(dice *dice.Dicepool, pop int) {
	gov := 6
	pop = utils.Min(pop, 8)
	p.comment += fmt.Sprintf("Colony")
	law := dice.Sroll("2d6-1")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) populate(dice *dice.Dicepool, mwuwp uwp.UWP, popType string) error {
	mwTL := mwuwp.TL()
	fmt.Println("===Populate planet", p.orbit, p.habZone, p.sizeType)
	maxPopDM := mwuwp.Pops() - 5
	switch popType {
	default:
		return fmt.Errorf("????")
	case PopulationOFF:
		p.uwpStr = fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode)
	case PopulationON:
		fmt.Println(p.comment, p.sizeType, p.habZone, "COMMENT")
		if strings.Contains(p.comment, "Mainworld") {
			fmt.Println("IS MW")
			p.comment = "Mainworld"
			p.injectUwpCodes(mwuwp)
			return nil
		}
		if mwTL < 8 {
			fmt.Println("NOT SETTLED low TL")
			return notSettledUWP(p)
		}
		p.comment = ""
		switch p.habZone {
		case habZoneInner:
			switch p.sizeType {
			case sizeDwarf, sizeMercurian:
				pop := maxPopDM + dice.Sroll("2d6-9")
				fmt.Println("Pop =", pop)
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3, 4, 5:
					p.defineAsCorporative(dice, pop)
				case 6, 7:
					p.defineAsResearch(dice, pop)
				case 8, 9:
					p.defineAsMilitary(dice, pop)
				case 10, 11:
					p.defineAsPrison(dice, pop)
				case 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeSubterran:
				pop := maxPopDM + dice.Sroll("2d6-8")
				fmt.Println("Pop =", pop)
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3, 4, 5, 6, 7:
					p.defineAsCorporative(dice, pop)
				case 8, 9:
					p.defineAsResearch(dice, pop)
				case 10, 11:
					p.defineAsMilitary(dice, pop)
				case 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeTerran, sizeSuperterran:
				pop := 0
				switch p.atmoCode {
				default:
					pop = dice.Sroll("2d6-4")
				case "0", "1", "2", "3", "4", "5":
					pop = dice.Sroll("2d6-8")
				case "A", "B", "C", "D", "E", "F":
					pop = dice.Sroll("2d6-9")
				}
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3, 4:
					p.defineAsCorporative(dice, pop)
				case 5, 6:
					p.defineAsResearch(dice, pop)
				case 7, 8:
					p.defineAsMilitary(dice, pop)
				case 9, 10, 11, 12:
					p.defineAsIndependent(dice, pop)
				}
			}
		case habZoneHabitable:
			switch p.sizeType {
			case sizeDwarf, sizeMercurian:
				pop := maxPopDM + dice.Sroll("2d6-9")

				fmt.Println("Pop =", pop)
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3, 4, 5:
					p.defineAsCorporative(dice, pop)
				case 6, 7:
					p.defineAsResearch(dice, pop)
				case 8, 9:
					p.defineAsMilitary(dice, pop)
				case 10, 11:
					p.defineAsPrison(dice, pop)
				case 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeSubterran:
				pop := maxPopDM + dice.Sroll("2d6-8")
				fmt.Println("Pop =", pop)
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3, 4, 5, 6, 7:
					p.defineAsCorporative(dice, pop)
				case 8, 9:
					p.defineAsResearch(dice, pop)
				case 10, 11:
					p.defineAsMilitary(dice, pop)
				case 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeTerran, sizeSuperterran:
				pop := 0
				switch p.atmoCode {
				default:
					pop = dice.Sroll("2d6-4")
				case "0", "1", "2", "3", "4", "5":
					pop = dice.Sroll("2d6-8")
				case "A", "B", "C", "D", "E", "F":
					pop = dice.Sroll("2d6-9")
				}
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3:
					p.defineAsCorporative(dice, pop)
				case 4, 5:
					p.defineAsResearch(dice, pop)
				case 6, 7, 8, 9:
					p.defineAsColony(dice, pop)
				case 10, 11, 12:
					p.defineAsIndependent(dice, pop)
				}
			}
		case habZoneOuter:
			fmt.Println("DO NOTHING YET")
			switch p.sizeType {
			case sizeDwarf, sizeMercurian:
				pop := maxPopDM + dice.Sroll("2d6-9")
				fmt.Println("Pop =", pop)
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3:
					p.defineAsCorporative(dice, pop)
				case 4, 5:
					p.defineAsResearch(dice, pop)
				case 6, 7:
					p.defineAsMilitary(dice, pop)
				case 8, 9:
					p.defineAsFuelBase(dice, pop)
				case 10, 11:
					p.defineAsPrison(dice, pop)
				case 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeSubterran:
				pop := maxPopDM + dice.Sroll("2d6-8")
				fmt.Println("Pop =", pop)
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3:
					p.defineAsCorporative(dice, pop)
				case 4, 5:
					p.defineAsResearch(dice, pop)
				case 6, 7:
					p.defineAsMilitary(dice, pop)
				case 8, 9:
					p.defineAsFuelBase(dice, pop)
				case 10, 11:
					p.defineAsPrison(dice, pop)
				case 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeTerran:
				pop := maxPopDM + dice.Sroll("2d6-8")
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3:
					p.defineAsCorporative(dice, pop)
				case 4, 5:
					p.defineAsResearch(dice, pop)
				case 6, 7:
					p.defineAsFuelBase(dice, pop)
				case 8, 9:
					p.defineAsMilitary(dice, pop)
				case 10, 11, 12:
					p.defineAsIndependent(dice, pop)
				}
			case sizeSuperterran:
				pop := maxPopDM + dice.Sroll("2d6-9")
				if pop < 0 {
					fmt.Println("NOT SETTLED")
					return notSettledUWP(p)
				}
				if pop == 0 {
					p.comment += "Constructing "
				}
				switch dice.Sroll("2d6") {
				case 2:
					return notSettledUWP(p)
				case 3, 4:
					p.defineAsCorporative(dice, pop)
				case 5, 6:
					p.defineAsResearch(dice, pop)
				case 7, 8:
					p.defineAsMilitary(dice, pop)
				case 9, 10, 11, 12:
					p.defineAsIndependent(dice, pop)
				}
			}
		}
	}
	return nil
}

func (b *belt) populate(dice *dice.Dicepool, mwuwp uwp.UWP, popType string) error {

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
	p.uwpStr = uwpS.Starport() + ehex.New().Set(uwpS.Size()).Code() + ehex.New().Set(uwpS.Atmo()).Code() + ehex.New().Set(uwpS.Hydr()).Code() + ehex.New().Set(uwpS.Pops()).Code() + ehex.New().Set(uwpS.Govr()).Code() + ehex.New().Set(uwpS.Laws()).Code() + "-" + ehex.New().Set(uwpS.TL()).Code()
	fmt.Println(p.uwpStr)
}
