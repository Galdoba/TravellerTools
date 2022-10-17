package planetarydetails

import (
	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/pkg/systemgeneration"
)

const (
	unatendedInt    = -999999
	unatendedFlt    = -999999.9
	unatendedStr    = "unatendedStr"
	CoreType_Molten = "Molten"
	CoreType_Rocky  = "Rocky"
	CoreType_Icy    = "Icy"
)

type PlanetaryDetails struct {
	dice    *dice.Dicepool
	uwpData uwp.UWP
	primary systemgeneration.Star
	local   systemgeneration.Star
	orbit   float64
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
	taint string
}

func NewPlanetaryDetails(dp *dice.Dicepool, uwpData uwp.UWP,
	PrimaryStar systemgeneration.Star,
	LocalStar systemgeneration.Star,
	orbit float64) PlanetaryDetails {
	pd := PlanetaryDetails{}
	pd.dice = dp
	pd.uwpData = uwpData
	pd.primary = PrimaryStar
	pd.local = LocalStar
	pd.diameter = unatendedInt
	pd.density = unatendedFlt
	pd.orbit = orbit
	pd.mass = unatendedFlt
	pd.surfaceGravity = unatendedFlt
	pd.orbitalPeriod = unatendedFlt
	pd.rotationPeriod = unatendedFlt
	pd.axialTilt = unatendedInt
	pd.taint = unatendedStr
	return pd
}
