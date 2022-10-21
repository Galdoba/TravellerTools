package planetarydetails

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

const (
	unatendedInt    = -999999
	unatendedFlt    = -999999.9
	unatendedStr    = "unatendedStr"
	CoreType_Molten = "Molten"
	CoreType_Rocky  = "Rocky"
	CoreType_Icy    = "Icy"
)

type Star interface {
	Class() string
	InnerLimit() float64
	HabitabilityLow() float64
	HabitabilityHi() float64
	SnowLine() float64
	OuterLimit() float64
	Mass() float64
	Luminocity() float64
}

type PlanetData interface {
	Orbit() float64
	SizeCode() string
	AtmoCode() string
	HydrCode() string
	Eccentricity() float64
	Habzone() string
	IsMW() bool
}

type PlanetaryDetails struct {
	dice *dice.Dicepool
	//uwpData uwp.UWP
	planet       PlanetData
	primary      Star
	orbit        float64
	size         int
	atmo         int
	hydr         int
	eccentricity float64
	habzone      string
	mw           bool
	////SIZE RELATED
	diameter       int //km
	density        float64
	coreType       string
	mass           float64
	surfaceGravity float64
	orbitalPeriod  float64 //years
	rotationPeriod float64 //hours
	tidalyLocked   bool
	axialTilt      int
	////ATMO RELATED
	taint           string
	atmoComposition string
	pressureCode    int
	pressure        float64
	////Climate
	albedo                    float64 //параметр поглощения радиации: 0 - вся радиация поглощается, 1 - вся радиация отражается
	greenhouseEffect          float64
	averageSurfaceTemperature float64
}

func NewPlanetaryDetails(dp *dice.Dicepool, planet PlanetData,
	PrimaryStar Star) PlanetaryDetails {
	pd := PlanetaryDetails{}
	pd.dice = dp

	//pd.uwpData = uwpData
	pd.size = ehex.New().Set(planet.SizeCode()).Value()
	pd.atmo = ehex.New().Set(planet.AtmoCode()).Value()
	pd.hydr = ehex.New().Set(planet.HydrCode()).Value()
	pd.primary = PrimaryStar
	pd.diameter = unatendedInt
	pd.density = unatendedFlt
	pd.orbit = planet.Orbit()
	pd.habzone = planet.Habzone()
	pd.eccentricity = planet.Eccentricity()
	pd.mw = planet.IsMW()
	pd.mass = unatendedFlt
	pd.surfaceGravity = unatendedFlt
	pd.orbitalPeriod = unatendedFlt
	pd.rotationPeriod = unatendedFlt
	pd.axialTilt = unatendedInt
	pd.taint = unatendedStr
	pd.coreType = unatendedStr
	pd.atmoComposition = unatendedStr
	pd.pressureCode = unatendedInt
	pd.pressure = unatendedFlt
	pd.greenhouseEffect = unatendedFlt
	pd.averageSurfaceTemperature = unatendedFlt
	pd.defineSizeRelatedDetails()
	pd.defineAtmosphereRelatedDetails()
	pd.defineClimate()
	return pd
}

func (pd *PlanetaryDetails) SizeRelatedString() string {
	str := ""
	str += fmt.Sprintf("Diameter: %v km\n", pd.diameter)
	str += fmt.Sprintf("Density : %v ED\n", pd.density)
	str += fmt.Sprintf("Core    : %v\n", pd.coreType)
	str += fmt.Sprintf("Mass    : %v EM\n", pd.mass)
	str += fmt.Sprintf("Gravity : %vg\n", pd.surfaceGravity)
	str += fmt.Sprintf("Orbital Period : %v years\n", pd.orbitalPeriod)
	str += fmt.Sprintf("Rotational Period : %v hours\n", pd.rotationPeriod)
	str += fmt.Sprintf("Axial Tilt : %v\n", pd.axialTilt)
	return str
}

func (pd *PlanetaryDetails) AtmoRelatedString() string {
	str := "Atmosphere\n"
	str += fmt.Sprintf("Composition: %v\n", pd.atmoComposition)
	str += fmt.Sprintf("Pressure   : %v\n", pd.pressure)
	str += fmt.Sprintf("Taint      : %v\n", pd.taint)
	return str
}
