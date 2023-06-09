package sizerelated

import (
	"fmt"
	"math"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/orbit"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/devtools/errmaker"
	"github.com/Galdoba/utils"
)

type SizeDetails struct {
	Star                         string
	Complete                     bool
	Diameter                     int
	Dessitytype                  string
	Dessity                      float64
	Mass                         float64
	Gravity                      float64
	OrbitalDistance              float64
	OrbitalPeriod                float64
	RotationPeriod               float64
	AxialTilt                    int
	OrbitalEccentricity          float64
	SeismicStressFactor          int
	IsBelt                       bool
	PredominatePlanetoidDiameter string
	MaximumPlanetoidDiameter     string
	PredominateBeltZone          string
	NZone                        int
	MZone                        int
	CZone                        int
	BeltOrbitWidth               float64
}

func New() *SizeDetails {
	sd := SizeDetails{}
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

func (sd *SizeDetails) GenerateDetails(dice *dice.Dicepool, prfl profile.Profile, star star.StarBody) error {
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
	planetOrbit := prfl.Data(profile.KEY_PLANETARY_ORBIT)
	if planetOrbit == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_PLANETARY_ORBIT)
	}
	sateliteOrbit := prfl.Data(profile.KEY_SATELITE_ORBIT)
	if sateliteOrbit == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_SATELITE_ORBIT)
	}

	if star == nil {
		return errmaker.ErrorFrom(fmt.Errorf("star class not provided"))
	}
	//////////////////
	sd.Star = star.Class()
	worldtypeCode := worldtype.Code()
	switch worldtypeCode {
	default:
		return fmt.Errorf("worldtype code '%v' unknown", worldtypeCode)
	case "B":
		sd.IsBelt = true
		for _, err := range []error{
			// sd.rollDiameter(dice, size),
			// sd.rollDesityType(dice, size, atmo, hzVar),
			// sd.rollDensity(dice),
			// sd.calculateWorldMass(size),
			// sd.calculateWorldGravity(size),
			// sd.rollOrbitalDistance(dice, planetOrbit),
			// sd.calculatePlanetaryOrbitalPeriod(star, sateliteOrbit),
			// sd.rollPlanetaryRotationPeriod(dice, star, sateliteOrbit),
			sd.rollPredominateDiameter(dice),
			sd.rollBeltZones(dice, hzVar),
			sd.rollBeltWidth(dice, planetOrbit),
		} {
			if err != nil {
				return errmaker.ErrorFrom(err)
			}
		}
		return nil
	case "K", "L", "M":
		return fmt.Errorf("worldtype code '%v' (gigants) unimplemented", worldtypeCode)
	case "A", "C", "D", "E", "F", "G", "H", "J":
		for _, err := range []error{
			sd.rollDiameter(dice, size),
			sd.rollDesityType(dice, size, atmo, hzVar),
			sd.rollDensity(dice),
			sd.calculateWorldMass(size),
			sd.calculateWorldGravity(size),
			sd.rollOrbitalDistance(dice, planetOrbit),
			sd.calculatePlanetaryOrbitalPeriod(star, sateliteOrbit),
			sd.rollPlanetaryRotationPeriod(dice, star, sateliteOrbit),
			sd.rollStressFactor(dice, star),
		} {
			if err != nil {
				return errmaker.ErrorFrom(err)
			}
		}
		sd.AxialTilt = rollAxialTilt(dice)
		sd.OrbitalEccentricity = rollOrbitalEccentricity(dice)
		return nil
	}

}

var ErrNoData = fmt.Errorf("profile have no data")

func (sd *SizeDetails) rollDiameter(dice *dice.Dicepool, size ehex.Ehex) error {
	r1 := dice.Sroll("2d6-7")
	r2 := dice.Sroll("2d6-7")
	r3 := dice.Sroll("2d6-7")
	sd.Diameter = size.Value()*1000 + (r1 * 100) + (r2 * 10) + (r3)
	for sd.Diameter <= 0 {
		sd.Diameter += 100
	}
	sd.Diameter = (sd.Diameter * 16) / 10
	return nil
}

