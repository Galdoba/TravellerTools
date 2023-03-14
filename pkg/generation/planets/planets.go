package planets

import (
	"fmt"
	"math"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/life"
	"github.com/Galdoba/utils"
)

/*
для генерации планет нужны данные о:
Star
Size
Atmo
HZ
Hydr
Life
*/

type planet struct {
	size    ehex.Ehex
	atmo    ehex.Ehex
	hydr    ehex.Ehex
	life    ehex.Ehex
	hz      int
	plType  int
	star    string
	comment string
}

type PlanetaryBody interface {
	Size() ehex.Ehex
	Atmo() ehex.Ehex
	Hydr() ehex.Ehex
	Life() ehex.Ehex
	HZ() int
	PlType() int
	Star() string
}

func (p *planet) Size() ehex.Ehex {
	return p.size
}
func (p *planet) Atmo() ehex.Ehex {
	return p.atmo
}
func (p *planet) Hydr() ehex.Ehex {
	return p.hydr
}
func (p *planet) Life() ehex.Ehex {
	return p.life
}
func (p *planet) HZ() int {
	return p.hz
}
func (p *planet) PlType() int {
	return p.plType
}
func (p *planet) Star() string {
	return p.star
}

// func BasicDataT5(dice *dice.Dicepool, hz int) planet {
// 	plType := planetType(dice, hz)
// 	pl := planet{}
// 	pl.hz = hz
// 	pl.plType = plType
// 	pl.comment = planetTypeStr(pl.plType)
// 	pl.size = ehex.New().Set(sizeT5(dice, plType))
// 	pl.atmo = ehex.New().Set(atmoT5(dice, plType, pl.size.Value()))
// 	pl.hydr = ehex.New().Set(hydrT5(dice, plType, pl.size.Value(), pl.atmo.Value()))
// 	return pl
// }

func PhysicalData_T5(dice *dice.Dicepool, hz int, star string) *planet {
	plType := planetType(dice, hz)
	pl := planet{}
	pl.hz = hz
	pl.plType = plType
	pl.star = star
	pl.comment = planetTypeStr(pl.plType)
	pl.size = ehex.New().Set(sizeT5(dice, plType))
	pl.atmo = ehex.New().Set(atmoT5(dice, plType, pl.size.Value()))
	pl.hydr = ehex.New().Set(hydrT5(dice, plType, pl.size.Value(), pl.atmo.Value()))
	pl.calculateLifeFactor(dice)

	return &pl
}

func SetupGasGigant(star string, size ehex.Ehex, gType string, hz int) *planet {
	pl := planet{}
	pl.star = star
	pl.size = size
	pl.atmo = ehex.New().Set(0)
	pl.hydr = ehex.New().Set(0)
	pl.life = ehex.New().Set(0)
	switch gType {
	case planetTypeStr(lgg):
		pl.plType = lgg
	case planetTypeStr(sgg):
		pl.plType = sgg
	case planetTypeStr(ig):
		pl.plType = ig
	}
	pl.hz = hz
	pl.comment = gType
	return &pl
}

func SetupBelt(star string, hz int) *planet {
	pl := planet{}
	pl.star = star
	pl.size = ehex.New().Set(0)
	pl.atmo = ehex.New().Set(0)
	pl.hydr = ehex.New().Set(0)
	pl.life = ehex.New().Set(0)
	pl.plType = planetoid
	pl.hz = hz
	pl.comment = "belt"
	return &pl
}

func SetupStar(starTp string) *planet {
	p := planet{}
	p.star = starTp
	p.comment = "star"
	p.plType = Star
	return &p
}

func (pl *planet) calculateLifeFactor(dice *dice.Dicepool) {
	life := life.DetermineDominantLife(dice, pl.atmo, pl.hydr, pl.hz, pl.star).Value()
	switch pl.plType {
	case planetoid, inferno, radworld:
		life += -10
	case iceworld, stormworld:
		life += -6
	case innerworld:
		life += -4
	}
	if life < 0 {
		life = 0
	}
	pl.life = ehex.New().Set(life)
}

