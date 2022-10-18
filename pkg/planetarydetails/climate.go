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
	} {
		if err != nil {
			return err
		}
	}
	pd.Biology()
	return nil
}

func (pd *PlanetaryDetails) setAlbedo() error {
	atmo := pd.uwpData.Atmo()
	hydr := pd.uwpData.Hydr()
	albRoll := pd.dice.Roll("2d10").Sum()
	albedoStep := 0.01
	albedoAdd := 0.0
	switch atmo {
	case 4, 5, 6, 7, 8, 9:
		switch hydr {
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
		switch hydr {
		case 0, 1, 2, 3, 4, 5:
			albedoAdd = 0.03
		case 6, 7, 8, 9, 10:
			albedoAdd = 0.45
		default:
			return fmt.Errorf("unexpected HydroCode")
		}
	}
	pd.albedo = albedoStep*float64(albRoll) + albedoAdd
	if pd.albedo > 1 {
		return fmt.Errorf("albedo can not be more than 1.0")
	}
	if pd.albedo < 0 {
		return fmt.Errorf("albedo can not be less than 0.0")
	}
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
	pd.averageSurfaceTemperature = utils.RoundFloat64(b*(math.Pow((1.0-a), 1/4.0))*(1+e), 1)
	if pd.averageSurfaceTemperature < 0 {
		return fmt.Errorf("cannot have negative temp")
	}
	return nil
}

func (pd *PlanetaryDetails) Biology() {
	dm := 0
	if !strings.Contains(pd.atmoComposition, "Oxygen") {
		dm -= 20
	}
	switch pd.uwpData.Atmo() {
	case 0, 1, 2, 3:
		dm -= 50
	case 4, 5:
		dm -= 5
	case 6, 7, 8, 9:
		dm += 10
	}
	switch pd.uwpData.Hydr() {
	case 1, 2:
		dm -= 15
	case 3, 4, 10:
		dm -= 5
	case 5, 6, 7, 8, 9:
		dm += 5
	}
	if pd.averageSurfaceTemperature < 0+273.25 {
		dm -= 10
	}
	if pd.averageSurfaceTemperature > 40+273.25 {
		dm -= 10
	}
	if pd.orbit > pd.local.HabitabilityHi() {
		dm -= 20
	}
	if pd.orbit < pd.local.HabitabilityLow() {
		dm -= 20
	}
	switch pd.primary.Class() {
	case "O", "B", "A", "M":
		dm -= 10
	case "G":
		dm += 10
	}
	bRoll := pd.dice.Roll("2d10").DM(dm).Sum()
	fmt.Printf("Biology Roll: %v (%v - %v)\n", bRoll, dm+2, dm+20)
	fmt.Printf("Avr. Temp: %v K\n", pd.averageSurfaceTemperature)
	fmt.Printf("Avr. Temp: %v C\n", pd.averageSurfaceTemperature-273.25)
}