const (
	DENSITY_HEAVY_CORE  = "Heavy Core"
	DENSITY_MOLTEN_CORE = "Molten Core"
	DENSITY_ROCKY_BODY  = "Rocky Body"
	DENSITY_ICY_BODY    = "Icy Body"
)

func (sd *SizeDetails) rollDesityType(dice *dice.Dicepool, size, atmo, hzVar ehex.Ehex) error {
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
	case r <= 1:
		sd.Dessitytype = DENSITY_HEAVY_CORE
	case r >= 2 && r <= 10:
		sd.Dessitytype = DENSITY_MOLTEN_CORE
	case r >= 11 && r <= 14:
		sd.Dessitytype = DENSITY_ROCKY_BODY
	case r >= 15:
		sd.Dessitytype = DENSITY_ICY_BODY
	}
	return nil
}

func (sd *SizeDetails) rollDensity(dice *dice.Dicepool) error {
	r := dice.Sroll("3d6")
	densitySl := []float64{}
	switch sd.Dessitytype {
	default:
		return errmaker.ErrorFrom(fmt.Errorf("sd.Dessitytype invalid"), sd.Dessitytype)
	case DENSITY_HEAVY_CORE:
		densitySl = []float64{1.10, 1.15, 1.20, 1.25, 1.30, 1.35, 1.40, 1.45, 1.50, 1.55, 1.60, 1.70, 1.80, 1.90, 2.00, 2.25}
	case DENSITY_MOLTEN_CORE:
		densitySl = []float64{0.82, 0.84, 0.86, 0.88, 0.90, 0.92, 0.94, 0.96, 0.98, 1.00, 1.02, 1.04, 1.06, 1.08, 1.10, 1.12}
	case DENSITY_ROCKY_BODY:
		densitySl = []float64{0.50, 0.52, 0.54, 0.56, 0.58, 0.60, 0.62, 0.64, 0.66, 0.68, 0.70, 0.72, 0.74, 0.76, 0.78, 0.80}
	case DENSITY_ICY_BODY:
		densitySl = []float64{0.18, 0.20, 0.22, 0.24, 0.26, 0.28, 0.30, 0.32, 0.34, 0.36, 0.38, 0.40, 0.42, 0.44, 0.46, 0.48}
	}
	sd.Dessity = densitySl[r-3]
	return nil
}

func (sd *SizeDetails) calculateWorldMass(size ehex.Ehex) error {
	if sd.Dessity <= 0 {
		return fmt.Errorf("density is <= 0")
	}
	r := float64(size.Value())
	m := sd.Dessity * (math.Pow((r / 8.0), 3))
	sd.Mass = m
	sd.Mass = utils.RoundFloat64(sd.Mass, 3)
	return nil
}

func (sd *SizeDetails) calculateWorldGravity(size ehex.Ehex) error {
	if sd.Mass <= 0 {
		return fmt.Errorf("Mass is <= 0")
	}
	r := float64(size.Value())
	m := sd.Mass * (64.0 / (math.Pow((r), 2)))
	sd.Gravity = m
	sd.Gravity = utils.RoundFloat64(sd.Gravity, 2)
	return nil
}

func (sd *SizeDetails) rollOrbitalDistance(dice *dice.Dicepool, planetOrbit ehex.Ehex) error {
	orbit := orbit.NewPlanetOrbit(dice, planetOrbit.Value())
	sd.OrbitalDistance = orbit.Distance()
	return nil
}

func (sd *SizeDetails) calculatePlanetaryOrbitalPeriod(star star.StarBody, sateliteOrbit ehex.Ehex) error {
	if sateliteOrbit.Code() != "*" {
		return nil
	}
	d := sd.OrbitalDistance
	m := star.Mass()
	p := (math.Sqrt(math.Pow(d, 3) / m)) * 365.25
	p = utils.RoundFloat64(p, 1)
	sd.OrbitalPeriod = p
	return nil
}

