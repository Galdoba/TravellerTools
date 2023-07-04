package ehex

import (
	"strings"
)

const (
	UNKNOWN   = "unknown"
	SPECIAL   = "special"
	ULTIMATE  = "ultimate"
	ANY_VALUE = "any value"
	UNSET     = "[No Data]"
)

type ehex struct {
	value     int
	code      string
	comment   string
	Container interface{}
	ContVal   func() *int
}

func New() *ehex {
	return &ehex{
		value:     0,
		code:      "0",
		comment:   UNSET,
		Container: nil,
		ContVal: func() *int {
			return nil
		},
	}
}

func (e *ehex) Put(i interface{}) Ehex {
	e.Container = i
	return e
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
	SetContainedValFunc(func() *int)
	ContainedVal() int
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

func (e *ehex) ContainedVal() int {
	if e.ContVal == nil {
		return e.value
	}
	return *e.ContVal()
}

func (e *ehex) SetContainedValFunc(f func() *int) {
	e.ContVal = f
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
	codeMap[34] = "Δ"
	codeMap[35] = "Λ"
	codeMap[36] = "Ξ"
	codeMap[37] = "Σ"
	codeMap[38] = "Φ"
	codeMap[39] = "Ψ"
	codeMap[40] = "Ω"
	codeMap[41] = "Ϣ"
	codeMap[42] = "Ϫ"
	codeMap[43] = "Ћ"
	codeMap[44] = "Џ"
	codeMap[45] = "Ы"
	codeMap[46] = "Э"
	codeMap[47] = "Ю"
	codeMap[48] = "Ѥ"
	codeMap[49] = "Ҧ"
	codeMap[50] = "Բ"
	codeMap[51] = "Գ"
	codeMap[52] = "Դ"
	codeMap[53] = "Ե"
	codeMap[54] = "Ը"
	codeMap[55] = "Թ"
	codeMap[56] = "Ժ"
	codeMap[57] = "Ի"
	codeMap[58] = "Հ"
	codeMap[59] = "Ն"
	codeMap[60] = "Ո"
	codeMap[61] = "Վ"
	codeMap[62] = "Ւ"
	codeMap[63] = "Ֆ"
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
	valueMap["Δ"] = 34
	valueMap["Λ"] = 35
	valueMap["Ξ"] = 36
	valueMap["Σ"] = 37
	valueMap["Φ"] = 38
	valueMap["Ψ"] = 39
	valueMap["Ω"] = 40
	valueMap["Ϣ"] = 41
	valueMap["Ϫ"] = 42
	valueMap["Ћ"] = 43
	valueMap["Џ"] = 44
	valueMap["Ы"] = 45
	valueMap["Э"] = 46
	valueMap["Ю"] = 47
	valueMap["Ѥ"] = 48
	valueMap["Ҧ"] = 49
	valueMap["Բ"] = 50
	valueMap["Գ"] = 51
	valueMap["Դ"] = 52
	valueMap["Ե"] = 53
	valueMap["Ը"] = 54
	valueMap["Թ"] = 55
	valueMap["Ժ"] = 56
	valueMap["Ի"] = 57
	valueMap["Հ"] = 58
	valueMap["Ն"] = 59
	valueMap["Ո"] = 60
	valueMap["Վ"] = 61
	valueMap["Ւ"] = 62
	valueMap["Ֆ"] = 63

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

/*
0
1
2
3
4
5
6
7
8
9
A
B
C
D
E
F
G
H
J
K
L
M
N
P
Q
R
S
T
U
V
W
X
Y
Z
Δ
Λ
Ξ
Σ
Φ
Ψ
Ω
Ϣ
Ϫ
Ћ
Џ
Ы
Э
Ю
Ѥ
Ҧ
Բ
Գ
Դ
Ե
Ը
Թ
Ժ
Ի
Հ
Ն
Ո
Վ
Ւ
Ֆ

///////////////
Բ Գ Դ Ե Ը Թ Ժ Ի
Ⴀ Ⴁ Ⴂ Ⴃ Ⴄ Ⴅ Ⴆ Ⴇ Ⴈ Ⴉ Ⴊ Ⴋ Ⴌ Ⴍ Ⴎ Ⴏ
ㄅ ㄆ ㄇ ㄈ ㄉ ㄊ
ㄋ ㄌ ㄍ ㄎ ㄏ ㄐ
ㄑ ㄓ ㄔ ㄕ ㄖ ㄙ
ㄜ ㄝ ㄞ ㄟ ㄠ ㄡ
ㄢ ㄤ ㄥ ㄦ ㄩ ㄪ
ㄫ ㄭ


Y Z % @ # $ y ~
_ $ a b c d e f
g h j k z m n p
q r s t u v w x
*/
