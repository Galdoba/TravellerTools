package size

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/devtools/errmaker"
)

type sizeDetails struct {
	worldTypeCode       string
	hzVar               int
	sizeVal             int
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

func New(worldType, hzVar, size ehex.Ehex) *sizeDetails {
	sd := sizeDetails{}
	sd.worldTypeCode = worldType.Code()
	sd.sizeVal = size.Value()
	sd.hzVar = hzVar.Value()
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
		return ErrNoDataByKey(profile.KEY_WORLDTYPE)
	}
	worldtypeCode := worldtype.Code()
	switch worldtypeCode {
	default:
		return fmt.Errorf("worldtype code '%v' unknown", worldtypeCode)
	case "B":
		return fmt.Errorf("worldtype code '%v' (planetoid) unimplemented", worldtypeCode)
	case "K", "L", "M":
		return fmt.Errorf("worldtype code '%v' (gigants) unimplemented", worldtypeCode)
	case "A", "C", "D", "E", "F", "G", "H", "J":
		size := prfl.Data(profile.KEY_SIZE)
		if size == nil {
			return ErrNoDataByKey(profile.KEY_SIZE)
		}
		atmo := prfl.Data(profile.KEY_ATMO)
		if atmo == nil {
			return ErrNoDataByKey(profile.KEY_ATMO)
		}
		for _, err := range []error{
			sd.rollDiameter(dice),
		} {
			if err != nil {
				return errmaker.ErrorFrom(err)
			}
		}
		return nil
	}

}

func ErrNoDataByKey(key string) error {
	return fmt.Errorf("profile have no data by key '%v'", key)
}

func (sd *sizeDetails) rollDiameter(dice *dice.Dicepool) error {
	r1 := dice.Sroll("2d6-7")
	r2 := dice.Sroll("2d6-7")
	r3 := dice.Sroll("2d6-7")
	sd.diameter = sd.sizeVal*1000 + (r1 * 100) + (r2 * 10) + (r3)
	for sd.diameter <= 0 {
		sd.diameter += 100
	}
	return nil
}

func (sd *sizeDetails) rollDesityType(dice *dice.Dicepool, size, atmo ehex.Ehex) error {
	dm := 0

}

type SizeRelatedDetails interface {
}
