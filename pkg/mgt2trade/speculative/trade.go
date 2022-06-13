package speculative

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/internal/dice"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/tradegoods"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/manifoldco/promptui"
)

/*
Trade Checklist
1. Find a supplier or local broker
2. Determine goods available
3. Determine purchase price
4. Purchase goods
5. Travel to another market
6. Find a buyer or local broker
7. Determine sale price

TODO: потенциальная проблема: купить на одной планете товар можно дешевле чем продать его же на этой же планете
*/

const (
	Market_Common = iota
	Market_Trade
	Market_Legal
	Market_Black
	Market_ALL
	Market_UserChose
	operationPurchase
	operationSale
)

type trader struct {
	brokerSkill int
}

type suplier struct {
	world               World
	market              int
	brokerSkill         int
	previousAttempts    int
	tradeGoodsAvailable map[string]*tradegoods.TradeGood
}

type World interface {
	MW_Name() string
	MW_UWP() string
	TravelZone() string
}

func FindSuplier(world World, suplierType int) (*suplier, error) {
	//chose suplier type if needed
	s := suplier{}
	s.world = world
	s.brokerSkill = 2
	s.tradeGoodsAvailable = make(map[string]*tradegoods.TradeGood)
	switch suplierType {
	default:
		return &s, fmt.Errorf("unknown suplierType '%v'", suplierType)
	case Market_UserChose:
		prompt := promptui.Select{
			Label: "Select Supplier",
			Items: []string{"Common Goods Suplier", "Trade Goods Suplier", "Legal Goods Suplier", "Black Market Suplier"},
		}
		res, result, err := prompt.Run()
		if err != nil {
			return &s, fmt.Errorf("\nprompt.Run() error: %v", err.Error())
		}
		suplierType = res
		fmt.Printf("You choose %q\n", result)
		s.market = suplierType
	case Market_Common, Market_Trade, Market_Black, Market_Legal, Market_ALL:
		s.market = suplierType
	}
	//body
	for _, ctg := range s.apopriateCodes() {
		tg, err := tradegoods.NewTradeGood(ctg)
		if err != nil {
			return &s, fmt.Errorf("tradegoods.NewTradeGood(%v) error: %v", ctg, err)
		}
		wrldTC, err := tradecodes.FromUWPstr(s.world.MW_UWP())
		wrldTC = append(wrldTC, s.world.TravelZone())
		if err != nil {
			return &s, fmt.Errorf("tradeCodes.FromUWPstr(%v) error: %v", s.world.MW_UWP(), err)
		}
		if tradeCodesOverlap(tg.Availability(), wrldTC) {
			s.tradeGoodsAvailable[ctg] = tg
		}
	}
	return &s, nil
}

func (s *suplier) apopriateCodes() []string {
	trGdCodes := []string{}
	switch s.market {
	case Market_Common:
		trGdCodes = tradegoods.CommonMarketCodes()
	case Market_Trade:
		trGdCodes = tradegoods.TradeMarketCodes()
	case Market_Legal:
		trGdCodes = tradegoods.LegalMarketCodes()
	case Market_Black:
		trGdCodes = tradegoods.IllegalMarketCodes()
	case Market_ALL:
		trGdCodes = tradegoods.AllCodes()
	}
	return trGdCodes
}

func (s *suplier) RollQuantity() {
	dm := 0
	dp := dice.New()
	uwpS, _ := uwp.FromString(s.world.MW_UWP())
	pop := uwpS.Pops()
	if pop < 4 {
		dm = -3
	}
	if pop > 8 {
		dm = 3
	}
	for code, tg := range s.tradeGoodsAvailable {
		d, f := tg.Tons()
		tons := dp.Roll(d).DM(dm).Sum() * f
		if tons < 0 {
			tons = 0
		}
		s.tradeGoodsAvailable[code].AddQuantity(tons)
	}
	for i := 0; i < pop; i++ {
		_, code := dp.PickStr(tradegoods.LegalMarketCodes())
		if _, ok := s.tradeGoodsAvailable[code]; !ok {
			newLot, _ := tradegoods.NewTradeGood(code)
			s.tradeGoodsAvailable[code] = newLot
		}
		d, f := s.tradeGoodsAvailable[code].Tons()
		tons := dp.Roll(d).DM(dm).Sum() * f
		if tons < 0 {
			tons = 0
		}
		s.tradeGoodsAvailable[code].AddQuantity(tons)
	}
	for code, tg := range s.tradeGoodsAvailable {
		if tg.Stored() == 0 {
			delete(s.tradeGoodsAvailable, code)
		}
	}
	s.previousAttempts++
}

