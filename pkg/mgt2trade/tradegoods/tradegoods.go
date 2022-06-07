package tradegoods

import (
	"fmt"
	"strconv"
)

/*
0123  4  5  6  7  8  9ABCDEF
-3   -2 -1  0  1  2  3
*/

type Source interface {
}

type TradeGood struct {
	code         string
	goodsType    string
	availability []string
	tonsDice     int
	tonsFactor   int
	basePrice    int
	purchaseDM   map[string]int
	saleDM       map[string]int
	example      string
	storedTons   int
}

func NewTradeGood(code string) (*TradeGood, error) {

	for _, tg := range goodsMap() {
		if tg.code == code {
			return &tg, nil
		}
	}
	return nil, fmt.Errorf("invalid code '%v'", code)
}

func AllCodes() []string {
	return []string{
		"11", "12", "13", "14", "15", "16",
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
		"61", "62", "63", "64", "65", "66",
	}
}

func CommonMarketCodes() []string {
	return []string{
		"11", "12", "13", "14", "15", "16",
	}
}

func TradeMarketCodes() []string {
	return []string{
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
	}
}

func LegalMarketCodes() []string {
	return []string{
		"11", "12", "13", "14", "15", "16",
		"21", "22", "23", "24", "25", "26",
		"31", "32", "33", "34", "35", "36",
		"41", "42", "43", "44", "45", "46",
		"51", "52", "53", "54", "55", "56",
	}
}

func IllegalMarketCodes() []string {
	return []string{
		"61", "62", "63", "64", "65", "66",
	}
}

