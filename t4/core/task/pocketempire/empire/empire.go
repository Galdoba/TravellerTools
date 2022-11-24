package empire

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/family"
	"github.com/Galdoba/utils"
)

type PocketEmpire struct {
	RulingFamily  *family.Family
	World         map[int]*worldCharacter
	Size          ehex.Ehex
	MilitaryPower ehex.Ehex
	EconomicPower ehex.Ehex
	Prestige      int //0-15
}

type worldCharacter struct {
	name              string
	uwp               uwp.UWP
	pbg               string
	tradecodes        []string
	culturalExtention []ehex.Ehex
	selfDetermination ehex.Ehex //0-10
	localPopularity   ehex.Ehex //0-15
	progression       int
	factions          []int //распределение по фракциям суммарно 100%
	hex               hexagon.Hexagon
}

type individualWorld interface {
	MW_Name() string
	MW_UWP() string
	PBG() string
	CoordX() int
	CoordY() int
}

func New() *PocketEmpire {
	empire := PocketEmpire{}
	empire.World = make(map[int]*worldCharacter)
	return &empire
}

func WorldCharacter(indWrld individualWorld) *worldCharacter {
	dice := dice.New().SetSeed(indWrld.MW_Name() + indWrld.MW_UWP() + indWrld.PBG())
	wc := worldCharacter{}
	wc.name = indWrld.MW_Name()
	wc.uwp = uwp.Inject(indWrld.MW_UWP())
	wc.pbg = indWrld.PBG()
	hex := hexagon.FromHex(indWrld)
	wc.hex = hex
	wc.tradecodes = setupTradeCodes(wc.uwp, dice)

	wc.selfDetermination = ehex.New().Set(dice.Sroll("2d6-2"))

	wc.progression = 0
	return &wc
}

func availableResources(u uwp.UWP, tc []string, dice *dice.Dicepool) []string {
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
	return tc
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

func setupTradeCodes(u uwp.UWP, dice *dice.Dicepool) []string {
	tc, _ := tradecodes.FromUWP(u)
	switch rollTemp(u, dice) {
	case 1:
		tc = append(tc, "Fr")
	case 2:
		tc = append(tc, "Co")
	case 4:
		tc = append(tc, "Ho")
	case 5:
		tc = append(tc, "Bo")
	}
	nl := nativeLifeRoll(u, dice, tc)
	fmt.Println(nl)
	if nl > 0 {
		tc = append(tc, fmt.Sprintf("Nl%v", nl))
	}
	return tc
}

func nativeLifeRoll(u uwp.UWP, dice *dice.Dicepool, tc []string) int {
	dm := 0
	switch u.Atmo() {
	case 0:
		dm = dm - 3
	case 4, 5, 6, 7, 8, 9:
		dm = dm + 4
	}
	switch u.Hydr() {
	case 0:
		dm = dm - 2
	case 2, 3, 4, 5, 6, 7, 8:
		dm = dm + 1
	}
	for _, t := range tc {
		switch t {
		case "Fr", "Co", "Ho", "Bo":
			dm = dm - 1
		}
	}
	return dice.Sroll("2d6") + dm - 10
}

type Constructor struct {
	dice *dice.Dicepool
}

func (c *Constructor) SetupEmpire(worlds ...individualWorld) (*PocketEmpire, error) {
	empire := New()

	return empire, nil
}

func (wc *worldCharacter) Descr() string {
	s := fmt.Sprintf("World: %v\n", wc.name)
	s += fmt.Sprintf("UWP  : %v\n", wc.uwp.String())
	s += fmt.Sprintf("HEX  : %v\n", wc.hex.String())
	s += fmt.Sprintf("PBG: : %v\n", wc.pbg)
	s += fmt.Sprintf("SD   : %v\n", wc.selfDetermination.Code())
	s += fmt.Sprintf("TC   : ")
	for _, tc := range wc.tradecodes {
		s += tc + " "
	}
	s += "\n"

	return s
}

func rollTemp(u uwp.UWP, dice *dice.Dicepool) int {
	dm := 0
	temperature := 0
	extreme := false
	size := ehex.ToCode(u.Size())
	switch size {
	case "0", "1":
		extreme = true
	case "2", "3":
		dm = -2
	case "4", "5", "E":
		dm = -1
	case "6", "7":
		dm = 0
	case "8", "9":
		dm = 1
	case "A", "D", "F":
		dm = 2
	case "B", "C":
		dm = 6
	}
	r := dice.Sroll("2d6") + dm
	switch r {
	case 3, 4:
		temperature = 2
		if extreme {
			temperature = 1
		}
	case 5, 6, 7, 8, 9:
		temperature = 3
		if extreme {
			switch dice.Sroll("1d2") {
			case 1:
				temperature = 1
			case 2:
				temperature = 5
			}
		}
	case 10, 11:
		temperature = 4
		if extreme {
			temperature = 5
		}
	}
	if r <= 2 {
		temperature = 1
	}
	if r >= 12 {
		temperature = 5
	}
	return temperature
}