func (sd *SizeDetails) rollPlanetaryRotationPeriod(dice *dice.Dicepool, star star.StarBody, sateliteOrbit ehex.Ehex) error {
	if sateliteOrbit.Code() != "*" {
		return nil
	}
	d := sd.OrbitalDistance
	m := star.Mass()
	w := float64(dice.Sroll("4d10+5"))
	p := w + (m / d)
	if p > 45 {
		r := dice.Sroll("2d6")
		switch r {
		case 2:
		case 3:
			p = float64(dice.Sroll("1d10") * 5)
		case 4:
			p = float64(dice.Sroll("1d10") * 10)
		case 5:
			p = float64(dice.Sroll("1d10") * 20)
		case 6:
			p = float64(dice.Sroll("1d10") * 30)
		case 7:
			p = float64(dice.Sroll("1d10")*24) + float64(dice.Sroll("1d24")-12) + float64(dice.Sroll("1d10")/10)
		case 8:
			p = float64(dice.Sroll("1d10")*24*5) + float64(dice.Sroll("1d24")-12) + float64(dice.Sroll("1d10")/10)
		case 9:
			p = float64(dice.Sroll("1d10")*24*10) + float64(dice.Sroll("1d24")-12) + float64(dice.Sroll("1d10")/10)
		case 10:
			p = float64(dice.Sroll("1d10")*24*20) + float64(dice.Sroll("1d24")-12) + float64(dice.Sroll("1d10")/10)
		case 11:
			p = float64(dice.Sroll("1d10")*24*30) + float64(dice.Sroll("1d24")-12) + float64(dice.Sroll("1d10")/10)
		case 12:
			p = float64(dice.Sroll("1d10")*24*20) + float64(dice.Sroll("1d24")-12) + float64(dice.Sroll("1d10")/10)
		}
		if p/24 >= sd.OrbitalPeriod {
			p = sd.OrbitalPeriod * 24
		}
	}
	p = utils.RoundFloat64(p, 1)
	sd.RotationPeriod = p
	return nil
}

func rollAxialTilt(dice *dice.Dicepool) int {
	switch dice.Sroll("2d10") {
	case 2, 3, 4, 5:
		return 0 + dice.Sroll("1d10-1")
	case 6, 7, 8, 9:
		return 10 + dice.Sroll("1d10-1")
	case 10, 11, 12, 13:
		return 20 + dice.Sroll("1d10-1")
	case 14, 15, 16, 17:
		return 30 + dice.Sroll("1d10-1")
	case 18, 19:
		return 40 + dice.Sroll("1d10-1")
	case 20:
		switch dice.Sroll("1d6") {
		case 1, 2:
			return 50 + dice.Sroll("1d10-1")
		case 3:
			return 60 + dice.Sroll("1d10-1")
		case 4:
			return 70 + dice.Sroll("1d10-1")
		case 5:
			return 80 + dice.Sroll("1d10-1")
		case 6:
			return 90
		}
	}
	return -1
}

func rollOrbitalEccentricity(dice *dice.Dicepool) float64 {
	if dice.Sroll("1d6") < 4 {
		return 0
	}
	oe := 0.0
	switch dice.Sroll("2d10") {
	case 2:
		oe = 0.002
	case 3:
		oe = 0.003
	case 4:
		oe = 0.004
	case 5:
		oe = 0.005
	case 6:
		oe = 0.006
	case 7:
		oe = 0.007
	case 8:
		oe = 0.008
	case 9:
		oe = 0.009
	case 10:
		oe = 0.010
	case 11:
		oe = 0.020
	case 12:
		oe = 0.030
	case 13:
		oe = 0.040
	case 14:
		oe = 0.050
	case 15:
		oe = 0.070
	case 16:
		oe = 0.100
	case 17:
		oe = 0.125
	case 18:
		oe = 0.150
	case 19:
		oe = 0.200
	case 20:
		oe = 0.250
	}
	return oe
}

func (sd *SizeDetails) rollStressFactor(dice *dice.Dicepool, star star.StarBody) error {
	x := dice.Sroll("1d6-3")
	p := 0
	switch sd.Dessitytype {
	case DENSITY_HEAVY_CORE:
		p = dice.Sroll("1d6-2")
	case DENSITY_MOLTEN_CORE:
		p = dice.Sroll("1d6-3")
	}
	s := int(star.Luminocity() / sd.OrbitalDistance)
	sd.SeismicStressFactor = x + p + s
	if sd.SeismicStressFactor < 1 {
		sd.SeismicStressFactor = 1
	}
	return nil
}

