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

func (st *stellar) String() string {
	str := ""
	for i := 1; i <= 8; i++ {
		if v, ok := st.systemstars[i]; ok == true {
			str += v.star + " "
		}
	}
	return str
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

func ConstructNew(paradigm string, knownData ...dataFeed) (*stellar, error) {
	err := (error)(nil)
	sector := ""
	subSect := ""
	hex := ""
	mwName := ""
	planets := ""
	pbg := ""
	for _, feed := range knownData {
		switch feed.key {
		default:
			return nil, fmt.Errorf("unknown datafeed key: '%v'", feed.key)
		case KEY_HEX:
			hex = feed.val
		case KEY_SECTOR:
			sector = feed.val
		case KEY_SUBSECT:
			subSect = feed.val
		case KEY_MW:
			mwName = feed.val
		case KEY_PLANETS:
			planets = feed.val
		case KEY_PBG:
			pbg = feed.val
		}
	}
	stlr := stellar{}
	stlr.systemstars = make(map[int]*star)
	seed := sector + "  " + subSect + "  " + hex + "  " + mwName + "  " + planets + "  " + pbg
	c := newConstructor(paradigm, seed)
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

func newConstructor(paradigm, seed string) *constructor {
	c := constructor{}
	switch paradigm {
	default:
		c.err = fmt.Errorf("paradigm '%v' is unknown/unimplemented", paradigm)
		return &c
	case CONSTRUCTOR_PARADIGM_T5:
		c.paradigm = paradigm
	}
	c.dice = dice.New()
	c.seed = seed
	c.dice.SetSeed(c.seed)
	return &c
}

func (c *constructor) newStar(pos int) *star {
	st := star{}
	return &st
}