func stptInt2Hex(st int) ehex.Ehex {
	switch st {
	case 0:
		return ehex.New().Set("X")
	case 1:
		return ehex.New().Set("E")
	case 2:
		return ehex.New().Set("D")
	case 3:
		return ehex.New().Set("C")
	case 4:
		return ehex.New().Set("B")
	case 5:
		return ehex.New().Set("A")
	}
	return nil
}

func (pl *planet) String() string {
	s := ""
	if pl.comment == "star" {
		return "star: " + pl.star
	}
	s += pl.size.Code()
	s += pl.atmo.Code()
	s += pl.hydr.Code()
	s += "-" + pl.life.Code()
	s += fmt.Sprintf("   (star:%v HZ:%v PlType:%v)", pl.star, pl.hz, pl.comment)
	res := pl.BaseResources()
	s += fmt.Sprintf(" [%v+%v+%v+%v = %v]", pl.size.Value(), pl.atmo.Value(), pl.hydr.Value(), pl.life.Value(), res)
	return s
}

func (pl *planet) BaseResources() int {
	res := 0
	switch pl.atmo.Value() {
	case 2, 4, 7, 9:
		res++
	case 14:
		res += -2
	case 15:
		res += -1
	}
	// switch pl.life.Value() {
	// case 0:
	// case 1:
	// 	res += 1
	// case 2, 3, 4, 5, 6, 7:
	// 	res += 2
	// case 8, 9, 10:
	// 	res += 3
	// }
	// res += (pl.size.Value() + pl.atmo.Value() + pl.hydr.Value()) / 3
	for _, val := range []int{pl.size.Value(), pl.atmo.Value(), pl.hydr.Value(), pl.life.Value()} {
		res += int(math.Sqrt(float64(val)))
	}

	return res
}

/*
A  S
B  L
C  F
D  G
E  H
X  Y
*/

func sizeT5(dice *dice.Dicepool, plType int) int {
	s := -1
	switch plType {
	case hospitable, iceworld, innerworld:
		s = stdSize(dice)
	case planetoid:
		s = 0
	case radworld, stormworld:
		s = dice.Sroll("2d6")
	case inferno:
		s = dice.Sroll("1d6+6")
	case bigwirld:
		s = dice.Sroll("2d6+7")
	case worldlet:
		s = dice.Sroll("1d6-3")
	}
	if s < 0 {
		s = 0
	}
	return s
}

func atmoT5(dice *dice.Dicepool, plType, size int) int {
	atmoDM := 0
	switch plType {
	case stormworld:
		atmoDM = 4
	case planetoid, inferno:
		return 0
	}
	atmo := dice.Flux() + size + atmoDM
	if atmo < 0 {
		return 0
	}
	if atmo > 15 {
		return 15
	}
	if size == 0 {
		return 0
	}
	return atmo
}

func hydrT5(dice *dice.Dicepool, plType, size, atmo int) int {
	hydrDM := 0
	switch plType {
	case innerworld, stormworld:
		hydrDM = -4
	case inferno, iceworld:
		return 0
	}
	if size < 2 {
		return 0
	}
	if atmo < 2 || atmo > 9 {
		hydrDM = hydrDM - 4
	}
	hydr := dice.Flux() + atmo + hydrDM
	if hydr < 0 {
		return 0
	}
	if hydr > 10 {
		return 10
	}
	return hydr
}

func popsT5(dice *dice.Dicepool, plType int, life ehex.Ehex) int {
	popsDM := 0
	if life != nil {
		popsDM = (-1) * (10 - life.Value())
	}
	switch plType {
	case iceworld, stormworld:
		popsDM = -6
	case innerworld:
		popsDM = -4
	}
	roll := dice.Sroll("2d6-2")
	if roll == 10 {
		roll = dice.Sroll("2d6+3")
	}
	pops := roll + popsDM
	if pops < 0 {
		pops = 0
	}
	return pops
}

