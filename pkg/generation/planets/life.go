package planets

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

func GenerateDominantLife(dice *dice.Dicepool, atmo ehex.Ehex, hydr ehex.Ehex, hz ehex.Ehex, star string) ehex.Ehex {
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
	tempDM := 0 - int(math.Abs(float64(hz.Value()-10))*5)
	starDM := -10
	if strings.Contains(star, "G") {
		starDM += 20
	}
	if strings.Contains(star, "F") || strings.Contains(star, "K") {
		starDM += 10
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

func GenerateNativeLife(dice *dice.Dicepool, atmo ehex.Ehex, hydr ehex.Ehex, hz ehex.Ehex, star string) ehex.Ehex {
	atmoDM := 0
	switch atmo.Value() {
	case 0:
		atmoDM = -3
	case 1:
		atmoDM = -2
	case 6, 7:
		atmoDM = 2
	case 4, 5, 8, 9:
		atmoDM = 1
	case 2, 3, 13, 14:
		atmoDM = -1
	}
	hydrDM := 0
	switch hydr.Value() {
	case 0:
		hydrDM = -2
	case 1, 2:
		hydrDM = -1
	case 3, 4, 10:
		hydrDM = 0
	case 5, 6, 7, 8, 9:
		hydrDM = 1
	}
	tempDM := 0 - int(math.Abs(float64(hz.Value()-10))*1)
	starDM := 0
	if strings.Contains(star, "G") || strings.Contains(star, "K") {
		starDM += 1
	}
	if strings.Contains(star, "F") || strings.Contains(star, "A") || strings.Contains(star, "B") {
		starDM -= 1
	}

	rollDM := atmoDM + hydrDM + tempDM + starDM

	r := dice.Sroll("2d6") + rollDM - 10
	fmt.Println("Roll1:", r-rollDM+10, "rollDM:", rollDM, atmoDM, hydrDM, tempDM, starDM)
	if r <= 0 {
		return ehex.New().Set(0)
	}
	r2 := dice.Sroll("2d6-2") + atmoDM + hydrDM + tempDM
	fmt.Println("Roll2:", r2-atmoDM-hydrDM-tempDM, "rollDM:", atmoDM, hydrDM, tempDM)
	if r2 < 0 {
		r2 = 0
	}
	if r2 > 10 {
		r2 = 10
	}
	return ehex.New().Set(r2)
}
