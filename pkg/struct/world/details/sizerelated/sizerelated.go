package sizerelated

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/devtools/errmaker"
)

type sizeDetails struct {
	complete            bool
	diameter            int
	dessitytype         string
	dessity             float64
	mass                float64
	gravity             float64
	orbitalPeriod       float64
	rotationPeriod      float64
	axialTilt           int
	seismicStressFactor int
}

func New() *sizeDetails {
	sd := sizeDetails{}
	return &sd
}

/*
Planet Type:
    Hospitable -    A
    Planetoid -     B
    Iceworld -      C
    RadWorld -      D
    Inferno -       E
    BigWorld -      F
    Worldlet -      G
    Inner World -   H
    Stormworld -    J
   	SSG -			K
	LGG - 			L
	IG -			M

*/

func (sd *sizeDetails) GenerateDetails(dice *dice.Dicepool, prfl profile.Profile) error {
	worldtype := prfl.Data(profile.KEY_WORLDTYPE)
	if worldtype == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_WORLDTYPE)
	}
	size := prfl.Data(profile.KEY_SIZE)
	if size == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_SIZE)
	}
	atmo := prfl.Data(profile.KEY_ATMO)
	if atmo == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_ATMO)
	}
	hzVar := prfl.Data(profile.KEY_HABITABLE_ZONE_VAR)
	if hzVar == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_HABITABLE_ZONE_VAR)
	}
	//////////////////
	worldtypeCode := worldtype.Code()
	switch worldtypeCode {
	default:
		return fmt.Errorf("worldtype code '%v' unknown", worldtypeCode)
	case "B":
		return fmt.Errorf("worldtype code '%v' (planetoid) unimplemented", worldtypeCode)
	case "K", "L", "M":
		return fmt.Errorf("worldtype code '%v' (gigants) unimplemented", worldtypeCode)
	case "A", "C", "D", "E", "F", "G", "H", "J":
		for _, err := range []error{
			sd.rollDiameter(dice, size),
			sd.rollDesityType(dice, size, atmo, hzVar),
			sd.rollDensity(dice),
		} {
			if err != nil {
				return errmaker.ErrorFrom(err)
			}
		}
		return nil
	}

}

var ErrNoData = fmt.Errorf("profile have no data")

func (sd *sizeDetails) rollDiameter(dice *dice.Dicepool, size ehex.Ehex) error {
	r1 := dice.Sroll("2d6-7")
	r2 := dice.Sroll("2d6-7")
	r3 := dice.Sroll("2d6-7")
	sd.diameter = size.Value()*1000 + (r1 * 100) + (r2 * 10) + (r3)
	for sd.diameter <= 0 {
		sd.diameter += 100
	}
	sd.diameter = (sd.diameter * 16) / 10
	return nil
}

const (
	DENSITY_HEAVY_CORE  = "Heavy Core"
	DENSITY_MOLTEN_CORE = "Molten Core"
	DENSITY_ROCKY_BODY  = "Rocky Body"
	DENSITY_ICY_BODY    = "Icy Body"
)

func (sd *sizeDetails) rollDesityType(dice *dice.Dicepool, size, atmo, hzVar ehex.Ehex) error {
	dm := 0
	s := size.Value()
	a := atmo.Value()
	o := hzVar.Value()
	switch {
	case s <= 4:
		dm += 1
	case s >= 6:
		dm -= 2
	}
	switch {
	case a <= 3:
		dm += 1
	case a >= 6:
		dm -= 2
	}
	if o > 11 {
		dm += 6
	}
	r := dice.Sroll("2d6") + dm
	switch {
	case r <= -1:
		sd.dessitytype = DENSITY_HEAVY_CORE
	case r >= 2 && r <= 10:
		sd.dessitytype = DENSITY_MOLTEN_CORE
	case r >= 11 && r <= 14:
		sd.dessitytype = DENSITY_ROCKY_BODY
	case r >= 15:
		sd.dessitytype = DENSITY_ICY_BODY
	}
	return nil
}

func (sd *sizeDetails) rollDensity(dice *dice.Dicepool) error {
	r := dice.Sroll("3d6")
	densitySl := []float64{}
	switch sd.dessitytype {
	default:
		return errmaker.ErrorFrom(fmt.Errorf("sd.dessitytype invalid"), sd.dessitytype)
	case DENSITY_HEAVY_CORE:
		densitySl = []float64{1.10, 1.15, 1.20, 1.25, 1.30, 1.35, 1.40, 1.45, 1.50, 1.55, 1.60, 1.70, 1.80, 1.90, 2.00, 2.25}
	case DENSITY_MOLTEN_CORE:
		densitySl = []float64{0.82, 0.84, 0.86, 0.88, 0.90, 0.92, 0.94, 0.96, 0.98, 1.00, 1.02, 1.04, 1.06, 1.08, 1.10, 1.12}
	case DENSITY_ROCKY_BODY:
		densitySl = []float64{0.50, 0.52, 0.54, 0.56, 0.58, 0.60, 0.62, 0.64, 0.66, 0.68, 0.70, 0.72, 0.74, 0.76, 0.78, 0.80}
	case DENSITY_ICY_BODY:
		densitySl = []float64{0.18, 0.20, 0.22, 0.24, 0.26, 0.28, 0.30, 0.32, 0.34, 0.36, 0.38, 0.40, 0.42, 0.44, 0.46, 0.48}
	}
	sd.dessity = densitySl[r-3]
	return nil
}

type SizeRelatedDetails interface {
}
