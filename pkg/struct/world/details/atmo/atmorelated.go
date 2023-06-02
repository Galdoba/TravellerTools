package atmorelated

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world/details/sizerelated"
	"github.com/Galdoba/devtools/errmaker"
	"github.com/Galdoba/utils"
)

type AtmoDetails struct {
	Composition        string
	Taint              int
	Pressure           float64
	Albedo             float64
	GreenhouseEffect   float64
	AverageTemp        float64
	EquatorAverageTemp float64
	EquatorTempSummer  float64
	EquatorTempWinter  float64
	PolarAverageTemp   float64
	PolarTempSummer    float64
	PolarTempWinter    float64
	DaySideTemp        float64
	NightSideTemp      float64
}

func New() *AtmoDetails {
	ad := AtmoDetails{}
	return &ad
}

var ErrNoData = fmt.Errorf("profile have no data")

func (ad *AtmoDetails) GenerateDetails(dice *dice.Dicepool, prfl profile.Profile, sd *sizerelated.SizeDetails, star star.StarBody) error {
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
	hydr := prfl.Data(profile.KEY_HYDR)
	if hydr == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_HYDR)
	}
	hzVar := prfl.Data(profile.KEY_HABITABLE_ZONE_VAR)
	if hzVar == nil {
		return errmaker.ErrorFrom(ErrNoData, profile.KEY_HABITABLE_ZONE_VAR)
	}

	for _, err := range []error{
		//ad.rollAtmosphericComposition(dice, atmo),
		ad.rollAtmosphericComposition(dice, atmo),
		ad.rollAlbedo(dice, atmo, hydr),
		ad.calculateTemp(dice, sd, star, size, atmo, hzVar),
	} {
		if err != nil {
			return errmaker.ErrorFrom(err)
		}
	}
	return nil

}

func (ad *AtmoDetails) rollAtmosphericComposition(dice *dice.Dicepool, atmo ehex.Ehex) error {
	ad.Composition = "Nitrogen, oxygen, argon and carbone dioxide"
	switch atmo.Value() {
	case 0:
		ad.Composition = "None"
		return nil
	case 1:
		ad.Composition = "Trace"
		ad.Pressure = pressureTrace(dice)
		return nil
	case 2:
		ad.Taint = taint(dice)
		ad.Pressure = pressureVThin(dice)
	case 3:
		ad.Pressure = pressureVThin(dice)
	case 4:
		ad.Taint = taint(dice)
		ad.Pressure = pressureThin(dice)
	case 5:
		ad.Pressure = pressureThin(dice)
	case 6:
		ad.Pressure = pressureStandard(dice)
	case 7:
		ad.Taint = taint(dice)
		ad.Pressure = pressureStandard(dice)
	case 8:
		ad.Pressure = pressureDense(dice)
	case 9:
		ad.Taint = taint(dice)
		ad.Pressure = pressureDense(dice)
	case 10:
		ad.Composition = "Exotic"
		ad.Pressure = pressureDense(dice)
	case 11:
		ad.Composition, ad.Pressure = corrosiveAtmosphere(dice)
	case 12:
		ad.Composition, ad.Pressure = insidiousAtmosphere(dice)
	case 13:
		ad.Pressure = pressureDenseHigh(dice)
	case 14:
		ad.Pressure = pressureThinLow(dice)
	case 15:
		ad.Composition = "Unusual"
		ad.Pressure = pressureStandard(dice)
	}
	return nil
}

func pressureTrace(dice *dice.Dicepool) float64 {
	return []float64{0.001, 0.002, 0.005, 0.007, 0.01, 0.02, 0.03, 0.05, 0.07, 0.08, 0.09}[dice.Sroll("2d6-2")]
}

func pressureVThin(dice *dice.Dicepool) float64 {
	return []float64{0.10, 0.12, 0.14, 0.16, 0.20, 0.22, 0.25, 0.30, 0.35, 0.40, 0.42}[dice.Sroll("2d6-2")]
}
func pressureThin(dice *dice.Dicepool) float64 {
	return []float64{0.43, 0.45, 0.47, 0.50, 0.52, 0.56, 0.60, 0.64, 0.66, 0.68, 0.70}[dice.Sroll("2d6-2")]
}
func pressureStandard(dice *dice.Dicepool) float64 {
	return []float64{0.71, 0.75, 0.80, 0.90, 1.00, 1.00, 1.10, 1.2, 1.3, 1.4, 1.49}[dice.Sroll("2d6-2")]
}
func pressureDense(dice *dice.Dicepool) float64 {
	return []float64{1.5, 1.6, 1.7, 1.8, 1.9, 2, 2.1, 2.2, 2.3, 2.4, 2.49}[dice.Sroll("2d6-2")]
}
func pressureDenseHigh(dice *dice.Dicepool) float64 {
	return []float64{2.5, 3, 5, 10, 20, 40, 80, 100, 150, 200, 250}[dice.Sroll("2d6-2")]
}
func pressureThinLow(dice *dice.Dicepool) float64 {
	return []float64{0.005, 0.007, 0.01, 0.03, 0.05, 0.07, 0.1, 0.2, 0.3, 0.4, 0.5}[dice.Sroll("2d6-2")]
}