func tradeCodes(world World) ([]string, error) {
	worldTradeCodes, err := tradecodes.FromUWPstr(world.MW_UWP())
	if err != nil {
		return nil, err
	}
	worldTradeCodes = append(worldTradeCodes, world.TravelZone())
	return worldTradeCodes, nil
}

func DetermineOffer(world World) ([]*tradegoods.TradeGood, error) {
	suggestedGoods := []*tradegoods.TradeGood{}
	worldTradeCodes, err := tradecodes.FromUWPstr(world.MW_UWP())
	if err != nil {
		return nil, err
	}
	worldTradeCodes = append(worldTradeCodes, world.TravelZone())
	for _, code := range tradegoods.AllCodes() {
		if tg, err := tradegoods.NewTradeGood(code); err == nil {
			dm := 0
			for k, v := range tg.PurchaseDM() {
				for i := range worldTradeCodes {
					switch {
					case worldTradeCodes[i] == k:
						dm += v
					}
				}
			}
			if dm > 0 {
				suggestedGoods = append(suggestedGoods, tg)
			}
		} else {
			return nil, fmt.Errorf("tradegoods.NewTradeGood(code) error: %v", err.Error())
		}

	}
	return suggestedGoods, nil
}

func DetermineDemand(world World) ([]*tradegoods.TradeGood, error) {
	requestedGoods := []*tradegoods.TradeGood{}
	worldTradeCodes, err := tradecodes.FromUWPstr(world.MW_UWP())
	if err != nil {
		return nil, err
	}
	worldTradeCodes = append(worldTradeCodes, world.TravelZone())
	for _, code := range tradegoods.AllCodes() {
		if tg, err := tradegoods.NewTradeGood(code); err == nil {
			dm := 0
			for k, v := range tg.SaleDM() {
				for i := range worldTradeCodes {
					switch {
					case worldTradeCodes[i] == k:
						dm += v
					}
				}
			}
			if dm > 0 {
				requestedGoods = append(requestedGoods, tg)
			}
		} else {
			return nil, fmt.Errorf("tradegoods.NewTradeGood(code) error: %v", err.Error())
		}

	}
	return requestedGoods, nil
}

func DetermineGoodsAvailable(world World) ([]*tradegoods.TradeGood, error) {
	availableGoods := []*tradegoods.TradeGood{}
	worldTradeCodes, err := tradecodes.FromUWPstr(world.MW_UWP())
	if err != nil {
		return nil, err
	}
	worldTradeCodes = append(worldTradeCodes, world.TravelZone())
	for _, code := range tradegoods.AllCodes() {
		if tg, err := tradegoods.NewTradeGood(code); err == nil {
			if tradeCodesOverlap(tg.Availability(), worldTradeCodes) {
				availableGoods = append(availableGoods, tg)
			}
		}
	}
	return availableGoods, nil
}

func TradeGoodsFlow(origin, destination World) ([]*tradegoods.TradeGood, error) {
	offerList, err := DetermineGoodsAvailable(origin)
	if err != nil {
		return nil, err
	}
	demandList, err := DetermineDemand(destination)
	if err != nil {
		return nil, err
	}
	tradeList := []*tradegoods.TradeGood{}
	for _, tgOffer := range offerList {
		for _, tgDemand := range demandList {
			if tgDemand.GoodsType() == tgOffer.GoodsType() {
				tradeList = append(tradeList, tgOffer)
			}
		}
	}
	return tradeList, nil
}

/*
Asim ---> Drinax
TG1
Tg2
...
TGn


*/

