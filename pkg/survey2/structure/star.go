package structure

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Class_Ia                = "Ia"
	Class_Ib                = "Ib"
	Class_II                = "II"
	Class_III               = "III"
	Class_IV                = "IV"
	Class_V                 = "V"
	Class_VI                = "VI"
	Class_BD                = "BD"
	Class_D                 = "D"
	Class_NS                = "NS"
	Class_PSR               = "PSR"
	Class_MGR               = "MGR"
	Class_BH                = "BH"
	Type_O                  = "O"
	Type_B                  = "B"
	Type_F                  = "F"
	Type_A                  = "A"
	Type_G                  = "G"
	Type_K                  = "K"
	Type_M                  = "M"
	Type_L                  = "L"
	Type_T                  = "T"
	Type_Y                  = "Y"
	SubType_0               = "0"
	SubType_1               = "1"
	SubType_2               = "2"
	SubType_3               = "3"
	SubType_4               = "4"
	SubType_5               = "5"
	SubType_6               = "6"
	SubType_7               = "7"
	SubType_8               = "8"
	SubType_9               = "9"
	Designation_Rogue       = "R"
	Designation_Primary     = "Aa"
	Designation_Primary_Cmp = "Ab"
	Designation_Close       = "Ba"
	Designation_Close_Cmp   = "Bb"
	Designation_Near        = "Ca"
	Designation_Near_Cmp    = "Cb"
	Designation_Far         = "Da"
	Designation_Far_Cmp     = "Db"
	tableType               = "tab_Type"
	tableHot                = "tab_Hot"
	tableSpecial            = "tab_Special"
	tableUnusual            = "tab_Unusual"
	tableGigants            = "tab_Gigants"
	tablePeculiar           = "tab_Peculiar"
	tableSubtNumeric        = "tab_Numeric"
	tableSubtMprim          = "tab_Mprimary"
)

type star struct {
	Class        string //Ia Ib ... BD D PSR NS BH --
	SubType      string //0123456789-
	Type         string //OBAFGKM+
	Mass         float64
	Diameter     float64
	Luminocity   float64
	Age          float64
	Designation  string //A Ab B Bb C Cb D Db
	Orbit        float64
	Eccentrisity float64
}

func (st *star) Normalize() {
	val := int(st.Mass * 1000)
	st.Mass = float64(val) / 1000
	val = int(st.Diameter * 1000)
	st.Diameter = float64(val) / 1000
	val = int(st.Luminocity * 1000)
	st.Luminocity = float64(val) / 1000
	val = int(st.Age * 1000)
	st.Age = float64(val) / 1000
	val = int(st.Orbit * 100)
	st.Orbit = float64(val) / 100
	val = int(st.Eccentrisity * 1000000)
	st.Eccentrisity = float64(val) / 1000000

}

func (st *star) Profile() string {
	t := st.Type + st.SubType
	c := st.Class
	tc := t + " " + c
	m := "-" + fmt.Sprintf("%v", st.Mass)
	d := "-" + fmt.Sprintf("%v", st.Diameter)
	l := "-" + fmt.Sprintf("%v", st.Luminocity)
	a := "-" + fmt.Sprintf("%v", st.Age)
	ds := ":" + st.Designation
	o := fmt.Sprintf("%v", st.Orbit)
	e := fmt.Sprintf("%v", st.Eccentrisity)
	switch c {
	case "BD", "D", "NS", "PSR", "MGR", "BH":
		tc = ""
	}
	switch ds {
	default:
		a = ""
	case ":Aa":
		ds = ""
		o = ""
		e = ""
	}
	return ds + o + e + tc + m + d + l + a
}

func FromProfile(prf string) (*star, error) {
	//ds + o + e + tc + m + d + l
	//             tc + m + d + l + a

	data := strings.Split(prf, "-")

	st := star{}
	switch len(data) {
	default:
		return nil, fmt.Errorf("bad input: (%v) %v", len(data), data)
	case 5:
		for i, prt := range data {
			switch i {
			case 0:
				st.Type, st.SubType, st.Class = disassemble(prt)
			case 1:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("mass: %v", err.Error())
				}
				st.Mass = fl
			case 2:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("diam: %v", err.Error())
				}
				st.Diameter = fl
			case 3:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("luma: %v", err.Error())
				}
				st.Luminocity = fl
			case 4:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("age: %v", err.Error())
				}
				st.Age = fl
			}
		}
	case 7:
		for i, prt := range data {
			switch i {
			case 0:
				st.Designation = prt
			case 1:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("orb: %v", err.Error())
				}
				st.Orbit = fl
			case 2:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("ecc: %v", err.Error())
				}
				st.Eccentrisity = fl
			case 3:
				st.Type, st.SubType, st.Class = disassemble(prt)
			case 4:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("mass: %v", err.Error())
				}
				st.Mass = fl
			case 5:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("diam: %v", err.Error())
				}
				st.Diameter = fl
			case 6:
				fl, err := strconv.ParseFloat(prt, 64)
				if err != nil {
					return nil, fmt.Errorf("luma: %v", err.Error())
				}
				st.Luminocity = fl

			}
		}
	}

	return &st, nil
}

func designations() []string {
	return []string{
		Designation_Rogue,
		Designation_Primary, Designation_Primary_Cmp,
		Designation_Close, Designation_Close_Cmp,
		Designation_Near, Designation_Near_Cmp,
		Designation_Far, Designation_Far_Cmp,
	}
}

//1-G7 V-0.929-0.967-0.738-6.336
//#-T# C-M    -D    -L    -A
//:D-O-E-T# C-M-D-L[-A]

//5-G7 V-0.929-0.967-0.738-6.336:Ab-0.09-0.11-G8 V-0.907-0.957-0.681:B-6.1-0.08-K8 V-0.626-0.777-0.136:Ca-12.1-0.47-M0 V-0.510-0.728-0.0895:Cb-0.21-0.24-D-0.490-0.017-0.000525
//#-T# C-M    -D    -L    -A    :D -O   -E   -T# C-M    -D    -L    :D-O   -E  -T# C-M    -D    -L...

func determine(table []string, i int, dm int) string {
	r := i + dm - 2
	if r < 0 {
		r = 0
	}
	if r > 10 {
		r = 10
	}
	return table[r]
}