func taint(dice *dice.Dicepool) int {
	return dice.Sroll("2d6-1")
}

func corrosiveAtmosphere(dice *dice.Dicepool) (string, float64) {
	switch dice.Sroll("2d6") {
	case 2:
		return "Helium, ammonia and methane", pressureThin(dice)
	case 3:
		return "Helium, ammonia and methane", pressureStandard(dice)
	case 4:
		return "Helium, ammonia and methane", pressureDense(dice)
	case 5:
		return "Nitrogen, ammonia and methane", pressureThin(dice)
	case 6:
		return "Nitrogen, ammonia and methane", pressureStandard(dice)
	case 7:
		return "Nitrogen, ammonia and methane", pressureDense(dice)
	case 8:
		return "Nitrogen, nitrogen dioxide and carbone dioxide", pressureThin(dice)
	case 9:
		return "Nitrogen, nitrogen dioxide and carbone dioxide", pressureStandard(dice)
	case 10:
		return "Nitrogen, nitrogen dioxide and carbone dioxide", pressureDense(dice)
	case 11:
		return "Carbone dioxide, sulfur dioxide and argon", pressureDense(dice)
	case 12:
		return "Phosphorous trioxide, phosphorous tricloride and carbone dioxide", pressureStandard(dice)
	}
	return "", 0
}

func insidiousAtmosphere(dice *dice.Dicepool) (string, float64) {
	switch dice.Sroll("2d6") {
	case 2:
		return "Methane, ammonia and hydrogen", pressureThin(dice)
	case 3:
		return "Methane, ammonia and hydrogen", pressureStandard(dice)
	case 4:
		return "Methane, ammonia and hydrogen", pressureDense(dice)
	case 5:
		return "Carbone dioxide, nitrogen and sulfur dioxide", pressureThin(dice)
	case 6:
		return "Carbone dioxide, nitrogen and sulfur dioxide", pressureStandard(dice)
	case 7:
		return "Carbone dioxide, nitrogen and sulfur dioxide", pressureDense(dice)
	case 8:
		return "Nitrogen, fluorine and carbone dioxide", pressureThin(dice)
	case 9:
		return "Nitrogen, fluorine and carbone dioxide", pressureStandard(dice)
	case 10:
		return "Nitrogen, fluorine and carbone dioxide", pressureDense(dice)
	case 11:
		return "Nitrogen, fluorine and carbone dioxide", pressureDenseHigh(dice)
	case 12:
		return "Nitrogen, methane and hydrogen", pressureDense(dice)
	}
	return "", 0
}

func (ad *AtmoDetails) rollAlbedo(dice *dice.Dicepool, atmo, hydr ehex.Ehex) error {
	switch atmo.Value() {
	case 4, 5, 6, 7, 8, 9:
		switch hydr.Value() {
		case 0, 1, 2:
			ad.Albedo = albedo(dice, 0)
		case 3, 4, 5:
			ad.Albedo = albedo(dice, 1)
		case 6, 7, 8:
			ad.Albedo = albedo(dice, 2)
		case 9, 10:
			ad.Albedo = albedo(dice, 3)
		}
	default:
		switch hydr.Value() {
		case 0, 1, 2, 3, 4:
			ad.Albedo = albedo(dice, 4)
		case 5, 6, 7, 8, 9, 10:
			ad.Albedo = albedo(dice, 5)
		}
	}
	return nil
}

func albedo(dice *dice.Dicepool, n int) float64 {
	albedoSt := []int{7, 13, 23, 29, 5, 47}
	albAdd := dice.Sroll("2d10-2")
	return utils.RoundFloat64(float64(albedoSt[n]+albAdd)*0.001, 3)
}

