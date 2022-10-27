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
		//fmt.Println(gs.System.populationType)
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
	for _, star := range gs.System.Stars {
		for _, orbit := range star.orbitDistances {
			if planet, ok := star.orbit[orbit].(*rockyPlanet); ok == true {
				if err := planet.populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
					return nil
				}
				for m, _ := range planet.moons {
					if err := planet.moons[m].populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
						return err
					}
				}
			}
			if gg, ok := star.orbit[orbit].(*ggiant); ok == true {
				for m, _ := range gg.moons {
					if err := gg.moons[m].populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
						return err
					}
				}
			}
			if blt, ok := star.orbit[orbit].(*belt); ok == true {
				if err := blt.populate(gs.Dice, mwuwp, gs.System.populationType); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func notSettledUWP(p *rockyPlanet) error {
	notSettled := fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode)
	p.comment = ""
	p.injectUwpCodes(uwp.Inject(notSettled))
	return nil
}

func notSettledUWPBelt(b *belt) error {
	notSettled := fmt.Sprintf("X000000-0")
	b.comment = b.composition
	b.injectUwpCodes(uwp.Inject(notSettled))
	return nil
}

func (p *rockyPlanet) defineAsCorporative(dice *dice.Dicepool, pop int) {
	gov := 1
	pop = utils.Min(pop, 5)
	p.comment += fmt.Sprintf("%v Corporation %v Facility", dice.PickStrOnly([]string{"Local", "Major"}), dice.PickStrOnly([]string{"Mining", "Science", "Storage"}))
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsMining(dice *dice.Dicepool, pop int, local bool) {
	gov := 1

	pop = utils.Min(pop, 5)
	p.comment += fmt.Sprintf("%v Corporation %v Facility", dice.PickStrOnly([]string{"Local", "Major"}), dice.PickStrOnly([]string{"Mining", "Science", "Storage"}))
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsResearch(dice *dice.Dicepool, pop int) {

	gov := 6
	pop = utils.Min(pop, 3)
	p.comment += fmt.Sprintf("Research Base")
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsMilitary(dice *dice.Dicepool, pop int) {

	gov := 6
	pop = utils.Min(pop, 4)
	p.comment += fmt.Sprintf("Military Base")
	law := dice.Sroll("1d6+5")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", p.sizeCode, p.atmoCode, p.hydrCode, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *rockyPlanet) defineAsPrison(dice *dice.Dicepool, pop int) {

	gov := 6
	pop = utils.Min(pop, 4)
	p.comment += fmt.Sprintf("Prison Facility")
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
	maxPopDM := mwuwp.Pops() - 5
	switch popType {
	default:
		return fmt.Errorf("????")
	case PopulationOFF:
		p.uwpStr = fmt.Sprintf("X%v%v%v000-0", p.sizeCode, p.atmoCode, p.hydrCode)
	case PopulationON:
		if strings.Contains(p.comment, "Mainworld") {
			p.comment = "Mainworld"
			p.injectUwpCodes(mwuwp)
			return nil
		}
		if mwTL < 8 {
			return notSettledUWP(p)
		}
		p.comment = ""
		switch p.habZone {
		case habZoneInner:
			switch p.sizeType {
			case sizeDwarf, sizeMercurian:
				pop := maxPopDM + dice.Sroll("2d6-9")

				if pop < 0 {

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

				if pop < 0 {

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

				if pop < 0 {

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

				if pop < 0 {

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
			switch p.sizeType {
			case sizeDwarf, sizeMercurian:
				pop := maxPopDM + dice.Sroll("2d6-9")

				if pop < 0 {

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

				if pop < 0 {

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
		p.tlCode = ehex.ToCode(mwTL - 1)
		p.updateUWP()
		p.depopulateEnviroment()
	}
	return nil
}

func (p *rockyPlanet) depopulate() {
	p.port = "X"
	p.popCode = "0"
	p.govCode = "0"
	p.lawCode = "0"
	p.tlCode = "0"
	p.comment = ""
	p.updateUWP()
}
func (b *belt) depopulate() {
	b.port = "X"
	b.popCode = "0"
	b.govCode = "0"
	b.lawCode = "0"
	b.tlCode = "0"
	b.comment = ""
	b.updateUWP()
}

func (p *rockyPlanet) depopulateEnviroment() {
	tl := ehex.ValueOf(p.tlCode)
	switch ehex.ValueOf(p.atmoCode) {
	case 0, 1:
		if tl < 8 {
			p.depopulate()
		}
	case 2, 3:
		if tl < 5 {
			p.depopulate()
		}
	case 4, 7, 9:
		if tl < 3 {
			p.depopulate()
		}
	case 10:
		if tl < 8 {
			p.depopulate()
		}
	case 11:
		if tl < 9 {
			p.depopulate()
		}
	case 12:
		if tl < 10 {
			p.depopulate()
		}
	case 13, 14:
		if tl < 5 {
			p.depopulate()
		}
	case 15:
		if tl < 8 {
			p.depopulate()
		}

	}
}

func (p *belt) depopulateEnviroment() {
	if ehex.ValueOf(p.tlCode) < 8 {
		p.depopulate()
	}
}

func (b *belt) populate(dice *dice.Dicepool, mwuwp uwp.UWP, popType string) error {
	mwTL := mwuwp.TL()
	switch popType {
	default:
		return fmt.Errorf("????")
	case PopulationOFF:
		b.uwpStr = fmt.Sprintf("X%v%v%v000-0", b.sizeCode, b.atmoCode, b.hydrCode)
	case PopulationON:
		if strings.Contains(b.comment, "Mainworld") {
			b.comment = "Mainworld"
			b.injectUwpCodes(mwuwp)
			return nil
		}
		if mwTL < 8 {
			//			b.injectUwpCodes("X000000-0")
			return notSettledUWPBelt(b)
		}
		b.comment = ""
		pop := 0
		switch {
		case strings.Contains(b.composition, "12% metal"):
			pop = dice.Sroll("2d6-8")
		case b.majorSizeAst <= 50:
			pop = dice.Sroll("2d6-10")
		default:
			pop = dice.Sroll("2d6-9")
		}
		if pop < 0 {
			pop = 0
		}
		switch dice.Sroll("2d6") {
		case 2:
			return notSettledUWPBelt(b)
		case 3, 4:
			b.defineCorpBase(dice, pop)
		case 5, 6:
			b.defineCompBase(dice, pop)
		case 7, 8:
			b.defineGovBase(dice, pop)
		case 9:
			b.defineFractions(dice, pop)
		case 10:
			b.defineMilitary(dice, pop)
		case 11:
			b.defineResearch(dice, pop)
		case 12:
			b.defineIndependent(dice, pop)
		}
		b.tlCode = ehex.ToCode(mwTL - 1)
		b.updateUWP()
		b.depopulateEnviroment()
	}
	return nil
}

func (p *rockyPlanet) updateUWP() {
	pop := ehex.ValueOf(p.popCode)
	switch pop {
	case 0:
		if p.comment != "" && !strings.Contains(p.comment, "Mainworld") {
			p.port = "Y"
		}
	case 1, 2:
		p.port = "H"
	case 3:
		p.port = "G"
	default:
		p.port = "F"
	}
	p.uwpStr = p.port + p.sizeCode + p.atmoCode + p.hydrCode + p.popCode + p.govCode + p.lawCode + "-" + p.tlCode
}

func (b *belt) updateUWP() {
	pop := ehex.ValueOf(b.popCode)
	switch pop {
	case 0, 1:
		if b.comment != "" && !strings.Contains(b.comment, "Mainworld") {
			b.port = "Y"
		}
	case 2:
		b.port = "H"
	case 3:
		b.port = "G"
	default:
		b.port = "F"
	}
	b.uwpStr = b.port + b.sizeCode + b.atmoCode + b.hydrCode + b.popCode + b.govCode + b.lawCode + "-" + b.tlCode
}

func (p *belt) defineCorpBase(dice *dice.Dicepool, pop int) {
	gov := 1
	pop = utils.Min(pop, 6)
	p.comment += fmt.Sprintf("Major Corporation Mining Operations ")
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *belt) defineCompBase(dice *dice.Dicepool, pop int) {
	gov := 1
	pop = utils.Min(pop, 6)
	p.comment += fmt.Sprintf("Local Company Mining Operations ")
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *belt) defineGovBase(dice *dice.Dicepool, pop int) {
	gov := 6
	pop = utils.Min(pop, 6)
	p.comment += fmt.Sprintf("Local Goverment Mining Operations ")
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *belt) defineFractions(dice *dice.Dicepool, pop int) {
	gov := 7
	pop = utils.Min(pop, 6)
	p.comment += fmt.Sprintf("The belt is mined by %v entities ", dice.Sroll("1d6+1"))
	law := dice.Sroll("1d6+4")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *belt) defineMilitary(dice *dice.Dicepool, pop int) {
	gov := 6
	pop = utils.Min(pop, 6)
	p.comment += fmt.Sprintf("Military base")
	law := dice.Sroll("1d6+5")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *belt) defineResearch(dice *dice.Dicepool, pop int) {
	gov := 6
	pop = utils.Min(pop, 6)
	p.comment += fmt.Sprintf("Research Station")
	law := dice.Sroll("1d6+3")
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
}

func (p *belt) defineIndependent(dice *dice.Dicepool, pop int) {
	gov := utils.Max(dice.Sroll("2d6-7")+pop, 0)
	pop = utils.Min(pop, 8)
	p.comment += fmt.Sprintf("Independent World")
	law := utils.Max(dice.Sroll("2d6-7")+gov, 0)
	planetUWP := uwp.Inject(fmt.Sprintf("X%v%v%v%v%v%v-0", 0, 0, 0, ehex.ToCode(pop), ehex.ToCode(gov), ehex.ToCode(law)))
	p.injectUwpCodes(planetUWP)
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
}

func (p *belt) injectUwpCodes(uwpS uwp.UWP) {
	p.port = uwpS.Starport()
	p.sizeCode = ehex.New().Set(uwpS.Size()).Code()
	p.atmoCode = ehex.New().Set(uwpS.Atmo()).Code()
	p.hydrCode = ehex.New().Set(uwpS.Hydr()).Code()
	p.popCode = ehex.New().Set(uwpS.Pops()).Code()
	p.govCode = ehex.New().Set(uwpS.Govr()).Code()
	p.lawCode = ehex.New().Set(uwpS.Laws()).Code()
	p.tlCode = ehex.New().Set(uwpS.TL()).Code()
	p.uwpStr = uwpS.Starport() + ehex.New().Set(uwpS.Size()).Code() + ehex.New().Set(uwpS.Atmo()).Code() + ehex.New().Set(uwpS.Hydr()).Code() + ehex.New().Set(uwpS.Pops()).Code() + ehex.New().Set(uwpS.Govr()).Code() + ehex.New().Set(uwpS.Laws()).Code() + "-" + ehex.New().Set(uwpS.TL()).Code()
}
