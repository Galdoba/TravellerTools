package profile

/*
World. A planet or satellite.
-----------------------------
Profile Keys:
Starport [ABCDEFGHXY]
Size [0-K]
Atmo [0-F]
Hydr [0-A]
HZvar [0-F] h
SateliteCode [0-2] h
SateliteOrb [0-F] h
Climate [0-6] h
Pops [0-F]
Govr [0-F]
Laws [0-J]
Tech [0-N] s
PopDigit [0-9] h
OtherBelts [0-3] h
GasG [0-4] h
LifeFactor [0-A] h
LifeCompatability [0-A] h (количество совпадений в геноме рандомного софонта с человеком минус 2)


Aralia (100-201 Seigor AB III 1d)

*/

const (
	KEY_MAINWORLD          = "MW"
	KEY_PORT               = "Starport"
	KEY_SIZE               = "Size"
	KEY_ATMO               = "Atmo"
	KEY_HYDR               = "Hydr"
	KEY_POPS               = "Pops"
	KEY_GOVR               = "Govr"
	KEY_LAWS               = "Laws"
	KEY_SEP                = "separator"
	KEY_TL                 = "Tech"
	KEY_WORLDTYPE          = "World Type"
	KEY_HABITABLE_ZONE_VAR = "HZvar"
	KEY_PLANETARY_ORBIT    = "Orbit"
	KEY_IS_SATELITE        = "Satelite?"
	KEY_SATELITE_ORBIT     = "Satelite Orbit"
	KEY_CLIMATE            = "Climate"
	KEY_LIFE_FACTOR        = "Life"
	KEY_LIFE_COMPATABILITY = "Life Compatability"
	KEY_POP_DIGIT          = "PopDigit"
	KEY_BELTS              = "Belts"
	KEY_GAS_GIANTS         = "Gigants"
	KEY_LIMIT_size         = "LIMIT_Size"
	KEY_LIMIT_pops         = "LIMIT_Pops"
	KEY_LIMIT_tl           = "LIMIT_tl"
	KEY_BASES              = "Bases"
)

func UWPkeys() []string {
	return []string{
		KEY_MAINWORLD,
		KEY_PLANETARY_ORBIT,
		KEY_SATELITE_ORBIT,
		KEY_HABITABLE_ZONE_VAR,
		KEY_WORLDTYPE,
		KEY_LIMIT_size,
		KEY_SIZE,
		KEY_ATMO,
		KEY_HYDR,
		KEY_LIFE_FACTOR,
		KEY_LIFE_COMPATABILITY,
		KEY_LIMIT_pops,
		KEY_POPS,

		KEY_GOVR,
		KEY_LAWS,
		KEY_LIMIT_tl,
		KEY_PORT,
		KEY_TL,

		KEY_POP_DIGIT,
		KEY_BASES,
	}
}

func UWPbasic() []string {
	return []string{
		KEY_PLANETARY_ORBIT,
		KEY_SATELITE_ORBIT,
		KEY_HABITABLE_ZONE_VAR,
		KEY_WORLDTYPE,
		KEY_LIMIT_size,
		KEY_SIZE,
		KEY_ATMO,
		KEY_HYDR,
		KEY_LIFE_FACTOR,
		KEY_LIFE_COMPATABILITY,
	}
}

func UWP(profile Profile) string {
	str := ""
	for _, key := range []string{KEY_PORT, KEY_SIZE, KEY_ATMO, KEY_HYDR, KEY_POPS, KEY_GOVR, KEY_LAWS, KEY_TL} {
		if key == KEY_TL {
			str += "-"
		}
		if val := profile.Data(key); val != nil {
			str += val.Code()
		} else {
			str += "?"
		}
	}
	return str
}
