package ehex

import (
	"strings"
)

const (
	UNKNOWN   = "unknown"
	SPECIAL   = "special"
	ULTIMATE  = "ultimate"
	ANY_VALUE = "any value"
)

type ehex struct {
	value   int
	code    string
	comment string
}

func New() *ehex {
	return &ehex{value: 0, code: "*", comment: ANY_VALUE}
}

func SetValue(data interface{}) *ehex {
	eh := New()
	eh.comment = "not assigned"
	eh.code = "no code"
	eh.value = -999
	switch data.(type) {
	default:
		eh.code = "unknown data type"
		eh.comment = "unknown data type"
	case string:
		eh.code = setStringCode(data.(string))
	case int:
		eh.code = hashValue(data.(int))
	}
	eh.value = hashCode(eh.code)
	eh.comment = defaultComment(eh.code)
	return eh
}

func (eh *ehex) Set(data interface{}) *ehex {
	eh.comment = "not assigned"
	eh.code = "no code"
	eh.value = -999
	switch data.(type) {
	default:
		eh.code = "unknown data type"
		eh.comment = "unknown data type"
	case string:
		eh.code = setStringCode(data.(string))
	case int:
		eh.code = hashValue(data.(int))
	}
	eh.value = hashCode(eh.code)
	eh.comment = defaultComment(eh.code)
	return eh
}

func defaultComment(code string) string {
	switch code {
	case "X", "?":
		return UNKNOWN
	case "Y":
		return SPECIAL
	case "Z":
		return ULTIMATE
	case "*":
		return ANY_VALUE
	default:
	}
	return ""
}

func setStringCode(code string) string {
	tryCode := strings.ToUpper(code)
	if hashCode(tryCode) != -1 {
		return tryCode
	}
	return ""
}

func (eh *ehex) Encode(meaning string) {
	eh.comment = meaning
}

// /////INTERFACE
type Ehex interface {
	Value() int
	Code() string
	Meaning() string
}

func (e *ehex) Value() int {
	return e.value
}

func (e *ehex) Code() string {
	return e.code
}

func (e *ehex) Meaning() string {
	return e.comment
}

func (e *ehex) String() string {
	return e.code
}

///////HASH

func hashValue(value int) string {
	codeMap := make(map[int]string)
	codeMap[-1001] = "-"
	codeMap[-3] = "*"
	codeMap[-2] = "?"
	codeMap[0] = "0"
	codeMap[1] = "1"
	codeMap[2] = "2"
	codeMap[3] = "3"
	codeMap[4] = "4"
	codeMap[5] = "5"
	codeMap[6] = "6"
	codeMap[7] = "7"
	codeMap[8] = "8"
	codeMap[9] = "9"
	codeMap[10] = "A"
	codeMap[11] = "B"
	codeMap[12] = "C"
	codeMap[13] = "D"
	codeMap[14] = "E"
	codeMap[15] = "F"
	codeMap[16] = "G"
	codeMap[17] = "H"
	codeMap[18] = "J"
	codeMap[19] = "K"
	codeMap[20] = "L"
	codeMap[21] = "M"
	codeMap[22] = "N"
	codeMap[23] = "P"
	codeMap[24] = "Q"
	codeMap[25] = "R"
	codeMap[26] = "S"
	codeMap[27] = "T"
	codeMap[28] = "U"
	codeMap[29] = "V"
	codeMap[30] = "W"
	codeMap[31] = "X"
	codeMap[32] = "Y"
	codeMap[33] = "Z"
	codeMap[34] = "%"
	codeMap[35] = "@"
	codeMap[36] = "#"
	codeMap[37] = "$"
	codeMap[38] = "y" //
	codeMap[39] = "~"
	codeMap[40] = "_"
	codeMap[41] = "$"
	codeMap[42] = "a"
	codeMap[43] = "b"
	codeMap[44] = "c"
	codeMap[45] = "d"
	codeMap[46] = "e"
	codeMap[47] = "f"
	codeMap[48] = "g"
	codeMap[49] = "h"
	codeMap[50] = "j"
	codeMap[51] = "k"
	codeMap[52] = "z" //
	codeMap[53] = "m"
	codeMap[54] = "n"
	codeMap[55] = "p"
	codeMap[56] = "q"
	codeMap[57] = "r"
	codeMap[58] = "s"
	codeMap[59] = "t"
	codeMap[60] = "u"
	codeMap[61] = "v"
	codeMap[62] = "w"
	codeMap[63] = "x"
	if val, ok := codeMap[value]; ok {
		return val
	}
	return "?"
}

