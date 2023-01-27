package life

import (
	"math"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
)

/*
2d10 + Atmo + Hydr + Temp + HZ + Star
*/

func DetermineDominantLife(dice *dice.Dicepool, atmo ehex.Ehex, hydr ehex.Ehex, hz int, star string) ehex.Ehex {
	atmoDM := 0
	switch atmo.Value() {
	case 0, 1, 2, 3:
		return ehex.New().Set(0)
	case 10, 11, 12, 15:
		atmoDM -= 20
	case 4, 5:
		atmoDM -= 5
	case 6, 7, 8, 9:
		atmoDM += 10
	}
	hydrDM := 0
	switch hydr.Value() {
	case 0:
		return ehex.New().Set(0)
	case 1, 2:
		hydrDM -= 15
	case 3, 4, 10:
		hydrDM -= 5
	case 5, 6, 7, 8, 9:
		hydrDM += 5
	}
	tempDM := 0 - int(math.Abs(float64(hz))*10)
	starDM := 0
	code, _, _ := stellar.Decode(star)
	switch code {
	default:
		starDM -= 10
	case "G":
		starDM += 10
	case "F", "K":
	}
	rollDM := atmoDM + hydrDM + tempDM + starDM
	r := dice.Sroll("2d10") + rollDM
	if r <= 0 {
		return ehex.New().Set(0)
	}
	switch r {
	case 1, 2, 3:
		return ehex.New().Set(1)
	case 4, 5:
		return ehex.New().Set(2)
	case 6, 7, 8:
		return ehex.New().Set(3)
	case 9, 10:
		return ehex.New().Set(4)
	case 11, 12:
		return ehex.New().Set(5)
	case 13, 14:
		return ehex.New().Set(6)
	case 15, 16:
		return ehex.New().Set(7)
	case 17, 18:
		return ehex.New().Set(8)
	case 19, 20:
		return ehex.New().Set(9)
	default:
		return ehex.New().Set(10)
	}

}
