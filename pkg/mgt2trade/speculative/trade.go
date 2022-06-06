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
*/

const (
	Market_Common = iota
	Market_Trade
	Market_Legal
	Market_Black
	Market_ALL
	Market_UserChose
)

type suplier struct {
	world               World
	market              int
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

func DetermineExport(world World) ([]*tradegoods.TradeGood, error) {
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
