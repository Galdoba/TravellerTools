package empire

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/empire/worldcharacter"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/family"
	"github.com/Galdoba/utils"
)

/*
GAME FLOW:

*/

type individualWorld interface {
	MW_Name() string
	MW_UWP() string
	PBG() string
	CoordX() int
	CoordY() int
}

type PocketEmpire struct {
	RulingFamily  *family.Family
	World         map[int]worldcharacter.World
	Size          ehex.Ehex
	MilitaryPower ehex.Ehex
	EconomicPower ehex.Ehex
	Prestige      int //0-15
}

func New() *PocketEmpire {
	empire := PocketEmpire{}
	empire.World = make(map[int]worldcharacter.World)
	return &empire
}

type economicExtention struct {
	resource       ehex.Ehex
	labor          ehex.Ehex
	infrastructure ehex.Ehex
	culture        ehex.Ehex
}

func (ee *economicExtention) String() string {
	return fmt.Sprintf("%v%v%v%v", ee.resource.Code(), ee.labor.Code(), ee.infrastructure.Code(), ee.culture.Code())
}

func haveLife(tc []string) bool {
	for _, t := range tc {
		if strings.Contains(t, "NL") {
			return true
		}
	}
	return false
}

func coreDensity(u uwp.UWP, tc []string, dice *dice.Dicepool) int {
	dm := 0
	if u.Size() < 5 {
		dm += 1
	}
	if u.Size() > 5 {
		dm -= 2
	}
	if u.Atmo() < 4 {
		dm += 1
	}
	if u.Atmo() < 5 {
		dm -= 2
	}
	if utils.ListContains(tc, "Fr") {
		dm += 6
	}
	r := dice.Sroll("2d6") + dm
	if r < 2 {
		return 0
	}
	if r < 11 {
		return 1
	}
	if r < 15 {
		return 2
	}
	return 3
}

type Constructor struct {
	dice *dice.Dicepool
}

func (c *Constructor) SetupEmpire(worlds ...individualWorld) (*PocketEmpire, error) {
	empire := New()

	return empire, nil
}

