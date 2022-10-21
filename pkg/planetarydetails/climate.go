package planetarydetails

import (
	"fmt"
	"math"
	"strings"

	"github.com/Galdoba/utils"
)

func (pd *PlanetaryDetails) defineClimate() error {
	for _, err := range []error{
		pd.setAlbedo(),
		pd.setGreenhouseEffect(),
		pd.setAveragetemperature(),
		pd.setOtherTemperatures(),
		pd.Biology(),
		pd.setHydrospere(),
	} {
		if err != nil {
			return err
		}
	}

	return nil
}

func (pd *PlanetaryDetails) setAlbedo() error {
	albRoll := pd.dice.Roll("2d10").Sum()
	albedoStep := 0.01
	albedoAdd := 0.0
	switch pd.atmo {
	case 4, 5, 6, 7, 8, 9:
		switch pd.hydr {
		case 0, 1, 2:
			albedoAdd = 0.05
		case 3, 4, 5:
			albedoAdd = 0.11
		case 6, 7, 8:
			albedoAdd = 0.21
		case 9, 10:
			albedoAdd = 0.27
		default:
			return fmt.Errorf("unexpected HydroCode")
		}
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15:
		switch pd.hydr {
		case 0, 1, 2, 3, 4, 5:
			albedoAdd = 0.03
		case 6, 7, 8, 9, 10:
			albedoAdd = 0.45
		default:
			return fmt.Errorf("unexpected HydroCode")
		}
	}
	pd.albedo = albedoStep*float64(albRoll) + albedoAdd
	//pd.albedo = 0.03 // при увеличении альбедо температура падает
	if pd.albedo > 1 {
		return fmt.Errorf("albedo can not be more than 1.0")
	}
	if pd.albedo < 0 {
		return fmt.Errorf("albedo can not be less than 0.0")
	}
	return nil
}

func (pd *PlanetaryDetails) setHydrospere() error {

	h := float64((pd.hydr*100)+pd.dice.Flux()*10+pd.dice.Flux()) / 10
	if h < 0 {
		h = 0
	}
	if h > 100 {
		h = 100
	}
	pd.hydrCover = h
	return nil
}

func (pd *PlanetaryDetails) setGreenhouseEffect() error {
	p := pd.pressure
	g := pd.surfaceGravity
	pd.greenhouseEffect = utils.RoundFloat64((p/6.0)*(p/g), 3)
	return nil
}

func (pd *PlanetaryDetails) setAveragetemperature() error {
	lum := pd.primary.Luminocity()
	dis := pd.orbit
	b := (271 * math.Pow(lum, 1/4.0)) / math.Pow(dis, 1/2.0)
	a := pd.albedo
	e := pd.greenhouseEffect
	switch pd.mw {
	case true:
		pd.RollForMW()
	case false:
		//fmt.Println(b, a, e, dis, lum)
		pd.averageSurfaceTemperature = utils.RoundFloat64((b*(math.Pow((1.0-a), 1/4.0))*(1+e))+273.15, 1)
		if pd.averageSurfaceTemperature < -273 {
			pd.averageSurfaceTemperature = -273
		}

	}
	if pd.averageSurfaceTemperature < 0 {
		return fmt.Errorf("cannot have negative temp")
	}
	return nil
}

func (pd *PlanetaryDetails) setOtherTemperatures() error {
	if pd.atmo == 0 {
		return nil
	}
	s := float64(pd.size)
	atM := 0.0
	switch {
	case pd.axialTilt >= 90:
		atM = 1.00
	case pd.axialTilt >= 85:
		atM = 0.90
	case pd.axialTilt >= 80:
		atM = 0.85
	case pd.axialTilt >= 75:
		atM = 0.80
	case pd.axialTilt >= 70:
		atM = 0.75
	case pd.axialTilt >= 65:
		atM = 0.65
	case pd.axialTilt >= 60:
		atM = 0.55
	case pd.axialTilt >= 55:
		atM = 0.50
	case pd.axialTilt >= 50:
		atM = 0.45
	case pd.axialTilt >= 45:
		atM = 0.40
	case pd.axialTilt >= 40:
		atM = 0.35
	case pd.axialTilt >= 35:
		atM = 0.25
	}
	pd.averageTemperatureEquatorSummer = (atM * (float64(pd.axialTilt) * 0.6)) + (3 * s) + pd.averageSurfaceTemperature
	pd.averageTemperatureEquatorWinter = pd.averageSurfaceTemperature + (3 * s) - (atM * (float64(pd.axialTilt) * 0.6))
	atM2 := 0.0
	switch {
	case pd.axialTilt == 5:
		atM2 = 0.80
	case pd.axialTilt == 4:
		atM2 = 0.70
	case pd.axialTilt == 3:
		atM2 = 0.50
	case pd.axialTilt == 2:
		atM2 = 0.40
	case pd.axialTilt == 1:
		atM2 = 0.30
	case pd.axialTilt == 0:
		atM2 = 0.20
	}
	pd.averageTemperaturePolarSummer = (atM2 * (float64(pd.axialTilt) * 0.6)) - (7 * s) + pd.averageSurfaceTemperature
	pd.averageTemperaturePolarWinter = pd.averageSurfaceTemperature - (7 * s) - (atM2 * (float64(pd.axialTilt) * 0.6))
	return nil
}

func (pd *PlanetaryDetails) RollForMW() {
	add := pd.dice.Roll("10d10").Sum()
	pd.averageSurfaceTemperature = 228 + float64(add) - 273
}

func (pd *PlanetaryDetails) Biology() error {
	dm := 0
	if !strings.Contains(pd.atmoComposition, "Oxygen") {
		dm -= 20
	}
	switch pd.atmo {
	case 0, 1:
		pd.lifeIndex = 0
		return nil
	case 2, 3:
		dm -= 10
	case 4, 5:
		dm -= 5
	case 6, 7, 8, 9:
		dm += 10
	}
	switch pd.hydr {
	case 0:
		pd.lifeIndex = 0
		return nil
	case 1, 2:
		dm -= 15
	case 3, 4, 10:
		dm -= 5
	case 5, 6, 7, 8, 9:
		dm += 5
	}
	if pd.averageSurfaceTemperature < 0 || pd.averageSurfaceTemperature > 40 {
		dm -= 10
	}
	if pd.habzone != "Habitable" {
		dm -= 20
	}
	switch pd.primary.Class() {

	case "O", "B", "A", "M":
		dm -= 10
	case "G":
		dm += 10
	}
	if pd.mw == true {
		dm += 10
	}
	pd.lifeIndex = pd.dice.Roll("2d10").DM(dm).Sum()
	if pd.lifeIndex < 0 {
		pd.lifeIndex = 0
	}
	switch pd.lifeIndex {
	default:
		pd.lifeIndex = 10
	case 0:
		pd.lifeIndex = 0
	case 1, 2, 3:
		pd.lifeIndex = 1
	case 4, 5:
		pd.lifeIndex = 2
	case 6, 7, 8:
		pd.lifeIndex = 3
	case 9, 10:
		pd.lifeIndex = 4
	case 11, 12:
		pd.lifeIndex = 5
	case 13, 14:
		pd.lifeIndex = 6
	case 15, 16:
		pd.lifeIndex = 7
	case 17, 18:
		pd.lifeIndex = 8
	case 19, 20:
		pd.lifeIndex = 9

	}
	if pd.averageSurfaceTemperature < -273.15 {
		return fmt.Errorf("Can have temperature below absolute zero")
	}
	return nil
}