func goodsMap() []TradeGood {
	tgList := []TradeGood{}
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
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Common Electrinics",
				availability: []string{"All"},
				tonsDice:     2,
				tonsFactor:   10,
				basePrice:    20000,
				purchaseDM:   map[string]int{"In": 2, "Ht": 3, "Ri": 1},
				saleDM:       map[string]int{"Ni": 2, "Lt": 1, "Po": 1},
				example:      "Simple electronics including basic computers up to TL10",
			})
		case "12":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Common Industrial Goods",
				availability: []string{"All"},
				tonsDice:     2,
				tonsFactor:   10,
				basePrice:    10000,
				purchaseDM:   map[string]int{"Na": 2, "In": 5},
				saleDM:       map[string]int{"Ni": 3, "Ag": 2},
				example:      "Machine components and spare parts for common machinery",
			})
		case "13":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Common Manufactured Goods",
				availability: []string{"All"},
				tonsDice:     2,
				tonsFactor:   10,
				basePrice:    20000,
				purchaseDM:   map[string]int{"Na": 2, "In": 5},
				saleDM:       map[string]int{"Ni": 3, "Hi": 2},
				example:      "Household appliances, clothing and so forth",
			})
		case "14":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Common Raw Materials",
				availability: []string{"All"},
				tonsDice:     2,
				tonsFactor:   20,
				basePrice:    5000,
				purchaseDM:   map[string]int{"Ag": 3, "Ga": 2},
				saleDM:       map[string]int{"In": 2, "Po": 2},
				example:      "Metal, plastics, chemicals and other basic materials",
			})
		case "15":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Common Consumables",
				availability: []string{"All"},
				tonsDice:     2,
				tonsFactor:   20,
				basePrice:    500,
				purchaseDM:   map[string]int{"Ag": 3, "Wa": 2, "Ga": 1, "As": -4},
				saleDM:       map[string]int{"As": 1, "Fl": 1, "Ic": 1, "Hi": 1},
				example:      "Food, drink and other agricultural products",
			})
		case "16":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Common Ore",
				availability: []string{"All"},
				tonsDice:     2,
				tonsFactor:   20,
				basePrice:    1000,
				purchaseDM:   map[string]int{"As": 4},
				saleDM:       map[string]int{"In": 3, "Na": -1},
				example:      "Ore bearing common metals",
			})
		case "21":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Advanced Electronics",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    100000,
				purchaseDM:   map[string]int{"In": 2, "Ht": 3},
				saleDM:       map[string]int{"Ni": 1, "Ri": 2, "As": 3},
				example:      "Advanced sensors, computers and other electronics up to TL15",
			})
		case "22":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Advanced Machine Parts",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    75000,
				purchaseDM:   map[string]int{"In": 2, "Ht": 1},
				saleDM:       map[string]int{"As": 2, "Ni": 1},
				example:      "Machine components and spare parts, including gravitic components",
			})
		case "23":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Advanced Manufactured Goods",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    100000,
				purchaseDM:   map[string]int{"In": 1},
				saleDM:       map[string]int{"Hi": 1, "Ri": 2},
				example:      "Devices and clothing incorporating advanced technologies",
			})
		case "24":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Advanced Weapons",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    150000,
				purchaseDM:   map[string]int{"Ht": 2},
				saleDM:       map[string]int{"Po": 1, "A": 2, "R": 4},
				example:      "Firearms, explosives, ammunition, artillery and other military-grade weaponry",
			})
		case "25":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Advanced Vehicles",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    180000,
				purchaseDM:   map[string]int{"Ht": 2},
				saleDM:       map[string]int{"As": 2, "Ri": 2},
				example:      "Air/rafts, spacecraft, grav tanks and other vehicles up to TL15",
			})
		case "26":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Biochemicals",
				availability: []string{"Ag", "Wa"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    50000,
				purchaseDM:   map[string]int{"Ag": 1, "Wa": 2},
				saleDM:       map[string]int{"In": 2},
				example:      "Biofuels, organic chemicals, extracts",
			})
		case "31":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Crystals & Gems",
				availability: []string{"As", "De", "Ic"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    20000,
				purchaseDM:   map[string]int{"As": 2, "De": 1, "Ic": 1},
				saleDM:       map[string]int{"In": 3, "Ri": 2},
				example:      "Diamonds, synthetic or natural gemstones",
			})
		case "32":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Cybernetics",
				availability: []string{"Ht"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    250000,
				purchaseDM:   map[string]int{"Ht": 1},
				saleDM:       map[string]int{"As": 1, "Ic": 1, "Ri": 2},
				example:      "Cybernetic components, replacement limbs",
			})
		case "33":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Live Animals",
				availability: []string{"Ag", "Ga"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    10000,
				purchaseDM:   map[string]int{"Ag": 2},
				saleDM:       map[string]int{"Lo": 3},
				example:      "Riding animals, beasts of burden, exotic pets",
			})
		case "34":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Luxury Consumables",
				availability: []string{"Ag", "Ga", "Wa"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    20000,
				purchaseDM:   map[string]int{"Ag": 2, "Wa": 1},
				saleDM:       map[string]int{"Ri": 2, "Hi": 2},
				example:      "Rare foods, fine liquors",
			})
		case "35":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Luxury Goods",
				availability: []string{"Hi"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    200000,
				purchaseDM:   map[string]int{"Hi": 1},
				saleDM:       map[string]int{"Ri": 4},
				example:      "Rare or extremely high-quality manufactured goods",
			})
		case "36":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Medical Supplies",
				availability: []string{"Ht", "Hi"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    50000,
				purchaseDM:   map[string]int{"Ht": 2},
				saleDM:       map[string]int{"In": 2, "Po": 1, "Ri": 1},
				example:      "Diagnostic equipment, basic drugs, cloning technology",
			})
		case "41":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Petrochemicals",
				availability: []string{"De", "Fl", "Ic", "Wa"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    10000,
				purchaseDM:   map[string]int{"De": 2},
				saleDM:       map[string]int{"In": 2, "Ag": 1, "Lt": 2},
				example:      "Oil, liquid fuels",
			})
		case "42":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Pharmaceuticals",
				availability: []string{"As", "De", "Hi", "Wa"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    100000,
				purchaseDM:   map[string]int{"As": 2, "Hi": 1},
				saleDM:       map[string]int{"Ri": 2, "Lt": 1},
				example:      "Drugs, medical supplies, anagathics, fast or slow drugs",
			})
		case "43":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Polymers",
				availability: []string{"In"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    7000,
				purchaseDM:   map[string]int{"In": 1},
				saleDM:       map[string]int{"Ri": 2, "Ni": 1},
				example:      "Plastics and other synthetics",
			})
		case "44":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Precious Metals",
				availability: []string{"As", "De", "Ic", "Fl"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    50000,
				purchaseDM:   map[string]int{"As": 3, "De": 1, "Ic": 2},
				saleDM:       map[string]int{"Ri": 3, "In": 2, "Ht": 1},
				example:      "Gold, silver, platinum, rare elements",
			})
		case "45":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Radioactives",
				availability: []string{"As", "De", "Lo"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    1000000,
				purchaseDM:   map[string]int{"As": 2, "Lo": 2},
				saleDM:       map[string]int{"In": 3, "Ht": 1, "Ni": -2, "Ag": -3},
				example:      "Uranium, plutonium, unobtanium, rare elements",
			})
		case "46":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Robots",
				availability: []string{"In"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    400000,
				purchaseDM:   map[string]int{"In": 1},
				saleDM:       map[string]int{"Ag": 2, "Ht": 1},
				example:      "Industial and personal robots and drones",
			})
		case "51":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Spices",
				availability: []string{"Ga", "De", "Wa"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    6000,
				purchaseDM:   map[string]int{"De": 2},
				saleDM:       map[string]int{"Hi": 2, "Ri": 3, "Po": 3},
				example:      "Preservatives, luxury food additives, natural drugs",
			})
		case "52":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Textiles",
				availability: []string{"Ag", "Ni"},
				tonsDice:     1,
				tonsFactor:   20,
				basePrice:    3000,
				purchaseDM:   map[string]int{"Ag": 7},
				saleDM:       map[string]int{"Hi": 3, "Na": 2},
				example:      "Clothing and fabrics",
			})
		case "53":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Uncommon Ore",
				availability: []string{"As", "Ic"},
				tonsDice:     1,
				tonsFactor:   20,
				basePrice:    5000,
				purchaseDM:   map[string]int{"As": 4},
				saleDM:       map[string]int{"In": 3, "Ni": 1},
				example:      "Ore containing precious or valuable metals",
			})
		case "54":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Uncommon Raw Materials",
				availability: []string{"Ag", "De", "Wa"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    20000,
				purchaseDM:   map[string]int{"Ag": 2, "Wa": 1},
				saleDM:       map[string]int{"In": 2, "Ht": 1},
				example:      "Valuable metals like titanium, rare elements",
			})
		case "55":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Wood",
				availability: []string{"Ag", "Ga"},
				tonsDice:     1,
				tonsFactor:   20,
				basePrice:    1000,
				purchaseDM:   map[string]int{"Ag": 6},
				saleDM:       map[string]int{"Ri": 2, "In": 1},
				example:      "Hard or beautiful woods and plant extracts",
			})
		case "56":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Vehicles",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   10,
				basePrice:    15000,
				purchaseDM:   map[string]int{"In": 2, "Ht": 1},
				saleDM:       map[string]int{"Ni": 2, "Hi": 1},
				example:      "Wheeled, tracked and other vehicles from TL10 or lower",
			})
		case "61":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Biochemicals (Illegal)",
				availability: []string{"Ag", "Wa"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    50000,
				purchaseDM:   map[string]int{"Wa": 2},
				saleDM:       map[string]int{"In": 6},
				example:      "Dangerous chemicals, extracts from endangered species",
			})
		case "62":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Cybernetics (Illegal)",
				availability: []string{"Ht"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    250000,
				purchaseDM:   map[string]int{"Ht": 1},
				saleDM:       map[string]int{"As": 4, "Ic": 4, "Ri": 8, "A": 6, "R": 6},
				example:      "Combat cybernetics, illegal enhancements",
			})
		case "63":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Drugs (Illegal)",
				availability: []string{"As", "De", "Hi", "Wa"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    100000,
				purchaseDM:   map[string]int{"As": 0, "De": 1, "Ga": 1, "Wa": 1},
				saleDM:       map[string]int{"Ri": 6, "Hi": 6},
				example:      "Addictive drugs, combat drugs",
			})
		case "64":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Luxuries (Illegal)",
				availability: []string{"Ag", "Ga", "Wa"},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    50000,
				purchaseDM:   map[string]int{"Ag": 2, "Wa": 1},
				saleDM:       map[string]int{"Ri": 6, "Hi": 4},
				example:      "Debauched or addictive luxuries",
			})
		case "65":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Weapons (Illegal)",
				availability: []string{"In", "Ht"},
				tonsDice:     1,
				tonsFactor:   5,
				basePrice:    150000,
				purchaseDM:   map[string]int{"Ht": 2},
				saleDM:       map[string]int{"Po": 6, "A": 8, "R": 10},
				example:      "Weapons of mass destruction, naval weapons",
			})
		case "66":
			tgList = append(tgList, TradeGood{
				code:         code,
				goodsType:    "Exotics",
				availability: []string{},
				tonsDice:     1,
				tonsFactor:   1,
				basePrice:    1,
				purchaseDM:   map[string]int{},
				saleDM:       map[string]int{},
				example:      "Alien relics, prototype technology, unique plant or animal life, priceless treasures and so forth",
			})
		}
	}

	return tgList
}

func (tg *TradeGood) Code() string {
	return tg.code
}
func (tg *TradeGood) GoodsType() string {
	return tg.goodsType
}
func (tg *TradeGood) Availability() []string {
	return tg.availability
}
func (tg *TradeGood) Tons() (string, int) {
	return strconv.Itoa(tg.tonsDice) + "d6", tg.tonsFactor
}
func (tg *TradeGood) BasePrice() int {
	return tg.basePrice
}
func (tg *TradeGood) PurchaseDM() map[string]int {
	return tg.purchaseDM
}
func (tg *TradeGood) SaleDM() map[string]int {
	return tg.saleDM
}
func (tg *TradeGood) Example() string {
	return tg.example
}
func (tg *TradeGood) Stored() int {
	return tg.storedTons
}

func (tg *TradeGood) AddQuantity(tons int) {
	tg.storedTons = tg.storedTons + tons
	if tg.storedTons < 0 {
		tg.storedTons = 0
	}
}
