package stations

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	Type_INITIAL_VALUE = iota
	Type_ANY
	Type_Naval
	Type_Naval_Defence
	Type_Naval_Fleet
	Type_Naval_Interdiction
	Type_Naval_Shipyard
	Type_Paramilitary
	Type_Paramilitary_Customs
	Type_Paramilitary_Mercenary
	Type_Paramilitary_Scout
	Type_Commercial
	Type_Commercial_Asteroid_Mine
	Type_Commercial_Manufacturing_Plant
	Type_Commercial_Refinery
	Type_Civilan
	Type_Civilan_Habitation
	Type_Civilan_Research
	Type_Civilan_Trade_Hub
	Type_Imperial
	Type_Imperial_Navy
	Type_Imperial_Consulate
	Type_Imperial_Xboat
	Type_Pirate
	Type_Pirate_Pirates_Paradise
	Type_Pirate_Smugglers_Heaven
	Type_Pirate_Wreckers_Yard

	Type_WRONG_VALUE
)

type World interface {
	MW_Name() string
	MW_UWP() string
	MW_Remarks() string
	TravelZone() string
	PBG() string
}

type SpaceStationsSystemReport struct {
	name       string
	uwp        uwp.UWP
	tradecodes []string
	popDigit   int
	belts      int
	ggPresent  int
	report     string
	Present    map[int]int
}

func GenerateSystemReport(world World) (*SpaceStationsSystemReport, error) {
	err := fmt.Errorf("err variable was not touched")
	rep := SpaceStationsSystemReport{}
	rep.name = world.MW_Name()
	rep.uwp, err = uwp.FromString(world.MW_UWP())
	if err != nil {
		return &rep, fmt.Errorf("uwp.FromString(world.MW_UWP()): %v", err.Error())
	}
	rep.tradecodes, err = tradecodes.Of(world)
	for i, v := range pbgDigit(world) {
		switch i {
		default:
			return &rep, fmt.Errorf("world PBG() data incorect")
		case 0:
			rep.popDigit = v
		case 1:
			rep.belts = v
		case 2:
			rep.ggPresent = v
		}
	}
	rep.Present = make(map[int]int)
	dp := dice.New().SetSeed(world.MW_Name() + world.MW_UWP() + world.MW_Remarks() + world.TravelZone() + world.PBG())
	switch {
	default:
		rep.rollNaval(dp)
		rep.rollParamilitary(dp)
		rep.rollCommercial(dp)
		rep.rollCivilan(dp)
		rep.rollImperial(dp)
	case rep.uwp.TL() < 7:
	}
	rep.rollPirate(dp)
	rep.fillReport()
	return &rep, err
}

func (rep *SpaceStationsSystemReport) rollNaval(dp *dice.Dicepool) {
	st := rep.uwp.Starport()
	pop := rep.uwp.Pops()
	tl := rep.uwp.TL()
	dm := 0
	if st == "A" || st == "B" {
		dm += +2
	}
	if pop > 5 {
		dm += +1
	}
	if pop > 8 {
		dm += +1
	}
	if tl >= 12 {
		dm += +1
	}
	arr := []int{}
	for i := 0; i < 4; i++ {
		arr = append(arr, dp.Roll("2d6").DM(dm).Sum())
	}
	for i, stationType := range []int{Type_Naval_Defence, Type_Naval_Fleet, Type_Naval_Interdiction, Type_Naval_Shipyard} {
		num := 0
		switch arr[i] {
		default:
			num = 0
			if arr[i] >= 17 {
				num = dp.Roll("4d6").Sum()
			}
		case 11:
			num = 1
		case 12:
			num = 2
		case 13:
			num = dp.Roll("1d6").Sum()
		case 14:
			num = dp.Roll("1d6").Sum() + 2
		case 15:
			num = dp.Roll("2d6").Sum()
		case 16:
			num = dp.Roll("3d6").Sum()
		}
		rep.Present[stationType] = num
		rep.Present[Type_Naval] += num
		rep.Present[Type_ANY] += num
	}
}

