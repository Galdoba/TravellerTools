package facilities

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

type World interface {
	MW_Name() string
	MW_UWP() string
}

type features struct {
	highport bool
	military bool
	naval    bool
	scout    bool
	corsair  bool
}

func GenerateFeatures(world World) (*features, error) {
	ft := features{}
	uwp, err := uwp.FromString(world.MW_UWP())
	if err != nil {
		return &ft, err
	}
	st := uwp.Starport()
	pop := uwp.Pops()
	law := uwp.Laws()
	tl := uwp.TL()
	featureDM := featuresDM(tl, pop, law)
	dp := dice.New().SetSeed(world.MW_Name() + world.MW_UWP())
	featuresTNmap := make(map[string][]int)
	featuresTNmap["A"] = []int{6, 8, 8, 10, 99}
	featuresTNmap["B"] = []int{8, 8, 8, 9, 99}
	featuresTNmap["C"] = []int{10, 99, 10, 9, 99}
	featuresTNmap["D"] = []int{12, 99, 99, 8, 12}
	featuresTNmap["E"] = []int{99, 99, 99, 99, 10}
	featuresTNmap["X"] = []int{99, 99, 99, 99, 10}
	for k, tnList := range featuresTNmap {
		if k != st {
			continue
		}
		for i, tn := range tnList {
			present := false
			if dp.Roll("2d6").DM(featureDM[i]).Sum() >= tn {
				present = true
			}
			switch i {
			case 0:
				ft.highport = present
			case 1:
				ft.military = present
			case 2:
				ft.naval = present
			case 3:
				ft.scout = present
			case 4:
				ft.corsair = present
			}
		}
	}
	return &ft, nil
}

func (ft *features) String() string {
	str := ""
	str += "highport =" + fmt.Sprint(ft.highport) + "\n"
	str += "military =" + fmt.Sprint(ft.military) + "\n"
	str += "naval =" + fmt.Sprint(ft.naval) + "\n"
	str += "scout =" + fmt.Sprint(ft.scout) + "\n"
	str += "corsair =" + fmt.Sprint(ft.corsair) + "\n"
	return str
}

func featuresDM(tl, pop, law int) []int {
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