func (sd *SizeDetails) rollPredominateDiameter(dice *dice.Dicepool) error {
	r1 := dice.Sroll("2d6") - 2
	r2 := dice.Sroll("1d6") - 1
	pbd := []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.3, 1, 5, 50, 500}
	mpd := []float64{0, 0, 1, 10, 100, 1000}
	bd := pbd[r1]
	bdm := mpd[r2]
	sd.PredominatePlanetoidDiameter = fmt.Sprintf("%vkm", int(bd))
	if bd < 0 {
		sd.PredominatePlanetoidDiameter = fmt.Sprintf("%vm", int(1000.0*bd))
	}
	sd.MaximumPlanetoidDiameter = fmt.Sprintf("%vkm", int(bdm))
	if bdm == 0 {
		sd.MaximumPlanetoidDiameter = ""
	}
	if bd < bdm {
		sd.PredominatePlanetoidDiameter, sd.MaximumPlanetoidDiameter = sd.MaximumPlanetoidDiameter, sd.PredominatePlanetoidDiameter
	}
	return nil
}

func (sd *SizeDetails) rollBeltZones(dice *dice.Dicepool, hzVar ehex.Ehex) error {
	dm := 0
	if hzVar.Value() < 9 {
		dm = -4
	}
	if hzVar.Value() > 11 {
		dm = 2
	}
	r1 := dice.Sroll("2d6") + dm
	switch r1 {
	default:
		return fmt.Errorf("invalid r1 = %v", r1)
	case -2, -1, 0, 1, 2, 3, 4:
		sd.PredominateBeltZone = "N"
	case 5, 6, 7, 8:
		sd.PredominateBeltZone = "M"
	case 9, 10, 11, 12, 13, 14:
		sd.PredominateBeltZone = "C"
	}
	r2 := dice.Sroll("2d6") - 2
	nAr := []int{}
	mAr := []int{}
	cAr := []int{}
	switch sd.PredominateBeltZone {
	case "N":
		nAr = []int{40, 40, 40, 40, 40, 50, 50, 50, 50, 60, 60}
		mAr = []int{30, 40, 40, 40, 40, 40, 40, 40, 30, 50, 40}
		cAr = []int{30, 20, 20, 20, 20, 10, 10, 10, 20, 10, 20}
	case "M":
		nAr = []int{20, 30, 20, 20, 30, 20, 10, 10, 10, 0, 0}
		mAr = []int{50, 50, 60, 60, 60, 70, 70, 80, 80, 80, 90}
		cAr = []int{30, 20, 20, 20, 10, 10, 20, 10, 10, 20, 10}
	case "C":
		nAr = []int{20, 20, 20, 10, 10, 10, 10, 10, 0, 0, 0}
		mAr = []int{30, 30, 30, 30, 30, 20, 20, 10, 10, 10, 20}
		cAr = []int{50, 50, 50, 60, 60, 70, 70, 80, 80, 80, 80}
	}
	sd.NZone = nAr[r2]
	sd.MZone = mAr[r2]
	sd.CZone = cAr[r2]
	sd.NZone += dice.Flux()
	sd.MZone += dice.Flux()
	if sd.NZone < 0 {
		sd.NZone = 0
	}
	sd.CZone = 100 - sd.MZone - sd.NZone
	return nil
}

func (sd *SizeDetails) rollBeltWidth(dice *dice.Dicepool, planetOrbit ehex.Ehex) error {
	dm := 0
	switch planetOrbit.Value() {
	case 0, 1, 2, 3, 4:
		dm = -3
	case 5, 6, 7, 8:
		dm = -1
	case 9, 10, 11, 12:
		dm = 1
	default:
		dm = 2
	}
	r1 := dice.Sroll("2d6") + dm - 2
	if r1 < 0 {
		r1 = 0
	}
	if r1 > 10 {
		r1 = 10
	}
	sd.BeltOrbitWidth = []float64{0.01, 0.05, 0.1, 0.1, 0.5, 0.5, 1.0, 1.5, 2.0, 5.0, 10.0}[r1]
	return nil
}

type SizeRelatedDetails interface {
}
