package profile

const (
	Port = "Starport"
	Size = "Size"
	Atmo = "Atmosphere"
	Hydr = "Hydrosphere"
	Pops = "Population"
	Govr = "Goverment"
	Laws = "Laws"
	TL   = "Tech Level"
)

func encodeUWPdata(aspect string, val string) ehex.Ehex {
	hex := ehex.New()
	hex.Set(val)
	hex.Encode("value is not correct")
	switch aspect {
	case Port:
		switch val {
		case "A":
			hex.Encode("Excellent Starport")
		case "B":
			hex.Encode("Good Starport")
		case "C":
			hex.Encode("Routine Starport")
		case "D":
			hex.Encode("Poor Starport")
		case "E":
			hex.Encode("Frontier Starport")
		case "X":
			hex.Encode("No Starport")
		case "F":
			hex.Encode("Routine Spaceport")
		case "G":
			hex.Encode("Poor Spaceport")
		case "H":
			hex.Encode("Primitive Spaceport")
		case "Y":
			hex.Encode("No Spaceport")
		}
	}