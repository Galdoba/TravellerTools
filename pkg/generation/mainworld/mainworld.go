package mainworld

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
	"github.com/Galdoba/TravellerTools/pkg/generation/table"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/utils"
)

type mainworldData struct {
	roller  *dice.Dicepool
	stellar string
	parent  string
	body    string
	hz      int
	size    int
	atmo    int
	hydr    int
	temp    int
	ggas    int
	belt    int
	pops    int
	popMult int
	govr    int
	laws    int
	stpt    int
	tech    int
	trco    []int
}

type MainWorld interface {
	UWP() uwp.UWP
	Stellar() string
	Classifications() []int
}

func New(dice *dice.Dicepool) *mainworldData {
	mw := mainworldData{}
	mw.stellar = "-1"
	mw.hz = -10
	mw.size = -1
	mw.atmo = -1
	mw.hydr = -1
	mw.temp = -10
	mw.ggas = -1
	mw.belt = -1
	mw.pops = -1
	mw.popMult = -1
	mw.govr = -1
	mw.laws = -1
	mw.stpt = -1
	mw.tech = -1
	mw.roller = dice
	return &mw
}

func (mw *mainworldData) DetermineSystemData(options ...string) error {
	for _, opt := range options {
		switch opt {
		default:
			if stellar.Valid(opt) {
				stars := stellar.Parse(opt)
				mw.stellar = strings.Join(stars, " ")
				return nil
			}
			return fmt.Errorf(" mw.DetermineSystemData(): invalid option '%v'", opt)
		}
	}
	if mw.stellar != "-1" {
		return fmt.Errorf(" mw.DetermineSystemData(): stellar was already asigned")
	}
	mw.stellar = stellar.GenerateStellar(mw.roller)
	stars := stellar.Parse(mw.stellar)
	if !stellar.Valid(stars[0]) {
		panic(fmt.Sprintf("TODO: type %v - %v", stars[0], stars))
	}
	switch mw.stellar {
	case "*RP":
		mw.hz = -2
		mw.temp = -2
		mw.ggas = 0
		mw.belt = 0
		return nil
	case "*RGG":
		mw.hz = -2
		mw.temp = -2
		mw.ggas = 1
		mw.belt = 0
		return nil
	case "NS":
		mw.hz = -2
		mw.temp = -2
		mw.atmo = 0
		mw.hydr = 0
		return nil
	}
	dm := 0
	switch {
	case strings.Contains(stars[0], "O"):
		dm = -2
	case strings.Contains(stars[0], "B"):
		dm = -2
	case strings.Contains(stars[0], "M"):
		dm = 2
	}
	switch mw.roller.Sroll("2d6") + dm {
	case 0, 1:
		mw.hz = -2
	case 2, 3, 4:
		mw.hz = -1
	case 5, 6, 7, 8, 9, 10:
		mw.hz = 0
	case 11, 12:
		mw.hz = 1
	case 13, 14:
		mw.hz = 2
	}
	return nil
}

