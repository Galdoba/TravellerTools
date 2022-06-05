package speculative

import (
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/tradegoods"
	"github.com/Galdoba/TravellerTools/pkg/mgt2trade/traffic/tradecodes"
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
	Market_Black
)

type suplier struct {
	world            World
	market           int
	previousAttempts int
}

type World interface {
	MW_Name() string
	MW_UWP() string
	TravelZone() string
}

func FindSuplier(world World) (*suplier, error) {
	//	err := fmt.Errorf("error value was not adressed")
	s := suplier{}
	return &s, nil
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
