package speculative

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TestTrade(t *testing.T) {
	var lastWorld *survey.SecondSurveyData
	for _, line := range []string{

		"|Khteaouw|0129|E531497-8|301| |  |AsT9|F8 V|I|{ -3 }|-3|(731-3)|[4158]||9 |-63  |8 |4|-64  |69   |Ni Po|  |Reaver's Deep|Keiar|Reav|Aslan Hierate, Tlaukhu control, Aokhalte (10), Sahao' (21), Ouokhoi (26)                     ",
		"|Oihoiei|0230|A8558A8-C|214| |R |AsWc|F9 V M1 V|I|{ 2 }|2 |(F7C+2)|[8A5C]||17|2520 |8 |4|-63  |70   |Ga Pa Ph|R |Reaver's Deep|Keiar|Reav|Aslan Hierate, single one-world clan dominates                                               ",
	} {

		wrld := survey.Parse(line)
		fmt.Println(" ")
		fmt.Println("test world:")
		fmt.Println(wrld.String())
		s, err := FindSuplier(wrld, Market_ALL)
		if s == nil {
			t.Errorf("FindSuplier() return no object")
		}
		if err != nil {
			t.Errorf("FindSuplier() returned error: '%v'", err.Error())
		}
		//s.RollQuantity()
		//for _, tg := range s.tradeGoodsAvailable {
		//	purchPrice := PricePurchaseAverage(s.world, tg)
		//	salePrice := PriceSaleAverage(s.world, tg)
		//	fmt.Println(tg.GoodsType(), tg.Stored(), purchPrice, salePrice, tg.BasePrice())
		//	if tg.Stored() < 1 {
		//		t.Errorf("must not be strored %v tons", tg.Stored())
		//	}
		//}
		// for k, v := range ListPrices(s.world) {
		// 	fmt.Printf("TG code %v, cost:%v\n", k, v)
		// }

		ga, _ := DetermineGoodsAvailable(wrld)
		fmt.Println("-Productiom----")
		for i := range ga {
			fmt.Println(ga[i].GoodsType())
		}
		fmt.Println("-Exported----")
		ex, _ := DetermineOffer(wrld)
		for i := range ex {
			fmt.Println(ex[i].GoodsType())
		}
		fmt.Println("-Demanded----")
		im, _ := DetermineDemand(wrld)
		for i := range im {
			fmt.Println(im[i].GoodsType())
		}
		//Transit: Khteaouw ---> Oihoiei ---> Atiyr ---> Hrike
		//Import: Khteaouw ---> Oihoiei
		// avGoods, err := DetermineGoodsAvailable(wrld)
		// if err != nil {
		// 	t.Errorf("DetermineGoodsAvailable() returned error: '%v'", err.Error())
		// }
		// for _, tGood := range avGoods {
		// 	fmt.Println(tGood.GoodsType())
		// }
		if lastWorld != nil {
			tgList, err := TradeGoodsFlow(lastWorld, wrld)
			if err != nil {
				t.Errorf("TradeGoodsFlow(lastWorld, wrld): %v", err.Error())
			}
			fmt.Printf("Trade: %v ---> %v\n", lastWorld.MW_Name(), wrld.MW_Name())
			if len(tgList) == 0 {
				fmt.Println("No Trade Goods")
			}
			for _, tg := range tgList {
				fmt.Println(tg.GoodsType(), PricePurchaseAverage(lastWorld, tg), "--->", PriceSaleAverage(wrld, tg))
			}
		}

		lastWorld = wrld
	}

}