func (rep *SpaceStationsSystemReport) rollParamilitary(dp *dice.Dicepool) {

	st := rep.uwp.Starport()
	law := rep.uwp.Laws()
	tl := rep.uwp.TL()
	dm := 0
	if st == "C" {
		dm += +1
	}
	if st == "D" {
		dm += +2
	}
	if law <= 4 {
		dm += +1
	}
	if tl >= 12 {
		dm += -1
	}
	if tl >= 15 {
		dm += -1
	}
	arr := []int{}
	for i := 0; i < 4; i++ {
		arr = append(arr, dp.Roll("2d6").DM(dm).Sum())
	}
	for i, stationType := range []int{Type_Paramilitary_Customs, Type_Paramilitary_Mercenary, Type_Paramilitary_Scout} {
		num := 0
		switch arr[i] {
		default:
			num = 0
			if arr[i] >= 15 {
				num = dp.Roll("2d6").Sum()
			}
		case 9:
			num = 1
		case 10:
			num = 2
		case 11:
			num = 3
		case 12:
			num = dp.Roll("1d6").Sum()
		case 13:
			num = dp.Roll("1d6").Sum() + 1
		case 14:
			num = dp.Roll("1d6").Sum() + 2
		}
		rep.Present[stationType] = num
		rep.Present[Type_Paramilitary] += num
		rep.Present[Type_ANY] += num
	}
}

func (rep *SpaceStationsSystemReport) rollCommercial(dp *dice.Dicepool) {
	st := rep.uwp.Starport()
	dm := 0
	switch st {
	case "A":
		dm += +2
	case "B", "C":
		dm += +1
	}
	for _, code := range rep.tradecodes {
		if code == "Hi" || code == "In" {
			dm += +2
		}
	}
	arr := []int{}
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			if rep.belts < 1 {
				dm += -1000
			} else {
				dm += rep.belts
			}
		case 1:
			for _, tc := range rep.tradecodes {
				if tc == "In" {
					dm += rep.uwp.Pops() / 2
				}
			}
		case 2:
			if rep.ggPresent < 1 && rep.uwp.Hydr() == 0 {
				dm += -1000
			} else {
				dm += rep.ggPresent
			}
		}
		arr = append(arr, dp.Roll("2d6").DM(dm).Sum())
	}
	for i, stationType := range []int{Type_Commercial_Asteroid_Mine, Type_Commercial_Manufacturing_Plant, Type_Commercial_Refinery} {
		num := 0
		switch arr[i] {
		default:
			num = 0
			if arr[i] >= 16 {
				num = dp.Roll("5d6").Sum()
			}
		case 10:
			num = 1
		case 11:
			num = 3
		case 12:
			num = dp.Roll("1d6").Sum()
		case 13:
			num = dp.Roll("2d6").Sum()
		case 14:
			num = dp.Roll("3d6").Sum()
		case 15:
			num = dp.Roll("4d6").Sum()
		}

		rep.Present[stationType] = num
		rep.Present[Type_Commercial] += num
		rep.Present[Type_ANY] += num
	}
}

