package tradegoods

/*
0123  4  5  6  7  8  9ABCDEF
-3   -2 -1  0  1  2  3
*/

type tradeGood struct {
	code       string
	goodsType  string
	available  []string
	tonsDice   int
	tonsFactor int
	basePrice  int
	purchaseDM map[string]int
	saleDM     map[string]int
	example    string
}

func goodsMap() []tradeGood {
	tgList := []tradeGood{}
	for _, code := range []string{
		"11", "12", "13", "14", "15", "16",
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
		"61", "62", "63", "64", "65", "66",
	} {
		switch code {
		case "11":
			tgList = append(tgList, tradeGood{
				code:       code,
				goodsType:  "Common Electrinics",
				available:  []string{"All"},
				tonsDice:   2,
				tonsFactor: 10,
				basePrice:  20000,
				purchaseDM: map[string]int{"In": 2, "Ht": 3, "Ri": 1},
				saleDM:     map[string]int{"Ni": 2, "Lt": 1, "Po": 1},
				example:    "Simple electronics including basic computers up to TL10",
			})
		case "12":
			tgList = append(tgList, tradeGood{
				code:       code,
				goodsType:  "Common Industrial Goods",
				available:  []string{"All"},
				tonsDice:   2,
				tonsFactor: 10,
				basePrice:  10000,
				purchaseDM: map[string]int{"Na": 2, "In": 5},
				saleDM:     map[string]int{"Ni": 3, "Ag": 2},
				example:    "Machine components and spare parts for common machinery",
			})
		}
	}

	return tgList
}