/*
TRADEGOODS
func availableResources(u uwp.UWP, tc []string, dice *dice.Dicepool) []string {
	availRes := []string{}
	cd := coreDensity(u, tc, dice)
	resourceMap := make(map[string][]int)
	resourceMap["Agricultural"] = []int{1, 4, 4, -4, 1, -3, 0, 0, 1, 0, -1, -2, 5, 0}
	resourceMap["Ores"] = []int{8, 7, 3, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0}
	resourceMap["Radioactives"] = []int{7, 5, 3, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0}
	resourceMap["Crystals"] = []int{6, 5, 2, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0}
	resourceMap["Compounds"] = []int{5, 6, 1, -4, 0, 1, 0, 0, 1, 0, 0, 1, 1, -1}
	resourceMap["Agroproducts"] = []int{0, 1, 1, 0, 2, 0, 1, 2, 1, 2, 1, 1, 5, 0}
	resourceMap["Metals"] = []int{2, 0, 0, -1, 0, 1, -1, 1, -1, 2, 4, 5, 0, 0}
	resourceMap["Non-Metals"] = []int{1, 0, 0, -1, 1, 1, 0, 1, 0, 2, 4, 6, 3, 0}
	resourceMap["Parts"] = []int{0, 1, -1, 1, 2, -1, 1, 2, 0, 0, 0, 2, 4, 1, 0}
	resourceMap["Durables"] = []int{0, 1, -1, 2, 3, -1, 1, 2, 0, 0, 1, 2, 3, 1, 0}
	resourceMap["Consumables"] = []int{0, 1, -1, 1, 4, -1, 1, 2, 1, 0, 1, 2, 4, 1, 0}
	resourceMap["Weapons"] = []int{0, 1, -1, 0, 1, 0, 1, 3, 1, 0, 1, 1, 2, 1, 0}
	resourceMap["Recordings"] = []int{0, 1, 2, 0, 1, 1, 2, 0, 1, 2, 3, -3, 1, 2, 3}
	resourceMap["Artforms"] = []int{0, 2, 3, 0, 1, 2, 0, 0, 0, 0, 0, 2, 1, 1, 1}
	resourceMap["Software"] = []int{0, 1, 4, 0, 1, 1, 1, 0, 1, 2, 3, -9, 0, 1, 4}
	resourceMap["Documents"] = []int{-1, 0, 1, 0, 1, 2, 4, 0, 2, 4, 6, 0, 1, 3, 1}
	res := []string{"Agricultural", "Ores", "Radioactives", "Crystals", "Compounds", "Agroproducts", "Metals", "Non-Metals", "Parts", "Durables", "Consumables", "Weapons", "Recordings", "Artforms", "Software", "Documents"}
	for _, rs := range res {
		tn := 0
		switch rs {
		case "Agricultural", "Ores", "Radioactives", "Crystals", "Compounds", "Agroproducts", "Metals", "Non-Metals":
			tn += resourceMap[rs][cd]
			switch u.Atmo() {
			case 4, 5, 6, 7, 8, 9:
				tn += resourceMap[rs][4]
			default:
				tn += resourceMap[rs][5]
			}
			switch u.Pops() {
			case 0, 1, 2, 3, 4:
				tn += resourceMap[rs][6]
			default:
				tn += resourceMap[rs][7]
			}
			switch u.TL() {
			case 0, 1, 2, 3:
				tn += resourceMap[rs][8]
			case 4, 5, 6:
				tn += resourceMap[rs][9]
			case 7, 8, 9, 10, 11:
				tn += resourceMap[rs][10]
			default:
				tn += resourceMap[rs][11]
			}
			switch haveLife(tc) {
			case true:
				tn += resourceMap[rs][12]
			case false:
				tn += resourceMap[rs][13]
			}
		case "Parts", "Durables", "Consumables", "Weapons":
			switch u.Atmo() {
			case 4, 5, 6, 7, 8, 9:
				tn += resourceMap[rs][0]
			default:
				tn += resourceMap[rs][1]
			}
			switch u.Pops() {
			case 0, 1, 2, 3, 4:
				tn += resourceMap[rs][2]
			case 5, 6, 7, 8:
				tn += resourceMap[rs][3]
			default:
				tn += resourceMap[rs][4]
			}
			switch u.Govr() {
			case 0, 1:
				tn += resourceMap[rs][5]
			case 2, 3, 4, 5, 6:
				tn += resourceMap[rs][6]
			case 7:
				tn += resourceMap[rs][7]
			default:
				tn += resourceMap[rs][8]
			}
			switch u.TL() {
			case 0, 1, 2, 3:
				tn += resourceMap[rs][9]
			case 4, 5, 6:
				tn += resourceMap[rs][10]
			case 7, 8, 9, 10, 11:
				tn += resourceMap[rs][11]
			default:
				tn += resourceMap[rs][12]
			}
			switch haveLife(tc) {
			case true:
				tn += resourceMap[rs][13]
			case false:
				tn += resourceMap[rs][14]
			}
		case "Recordings", "Artforms", "Software", "Documents":
			switch u.Pops() {
			case 0, 1, 2, 3, 4:
				tn += resourceMap[rs][0]
			case 5, 6, 7, 8:
				tn += resourceMap[rs][1]
			default:
				tn += resourceMap[rs][2]
			}
			switch u.Govr() {
			case 0, 1:
				tn += resourceMap[rs][3]
			case 2, 3, 4, 5, 6:
				tn += resourceMap[rs][4]
			case 7:
				tn += resourceMap[rs][5]
			default:
				tn += resourceMap[rs][6]
			}
			switch u.Laws() {
			case 0, 1, 2:
				tn += resourceMap[rs][7]
			case 3, 4, 5, 6:
				tn += resourceMap[rs][8]
			case 7, 8, 9:
				tn += resourceMap[rs][9]
			default:
				tn += resourceMap[rs][10]
			}
			switch u.TL() {
			case 0, 1, 2, 3:
				tn += resourceMap[rs][11]
			case 4, 5, 6:
				tn += resourceMap[rs][12]
			case 7, 8, 9, 10, 11:
				tn += resourceMap[rs][13]
			default:
				tn += resourceMap[rs][14]
			}
		}
		if dice.Sroll("2d6") <= tn {
			availRes = append(availRes, rs)
		}
	}
	return availRes
}
*/
