package uwp

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile"
)

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

type uwp struct {
	dataType map[string]ehex.Ehex
	descr    string
}

type UWP interface {
	profile.Profile
	fmt.Stringer
	Starport() string
	Size() int
	Atmo() int
	Hydr() int
	Pops() int
	Govr() int
	Laws() int
	TL() int
}

func New() *uwp {
	u := uwp{}
	u.dataType = make(map[string]ehex.Ehex)
	u.descr = "UWP describes World Characteristics"
	return &u
}

func Inject(uwpString string) *uwp {
	u, err := FromString(uwpString)

	if err != nil {
		fmt.Println(err.Error(), uwpString)
		return nil
	}
	return u
}

func FromString(uwpString string) (*uwp, error) {
	u := New()
	s := strings.Split(uwpString, "")
	if len(s) != 9 {
		return nil, fmt.Errorf("invalid uwp string (%v)", uwpString)
	}

	for i, hex := range s {
		val := ehex.New().Set(hex)
		switch i {
		case 0:
			err := u.Encode(Port, val)
			if err != nil {
				return nil, err
			}
		case 1:
			err := u.Encode(Size, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		case 2:
			err := u.Encode(Atmo, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		case 3:
			err := u.Encode(Hydr, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		case 4:
			err := u.Encode(Pops, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		case 5:
			err := u.Encode(Govr, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		case 6:
			err := u.Encode(Laws, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		case 8:
			err := u.Encode(TL, ehex.New().Set(hex))
			if err != nil {
				return nil, err
			}
		}
	}
	return u, nil
}

func (u *uwp) Encode(aspect string, val ehex.Ehex) error {
	switch aspect {
	default:
		return fmt.Errorf("aspect '%v' is not valid for this type of profile", aspect)
	case Port, Size, Atmo, Hydr, Pops, Govr, Laws, TL:
		u.dataType[aspect] = encodeUWPdata(aspect, val.Code())
	}
	if u.dataType[aspect].Meaning() == "value is not correct" {
		return fmt.Errorf("value '%v' is not correct for aspect '%v'", val, aspect)
	}
	return nil
}

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
	case Size:
		switch val {
		case "0":
			hex.Encode("Asteroid Belt")
		case "1":
			hex.Encode("1,600 km")
		case "2":
			hex.Encode("3,200 km")
		case "3":
			hex.Encode("4,800 km")
		case "4":
			hex.Encode("6,400 km")
		case "5":
			hex.Encode("8,000 km")
		case "6":
			hex.Encode("9,600 km")
		case "7":
			hex.Encode("11,200 km")
		case "8":
			hex.Encode("12,800 km")
		case "9":
			hex.Encode("14,400 km")
		case "A":
			hex.Encode("16,000 km")
		case "B":
			hex.Encode("17,600 km")
		case "C":
			hex.Encode("19,200 km")
		case "D":
			hex.Encode("20,800 km")
		case "E":
			hex.Encode("22,400 km")
		case "F":
			hex.Encode("24,000 km")
		case "G":
			hex.Encode("24.944-26.553 km")
		case "H":
			hex.Encode("26.554-28.162 km")
		case "J":
			hex.Encode("28.163-29.771 km")
		case "K":
			hex.Encode("29.772-31.381 km")
		case "L":
			hex.Encode("31.382-32.990 km")
		case "M":
			hex.Encode("32.991-34.599 km")
		case "N":
			hex.Encode("34.600-36.209 km")
		case "P":
			hex.Encode("36.201-37.818 km")
		case "Q":
			hex.Encode("37.819-39.427 km")

		}
	case Atmo:
		switch val {
		case "0":
			hex.Encode("Vacuum")
		case "1":
			hex.Encode("Trace")
		case "2":
			hex.Encode("Very Thin Tainted")
		case "3":
			hex.Encode("Very Thin")
		case "4":
			hex.Encode("Thin Tainted")
		case "5":
			hex.Encode("Thin")
		case "6":
			hex.Encode("Standard")
		case "7":
			hex.Encode("Standard Tainted")
		case "8":
			hex.Encode("Dense")
		case "9":
			hex.Encode("Dense Tainted")
		case "A":
			hex.Encode("Exotic")
		case "B":
			hex.Encode("Corrosive")
		case "C":
			hex.Encode("Insidious")
		case "D":
			hex.Encode("Dense High")
		case "E":
			hex.Encode("Thin Low")
		case "F":
			hex.Encode("Unusual")
		}
	case Hydr:
		switch val {
		case "0":
			hex.Encode("Desert World")
		case "1":
			hex.Encode("10% Water")
		case "2":
			hex.Encode("20% Water")
		case "3":
			hex.Encode("30% Water")
		case "4":
			hex.Encode("40% Water")
		case "5":
			hex.Encode("50% Water")
		case "6":
			hex.Encode("60% Water")
		case "7":
			hex.Encode("70% Water")
		case "8":
			hex.Encode("80% Water")
		case "9":
			hex.Encode("90% Water")
		case "A":
			hex.Encode("Water World")
		}
	case Pops:
		switch val {
		case "0":
			hex.Encode("Unpopulated")
		case "1":
			hex.Encode("Tens")
		case "2":
			hex.Encode("Hundreds")
		case "3":
			hex.Encode("Thousands")
		case "4":
			hex.Encode("Tens of Thousands")
		case "5":
			hex.Encode("100,000's")
		case "6":
			hex.Encode("Millions")
		case "7":
			hex.Encode("Ten Millions")
		case "8":
			hex.Encode("Hundred Millions")
		case "9":
			hex.Encode("Billions")
		case "A":
			hex.Encode("Tens of Billions")
		case "B":
			hex.Encode("Hundred Billions")
		case "C":
			hex.Encode("Trillions")
		case "D":
			hex.Encode("Ten Trillions")
		case "E":
			hex.Encode("Hundred Trillions")
		case "F":
			hex.Encode("Quadrillions")
		}
	case Govr:
		switch val {
		case "0":
			hex.Encode("No Government Structure")
		case "1":
			hex.Encode("Company/Corporation")
		case "2":
			hex.Encode("Participating Democracy")
		case "3":
			hex.Encode("Self-Perpetuating Oligarchy")
		case "4":
			hex.Encode("Representative Democracy")
		case "5":
			hex.Encode("Feudal Technocracy")
		case "6":
			hex.Encode("Captive Government/Colony")
		case "7":
			hex.Encode("Balkanization")
		case "8":
			hex.Encode("Civil Service Bureaucracy")
		case "9":
			hex.Encode("Impersonal Bureaucracy")
		case "A":
			hex.Encode("Charismatic Dictatorship")
		case "B":
			hex.Encode("Non-Charismatic Dictatorship")
		case "C":
			hex.Encode("Charismatic Oligarchy")
		case "D":
			hex.Encode("Religious Dictatorship")
		case "E":
			hex.Encode("Religious Autocracy")
		case "F":
			hex.Encode("Totalitarian Oligarchy")
		}
	case Laws:
		switch val {
		case "0":
			hex.Encode("No Law: No prohibitions")
		case "1":
			hex.Encode("Low Law: Prohibition of WMD, Psi weapons")
		case "2":
			hex.Encode("Low Law: Prohibition of “Portable” Weapons")
		case "3":
			hex.Encode("Low Law: Prohibition of Acid, Fire, Gas")
		case "4":
			hex.Encode("Moderate Law: Prohibition of Laser, Beam")
		case "5":
			hex.Encode("Moderate Law: No Shock,EMP,Rad, Mag, Grav")
		case "6":
			hex.Encode("Moderate Law: Prohibition of MachineGuns")
		case "7":
			hex.Encode("Moderate Law: Prohibition of Pistols")
		case "8":
			hex.Encode("High Law: Open display of weapons prohibited")
		case "9":
			hex.Encode("High Law: No weapons outside the home")
		case "A":
			hex.Encode("Extreme Law: All weapons prohibited")
		case "B":
			hex.Encode("Extreme Law: Continental passports required")
		case "C":
			hex.Encode("Extreme Law: Unrestricted invasion of privacy")
		case "D":
			hex.Encode("Extreme Law: Paramilitary law enforcement")
		case "E":
			hex.Encode("Extreme Law: Full-fledged police state")
		case "F":
			hex.Encode("Extreme Law: Daily life rigidly controlled")
		case "G":
			hex.Encode("Extreme Law: Disproportionate punishment")
		case "H":
			hex.Encode("Extreme Law: Legalized oppressive practices")
		case "J":
			hex.Encode("Extreme Law: Routine oppression")
		}
	case TL:
		switch val {
		case "0":
			hex.Encode("TL0")
		case "1":
			hex.Encode("TL1")
		case "2":
			hex.Encode("TL2")
		case "3":
			hex.Encode("TL3")
		case "4":
			hex.Encode("TL4")
		case "5":
			hex.Encode("TL5")
		case "6":
			hex.Encode("TL6")
		case "7":
			hex.Encode("TL7")
		case "8":
			hex.Encode("TL8")
		case "9":
			hex.Encode("TL9")
		case "A":
			hex.Encode("TL10")
		case "B":
			hex.Encode("TL11")
		case "C":
			hex.Encode("TL12")
		case "D":
			hex.Encode("TL13")
		case "E":
			hex.Encode("TL14")
		case "F":
			hex.Encode("TL15")
		case "G":
			hex.Encode("TL16")
		case "H":
			hex.Encode("TL17")
		case "J":
			hex.Encode("TL18")
		case "K":
			hex.Encode("TL19")
		case "L":
			hex.Encode("TL20")
		}
	}
	return hex
}

func (u *uwp) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v-%v", u.dataType[Port].Code(), u.dataType[Size].Code(), u.dataType[Atmo].Code(), u.dataType[Hydr].Code(),
		u.dataType[Pops].Code(), u.dataType[Govr].Code(), u.dataType[Laws].Code(), u.dataType[TL].Code())
}

func (u *uwp) Starport() string {
	return u.dataType[Port].Code()
}

func (u *uwp) Size() int {
	return u.dataType[Size].Value()
}
func (u *uwp) Atmo() int {
	return u.dataType[Atmo].Value()
}
func (u *uwp) Hydr() int {
	return u.dataType[Hydr].Value()
}
func (u *uwp) Pops() int {
	return u.dataType[Pops].Value()
}
func (u *uwp) Govr() int {
	return u.dataType[Govr].Value()
}
func (u *uwp) Laws() int {
	return u.dataType[Laws].Value()
}
func (u *uwp) TL() int {
	return u.dataType[TL].Value()
}

func (u *uwp) Describe(aspect string) string {
	switch aspect {
	default:
		return "description error"
	case Port:
		return u.dataType[Port].Meaning()
	case Size:
		return u.dataType[Size].Meaning()
	case Atmo:
		return u.dataType[Atmo].Meaning()
	case Hydr:
		return u.dataType[Hydr].Meaning()
	case Pops:
		return u.dataType[Pops].Meaning()
	case Govr:
		return u.dataType[Govr].Meaning()
	case Laws:
		return u.dataType[Laws].Meaning()
	case TL:
		return u.dataType[TL].Meaning()
	case "All":
		str := ""
		str += Port + "    : " + u.dataType[Port].Code() + " (" + u.dataType[Port].Meaning() + ")" + "\n"
		str += Size + "        : " + u.dataType[Size].Code() + " (" + u.dataType[Size].Meaning() + ")" + "\n"
		str += Atmo + "  : " + u.dataType[Atmo].Code() + " (" + u.dataType[Atmo].Meaning() + ")" + "\n"
		str += Hydr + " : " + u.dataType[Hydr].Code() + " (" + u.dataType[Hydr].Meaning() + ")" + "\n"
		str += Pops + "  : " + u.dataType[Pops].Code() + " (" + u.dataType[Pops].Meaning() + ")" + "\n"
		str += Govr + "   : " + u.dataType[Govr].Code() + " (" + u.dataType[Govr].Meaning() + ")" + "\n"
		str += Laws + "        : " + u.dataType[Laws].Code() + " (" + u.dataType[Laws].Meaning() + ")" + "\n"
		str += TL + "  : " + u.dataType[TL].Code() + " (" + u.dataType[TL].Meaning() + ")"
		return str
	}
}

func (u *uwp) Decode(data string) (string, int, string, error) {
	return "", 0, "", fmt.Errorf("abstract method - shouldnot be used %v", data)
}

func (u *uwp) callData(data string) ehex.Ehex {
	return u.dataType[data]
}

func (u *uwp) AddValue(data string, val int) error {
	switch data {
	default:
		return fmt.Errorf("unknown data type '%v'", data)
	case Size, Atmo, Hydr, Pops, Govr, Laws, TL:
		oldVal := u.callData(data)
		newValInt := oldVal.Value() + val
		newVal := ehex.New().Set(newValInt)
		err := u.Encode(data, newVal)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *uwp) ChangeValue(data string, val int) error {
	switch data {
	default:
		return fmt.Errorf("unknown data type '%v'", data)
	case Size, Atmo, Hydr, Pops, Govr, Laws, TL:
		newVal := ehex.New().Set(val)
		err := u.Encode(data, newVal)
		if err != nil {
			return err
		}
	}
	return nil
}

//err := u.Encode(Size, ehex.New().Set(hex))

func RecomendTravelZone(u UWP) string {
	warn := 0
	if u.Atmo() >= 10 {
		warn++
	}
	switch u.Govr() {
	case 0, 7, 10:
		warn++
	}
	if u.Laws() == 0 {
		warn++
	}
	if u.Laws() >= 9 {
		warn++
	}
	if u.Govr()+u.Laws() >= 20 {
		warn++
	}
	if u.Govr()+u.Laws() >= 22 {
		warn++
	}
	if u.Starport() == "X" || u.Starport() == "Y" {
		warn++
	}
	switch warn {
	default:
		return ""
	case 1, 2:
		return "A"
	case 3, 4, 5, 6, 7:
		return "R"
	}
}