func hashCode(code string) int {
	// valMap := make(map[string]int)
	// valMap["0"] = 0
	// valMap["1"] = 1
	// valMap["2"] = 2
	// valMap["3"] = 3
	// valMap["4"] = 4
	// valMap["5"] = 5
	// valMap["6"] = 6
	// valMap["7"] = 7
	// valMap["8"] = 8
	// valMap["9"] = 9
	// valMap["A"] = 10
	// valMap["B"] = 11
	// valMap["C"] = 12
	// valMap["D"] = 13
	// valMap["E"] = 14
	// valMap["F"] = 15
	// valMap["G"] = 16
	// valMap["H"] = 17
	// valMap["J"] = 18
	// valMap["K"] = 19
	// valMap["L"] = 20
	// valMap["M"] = 21
	// valMap["N"] = 22
	// valMap["P"] = 23
	// valMap["Q"] = 24
	// valMap["R"] = 25
	// valMap["S"] = 26
	// valMap["T"] = 27
	// valMap["U"] = 28
	// valMap["V"] = 29
	// valMap["W"] = 30
	// valMap["X"] = 31
	// valMap["Y"] = 32
	// valMap["Z"] = 33
	// valMap["?"] = -2
	// valMap["*"] = -3
	// valMap["-"] = -1001
	valueMap := make(map[string]int)
	valueMap["-"] = -1001
	valueMap["*"] = -3
	valueMap["?"] = -2
	valueMap["0"] = 0
	valueMap["1"] = 1
	valueMap["2"] = 2
	valueMap["3"] = 3
	valueMap["4"] = 4
	valueMap["5"] = 5
	valueMap["6"] = 6
	valueMap["7"] = 7
	valueMap["8"] = 8
	valueMap["9"] = 9
	valueMap["A"] = 10
	valueMap["B"] = 11
	valueMap["C"] = 12
	valueMap["D"] = 13
	valueMap["E"] = 14
	valueMap["F"] = 15
	valueMap["G"] = 16
	valueMap["H"] = 17
	valueMap["J"] = 18
	valueMap["K"] = 19
	valueMap["L"] = 20
	valueMap["M"] = 21
	valueMap["N"] = 22
	valueMap["P"] = 23
	valueMap["Q"] = 24
	valueMap["R"] = 25
	valueMap["S"] = 26
	valueMap["T"] = 27
	valueMap["U"] = 28
	valueMap["V"] = 29
	valueMap["W"] = 30
	valueMap["X"] = 31
	valueMap["Y"] = 32
	valueMap["Z"] = 33
	valueMap["%"] = 34
	valueMap["@"] = 35
	valueMap["#"] = 36
	valueMap["$"] = 37
	valueMap["y"] = 38
	valueMap["~"] = 39
	valueMap["_"] = 40
	valueMap["$"] = 41
	valueMap["a"] = 42
	valueMap["b"] = 43
	valueMap["c"] = 44
	valueMap["d"] = 45
	valueMap["e"] = 46
	valueMap["f"] = 47
	valueMap["g"] = 48
	valueMap["h"] = 49
	valueMap["j"] = 50
	valueMap["k"] = 51
	valueMap["z"] = 52
	valueMap["m"] = 53
	valueMap["n"] = 54
	valueMap["p"] = 55
	valueMap["q"] = 56
	valueMap["r"] = 57
	valueMap["s"] = 58
	valueMap["t"] = 59
	valueMap["u"] = 60
	valueMap["v"] = 61
	valueMap["w"] = 62
	valueMap["x"] = 63

	if val, ok := valueMap[code]; ok {
		return val
	}
	return -1
}

func ToCode(i int) string {
	return New().Set(i).Code()
}

func ValueOf(s string) int {
	return New().Set(s).Value()
}