func tradeCodesOverlap(sl1, sl2 []string) bool {
	for _, s1 := range sl1 {
		for _, s2 := range sl2 {
			if s1 == s2 {
				return true
			}
			if s1 == "All" || s2 == "All" {
				return true
			}
		}
	}
	return false
}

func PricePurchaseAverage(wrld World, tg *tradegoods.TradeGood) int {
	tc, _ := tradeCodes(wrld)

	purchDM := maxDM(tg.PurchaseDM(), tc)
	saleDM := maxDM(tg.SaleDM(), tc)
	dm := purchDM - saleDM
	index1 := dm + 10
	index2 := dm + 11
	price1, _ := modifiedPrice(operationPurchase, index1, tg.BasePrice())
	price2, _ := modifiedPrice(operationPurchase, index2, tg.BasePrice())
	return (price1 + price2) / 2
}

func PriceSaleAverage(wrld World, tg *tradegoods.TradeGood) int {
	tc, _ := tradeCodes(wrld)
	purchDM := maxDM(tg.PurchaseDM(), tc)
	saleDM := maxDM(tg.SaleDM(), tc)
	dm := saleDM - purchDM
	index1 := dm + 10
	index2 := dm + 11
	price1, _ := modifiedPrice(operationSale, index1, tg.BasePrice())
	price2, _ := modifiedPrice(operationSale, index2, tg.BasePrice())
	return (price1 + price2) / 2
}

func ListPrices(w World) map[string][]int {
	s, _ := FindSuplier(w, Market_ALL)
	priceMap := make(map[string][]int)
	for _, code := range tradegoods.AllCodes() {
		s.tradeGoodsAvailable[code], _ = tradegoods.NewTradeGood(code)
		priceMap[code] = []int{s.tradeGoodsAvailable[code].BasePrice(), PricePurchaseAverage(s.world, s.tradeGoodsAvailable[code]), PriceSaleAverage(s.world, s.tradeGoodsAvailable[code])}
	}
	return priceMap
}

func modifiedPrice(operation, index, basePrise int) (int, error) {
	mod := []int{}
	modMap := make(map[int][]int)
	modMap[-2] = []int{250, 20}
	modMap[-1] = []int{200, 30}
	modMap[0] = []int{175, 40}
	modMap[1] = []int{150, 45}
	modMap[2] = []int{135, 50}
	modMap[3] = []int{125, 55}
	modMap[4] = []int{120, 60}
	modMap[5] = []int{115, 65}
	modMap[6] = []int{110, 70}
	modMap[7] = []int{105, 75}
	modMap[8] = []int{100, 80}
	modMap[9] = []int{95, 85}
	modMap[10] = []int{90, 90}
	modMap[11] = []int{85, 100}
	modMap[12] = []int{80, 105}
	modMap[13] = []int{75, 110}
	modMap[14] = []int{70, 115}
	modMap[15] = []int{65, 120}
	modMap[16] = []int{60, 125}
	modMap[17] = []int{55, 130}
	modMap[18] = []int{50, 140}
	modMap[19] = []int{45, 150}
	modMap[20] = []int{40, 160}
	modMap[21] = []int{35, 175}
	modMap[22] = []int{30, 200}
	modMap[23] = []int{25, 250}
	modMap[24] = []int{20, 300}
	mod = modMap[index]
	if index <= -3 {
		mod = []int{300, 10}
	}
	if index >= 25 {
		mod = []int{15, 400}
	}
	pct := 0
	switch operation {
	default:
		return 0, fmt.Errorf("modifiedPrice(): unknown operation '%v'", operation)
	case operationPurchase:
		pct = mod[0]
	case operationSale:
		pct = mod[1]
	}
	return (basePrise * pct) / 100, nil

}

func maxDM(dmMap map[string]int, tags []string) int {
	//matched := []int{}
	dm := 0
	for _, tc := range tags {
		if val, ok := dmMap[tc]; ok {
			//matched = append(matched, val)
			dm += val
		}
	}
	//return maxFromSl(matched)
	return dm
}

func maxFromSl(sl []int) int {
	if len(sl) == 0 {
		return 0
	}
	m := -999999
	for _, i := range sl {
		if i > m {
			m = i
		}
	}
	return m
}