func govrT5(dice *dice.Dicepool, pops int) int {
	govr := dice.Flux() + pops
	if govr > 15 {
		return 15
	}
	if govr < 0 {
		return 0
	}
	return govr
}

func lawsT5(dice *dice.Dicepool, govr int) int {
	laws := dice.Flux() + govr
	if laws > 18 {
		return 18
	}
	if laws < 0 {
		return 0
	}
	return laws
}

func spaceportT5NMW(dice *dice.Dicepool, pops int) int {
	stpt := pops - dice.Sroll("1d6")
	if stpt <= 0 {
		stpt = 0
	}
	switch stpt {
	case 0:
		return 0
	case 1, 2:
		return 1
	case 3:
		return 2
	case 4, 5:
		return 3
	case 6, 7:
		return 4
	default:
		return 5
	}

}

func stdSize(dice *dice.Dicepool) int {
	s := dice.Sroll("2d6-2")
	if s == 10 {
		s = 9 + dice.Sroll("1d6")
	}
	return s
}

// func (pl *planet) rollTech(dice *dice.Dicepool) {
// 	tl := dice.Sroll("1d6")
// 	switch pl.stpt.Code() {
// 	case "A":
// 		tl += 6
// 	case "B":
// 		tl += 4
// 	case "C":
// 		tl += 2
// 	case "X":
// 		tl += -4
// 	}
// 	switch pl.size.Code() {
// 	case "0", "1":
// 		tl += 2
// 	case "2", "3", "4":
// 		tl += 1
// 	}
// 	switch pl.atmo.Code() {
// 	case "0", "1", "2", "3", "A", "B", "C", "D", "E", "F":
// 		tl += 1
// 	}
// 	switch pl.hydr.Code() {
// 	case "A":
// 		tl += 2
// 	case "9":
// 		tl += 1
// 	}
// 	switch pl.pops.Code() {
// 	case "1", "2", "3", "4", "5":
// 		tl += 1
// 	case "9":
// 		tl += 2
// 	case "A", "B", "C", "D", "E", "F":
// 		tl += 4
// 	}
// 	switch pl.govr.Code() {
// 	case "0", "5":
// 		tl += 1
// 	case "D":
// 		tl += -2
// 	}
// 	if tl < 0 {
// 		tl = 0
// 	}
// 	pl.tech = ehex.New().Set(tl)
// }

const (
	hospitable = iota
	planetoid
	iceworld
	radworld
	inferno
	bigwirld
	worldlet
	innerworld
	stormworld
	lgg
	sgg
	ig
	Star
)

func planetTypeStr(pt int) string {
	switch pt {
	case hospitable:
		return "Hospitable"
	case planetoid:
		return "Planetoid"
	case iceworld:
		return "Iceworld"
	case radworld:
		return "Radworld"
	case inferno:
		return "Inferno"
	case bigwirld:
		return "Bigworld"
	case worldlet:
		return "Worldlet"
	case innerworld:
		return "Innerworld"
	case stormworld:
		return "Stormworld"
	case lgg:
		return "Large Gas Gigant"
	case sgg:
		return "Small Gas Gigant"
	case ig:
		return "Ice Gigant"
	case Star:
		return "Star"
	}
	return ""
}

func planetType(dice *dice.Dicepool, hz int) int {
	if hz > 1 {
		hz = 2
	}
	switch hz {
	default:
		switch dice.Sroll("1d6") {
		case 1:
			return inferno
		case 2:
			return innerworld
		case 3:
			return bigwirld
		case 4:
			return stormworld
		case 5:
			return radworld
		case 6:
			return hospitable
		}
	case 2:
		switch dice.Sroll("1d6") {
		case 1:
			return worldlet
		case 2:
			return iceworld
		case 3:
			return bigwirld
		case 4:
			return iceworld
		case 5:
			return radworld
		case 6:
			return iceworld
		}
	}
	return -1
}