func (mw *mainworldData) DeterminePhysicalCharacteristics(options ...string) error {
	for _, opt := range options {
		switch {
		default:
			return fmt.Errorf(" mw.DeterminePhysicalCharacteristics(): invalid option '%v'", opt)
		case strings.HasPrefix(opt, "size:"):
			code := strings.TrimSuffix(opt, "size:")
			val := ehex.New().Set(code)
			if utils.InRange(val.Value(), 0, 24) {
				mw.size = val.Value()
			}
		case strings.HasPrefix(opt, "atmo:"):
			code := strings.TrimSuffix(opt, "atmo:")
			val := ehex.New().Set(code)
			if utils.InRange(val.Value(), 0, 15) {
				mw.atmo = val.Value()
			}
		case strings.HasPrefix(opt, "hydr:"):
			code := strings.TrimSuffix(opt, "hydr:")
			val := ehex.New().Set(code)
			if utils.InRange(val.Value(), 0, 10) {
				mw.hydr = val.Value()
			}
		case strings.HasPrefix(opt, "temp:"):
			code := strings.TrimSuffix(opt, "temp:")
			val, err := strconv.Atoi(code)
			if err != nil {
				return fmt.Errorf("mw.DeterminePhysicalCharacteristics(): invalid option '%v'", opt)
			}
			if utils.InRange(val, -2, 2) {
				mw.temp = val
			}
		case strings.HasPrefix(opt, "ggas:"):
			code := strings.TrimSuffix(opt, "ggas:")
			if code == "G" {
				mw.ggas = utils.BoundInt((mw.roller.Sroll("2d6")/2 - 2), 0, 4)
				continue
			}
			val, err := strconv.Atoi(code)
			if err != nil {
				return fmt.Errorf("mw.DeterminePhysicalCharacteristics(): invalid option '%v'", opt)
			}
			if utils.InRange(val, 0, 5) {
				mw.ggas = val
			}
		case strings.HasPrefix(opt, "belt:"):
			code := strings.TrimSuffix(opt, "belt:")
			val, err := strconv.Atoi(code)
			if err != nil {
				return fmt.Errorf("mw.DeterminePhysicalCharacteristics(): invalid option '%v'", opt)
			}
			if utils.InRange(val, 0, 3) {
				mw.belt = val
			}
			// case strings.HasPrefix(opt, "popMult:"):
			// 	code := strings.TrimSuffix(opt, "popMult:")
			// 	val, err := strconv.Atoi(code)
			// 	if err != nil {
			// 		return fmt.Errorf("mw.DeterminePhysicalCharacteristics(): invalid option '%v'", opt)
			// 	}
			// 	if utils.InRange(val, 0, 9) {
			// 		mw.popMult = val
			// 	}
		}
	}
	if mw.size == -1 {
		mw.size = mw.roller.Sroll("2d6-2")
		inc := true
		for inc && mw.size <= 19 && mw.size >= 10 {
			switch mw.roller.Sroll("1d2") {
			case 1:
				inc = false
			case 2:
				mw.size++
			}
		}
		mw.size = utils.BoundInt(mw.size, 0, 19)
	}
	if mw.atmo == -1 {
		mw.atmo = mw.roller.Sroll("2d6") - 7 + mw.size
		switch mw.size {
		case 0:
			mw.atmo = 0
		case 1:
			mw.atmo = mw.atmo - 5
		case 3, 4:
			switch mw.atmo {
			case 4, 5, 8, 9:
				mw.atmo = mw.atmo - 2
			case 7:
				mw.atmo = 4
			case 6:
				mw.atmo = 5
			case 2, 3:
				mw.atmo = 1

			}
		}
		mw.atmo = utils.BoundInt(mw.atmo, 0, 15)
	}
	if mw.temp == -10 {
		r := mw.roller.Sroll("2d6")
		switch mw.atmo {
		case 2, 3:
			r = r - 2
		case 4, 5, 14:
			r = r - 1
		case 8, 9:
			r = r + 1
		case 10, 13, 15:
			r = r + 2
		case 11, 12:
			r = r + 6
		}
		r = r + (mw.hz * -4)
		tmp := table.DiceChart(
			table.Row("2-", "-2"),
			table.Row("3-4", "-1"),
			table.Row("5-9", "0"),
			table.Row("10-11", "1"),
			table.Row("12+", "2"),
		).Result(r)
		mw.temp, _ = strconv.Atoi(tmp)
	}
	if mw.hydr == -1 {
		//maxStat := utils.Max(mw.size, mw.atmo)
		//dm := utils.Max(mw.size, mw.atmo)
		dm := mw.atmo
		switch mw.atmo {
		case 0, 1, 9, 10, 11, 12, 13, 14, 15:
			dm = -4
		}
		if mw.temp == 1 {
			dm = dm - 2
		}
		if mw.temp == 2 {
			dm = dm - 6
		}
		mw.hydr = utils.BoundInt(mw.roller.Sroll("2d6")-7+dm, 0, 10)
	}
	if mw.ggas == -1 {
		mw.ggas = utils.BoundInt((mw.roller.Sroll("2d6")/2 - 2), 0, 4)
	}
	if mw.belt == -1 {
		mw.belt = utils.BoundInt((mw.roller.Sroll("1d6") - 3), 0, 3)
	}
	return nil
}

func (mw *mainworldData) DetermineSocialCharacteristics(options ...string) error {
	patern := 0
	for _, opt := range options {
		switch {
		default:
			return fmt.Errorf(" mw.DetermineSystemData(): invalid option '%v'", opt)
		case strings.HasPrefix(opt, "patern:"):
			code := strings.TrimSuffix(opt, "patern:")
			val := ehex.New().Set(code)
			if utils.InRange(val.Value(), 1, 5) {
				patern = val.Value()
			}
		}
	}
	if mw.pops == -1 {
		switch patern {
		case 0, 1:
			mw.pops = mw.roller.Sroll("2d6") - 2
			for mw.pops >= 10 && mw.pops <= 15 {
				switch mw.roller.Sroll("1d6") {
				default:
					break
				case 6:
					mw.pops++
				}
			}
		}
	}
	mw.pops = utils.BoundInt(mw.pops, 0, 15)
	if mw.govr == -1 {
		mw.govr = utils.BoundInt(mw.roller.Sroll("2d6")-7+mw.pops, 0, 15)
	}
	if mw.laws == -1 {
		mw.laws = utils.BoundInt(mw.roller.Sroll("2d6")-7+mw.govr, 0, 15)
	}

	return nil
}

