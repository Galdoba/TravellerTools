package planetarydetails

import "fmt"

const (
	TAINT_NONE    = "None"
	Atmo_Standard = "Nitrogen, Oxigen, Argon, Carbon Dioxide"
	Atmo_1        = "Hellium, Ammonia, Methane"
	Atmo_2        = "Nitrogen, Ammonia, Methane"
	Atmo_3        = "Nitrogen, Nitrogen Dioxide, Nitrous Oxide, Carbon Dioxide"
	Atmo_4        = "Carbon Dioxide, Sulfur Dioxide, Argon"
	Atmo_5        = "Phosphorous Trioxide, Phosporous Tricloride, Carbone Dioxide"
	Atmo_6        = "Methane, Ammonia, Hydrogen"
	Atmo_7        = "Carbon Dioxide, Nitrogen, Sulfur Dioxide"
	Atmo_8        = "Nitrogen, Flourine, Carbon Dioxide"
	Atmo_9        = "Nitrogen, Methan, Hydrogen"
)

func (pd *PlanetaryDetails) setTaint() error {
	switch pd.atmo {
	default:
		pd.taint = TAINT_NONE
	case 2, 4, 7, 9, 10:
		tRoll := pd.dice.Roll("2d6").DM(-2).Sum()
		switch tRoll {
		case 2:
			pd.taint = "Clorine"
		case 3:
			pd.taint = "Fluorine"
		case 4:
			pd.taint = "Sulfur"
		case 5:
			pd.taint = "High Oxygen"
		case 6:
			pd.taint = "Desease"
		case 7:
			pd.taint = "Pollen/Spores"
		case 8:
			pd.taint = "Biotoxins"
		case 9:
			pd.taint = "Dust"
		case 10:
			pd.taint = "Volcanic Ash"
		case 11:
			pd.taint = "Low Oxygen"
		case 12:
			pd.taint = "Nitrogen Oxides"
		}
	}
	return nil
}

func (pd *PlanetaryDetails) defineAtmosphereRelatedDetails() error {
	pd.pressureCode = pd.atmo
	switch pd.pressureCode {
	case 0:
		pd.atmoComposition = "None"
	case 1:
		pd.atmoComposition = "Trace"
	case 3, 5, 6, 8, 13, 14:
		pd.atmoComposition = Atmo_Standard
	case 2, 4, 7, 9:
		pd.atmoComposition = Atmo_Standard
	case 11:
		if err := pd.setCorrosiveAtmoType(); err != nil {
			return err
		}
	case 12:
		if err := pd.setInsidiousAtmoType(); err != nil {
			return err
		}
	}
	if err := pd.setTaint(); err != nil {
		return err
	}
	if err := pd.setPresure(); err != nil {
		return err
	}
	return nil
}

func (pd *PlanetaryDetails) setPresure() error {
	presureRoll := pd.dice.Roll("2d6").Sum()
	switch pd.pressureCode {
	case 0:
		pd.pressure = 0
	case 1:
		presure := []float64{0.001, 0.002, 0.005, 0.007, 0.01, 0.02, 0.03, 0.05, 0.07, 0.08, 0.09}
		pd.pressure = presure[presureRoll-2]
	case 2, 3:
		presure := []float64{0.10, 0.12, 0.14, 0.16, 0.20, 0.22, 0.25, 0.30, 0.35, 0.40, 0.42}
		pd.pressure = presure[presureRoll-2]
	case 4, 5:
		presure := []float64{0.43, 0.45, 0.47, 0.50, 0.52, 0.56, 0.60, 0.64, 0.66, 0.68, 0.70}
		pd.pressure = presure[presureRoll-2]
	case 6, 7:
		presure := []float64{0.71, 0.75, 0.80, 0.90, 1.00, 1.00, 1.00, 1.10, 1.20, 1.30, 1.40, 1.49}
		pd.pressure = presure[presureRoll-2]
	case 8, 9, 10:
		presure := []float64{1.50, 1.60, 1.70, 1.80, 1.90, 2.00, 2.10, 2.20, 2.30, 2.40, 2.49}
		pd.pressure = presure[presureRoll-2]
	case 13:
		presure := []float64{2.50, 3, 5, 10, 20, 40, 80, 100, 150, 200, 250}
		pd.pressure = presure[presureRoll-2]
	case 14:
		presure := []float64{0.005, 0.007, 0.01, 0.03, 0.05, 0.07, 0.1, 0.2, 0.3, 0.4, 0.5}
		pd.pressure = presure[presureRoll-2]
	default:
		return fmt.Errorf("case 15 not desided")
	}
	return nil
}

func (pd *PlanetaryDetails) setCorrosiveAtmoType() error {
	if pd.atmo != 11 {
		return nil
	}
	cRoll := pd.dice.Roll("2d6").Sum()
	switch cRoll {
	default:
		return fmt.Errorf("unexpected roll result on setCorrosiveAtmoType()")
	case 2:
		pd.atmoComposition = Atmo_1
		pd.pressureCode = 4
	case 3:
		pd.atmoComposition = Atmo_1
		pd.pressureCode = 6
	case 4:
		pd.atmoComposition = Atmo_1
		pd.pressureCode = 8
	case 5:
		pd.atmoComposition = Atmo_2
		pd.pressureCode = 4
	case 6:
		pd.atmoComposition = Atmo_2
		pd.pressureCode = 6
	case 7:
		pd.atmoComposition = Atmo_2
		pd.pressureCode = 8
	case 8:
		pd.atmoComposition = Atmo_3
		pd.pressureCode = 4
	case 9:
		pd.atmoComposition = Atmo_3
		pd.pressureCode = 6
	case 10:
		pd.atmoComposition = Atmo_3
		pd.pressureCode = 8
	case 11:
		pd.atmoComposition = Atmo_4
		pd.pressureCode = 8
	case 12:
		pd.atmoComposition = Atmo_5
		pd.pressureCode = 6
	}
	return nil
}

func (pd *PlanetaryDetails) setInsidiousAtmoType() error {
	if pd.atmo != 12 {
		return nil
	}
	cRoll := pd.dice.Roll("2d6").Sum()
	switch cRoll {
	default:
		return fmt.Errorf("unexpected roll result on setCorrosiveAtmoType()")
	case 2:
		pd.atmoComposition = Atmo_6
		pd.pressureCode = 4
	case 3:
		pd.atmoComposition = Atmo_6
		pd.pressureCode = 6
	case 4:
		pd.atmoComposition = Atmo_6
		pd.pressureCode = 8
	case 5:
		pd.atmoComposition = Atmo_7
		pd.pressureCode = 4
	case 6:
		pd.atmoComposition = Atmo_7
		pd.pressureCode = 6
	case 7:
		pd.atmoComposition = Atmo_7
		pd.pressureCode = 8
	case 8:
		pd.atmoComposition = Atmo_8
		pd.pressureCode = 4
	case 9:
		pd.atmoComposition = Atmo_8
		pd.pressureCode = 6
	case 10:
		pd.atmoComposition = Atmo_8
		pd.pressureCode = 8
	case 11:
		pd.atmoComposition = Atmo_8
		pd.pressureCode = 13
	case 12:
		pd.atmoComposition = Atmo_9
		pd.pressureCode = 8
	}
	return nil
}
