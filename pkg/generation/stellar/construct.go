package stellar

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar/t5"
)

type stellar struct {
	//constructor *constructor
	systemstars map[int]*star
	lo          []int
}

type Stellar interface {
	String() string
	Layout() []int
	Stars() []string
	//StarMap() map[int]string
}

func StarMap(st Stellar) map[int]string {
	strPos := st.Layout()[2:]
	stars := st.Stars()
	strMap := make(map[int]string)
	for i, v := range strPos {
		strMap[v] = stars[i]
	}
	return strMap
}

func met(sl []int, elem int) bool {
	for _, v := range sl {
		if v == elem {
			return true
		}
	}
	return false
}

func (st *stellar) String() string {
	str := ""
	for i := 1; i <= 8; i++ {
		if v, ok := st.systemstars[i]; ok == true {
			str += v.star + " "
		}
	}
	return str
}

func (st *stellar) Layout() []int {
	return st.lo
}

func (st *stellar) Stars() []string {
	strs := []string{}
	for _, l := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		if str, ok := st.systemstars[l]; ok {
			strs = append(strs, str.star)
		}
	}
	return strs
}

type dataFeed struct {
	key string
	val string
}

const (
	KEY_HEX     = "Hex"
	KEY_SECTOR  = "Sector"
	KEY_SUBSECT = "Subsector"
	KEY_MW      = "Mainworld"
	KEY_PLANETS = "Planets"
	KEY_PBG     = "PBG"
)

func Instruction(key, val string) dataFeed {
	return dataFeed{key, val}
}

func starLayout(dice *dice.Dicepool) []int {
	lo := []int{dice.Flux() + dice.Sroll("1d3-2")}
	lo = append(lo, dice.Flux()) //2 primeFlux (первый для спектра второй для размеров)
	lo = append(lo, posPrimary)
	if dice.Flux() >= 3 {
		lo = append(lo, posPrimaryComp)
	}
	for _, pos := range []int{posClose, posNear, posFar} {
		if dice.Flux() >= 3 {
			lo = append(lo, pos)
			if dice.Flux() >= 3 {
				lo = append(lo, pos+1)
			}
		}
	}
	return lo
}

func ConstructNew(paradigm string, dice *dice.Dicepool) (*stellar, error) {
	err := (error)(nil)
	stlr := stellar{}
	stlr.systemstars = make(map[int]*star)
	c := newConstructor(paradigm, dice)
	if c.err != nil {
		return &stlr, fmt.Errorf("constructor: %v", c.err)
	}
	stlr.lo = starLayout(c.dice)
	for i, pos := range stlr.lo {
		if i == 0 || i == 1 {
			continue
		}
		spec, dec, size := t5.StarTypeAndSize(c.dice, stlr.lo[0], stlr.lo[1], pos)
		if spec == "BD" {
			spec = brownDwarfSpectral(c.dice)
		}
		str, err := NewStar(spec, dec, size)
		str.orbitingOn, str.maxHighOrbit = t5.StarOrbit(c.dice, pos)

		if err != nil {
			return &stlr, fmt.Errorf("newStar: %v (%v-%v-%v) [%v]", err.Error(), spec, dec, size, pos)
		}
		//fmt.Println(str)

		stlr.systemstars[pos] = str
	}
	for a := 0; a < 11; a++ {
		if val, ok := stlr.systemstars[a]; ok == true {
			if val == nil {
				return &stlr, fmt.Errorf("star with position %v added but is <nil>", a)
			}
		}
	}
	return &stlr, err
}

const (
	posUndefined = iota
	posPrimary
	posPrimaryComp
	posClose
	posCloseComp
	posNear
	posNearComp
	posFar
	posFarComp
	posWRONG
)

type constructor struct {
	paradigm string
	dice     *dice.Dicepool
	seed     string
	err      error
}

const (
	CONSTRUCTOR_PARADIGM_T5 = "T5"
)

func newConstructor(paradigm string, dice *dice.Dicepool) *constructor {
	c := constructor{}
	switch paradigm {
	default:
		c.err = fmt.Errorf("paradigm '%v' is unknown/unimplemented", paradigm)
		return &c
	case CONSTRUCTOR_PARADIGM_T5:
		c.paradigm = paradigm
	}
	c.dice = dice
	return &c
}

func (c *constructor) newStar(pos int) *star {
	st := star{}
	return &st
}