func (rep *SpaceStationsSystemReport) rollCivilan(dp *dice.Dicepool) {
	st := rep.uwp.Starport()
	pop := rep.uwp.Pops()
	tl := rep.uwp.TL()
	dm := 0
	switch st {
	case "A":
		dm += +2
	case "B":
		dm += +1
	}
	if pop > 7 {
		dm += +1
	}
	if tl >= 12 {
		dm += +1
	}
	arr := []int{}
	for i := 0; i < 4; i++ {
		arr = append(arr, dp.Roll("2d6").DM(dm).Sum())
	}
	for i, stationType := range []int{Type_Civilan_Habitation, Type_Civilan_Research, Type_Civilan_Trade_Hub} {
		num := 0
		switch arr[i] {
		default:
			num = 0
			if arr[i] >= 16 {
				num = dp.Roll("3d6").Sum()
			}
		case 10:
			num = 1
		case 11:
			num = 3
		case 12:
			num = dp.Roll("1d6").Sum()
		case 13:
			num = dp.Roll("1d6").Sum() + 3
		case 14:
			num = dp.Roll("2d6").Sum()
		case 15:
			num = dp.Roll("2d6").Sum() + 3
		}
		rep.Present[stationType] = num
		rep.Present[Type_Civilan] += num
		rep.Present[Type_ANY] += num
	}
}
func (rep *SpaceStationsSystemReport) rollImperial(dp *dice.Dicepool) {
	st := rep.uwp.Starport()
	dm := 0
	switch st {
	case "A":
		dm += +2
	case "B":
		dm += +1
	}
	arr := []int{}
	for i := 0; i < 4; i++ {
		arr = append(arr, dp.Roll("2d6").DM(dm).Sum())
	}
	for i, stationType := range []int{Type_Imperial_Consulate, Type_Imperial_Navy, Type_Imperial_Xboat} {
		num := 0
		switch arr[i] {
		default:
			num = 0
			if arr[i] >= 14 {
				num = 2
			}
		case 12, 13:
			num = 1
		}
		rep.Present[stationType] = num
		rep.Present[Type_Imperial] += num
		rep.Present[Type_ANY] += num
	}
}
func (rep *SpaceStationsSystemReport) rollPirate(dp *dice.Dicepool) {
	st := rep.uwp.Starport()
	law := rep.uwp.Laws()

	dm := 0
	switch st {
	case "B", "D", "E":
		dm += +1
	case "C":
		dm += +2
	}
	switch law {
	case 0:
		dm += +2
	case 1, 2:
		dm += 1
	}
	arr := []int{}
	for i := 0; i < 4; i++ {
		arr = append(arr, dp.Roll("2d6").DM(dm).Sum())
	}
	for i, stationType := range []int{Type_Pirate_Pirates_Paradise, Type_Pirate_Smugglers_Heaven, Type_Pirate_Wreckers_Yard} {
		num := 0
		switch arr[i] {
		default:
			num = 0
			if arr[i] >= 16 {
				num = dp.Roll("1d6").Sum()
			}
		case 10, 11:
			num = 1
		case 12, 13:
			num = 2
		case 14, 15:
			num = 3
		}
		rep.Present[stationType] = num
		rep.Present[Type_Pirate] += num
		rep.Present[Type_ANY] += num
	}
}

