package planetarydetails

import (
	"fmt"
	"math"

	"github.com/Galdoba/utils"
)

func (pd *PlanetaryDetails) defineSizeRelatedDetails() error {
	for _, err := range []error{
		pd.setDiameter(),
		pd.setDensity(),
		pd.setMass(),
		pd.setGravity(),
		pd.setOrbiatlPeriod(),
		pd.setRotationPeriod(),
		pd.setAxialTilt(),
	} {
		if err != nil {
			return err
		}
	}
	return nil
}

func (pd *PlanetaryDetails) setDiameter() error {
	add := pd.dice.Roll("1d10").DM(-1).Sum()*100 + pd.dice.Roll("1d10").DM(-1).Sum()*10 + pd.dice.Roll("1d10").DM(-1).Sum()*1
	sizeCode := pd.uwpData.Size()
	switch sizeCode {
	case 0:
		pd.diameter = 0 + add/2
		return nil
	default:
		if sizeCode > 0 && sizeCode < 25 {
			pd.diameter = (sizeCode * 1000) - 500 + add
			pd.diameter = int(float64(pd.diameter) * 1.6)
			return nil
		}
		return fmt.Errorf("sizeCode invalid (%v)", sizeCode)
	}
}

func (pd *PlanetaryDetails) setDensity() error {
	ctDM := 0
	if pd.uwpData.Size() <= 5 && pd.orbit > pd.local.InnerLimit() && pd.orbit < pd.local.SnowLine() {
		ctDM += 3
	}
	if pd.uwpData.Size() >= 6 && pd.orbit > pd.local.InnerLimit() && pd.orbit < pd.local.HabitabilityLow() {
		ctDM -= 2
	}
	if pd.uwpData.Size() >= 6 && pd.orbit >= pd.local.HabitabilityLow() && pd.orbit <= pd.local.HabitabilityHi() {
		ctDM -= 4
	}
	if pd.uwpData.Size() <= 5 && pd.orbit > pd.local.SnowLine() && pd.orbit <= pd.local.OuterLimit() {
		ctDM += 9
	}
	if pd.uwpData.Size() >= 6 && pd.orbit > pd.local.SnowLine() && pd.orbit <= pd.local.OuterLimit() {
		ctDM += 3
	}
	ctRoll := pd.dice.Roll("2d6").DM(ctDM).Sum()
	if ctRoll <= 6 {
		pd.coreType = CoreType_Molten
	}
	if ctRoll >= 7 && ctRoll <= 15 {
		pd.coreType = CoreType_Rocky
	}
	if ctRoll >= 16 {
		pd.coreType = CoreType_Icy
	}
	denRoll := pd.dice.Roll("2d10").Sum()
	moltenDencity := []float64{0.86, 0.88, 0.90, 0.92, 0.94, 0.96, 0.98, 1, 1, 1, 1.02, 1.04, 1.06, 1.08, 1.10, 1.12, 1.14, 1.16, 1.18}
	rockyDencity := []float64{0.5, 0.52, 0.54, 0.56, 0.58, 0.60, 0.62, 0.64, 0.64, 0.66, 0.68, 0.7, 0.72, 0.74, 0.76, 0.78, 0.8, 0.82, 0.84}
	icyDencity := []float64{0.12, 0.14, 0.16, 0.18, 0.20, 0.22, 0.24, 0.26, 0.28, 0.30, 0.32, 0.34, 0.36, 0.38, 0.40, 0.42, 0.44, 0.46, 0.48}
	switch pd.coreType {
	case CoreType_Molten:
		pd.density = moltenDencity[denRoll-2]
	case CoreType_Rocky:
		pd.density = rockyDencity[denRoll-2]
	case CoreType_Icy:
		pd.density = icyDencity[denRoll-2]
	}

	return nil
}

func (pd *PlanetaryDetails) setMass() error {
	size := float64(pd.uwpData.Size())
	if size < 1 {
		size = 0.5
	}
	pd.mass = utils.RoundFloat64(pd.density*((size/8.0)*(size/8.0)*(size/8.0)), 3)
	return nil
}

func (pd *PlanetaryDetails) setGravity() error {
	size := float64(pd.uwpData.Size())
	if size < 1 {
		size = 0.5
	}
	m := pd.mass
	pd.surfaceGravity = utils.RoundFloat64(m*(64/(size*size)), 2)
	return nil
}

func (pd *PlanetaryDetails) setOrbiatlPeriod() error {
	d := pd.orbit
	m := pd.local.Mass()
	pd.orbitalPeriod = utils.RoundFloat64(math.Sqrt((d*d*d)/m), 3)
	return nil
}

func (pd *PlanetaryDetails) setRotationPeriod() error {
	w := float64(pd.dice.Roll("4d10").DM(5).Sum())
	m := pd.local.Mass()
	d := pd.orbit
	pd.rotationPeriod = utils.RoundFloat64(w+(m/d), 2)
	if pd.rotationPeriod > 45 {
		switch pd.dice.Roll("2d6").Sum() {
		case 2:
			return nil
		case 3:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 5)
		case 4:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 10)
		case 5:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 20)
		case 6:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 30)
		case 7:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 24)
		case 8:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 120)
		case 9:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 240)
		case 10:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 480)
		case 11:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * 720)
		case 12:
			pd.rotationPeriod = float64(pd.dice.Roll("1d10").Sum() * -480)
		}
		if pd.rotationPeriod > (pd.orbitalPeriod / 8766) {
			pd.rotationPeriod = utils.RoundFloat64(pd.orbitalPeriod/8766, 2)
			pd.tidalyLocked = true
		}
	}
	return nil
}

func (pd *PlanetaryDetails) setAxialTilt() error {
	r1 := pd.dice.Roll("2d10").Sum()
	r2 := pd.dice.Roll("1d6").Sum()
	dm := 0
	switch r1 {
	case 2, 3, 4, 5:
		dm = 0
	case 6, 7, 8, 9:
		dm = 10
	case 10, 11, 12, 13:
		dm = 20
	case 14, 15, 16, 17:
		dm = 30
	case 18, 19:
		dm = 40
	case 20:
		switch r2 {
		case 1, 2:
			dm = 50
		case 3:
			dm = 60
		case 4:
			dm = 70
		case 5:
			dm = 80
		case 6:
			pd.axialTilt = 90
			return nil
		}
	}
	pd.axialTilt = pd.dice.Roll("1d10").DM(-1).DM(dm).Sum()
	return nil
}
