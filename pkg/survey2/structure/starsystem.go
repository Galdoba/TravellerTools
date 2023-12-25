package structure

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StarSystem struct {
	Qty         int              `json:"Number of Stars"`
	Gyrs        float64          `json:"Star System Age"`
	Stars       map[string]*Star `json:"Stars"`
	TestOrbital *Orbital         `json:"HEX POPULATION"`
}

func (ss *StarSystem) MarshalStarSystem() ([]byte, error) {
	fmt.Println("MarshalJSON()")
	return json.MarshalIndent(ss, "", "    ")
}

//5-G7 V-0.929-0.967-0.738-6.336:Ab-0.09-0.11-G8 V-0.907-0.957-0.681:B-6.1-0.8-K8 V-0.626-0.777-0.136:Ca-12.1-0.47-Mo V-0.510-0.728-0.0895:Cb-0.21-0.24-D-0.490-0.017-0.000525

func SystemFromProfile(prf string) (*StarSystem, error) {
	ss := StarSystem{}
	ss.Stars = make(map[string]*Star)
	starData := strings.Split(prf, ":")
	for i, data := range starData {

		switch i {
		case 0:
			star, qty, age, err := PrimaryFromProfile(data)
			if err != nil {
				return nil, err
			}
			ss.Stars["Aa"] = star
			ss.Qty = qty
			ss.Gyrs = age
		default:
			star, err := SecondaryFromProfile(data)
			if err != nil {
				return nil, err
			}
			ss.Stars[star.Designation] = star

		}
	}
	for k, str := range ss.Stars {
		switch k {
		case "Ab":
			ss.Stars["Aa"].Companion = str
			delete(ss.Stars, k)
		case "Bb":
			ss.Stars["Ba"].Companion = str
			delete(ss.Stars, k)
		case "Cb":
			ss.Stars["Ca"].Companion = str
			delete(ss.Stars, k)
		case "Db":
			ss.Stars["Da"].Companion = str
			delete(ss.Stars, k)
		}
	}
	return &ss, nil
}

// type Star struct {
// 	Type         string  `json:"Type,omitempty"`     //OBAFGKM+
// 	SubType      string  `json:"Sub Type,omitempty"` //0123456789-
// 	Class        string  `json:"Class,omitempty"`    //Ia Ib ... BD D PSR NS BH --
// 	Mass         float64 `json:"Mass,omitempty"`
// 	Diameter     float64 `json:"Diameter,omitempty"`
// 	Luminocity   float64 `json:"Luminocity,omitempty"`
// 	Age          float64 `json:"Age,omitempty"`
// 	Designation  string  `json:"System Position,omitempty"` //A Ab B Bb C Cb D Db
// 	Orbit        float64 `json:"Orbit#,omitempty"`
// 	Eccentrisity float64 `json:"Eccentrisity,omitempty"`
// }

type Orbital struct {
	OrbitalType       string     `json:"Type"` // Star/Terrestrial/Belt/Gigant/Void
	OrbNum            float64    `json:"Orbit#"`
	ObjectStar        *Star      `json:"Star,omitempty"`
	ObjectTerrestrial string     `json:"Terrestrial,omitempty"`
	ObjectBelt        *Belt      `json:"Belt,omitempty"`
	ObjectGigant      *GasGigant `json:"Gas Gigant,omitempty"`
	//ObjectVoid        GasGigant `json:"Gas Gigant,omitempty"`
	parent   *Orbital
	Satelite map[string]*Orbital `json:"Filled Orbits,omitempty"`
}

type SystemModel struct {
	CentralObject *Orbital
}

/*
System
  CentralObj // Void/Star
    Orbitals1
	  Orbitals2
	    ...
		  OrbitalsN
  Rogues
   -Rogue1
   -Rogue2
   -...
   -RogueN
AS STRUCT
System
  PrimaryStar or Void
    Star
	| Terrestrial
	| | Ring
	| | Moon
	| Belt
	|  +--Planetoids
	| Gigant
	| | Ring
	| | Moon
	Terrestrial
	| Ring
	| Moon
	Belt
	  --Planetoids
	Gigant
	  Ring
	  Moon
  Rogues
    Rogue Dwarf
	Rogue GG
	Rogue Terrestrial
	Rogue Planetoid



*/

type OrbitalBody interface {
	Parent() OrbitalBody
	OrbitN() float64
	BodyType() string
}

type Void struct {
	objType   string
	Satelites map[string]Orbital
}

func (v *Void) Parent() OrbitalBody {
	return nil
}
func (v *Void) OrbitN() float64 {
	return 0.0
}
func (v *Void) BodyType() string {
	return "Void"
}

type Terrestrial struct {
	planetName string
	orbit      float64
	parent     OrbitalBody
}

func newBelt(st *Orbital) *Orbital {
	orb := Orbital{}
	orb.OrbitalType = "Belt"
	orb.ObjectBelt = &Belt{
		Profile:  "a-a-a",
		Nomena:   "belt1",
		Span:     0,
		Mtype:    10,
		Stype:    40,
		Ctype:    30,
		Otype:    20,
		Bulk:     3,
		Resource: 5,
	}
	orb.parent = st
	return &orb
}

func newStar() *Orbital {
	orb := Orbital{}
	orb.OrbitalType = "Star"
	orb.ObjectStar = &Star{
		Type:         "G",
		SubType:      "2",
		Class:        "V",
		Mass:         1,
		Diameter:     1,
		Luminocity:   1,
		Age:          4.5,
		Designation:  "Aa",
		Orbit:        0,
		Eccentrisity: 0,
		Companion:    nil,
	}
	orb.Satelite = make(map[string]*Orbital)
	return &orb
}

func newVoid() *Orbital {
	v := Orbital{}
	v.OrbitalType = "Void"
	v.Satelite = make(map[string]*Orbital)
	v.Satelite["0.0"] = newStar()
	str := v.Satelite["0.0"]
	str.Satelite["2.1"] = newBelt(str)
	str.Satelite["3.1"] = newBelt(str)
	return &v
}

func (b *Orbital) BodyType() string {
	if b.ObjectStar != nil {
		return "Star"
	}
	if b.ObjectBelt != nil {
		return "Belt"
	}
	return "Void"
}