func (r *SpaceStationsSystemReport) fillReport() {
	r.report = ""
	r.report += fmt.Sprintf("%v system %v %v \nSpace stations report:\n", r.name, r.uwp, r.tradecodes)
	if r.Present[Type_ANY] > 0 {
		r.report += fmt.Sprintf("There are total %v space stations in the system.\n", r.Present[Type_ANY])
	} else {
		r.report += fmt.Sprintf("No space stations in the system.")
		return
	}
	if r.Present[Type_Naval] > 0 {
		r.report += fmt.Sprintf("Naval Stations (%v total):\n", r.Present[Type_Naval])
		for _, val := range []int{Type_Naval_Defence, Type_Naval_Fleet, Type_Naval_Interdiction, Type_Naval_Shipyard} {
			if r.Present[val] == 0 {
				continue
			}
			switch val {
			case Type_Naval_Defence:
				r.report += fmt.Sprintf(" %v Defence\n", r.Present[Type_Naval_Defence])
			case Type_Naval_Fleet:
				r.report += fmt.Sprintf(" %v Fleet\n", r.Present[Type_Naval_Fleet])
			case Type_Naval_Interdiction:
				r.report += fmt.Sprintf(" %v Interdiction\n", r.Present[Type_Naval_Interdiction])
			case Type_Naval_Shipyard:
				r.report += fmt.Sprintf(" %v Shipyard\n", r.Present[Type_Naval_Shipyard])
			}
		}
	}
	if r.Present[Type_Paramilitary] > 0 {
		r.report += fmt.Sprintf("Paramilitary Stations (%v total):\n", r.Present[Type_Paramilitary])
		for _, val := range []int{Type_Paramilitary_Customs, Type_Paramilitary_Mercenary, Type_Paramilitary_Scout} {
			if r.Present[val] == 0 {
				continue
			}
			switch val {
			case Type_Paramilitary_Customs:
				r.report += fmt.Sprintf(" %v Customs\n", r.Present[Type_Paramilitary_Customs])
			case Type_Paramilitary_Mercenary:
				r.report += fmt.Sprintf(" %v Mercenary\n", r.Present[Type_Paramilitary_Mercenary])
			case Type_Paramilitary_Scout:
				r.report += fmt.Sprintf(" %v Scout\n", r.Present[Type_Paramilitary_Scout])

			}
		}
	}
	if r.Present[Type_Commercial] > 0 {
		r.report += fmt.Sprintf("Commercial Stations (%v total):\n", r.Present[Type_Commercial])
		for _, val := range []int{Type_Commercial_Asteroid_Mine, Type_Commercial_Manufacturing_Plant, Type_Commercial_Refinery} {
			if r.Present[val] == 0 {
				continue
			}
			switch val {
			case Type_Commercial_Asteroid_Mine:
				r.report += fmt.Sprintf(" %v Asteroid Mine\n", r.Present[Type_Commercial_Asteroid_Mine])
			case Type_Commercial_Manufacturing_Plant:
				r.report += fmt.Sprintf(" %v Manufacturing Plant\n", r.Present[Type_Commercial_Manufacturing_Plant])
			case Type_Commercial_Refinery:
				r.report += fmt.Sprintf(" %v Refinery\n", r.Present[Type_Commercial_Refinery])

			}
		}
	}
	if r.Present[Type_Civilan] > 0 {
		r.report += fmt.Sprintf("Civilian Stations (%v total):\n", r.Present[Type_Civilan])
		for _, val := range []int{Type_Civilan_Habitation, Type_Civilan_Research, Type_Civilan_Trade_Hub} {
			if r.Present[val] == 0 {
				continue
			}
			switch val {
			case Type_Civilan_Habitation:
				r.report += fmt.Sprintf(" %v Habitation\n", r.Present[Type_Civilan_Habitation])
			case Type_Civilan_Research:
				r.report += fmt.Sprintf(" %v Research\n", r.Present[Type_Civilan_Research])
			case Type_Civilan_Trade_Hub:
				r.report += fmt.Sprintf(" %v Trade Hub\n", r.Present[Type_Civilan_Trade_Hub])

			}
		}
	}
	if r.Present[Type_Imperial] > 0 {
		r.report += fmt.Sprintf("Imperial Stations (%v total):\n", r.Present[Type_Imperial])
		for _, val := range []int{Type_Imperial_Consulate, Type_Imperial_Navy, Type_Imperial_Xboat} {
			if r.Present[val] == 0 {
				continue
			}
			switch val {
			case Type_Imperial_Consulate:
				r.report += fmt.Sprintf(" %v Consulate\n", r.Present[Type_Imperial_Consulate])
			case Type_Imperial_Navy:
				r.report += fmt.Sprintf(" %v Navy\n", r.Present[Type_Imperial_Navy])
			case Type_Imperial_Xboat:
				r.report += fmt.Sprintf(" %v X-boat Waystation\n", r.Present[Type_Imperial_Xboat])

			}
		}
	}
	if r.Present[Type_Pirate] > 0 {
		r.report += fmt.Sprintf("Pirate Stations (%v total):\n", r.Present[Type_Pirate])
		for _, val := range []int{Type_Pirate_Pirates_Paradise, Type_Pirate_Smugglers_Heaven, Type_Pirate_Wreckers_Yard} {
			if r.Present[val] == 0 {
				continue
			}
			switch val {
			case Type_Pirate_Pirates_Paradise:
				r.report += fmt.Sprintf(" %v Pirate's Paradise\n", r.Present[Type_Pirate_Pirates_Paradise])
			case Type_Pirate_Smugglers_Heaven:
				r.report += fmt.Sprintf(" %v Smuggler's Heaven\n", r.Present[Type_Pirate_Smugglers_Heaven])
			case Type_Pirate_Wreckers_Yard:
				r.report += fmt.Sprintf(" %v Wrecker's Yard\n", r.Present[Type_Pirate_Wreckers_Yard])

			}
		}
	}

}

func (r *SpaceStationsSystemReport) Summary() string {
	return r.report
}

func pbgDigit(world World) []int {
	pbg := strings.Split(world.PBG(), "")
	pbgInt := []int{}
	for _, v := range pbg {
		value := ehex.New().Set(v).Value()
		pbgInt = append(pbgInt, value)
	}
	return pbgInt
}

func tradeCodes(world World) []string {
	tc := strings.Fields(world.MW_Remarks())
	uwp, _ := uwp.FromString(world.MW_UWP())
	if uwp.TL() >= 12 {
		tc = append(tc, world.TravelZone())
	}
	tc = append(tc, world.TravelZone())
	return tc
}
