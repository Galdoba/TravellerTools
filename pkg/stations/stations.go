package stations

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
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
	Type_Imperial
	Type_Pirate
	Type_WRONG_VALUE
)

type World interface {
	MW_Name() string
	MW_UWP() string
	MW_Remarks() string
	TravelZone() string
}

type SpaceStationsSystemReport struct {
	name    string
	uwp     string
	report  string
	Present map[int]int
}

func GenerateSystemReport(world World) (*SpaceStationsSystemReport, error) {
	err := fmt.Errorf("err variable was not touched")
	rep := SpaceStationsSystemReport{}
	rep.name = world.MW_Name()
	rep.uwp = world.MW_UWP()
	rep.Present = make(map[int]int)
	dp := dice.New().SetSeed(world.MW_Name() + world.MW_UWP())
	dp.Vocal()
	if err = rep.rollNaval(dp); err != nil {
		return &rep, err
	}

	if err = rep.rollParamilitary(dp); err != nil {
		return &rep, err
	}
	rep.fillReport()
	return &rep, err
}

func (rep *SpaceStationsSystemReport) rollNaval(dp *dice.Dicepool) error {
	uwpS, err := uwp.FromString(rep.uwp)
	if err != nil {
		return err
	}
	st := uwpS.Starport()
	pop := uwpS.Pops()
	tl := uwpS.TL()
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
			if arr[i] > 16 {
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
	return nil
}

func (rep *SpaceStationsSystemReport) rollParamilitary(dp *dice.Dicepool) error {
	uwpS, err := uwp.FromString(rep.uwp)
	if err != nil {
		return err
	}
	st := uwpS.Starport()
	law := uwpS.Laws()
	tl := uwpS.TL()
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
	return nil
}

func (rep *SpaceStationsSystemReport) rollCommercial(dp *dice.Dicepool) error {
	uwpS, err := uwp.FromString(rep.uwp)
	if err != nil {
		return err
	}
	st := uwpS.Starport()
	pop := uwpS.Pops()
	tl := uwpS.TL()
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
			if arr[i] > 16 {
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
	return nil
}

func (r *SpaceStationsSystemReport) fillReport() {
	r.report = ""
	r.report += fmt.Sprintf("%v system space stations report:\n", r.name)
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
	return
}

func (r *SpaceStationsSystemReport) Summary() string {
	return r.report
}