func (ad *AtmoDetails) calculateTemp(dice *dice.Dicepool, sd *sizerelated.SizeDetails, star star.StarBody, size, atmo, hzVar ehex.Ehex) error {
	//ge = (p/6)*(p/g)
	e := (ad.Pressure / 6) * (ad.Pressure / sd.Gravity)
	ad.GreenhouseEffect = utils.RoundFloat64(e, 3)
	// b := (271.0 * math.Pow(star.Luminocity(), 0.25)) / math.Sqrt(sd.OrbitalDistance)
	// b = (float64(((hzVar.Value() - 10) * 50) + dice.Sroll("1d40")))

	// t := b * math.Pow(1-ad.Albedo, 0.25) * (1 + ad.GreenhouseEffect)
	avTemp := 0.0
	if hzVar.Value() > 11 {
		avTemp = float64(dice.Sroll("1d223") - 223)
	}
	if hzVar.Value() == 11 {
		avTemp = float64(273 - dice.Sroll("1d50"))
	}
	if hzVar.Value() == 10 {
		avTemp = float64(273 + dice.Sroll("1d30"))
	}
	if hzVar.Value() == 9 {
		avTemp = float64(303 + dice.Sroll("1d50"))
	}
	if hzVar.Value() < 9 {
		avTemp = float64(353 + dice.Sroll("1d350"))
	}

	ad.AverageTemp = utils.RoundFloat64(float64(avTemp), 0)

	switch size.Value() {
	case 0, 1:
		// ad.EquatorAverageTemp = ad.AverageTemp
		// ad.EquatorTempSummer = ad.AverageTemp
		// ad.EquatorTempWinter = ad.AverageTemp
	default:
		ad.EquatorAverageTemp = utils.RoundFloat64(ad.AverageTemp+float64(3*size.Value()), 2)
		qs := ad.EquatorAverageTemp + (axialTiltEquatorModifier(sd.AxialTilt) * (float64(sd.AxialTilt) * 0.6))
		ad.EquatorTempSummer = utils.RoundFloat64(qs, 2)
		qw := ad.EquatorAverageTemp + (axialTiltEquatorModifier(sd.AxialTilt) * (float64(sd.AxialTilt) * -1))
		ad.EquatorTempWinter = utils.RoundFloat64(qw, 2)
		ad.PolarAverageTemp = utils.RoundFloat64(ad.AverageTemp-float64(7*size.Value()), 2)
		ps := ad.PolarAverageTemp + (axialTiltPolarModifier(sd.AxialTilt) * float64(sd.AxialTilt) * 0.6)
		pw := ad.PolarAverageTemp + (axialTiltPolarModifier(sd.AxialTilt) * float64(sd.AxialTilt) * -1)
		ad.PolarTempSummer = utils.RoundFloat64(ps, 2)
		ad.PolarTempWinter = utils.RoundFloat64(pw, 2)

	}
	ad.fixLowTemp()
	// ad.AverageTemp = utils.RoundFloat64(ad.AverageTemp-273.15, 2)
	// ad.EquatorAverageTemp = utils.RoundFloat64(ad.EquatorAverageTemp-273.15, 2)
	// ad.PolarAverageTemp = utils.RoundFloat64(ad.PolarAverageTemp-273.15, 2)
	// ad.EquatorTempSummer = utils.RoundFloat64(ad.EquatorTempSummer-273.15, 2)
	// ad.EquatorTempWinter = utils.RoundFloat64(ad.EquatorTempWinter-273.15, 2)
	// ad.PolarTempSummer = utils.RoundFloat64(ad.PolarTempSummer-273.15, 2)
	// ad.PolarTempWinter = utils.RoundFloat64(ad.PolarTempWinter-273.15, 2)
	// ad.fixLowTemp()
	return nil
}

func (ad *AtmoDetails) fixLowTemp() {
	if ad.AverageTemp < 0 {
		ad.AverageTemp = 0
	}
	if ad.EquatorAverageTemp < 0 {
		ad.EquatorAverageTemp = 0
	}
	if ad.EquatorTempSummer < 0 {
		ad.EquatorTempSummer = 0
	}
	if ad.EquatorTempWinter < 0 {
		ad.EquatorTempWinter = 0
	}
	if ad.PolarAverageTemp < 0 {
		ad.PolarAverageTemp = 0
	}
	if ad.PolarTempSummer < 0 {
		ad.PolarTempSummer = 0
	}
	if ad.PolarTempWinter < 0 {
		ad.PolarTempWinter = 0
	}
}

func axialTiltEquatorModifier(tilt int) float64 {
	switch tilt / 5 {
	case 7:
		return 0.25
	case 8:
		return 0.35
	case 9:
		return 0.40
	case 10:
		return 0.45
	case 11:
		return 0.50
	case 12:
		return 0.55
	case 13:
		return 0.65
	case 14:
		return 0.75
	case 15:
		return 0.80
	case 16:
		return 0.85
	case 17:
		return 0.90
	case 18:
		return 1
	}
	return 0
}

func axialTiltPolarModifier(tilt int) float64 {
	switch tilt {
	case 5:
		return 0.8
	case 4:
		return 0.7
	case 3:
		return 0.5
	case 2:
		return 0.4
	case 1:
		return 0.3
	case 0:
		return 0.2
	}
	return 0
}

func modR(atmo ehex.Ehex) float64 {
	switch atmo.Value() {
	case 0:
		return 1.4
	case 1:
		return 1.3
	case 2, 3, 14:
		return 1.2
	case 4, 5:
		return 1.15
	case 6, 7:
		return 1.1
	case 8, 9, 10:
		return 1.09
	case 13:
		return 1.02
	default:
		return 1

	}
}

func modZ(atmo ehex.Ehex) float64 {
	switch atmo.Value() {
	case 0:
		return 0.1
	case 1:
		return 0.1
	case 2, 3, 14:
		return 0.2
	case 4, 5:
		return 0.66
	case 6, 7:
		return 0.75
	case 8, 9, 10:
		return 0.85
	case 13:
		return 0.95
	default:
		return 1

	}
}
