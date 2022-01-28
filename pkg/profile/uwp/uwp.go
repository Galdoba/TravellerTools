package uwp

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/ehex"
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
	aspect *profile.ProfileData
	descr  string
}

func New() *uwp {
	u := uwp{}
	u.aspect = profile.New(Port, Size, Atmo, Hydr, Pops, Govr, Laws, TL)
	u.descr = "UWP describes World Characteristics"
	return &u
}

func (u *uwp) SetString(str string) error {
	s := strings.Split(str, "")
	if len(s) != 9 {
		return fmt.Errorf("invalid uwp string (%v)", str)
	}
	for i, hex := range s {

		switch i {
		case 0:
			err := u.Set(Port, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 1:
			err := u.Set(Size, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 2:
			err := u.Set(Atmo, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 3:
			err := u.Set(Hydr, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 4:
			err := u.Set(Pops, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 5:
			err := u.Set(Govr, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 6:
			err := u.Set(Laws, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		case 8:
			err := u.Set(TL, ehex.New().Set(hex).Value())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (u *uwp) Set(aspect string, val int) error {
	switch aspect {
	default:
		return fmt.Errorf("aspect '%v' is not valid for this type of profile", aspect)
	case Port, Size, Atmo, Hydr, Pops, Govr, Laws, TL:
		u.aspect.Data[aspect] = setUWPdata(aspect, val)
	}
	if u.aspect.Data[aspect].Meaning() == "value is not correct" {
		return fmt.Errorf("value '%v' is not correct for aspect '%v'", val, aspect)
	}
	return nil
}

func setUWPdata(aspect string, val int) ehex.Ehex {
	hex := ehex.New()
	hex.Set(val)
	hex.Encode("value is not correct")
	switch aspect {
	case Port:
		switch ehex.New().Set(val).Code() {
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
		switch ehex.New().Set(val).Code() {
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
		}
	case Atmo:
		switch ehex.New().Set(val).Code() {
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
		switch ehex.New().Set(val).Code() {
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
		switch ehex.New().Set(val).Code() {
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
		switch ehex.New().Set(val).Code() {
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
		switch ehex.New().Set(val).Code() {
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
		switch ehex.New().Set(val).Code() {
		case "0":
			hex.Encode("TL 0")
		case "1":
			hex.Encode("TL 1")
		case "2":
			hex.Encode("TL 2")
		case "3":
			hex.Encode("TL 3")
		case "4":
			hex.Encode("TL 4")
		case "5":
			hex.Encode("TL 5")
		case "6":
			hex.Encode("TL 6")
		case "7":
			hex.Encode("TL 7")
		case "8":
			hex.Encode("TL 8")
		case "9":
			hex.Encode("TL 9")
		case "A":
			hex.Encode("TL 10")
		case "B":
			hex.Encode("TL 11")
		case "C":
			hex.Encode("TL 12")
		case "D":
			hex.Encode("TL 13")
		case "E":
			hex.Encode("TL 14")
		case "F":
			hex.Encode("TL 15")
		case "G":
			hex.Encode("TL 16")
		case "H":
			hex.Encode("TL 17")
		case "J":
			hex.Encode("TL 18")
		case "K":
			hex.Encode("TL 19")
		case "L":
			hex.Encode("TL 20")
		}
	}
	return hex
}

func (u *uwp) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v-%v", u.aspect.Data[Port].Code(), u.aspect.Data[Size].Code(), u.aspect.Data[Atmo].Code(), u.aspect.Data[Hydr].Code(),
		u.aspect.Data[Pops].Code(), u.aspect.Data[Govr].Code(), u.aspect.Data[Laws].Code(), u.aspect.Data[TL].Code())
}

func (u *uwp) Starport() string {
	return u.aspect.Data[Port].Code()
}

func (u *uwp) Size() int {
	return u.aspect.Data[Size].Value()
}
func (u *uwp) Atmo() int {
	return u.aspect.Data[Atmo].Value()
}
func (u *uwp) Hydr() int {
	return u.aspect.Data[Hydr].Value()
}
func (u *uwp) Pops() int {
	return u.aspect.Data[Pops].Value()
}
func (u *uwp) Govr() int {
	return u.aspect.Data[Govr].Value()
}
func (u *uwp) Laws() int {
	return u.aspect.Data[Laws].Value()
}
func (u *uwp) TL() int {
	return u.aspect.Data[TL].Value()
}

func (u *uwp) Describe(aspect string) string {
	switch aspect {
	default:
		return "description error"
	case Port:
		return u.aspect.Data[Port].Meaning()
	case Size:
		return u.aspect.Data[Size].Meaning()
	case Atmo:
		return u.aspect.Data[Atmo].Meaning()
	case Hydr:
		return u.aspect.Data[Hydr].Meaning()
	case Pops:
		return u.aspect.Data[Pops].Meaning()
	case Govr:
		return u.aspect.Data[Govr].Meaning()
	case Laws:
		return u.aspect.Data[Laws].Meaning()
	case TL:
		return u.aspect.Data[TL].Meaning()
	case "All":
		str := ""
		str += Port + "    : " + u.aspect.Data[Port].Code() + " (" + u.aspect.Data[Port].Meaning() + ")" + "\n"
		str += Size + "        : " + u.aspect.Data[Size].Code() + " (" + u.aspect.Data[Size].Meaning() + ")" + "\n"
		str += Atmo + "  : " + u.aspect.Data[Atmo].Code() + " (" + u.aspect.Data[Atmo].Meaning() + ")" + "\n"
		str += Hydr + " : " + u.aspect.Data[Hydr].Code() + " (" + u.aspect.Data[Hydr].Meaning() + ")" + "\n"
		str += Pops + "  : " + u.aspect.Data[Pops].Code() + " (" + u.aspect.Data[Pops].Meaning() + ")" + "\n"
		str += Govr + "   : " + u.aspect.Data[Govr].Code() + " (" + u.aspect.Data[Govr].Meaning() + ")" + "\n"
		str += Laws + "        : " + u.aspect.Data[Laws].Code() + " (" + u.aspect.Data[Laws].Meaning() + ")" + "\n"
		str += TL + "  : " + u.aspect.Data[TL].Code() + " (" + u.aspect.Data[TL].Meaning() + ")"
		return str
	}
}
