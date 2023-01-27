package planets

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/life"
	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
	"github.com/Galdoba/utils"
)

/*
для генерации планет нужны данные о:
Size
Atmo
HZ
Hydr

*/

type planet struct {
	dice         *dice.Dicepool
	planetType   string
	parentStar   string
	physicalData []ehex.Ehex
	hz           int
	dominantLife ehex.Ehex
}

func New(dice *dice.Dicepool) *planet {
	p := planet{}
	p.dice = dice
	return &p
}

func (p *planet) SetParentStar(star string) error {
	if !stellar.Valid(star) {
		return fmt.Errorf("star '%v' is invalid", star)
	}
	p.parentStar = star
	return nil
}

func (p *planet) String() string {
	return fmt.Sprintf("%v %v %v%v%v %v %v", p.parentStar, p.planetType, p.physicalData[0], p.physicalData[1], p.physicalData[2], p.dominantLife, p.hz)
}

func (p *planet) GenerateBasic() error {
	if p.parentStar == "" {
		p.setRandomParentStar()
	}
	for _, err := range []error{
		p.determineHZ(),
		p.determineSize(),
		p.determineAtmo(),
		p.determineHydr(),
	} {
		if err != nil {
			return err
		}
	}
	p.dominantLife = life.DetermineDominantLife(p.dice, p.physicalData[1], p.physicalData[2], p.hz, p.parentStar)
	return nil
}

func (p *planet) setRandomParentStar() {
	p.parentStar = stellar.GenerateStellarOneStar(p.dice)
}

func (p *planet) determineHZ() error {
	starDM := 0
	class, _, _ := stellar.Decode(p.parentStar)
	switch class {
	case "M":
		starDM = 2
	case "O", "B":
		starDM = -2
	case "L", "T", "Y":
		starDM = -4
	}
	r := p.dice.Flux() + starDM
	switch r {
	default:
		if r < -5 {
			p.hz = -2
		}
		if r < -5 {
			p.hz = -2
		}
	case -5, -4, -3:
		p.hz = -1
	case -2, -1, 0, 1, 2:
		p.hz = 0
	case 3, 4, 5:
		p.hz = 1
	}
	return nil
}

func (p *planet) determineSize() error {
	size := p.dice.Sroll("2d6-2")
	inc := true
	for inc && size <= 15 && size >= 10 {
		switch p.dice.Sroll("1d2") {
		case 1:
			inc = false
		case 2:
			size++
		}
	}
	size = utils.BoundInt(size, 0, 15)
	p.physicalData = append(p.physicalData, ehex.New().Set(size))
	return nil
}

func (p *planet) determineAtmo() error {
	if len(p.physicalData) < 1 {
		return fmt.Errorf("size was not determined")
	}
	size := p.physicalData[0].Value()
	atmo := p.dice.Sroll("2d6") - 7 + size
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
	p.physicalData = append(p.physicalData, ehex.New().Set(atmo))
	return nil
}

func (p *planet) determineHydr() error {
	if len(p.physicalData) < 2 {
		return fmt.Errorf("atmo was not determined")
	}
	atmo := p.physicalData[1].Value()
	dm := atmo
	switch atmo {
	case 0, 1, 9, 10, 11, 12, 14:
		dm = -4
		if p.hz == -1 {
			dm = dm - 2
		}
		if p.hz == -2 {
			dm = dm - 6
		}
	case 13, 15:
		dm = -4
	}
	hydr := utils.BoundInt(p.dice.Sroll("2d6")-7+dm, 0, 10)
	p.physicalData = append(p.physicalData, ehex.New().Set(hydr))
	return nil
}
