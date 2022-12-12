package empire

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/empire/worldcharacter"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/family"
	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/military"
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

// type PocketEmpire struct {
// 	RulingFamily  *family.Family
// 	World         map[int]worldcharacter.World
// 	Size          ehex.Ehex
// 	MilitaryPower ehex.Ehex
// 	EconomicPower ehex.Ehex
// 	Prestige      int //0-15
// }

type PocketEmpire struct {
	Name          string
	RullingFamily *family.Family //перевести в интерфейс
	World         []worldcharacter.World
	Military      []military.Unit
	//
	selfDetermination int
	popularity        int
	prestige          int
	totalPopulation   float64
}

func New() *PocketEmpire {
	empire := PocketEmpire{}
	empire.AddMilitary(military.NewUnit(hexagon.Global(0, 0), military.TYPE_GROUND, 5).SetName("test Unit"))
	return &empire
}

func (e *PocketEmpire) AddMilitary(unit *military.Unit) error {
	for _, wHave := range e.Military {
		if military.AreSame(unit, &wHave) {
			return fmt.Errorf("unit was already integrated")
		}
	}
	e.Military = append(e.Military, *unit)
	return nil
}

//integrateWorld - добавляет мир в империю, если его нет
func (e *PocketEmpire) integrateWorld(wIntegrated worldcharacter.World) error {
	for _, wHave := range e.World {
		if worldcharacter.AreSame(wIntegrated, wHave) {
			return fmt.Errorf("%v was already integrated", wIntegrated.Name())
		}
	}
	e.World = append(e.World, wIntegrated)
	return nil
}

func (e *PocketEmpire) UEP() string {
	p := ""
	p += ehex.ToCode(e.selfDetermination)
	p += ehex.ToCode(e.popularity)
	p += ehex.ToCode(populationCode(e))
	p += "-"
	p += e.GovermentCode()
	p += e.LawLevelCode()
	p += e.TechLevelCode()
	p += "-"
	p += ehex.ToCode(sizeCode(len(e.World)))
	p += ehex.ToCode(sizeCode(len(e.Military)))
	p += e.EconomicPowerCode()
	p += ehex.ToCode(e.Prestige())
	return p
}

func (e *PocketEmpire) GovermentCode() string {
	maxPop := 0.0
	wrld := 0
	for i, w := range e.World {
		pop := w.EstimatedPopulation()
		if maxPop < pop {
			maxPop = pop
			wrld = i
		}
	}
	uwps, _ := uwp.FromString(e.World[wrld].UWPs())
	return ehex.ToCode(uwps.Govr())
}

func (e *PocketEmpire) LawLevelCode() string {
	sumLaw := 0
	for _, w := range e.World {
		uwps, _ := uwp.FromString(w.UWPs())
		sumLaw += uwps.Laws()
	}
	return ehex.ToCode(sumLaw / len(e.World))
}

func (e *PocketEmpire) TechLevelCode() string {
	tl := 0
	for _, w := range e.World {
		uwps, _ := uwp.FromString(w.UWPs())
		tl += uwps.TL()
	}
	return ehex.ToCode(tl / len(e.World))
}

func sizeCode(i int) int {
	val := 0
	switch {
	case i == 1:
		val = 1
	case i == 2:
		val = 2
	case utils.InRange(i, 3, 4):
		val = 3
	case utils.InRange(i, 5, 8):
		val = 4
	case utils.InRange(i, 9, 16):
		val = 5
	case utils.InRange(i, 17, 32):
		val = 6
	case utils.InRange(i, 33, 64):
		val = 7
	case utils.InRange(i, 65, 128):
		val = 8
	case utils.InRange(i, 129, 256):
		val = 9
	case utils.InRange(i, 257, 512):
		val = 10
	case utils.InRange(i, 513, 1024):
		val = 11
	case utils.InRange(i, 1025, 2048):
		val = 12
	case utils.InRange(i, 2049, 4096):
		val = 13
	case utils.InRange(i, 4097, 8192):
		val = 14
	case utils.InRange(i, 8193, 16392):
		val = 15

	}
	return val
}

func (e *PocketEmpire) EconomicPowerCode() string {
	ru := 0
	for _, w := range e.World {
		ru += w.GetGWP()
	}
	return GEPcode(ru)
}

func GEPcode(ru int) string {
	val := 0
	switch {
	case ru == 0:
		val = 0
	case ru == 1:
		val = 1
	case utils.InRange(ru, 2, 3):
		val = 2
	case utils.InRange(ru, 4, 10):
		val = 3
	case utils.InRange(ru, 11, 30):
		val = 4
	case utils.InRange(ru, 31, 100):
		val = 5
	case utils.InRange(ru, 101, 250):
		val = 6
	case utils.InRange(ru, 251, 1000):
		val = 7
	case utils.InRange(ru, 1001, 2500):
		val = 8
	case utils.InRange(ru, 2501, 10000):
		val = 9
	case utils.InRange(ru, 10001, 25000):
		val = 10
	case utils.InRange(ru, 25001, 60000):
		val = 11
	case utils.InRange(ru, 60001, 180000):
		val = 12
	case utils.InRange(ru, 180001, 600000):
		val = 13
	case utils.InRange(ru, 600001, 1800000):
		val = 14
	case utils.InRange(ru, 1800001, 1800000000):
		val = 15
	}

	return ehex.ToCode(val)
}

func (e *PocketEmpire) calculatePopularity() {
	pp := 0.0
	for _, w := range e.World {
		pp += float64(w.Popularity())
	}
	pp = pp / float64(len(e.World))
	e.popularity = int(math.Round(pp))
}

func (e *PocketEmpire) calculateSelfDetermination() {
	sd := 0.0
	for _, w := range e.World {
		sd += float64(w.SelfDetermination())
	}
	sd = sd / float64(len(e.World))
	e.selfDetermination = int(math.Round(sd))
}

func populationCode(e *PocketEmpire) int {
	tp := int(e.totalPopulation)
	f := 1
	for tp > 0 {
		tp = tp / 10
		f++
	}
	return f
}

func (e *PocketEmpire) calculatePopulation() {
	totalPop := 0.0
	for _, w := range e.World {
		totalPop += w.EstimatedPopulation()
	}
	e.totalPopulation = totalPop
}

//excludeWorld - исключает мир из империи, если его нет
func (e *PocketEmpire) excludeWorld(wExcluded worldcharacter.World) error {
	err := fmt.Errorf("%v is not belong to %v", wExcluded.Name(), e.Name)
	var wrldsLeft []worldcharacter.World
	for _, wHave := range e.World {
		if worldcharacter.AreSame(wExcluded, wHave) {
			err = nil
			continue
		}
		wrldsLeft = append(wrldsLeft, wHave)
	}
	e.World = wrldsLeft
	return err
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
func (e *PocketEmpire) Prestige() int {
	return e.prestige
}

func (e *PocketEmpire) calculatePrestige() {
	avp := 0
	for _, w := range e.World {
		avp += w.Popularity()
	}
	avp = avp / len(e.World)
	p := populationCode(e)
	n := sizeCode(len(e.World))
	m := sizeCode(len(e.Military))
	ec := ehex.ValueOf(e.EconomicPowerCode())
	base := (avp + p + n + m + ec) / 5
	ab := 0
	e.prestige = utils.BoundInt(base+ab, 0, 15)
}
