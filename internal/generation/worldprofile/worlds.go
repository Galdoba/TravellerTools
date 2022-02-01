package worldprofile

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
)

/*
NewMain(seed string) string
NewSecondary(ssd survey.SecondSurveyData, otherType int) string



*/
//NewMain - New mainworld UWP
func NewMain(seed string) string {
	dp := dice.New().SetSeed(seed)
	///////
	sz := dp.Roll("2d6").DM(-2).Sum()
	//////
	at := dp.Flux() + sz
	switch {
	case at < 0 || sz == 0:
		at = 0
	case at > 15:
		at = 15
	}
	//////
	dm := 0
	switch at {
	case 0, 1, 10, 11, 12, 13, 14, 15:
		dm = -4
	}
	hd := dp.Flux() + at + dm
	if sz < 2 {
		hd = 0
	}
	///////
	pp := dp.Roll("2d6").DM(-2).Sum()
	if pp == 10 {
		pp = dp.Roll("2d6").DM(3).Sum()
	}
	///////
	gv := dp.Flux() + pp
	switch {
	case gv < 0:
		gv = 0
	case gv > 15:
		gv = 15
	}
	lw := dp.Flux() + gv
	if lw > 18 {
		lw = 18
	}
	stDM := 0
	switch {
	case pp == 8 || pp == 9:
		stDM = 1
	case pp > 9:
		stDM = 2
	case pp < 3:
		stDM = -2
	case pp == 3 || pp == 4:
		stDM = -1
	}
	st := "X"
	stR := dp.Roll("2d6").DM(stDM).Sum()
	switch stR {
	case 3, 4:
		st = "E"
	case 5, 6:
		st = "D"
	case 7, 8:
		st = "C"
	case 9, 10:
		st = "B"
	}
	if stR > 10 {
		st = "A"
	}

	tl := dp.Roll("1d6").Sum()
	switch st {
	case "A":
		tl = tl + 6
	case "B":
		tl = tl + 4
	case "C":
		tl = tl + 2
	case "X":
		tl = tl - 4
	}
	switch sz {
	case 0, 1:
		tl = tl + 2
	case 2, 3, 4:
		tl = tl + 1
	}
	switch at {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		tl = tl + 1
	}
	switch hd {
	case 9:
	}

	res := st
	res += fmt.Sprintf(ehex.New().Set(sz).Code())
	res += fmt.Sprintf(ehex.New().Set(at).Code())
	res += fmt.Sprintf(ehex.New().Set(hd).Code())
	res += fmt.Sprintf(ehex.New().Set(pp).Code())
	res += fmt.Sprintf(ehex.New().Set(gv).Code())
	res += fmt.Sprintf(ehex.New().Set(lw).Code())
	res += "-"
	res += fmt.Sprintf(ehex.New().Set(sz).Code())

	return res
}