func sizeUF(dice *dice.Dicepool, hz int) int {
	if hz < 0 {
		hz = -1
	}
	if hz > 0 {
		hz = 1
	}
	sizeGrade := 0
	sizeRoll := dice.Sroll("2d6")
	switch hz {
	case -1:
		switch sizeRoll {
		case 2, 3, 4:
			sizeGrade = 1
		case 5, 6, 7:
			sizeGrade = 2
		case 8, 9:
			sizeGrade = 3
		case 10, 11:
			sizeGrade = 4
		case 12:
			sizeGrade = 5
		}
	case 0:
		switch sizeRoll {
		case 2:
			sizeGrade = 1
		case 3:
			sizeGrade = 2
		case 4, 5, 6:
			sizeGrade = 3
		case 7, 8, 9, 10, 11:
			sizeGrade = 4
		case 12:
			sizeGrade = 5
		}
	case 1:
		switch sizeRoll {
		case 2, 3, 4, 5, 6:
			sizeGrade = 1
		case 7, 8, 9:
			sizeGrade = 2
		case 10:
			sizeGrade = 3
		case 11:
			sizeGrade = 4
		case 12:
			sizeGrade = 5
		}
	}
	size := 0
	switch sizeGrade {
	case 1:
		switch dice.Sroll("1d2") {
		case 1:
			size = 0
		case 2:
			size = 1
		}
	case 2:
		switch dice.Sroll("1d2") {
		case 1:
			size = 2
		case 2:
			size = 3
		}
	case 3:
		switch dice.Sroll("1d3") {
		case 1:
			size = 4
		case 2:
			size = 5
		case 3:
			size = 6
		}
	case 4:
		switch dice.Sroll("2d6") {
		case 2, 3, 4, 5:
			size = 7
		case 6, 7, 8, 9:
			size = 8
		case 10, 11, 12:
			size = 9
		}
	case 5:
		size = 8 + dice.Sroll("2d6")
	}
	return size
}

func BasicDataMGT2(dice *dice.Dicepool, hz int) planet {
	p := planet{}
	p.hz = hz
	p.size = determineSize(dice)
	p.atmo = determineAtmo(dice, p.size.Value())
	p.hydr = determineHydr(dice, p.size.Value(), p.atmo.Value(), p.hz)
	return p
}

func determineSize(dice *dice.Dicepool) ehex.Ehex {
	size := dice.Sroll("2d6-2")
	inc := true
	for inc && size <= 15 && size >= 10 {
		switch dice.Sroll("1d2") {
		case 1:
			inc = false
		case 2:
			size++
		}
	}
	size = utils.BoundInt(size, 0, 15)
	return ehex.New().Set(size)
}

func determineAtmo(dice *dice.Dicepool, size int) ehex.Ehex {
	atmo := dice.Sroll("2d6") - 7 + size
	switch size {
	case 0:
		atmo = 0
	case 1:
		atmo = atmo - 5
	case 3, 4:
		switch atmo {
		case 4, 5, 8, 9:
			atmo = atmo - 2
		case 7:
			atmo = 4
		case 6:
			atmo = 5
		case 2, 3:
			atmo = 1

		}
	}
	atmo = utils.BoundInt(atmo, 0, 15)
	return ehex.New().Set(atmo)
}

func determineHydr(dice *dice.Dicepool, size, atmo, hz int) ehex.Ehex {
	dm := atmo
	switch size {
	case 0, 1:
		return ehex.New().Set(0)
	}
	switch atmo {
	case 0, 1, 9, 10, 11, 12, 14:
		dm = -4
		if hz == -1 {
			dm = dm - 2
		}
		if hz <= -2 {
			dm = dm - 6
		}
	case 13, 15:
		dm = -4
	}
	switch hz {
	case 1, -1:
		dm = dm - 2
	case 0:
	default:
		dm = dm - 6
	}
	hydr := utils.BoundInt(dice.Sroll("2d6")-7+dm, 0, 10)
	return ehex.New().Set(hydr)
}

func OfferWorldOrbit(i int) int {
	return []int{12, 11, 10, 8, 6, 4, 2, 0, 1, 3, 5, 7, 9}[i]
}

func OfferWorldOrbit2(i int) int {
	return []int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7}[i]
}
