package calculations

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/internal/ehex"
)

func EconomicExtention(uwp, pbg string, ix int) string {

	hex := strings.Split(uwp, "")
	tl := ehex.New().Set(hex[8])
	pop := ehex.New().Set(hex[4])
	hex2 := strings.Split(pbg, "")
	b := ehex.New().Set(hex2[1])
	g := ehex.New().Set(hex2[2])
	res, lab, inf, eff := 0, 0, 0, 0
	d := dice.New().SetSeed(uwp + pbg)
	//////////////////////////////////
	res = d.Roll("2d6").Sum()
	if tl.Value() >= 8 {
		res += b.Value() + g.Value()
	}
	resH := ehex.New().Set(res)
	//////////////////////////////////
	lab = pop.Value() - 1
	if lab == 0 {
		lab = 1
	}
	labH := ehex.New().Set(lab)
	//////////////////////////////////
	switch pop.Value() {
	case 0:
		inf = 0
	case 1, 2, 3:
		inf = ix
	case 4, 5, 6:
		inf = d.Roll("1d6").Sum() + ix
	default:
		inf = d.Roll("2d6").Sum() + ix
	}
	infH := ehex.New().Set(inf)
	//////////////////////////////////
	eff = d.Flux()
	if eff == 0 {
		eff = 1
	}
	sep := ""
	if eff > -1 {
		sep = "+"
	}
	//////////////////////////////////
	return fmt.Sprintf("(%v%v%v%v%v)", resH.String(), labH.String(), infH.String(), sep, eff)
}

func ExValid(eX string) bool {
	hex := strings.Split(eX, "")
	if len(hex) != 7 {
		return false
	}
	r := ehex.New().Set(hex[1])
	l := ehex.New().Set(hex[2])
	i := ehex.New().Set(hex[3])
	e := hex[4] + hex[5]
	switch {
	case !resValid(r):
		return false
	case !labValid(l):
		return false
	case !infValid(i):
		return false
	case !effValid(e):
		return false
	}
	return true
}

func FixEconomicExtention(eX, uwp, pbg, seed string, ix int) string {
	var hex []string
	hex = strings.Split(eX, "")
	for len(hex) < 7 {
		hex = append(hex, "?")
	}
	hex[0] = "("
	r := ehex.New().Set(hex[1])
	l := ehex.New().Set(hex[2])
	i := ehex.New().Set(hex[3])
	e := hex[4] + hex[5]
	u := strings.Split(uwp, "")
	tl := ehex.New().Set(u[8])
	pop := ehex.New().Set(u[4])
	d := dice.New().SetSeed(seed)
	if !resValid(r) {
		hex[1] = reRollResourses(d, tl, pbg).Code()
	}
	if !labValid(l) {
		hex[2] = reRollLabor(d, pop).Code()
	}
	if !infValid(i) {
		hex[3] = reRollInfrastructure(d, pop, ix).Code()
	}
	if !effValid(e) {
		effHx := strings.Split(rerollEfficiency(d), "")
		hex[4] = effHx[0]
		hex[5] = effHx[1]
	}
	hex[6] = ")"
	return strings.Join(hex, "")
}

func resValid(r ehex.Ehex) bool {
	switch {
	case r.Value() < 0:
		return false
	case r.Value() > 19:
		return false
	}
	return true
}

func reRollResourses(d *dice.Dicepool, tl ehex.Ehex, pbg string) ehex.Ehex {
	hex2 := strings.Split(pbg, "")
	b := ehex.New().Set(hex2[1])
	g := ehex.New().Set(hex2[2])
	res := d.Roll("2d6").Sum()
	if tl.Value() >= 8 {
		res += b.Value() + g.Value()
	}
	if res < 0 {
		res = 0
	}
	return ehex.New().Set(res)
}

func labValid(l ehex.Ehex) bool {
	switch {
	case l.Code() == " ":
		return false
	case l.Value() < 0:
		return false
	case l.Value() > 14:
		return false
	}
	return true
}

func reRollLabor(d *dice.Dicepool, pop ehex.Ehex) ehex.Ehex {
	lab := pop.Value() - 1
	if lab == -1 {
		lab = 0
	}
	return ehex.New().Set(lab)
}

func infValid(i ehex.Ehex) bool {
	switch {
	case i.Value() < 0:
		return false
	case i.Value() > 20:
		return false
	}
	return true
}

func reRollInfrastructure(d *dice.Dicepool, pop ehex.Ehex, ix int) ehex.Ehex {
	inf := 0
	switch pop.Value() {
	case 0:
		inf = 0
	case 1, 2, 3:
		inf = ix
	case 4, 5, 6:
		inf = d.Roll("1d6").Sum() + ix
	default:
		inf = d.Roll("2d6").Sum() + ix
	}
	if inf < 0 {
		inf = 0
	}
	return ehex.New().Set(inf)
}

func effValid(eff string) bool {
	switch eff {
	case "+1", "+2", "+3", "+4", "+5", "-1", "-2", "-3", "-4", "-5":
		return true
	}
	return false
}

func rerollEfficiency(d *dice.Dicepool) string {
	eff := d.Flux()
	if eff == 0 {
		eff = 1
	}
	sep := ""
	if eff > 0 {
		sep = "+"
	}
	return sep + strconv.Itoa(eff)
}
