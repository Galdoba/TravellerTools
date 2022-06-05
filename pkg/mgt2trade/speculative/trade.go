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
}

func FindSuplier(world World) (*suplier, error) {
	//	err := fmt.Errorf("error value was not adressed")
	s := suplier{}
	return &s, nil
}

func DetermineGoodsAvailable(world World) []*tradegoods.TradeGood {
	tradeCodes, err := tradecodes.FromUWPstr(world.MW_UWP())
	for i, tc := range tradeCodes {

	}
	return nil, err
}