func (mw *mainworldData) DetermineAdditionalCharacteristics(options ...string) error {
	dm := 0
	switch mw.pops {
	case 0, 1, 2:
		dm = -2
	case 3, 4:
		dm = -1
	case 8, 9:
		dm = 1
	}
	if mw.pops > 9 {
		dm = 1 + ((mw.pops - 10) / 3)
	}
	code := "X"
	r := mw.roller.Sroll("2d6") + dm
	switch r {
	case 3, 4:
		code = "E"
	case 5, 6:
		code = "D"
	case 7, 8:
		code = "C"
	case 9, 10:
		code = "B"
	}
	if r >= 11 {
		code = "A"
	}
	mw.stpt = ehex.ValueOf(code)
	//////////////
	tl := mw.roller.Sroll("1d6")
	switch mw.size { //Size
	case 0, 1:
		tl = tl + 2
	case 2, 3, 4:
		tl++
	}
	switch mw.atmo { //Atmo
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		tl++
	}
	switch mw.hydr { //Hydro
	case 0, 9:
		tl++
	case 10:
		tl = tl + 2
	}
	switch mw.pops { //pop
	case 1, 2, 3, 4, 5, 8:
		tl++
	case 9:
		tl = tl + 2
	case 10:
		tl = tl + 4
	case 11:
		tl = tl + 3
	case 12:
		tl = tl + 2
	case 13:
		tl = tl + 1
	}
	switch mw.govr { //gov
	case 0, 5:
		tl++
	case 7:
		tl = tl + 2
	case 13, 14:
		tl = tl - 2
	}
	switch ehex.ToCode(mw.stpt) {
	case "X":
		tl = tl - 4
	case "C":
		tl = tl + 2
	case "B":
		tl = tl + 4
	case "A":
		tl = tl + 6
	}
	if tl < 0 {
		tl = 0
	}
	mw.tech = tl
	mw.enviromentalLimits()
	u := mw.UWP()
	tc, _ := classifications.FromUWP(u)
	mw.trco = classifications.Values(tc)
	return nil
}

func (mw *mainworldData) enviromentalLimits() error {
	if mw.pops == 0 {
		mw.govr = 0
		mw.laws = 0
		mw.stpt = ehex.ValueOf("X")
		if mw.roller.Sroll("2d6") != 12 {
			mw.tech = 0
		}
	}
	tlLimit := 0
	switch mw.atmo {
	case 5, 6, 8:
	case 4, 7, 9:
		tlLimit = 3
	case 10:
		tlLimit = 8
	case 11:
		tlLimit = 9
	case 12:
		tlLimit = 10
	case 13, 14:
		tlLimit = 5
	case 15:
		tlLimit = 8
	}

	if mw.tech < tlLimit {
		mw.pops = 0
		mw.govr = 0
		mw.laws = 0
		mw.stpt = ehex.ValueOf("X")
		mw.tech = 0
	}
	switch {
	case mw.tech < 5 || mw.pops+(mw.tech/2) < 3:
		switch ehex.ToCode(mw.stpt) {
		case "A", "B", "C", "D":
			mw.stpt = ehex.ValueOf("E")
		}
	case mw.tech < 7 || mw.pops+(mw.tech/2) < 5:
		switch ehex.ToCode(mw.stpt) {
		case "A", "B", "C":
			mw.stpt = ehex.ValueOf("D")
		}
	case mw.tech < 8 || mw.pops+(mw.tech/2) < 6:
		switch ehex.ToCode(mw.stpt) {
		case "A", "B":
			mw.stpt = ehex.ValueOf("C")
		}
	case mw.tech < 9 || mw.pops+(mw.tech/2) < 7:
		switch ehex.ToCode(mw.stpt) {
		case "A":
			mw.stpt = ehex.ValueOf("B")
		}
	}
	return nil
}

/////////////////
func (mw *mainworldData) ProfileUncharted() string {
	s := "_" + ehex.ToCode(mw.size) + ehex.ToCode(mw.atmo) + ehex.ToCode(mw.hydr) + "___-_   _" + ehex.ToCode(mw.belt) + ehex.ToCode(mw.ggas) + " temp: " + fmt.Sprintf("%v", mw.temp) + " hz: " + fmt.Sprintf("%v", mw.hz) + " " + mw.stellar
	return s
}

func (mw *mainworldData) UWP() uwp.UWP {
	str := ehex.ToCode(mw.stpt) + ehex.ToCode(mw.size) + ehex.ToCode(mw.atmo) + ehex.ToCode(mw.hydr) + ehex.ToCode(mw.pops) + ehex.ToCode(mw.govr) + ehex.ToCode(mw.laws) + "-" + ehex.ToCode(mw.tech)
	fmt.Println("****", str)
	u, _ := uwp.FromString(str)
	return u
}
