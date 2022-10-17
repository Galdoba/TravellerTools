package planetarydetails

const (
	TAINT_NONE = "None"
)

func (pd *PlanetaryDetails) setTaint() error {
	switch pd.uwpData.Atmo() {
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
