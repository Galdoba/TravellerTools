package facilities

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

const (
	Highport = iota
	Military
	Naval
	Scout
	Corsair
	//Security Aspects
	LawLevel
	LawEnforcement
	Defences
	//Instalations
	ArmyBase
	DefenceBase
	MaintainanceFacilities
	NavalBase
	NavalDepot
	ResearchInstallation
	ScoutBase
	Shipyard
	WayStation
)

type World interface {
	MW_Name() string
	MW_UWP() string
}

type facilities struct {
	ByType map[int]bool
	Rating map[int]int
}

func GenerateFacilities(world World) (*facilities, error) {
	ft := facilities{}
	ft.ByType = make(map[int]bool)
	uwp, err := uwp.FromString(world.MW_UWP())
	if err != nil {
		return &ft, err
	}
	st := uwp.Starport()
	pop := uwp.Pops()
	law := uwp.Laws()
	tl := uwp.TL()
	featureDM := facilitiesDM(tl, pop, law)
	dp := dice.New().SetSeed(world.MW_Name() + world.MW_UWP())
	facilitiesTNmap := make(map[string][]int)
	facilitiesTNmap["A"] = []int{6, 8, 8, 10, 99}
	facilitiesTNmap["B"] = []int{8, 8, 8, 9, 99}
	facilitiesTNmap["C"] = []int{10, 99, 10, 9, 99}
	facilitiesTNmap["D"] = []int{12, 99, 99, 8, 12}
	facilitiesTNmap["E"] = []int{99, 99, 99, 99, 10}
	facilitiesTNmap["X"] = []int{99, 99, 99, 99, 10}
	for k, tnList := range facilitiesTNmap {
		if k != st {
			continue
		}
		for i, tn := range tnList {
			present := false
			if dp.Roll("2d6").DM(featureDM[i]).Sum() >= tn {
				present = true
			}
			switch i {
			default:
				return &ft, fmt.Errorf("unknown facility type code %v", i)
			case Highport, Military, Naval, Scout, Corsair:
				ft.ByType[i] = present
			}
		}
	}
	ft.Rating = make(map[int]int)
	for _, aspect := range []int{LawLevel, LawEnforcement, Defences} {
		rr := dp.Roll("2d6").Sum()
		dm := 0
		switch aspect {
		case LawLevel:
			switch st {
			case "A":
				dm = 7
			case "B":
				dm = 5
			case "C":
				dm = 3
			case "D":
				dm = 1
			}
			ft.Rating[LawLevel] = rr + dm - law
			if ft.Rating[LawLevel] < 0 {
				ft.Rating[LawLevel] = 0
			}
			fmt.Println(ft.Rating[LawLevel])
		}
	}
	return &ft, nil
}

func (ft *facilities) String() string {
	str := ""
	str += "highport =" + fmt.Sprint(ft.ByType[Highport]) + "\n"
	str += "military =" + fmt.Sprint(ft.ByType[Military]) + "\n"
	str += "naval =" + fmt.Sprint(ft.ByType[Naval]) + "\n"
	str += "scout =" + fmt.Sprint(ft.ByType[Scout]) + "\n"
	str += "corsair =" + fmt.Sprint(ft.ByType[Corsair]) + "\n"
	return str
}

func facilitiesDM(tl, pop, law int) []int {
	dms := []int{}
	dms = append(dms, stdValDM(tl)+stdValDM(pop))
	dms = append(dms, 0)
	dms = append(dms, 0)
	dms = append(dms, 0)
	switch law {
	case 0:
		dms = append(dms, 2)
	case 1:
		dms = append(dms, 0)
	default:
		dms = append(dms, -2)
	}
	return dms
}

func stdValDM(chr int) int {
	switch chr {
	case 0:
		return -3
	case 1, 2:
		return -2
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	default:
		return chr/3 - 2
	}
}

/*
facilities list:
Highport
Military
Naval
Scout
Corsair

*/
