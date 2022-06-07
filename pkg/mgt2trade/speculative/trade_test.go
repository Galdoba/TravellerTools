package speculative

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TestTrade(t *testing.T) {
	for _, line := range []string{
		"|Thekar|0626|C553400-9|702|||NaHu|F9 V|I|{ -1 }|-1|(832-5)|[1314]|B|10|-240|8|4|-59|66|Ni Po Asla0||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		"|Abramo|0630|B200722-B|323||KM|NaHu|G3 V M6 V|I|{ +3 }|3|(E6D-1)|[3A17]|BD|11|-1092|8|4|-59|70|Na Va Pi Asla4|F|Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		"|New Covenant|0722|A5579DE-9|413|A|N|CsIm|K5 V|I|{ +2 }|2|(F8C+5)|[DB9D]|BE|15|7200|8|4|-58|62|Hi Pz|N|Reaver's Deep|Keiar|Reav|Client state, Third Imperium",
		"|Icarus|0729|C759855-5|223|||NaHu|G0 V|I|{ -1 }|-1|(A76-3)|[6733]|Be|15|-1260|8|4|-58|69|Ph Asla1||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		"|Tulena|0827|D746300-7|523|||NaHu|F6 V M6 V|I|{ -3 }|-3|(520-5)|[1112]|B|12|0|8|4|-57|67|Lo||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		"|Dunmarrow|0921|B544653-A|203||S|CsIm|G0 V|J|{ +2 }|2|(B56-1)|[3827]|BC|14|-330|9|4|-56|61|Ag Ni|S|Reaver's Deep|Ea|Reav|Client state, Third Imperium",
		"|Hrou|0923|D200579-8|314|||NaHu|F3 V|J|{ -3 }|-3|(C41-2)|[6269]|B|13|-96|9|4|-56|63|Ni Va Asla9||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Lestrow|0926|C798784-8|813|A||GdMh|M0 V|J|{ +0 }|0|(D68-2)|[5736]|BC|12|-1248|9|4|-56|66|Ag Pi Pz||Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		"|Laroaetea|1024|E556555-6|512|||NaHu|M2 V M5 V|J|{ -2 }|-2|(742-4)|[3334]|BC|14|-224|9|4|-55|64|Ag Ni Asla8||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Fask|1028|C9868AA-8|321|A||GdMh|M1 V M6 V|J|{ +0 }|0|(D78+2)|[A87A]|Bc|10|1456|9|4|-55|68|Ri Pa Ph Pz||Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		"|Theodora|1030|B857563-A|104|A|KM|GdMh|G9 V M7 V|J|{ +3 }|3|(B47+1)|[2827]|BC|11|308|9|4|-55|70|Ag Ni Ga Da Mr|F|Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		"|Tearlach|1121|E569749-7|413|||NaHu|G9 V M4 V|J|{ -1 }|-1|(967+1)|[8668]|BC|14|378|9|4|-54|61|Ri||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Gaajpadje|1124|E667874-4|904|||NaHu|G8 V|J|{ -1 }|-1|(A75-3)|[6732]|Bc|10|-1050|9|4|-54|64|Ga Ri Pa Ph (J'aadje)6||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Earlo|1125|D542102-7|404|A||NaHu|F7 V|J|{ -3 }|-3|(300-5)|[1113]|B|13|0|9|4|-54|65|He Lo Po Da Asla7||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Mirak|1127|C766763-8|702|A|M|GdMh|M1 V|J|{ +1 }|1|(B69-2)|[4825]|BC|10|-1188|9|4|-54|67|Ag Ga Ri Pz Mr|M|Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		"|Dran|1129|C551566-9|123|A||GdMh|K9 V M5 V|J|{ -1 }|-1|(C43-2)|[4448]|B|12|-288|9|4|-54|69|Ni Po Da O:1230||Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		"|Marlheim|1230|A5759A8-B|303|A|KM|GdMh|G2 V M9 V|J|{ +5 }|5|(E8G+5)|[9E5B]|BE|14|8960|9|4|-53|70|Hi In Cx Pz|F|Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		"|Leaa|1222|E100488-9|312|||NaAs|G7 V|J|{ -2 }|-2|(931-2)|[4259]|B|13|-54|9|4|-53|62|Ni Va AslaW||Reaver's Deep|Ea|Reav|Non-Aligned, Aslan-dominated",
		"|Roakhoi|1224|C969543-5|702|||NaHu|F5 V|J|{ -2 }|-2|(742-5)|[2322]|Bc|8|-280|9|4|-53|64|Ni Pr Asla5||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Ea|1225|C7586AA-7|214|A||NaHu|M0 V M1 V|J|{ -1 }|-1|(853+1)|[8579]|BC|14|120|9|4|-53|65|Ag Ni Da Asla8||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Htalrea|1226|E767610-1|103|||NaHu|M1 V|J|{ -1 }|-1|(853-5)|[1511]|BC|10|-600|9|4|-53|66|Ag Ni Ga Ri (Polyphemes)||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Shamas|1321|E556305-6|703|||NaHu|F7 V|J|{ -3 }|-3|(520-5)|[1134]|B|13|0|9|4|-52|61|Lo||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Vincit|1327|C8987A9-8|403||M|DuCf|M1 V|J|{ +0 }|0|(C68+1)|[8769]|BC|15|576|9|4|-52|67|Ag Pi|M|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Andiros|1328|C799566-8|601|||NaHu|M3 V|J|{ -2 }|-2|(C42-3)|[4347]|B|12|-288|9|4|-52|68|Ni O:1428||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Kingston|1428|B764994-C|213||KM|NaHu|M0 V|J|{ +4 }|4|(F8F+2)|[7D3A]|Bc|15|3600|9|4|-51|68|Hi Pr|F|Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Fort William|1521|C540467-9|601||M|DuCf|M2 V M4 V|J|{ -1 }|-1|(732-1)|[4359]|B|8|-42|9|4|-50|61|De He Ni Po Mr|M|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Fulton|1524|C98A788-9|102|||DuCf|G3 V M5 V|J|{ +1 }|1|(B6A+1)|[7859]|BC|12|660|9|4|-50|64|Ri Wa||Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Ranald|1526|C556544-9|622|||DuCf|M0 V|J|{ +0 }|0|(B44-2)|[3537]|BC|10|-352|9|4|-50|66|Ag Ni||Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Invermory|1622|B584789-A|522||KM|DuCf|M3 V M5 V|J|{ +5 }|5|(G6E+5)|[8C6B]|BC|10|6720|9|4|-49|62|Ag Ri|F|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Duncinae|1624|A686648-B|502||KM|DuCf|G0 V|J|{ +4 }|4|(A58+4)|[6A5B]|BC|10|1600|9|4|-49|64|Ag Ni Ga Ri Cx|F|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Just|1625|C7487AA-5|420|A|M|DuCf|K4 V|J|{ +0 }|0|(967+2)|[9777]|BC|7|756|9|4|-49|65|Ag Pi Pz|M|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		"|Lajanjigal|1721|DAB6513-8|801|A||NaHu|G8 V|K|{ -3 }|-3|(C41-5)|[2225]|B|14|-240|10|5|-48|61|Fl Ni Da (Languljigee)||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Coventry|1723|X565733-2|404|R||NaHu|G8 V|K|{ +0 }|0|(965-3)|[4721]|BC|13|-810|10|5|-48|63|Ag Ri Fo Px||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Traneer|1727|E576679-6|323|||NaHu|G5 V|K|{ -2 }|-2|(852-1)|[7467]|BC|13|-80|10|5|-48|67|Ag Ni||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Dakaar|1821|B425612-B|202||KM|NaHu|K0 V D|K|{ +2 }|2|(A56-2)|[2817]|B|6|-600|10|5|-47|61|Ni|F|Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Drexilthar|1826|B56969D-7|914|A|S|CsIm|G4 V|K|{ +0 }|0|(854+4)|[A69B]|BC|14|640|10|5|-47|66|Ni Ri Da (Iltharans)|S|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		"|Kraan|1828|C501456-8|223|||NaHu|M3 V|K|{ -2 }|-2|(B31-3)|[3247]|B|11|-99|10|5|-47|68|Ic Ni Va||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Daken|1830|C631233-9|902|||NaHu|M2 V M5 V|K|{ -1 }|-1|(610-4)|[1126]|B|7|0|10|5|-47|70|Lo Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Cassandra|1924|B000538-C|914|||NaHu|F0 V|K|{ +1 }|1|(C45+1)|[565C]|B|15|240|10|5|-46|64|As Ni Va||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Outpost|1926|B310442-D|413||N|CsIm|K7 V M5 V|K|{ +1 }|1|(A34-3)|[1519]|B|15|-360|10|5|-46|66|Ni|N|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		"|Tashrakaar|1927|D651695-6|213|||NaHu|M2 V|K|{ -3 }|-3|(851-5)|[4334]|B|14|-200|10|5|-46|67|Ni Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Luushakaan|2021|D541513-5|811|||NaHu|M0 V|K|{ -3 }|-3|(741-5)|[2222]|B|11|-140|10|5|-45|61|He Ni Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Dutrissal|2027|CAC5235-9|802|||NaHu|F4 V M8 V|K|{ -1 }|-1|(610-3)|[1137]|B|9|0|10|5|-45|67|Fl Lo||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Drellesarr|2029|B310550-A|502|A||NaHu|K5 V|K|{ +1 }|1|(945-3)|[1615]|B|6|-540|10|5|-45|69|Ni Da||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Drenslaar|2030|D553694-6|313|||NaHu|M2 V M5 V|K|{ -3 }|-3|(851-5)|[4334]|B|11|-200|10|5|-45|70|Ni Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Grendal|2127|C889855-A|613||S|CsIm|F7 V M6 V|K|{ +2 }|2|(E7B+1)|[6A38]|BC|15|1078|10|5|-44|67|Ri Ph|S|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		"|Tharrill|2128|C885741-8|722|||NaHu|M0 V M3 V|K|{ +1 }|1|(D69-3)|[3814]|BC|14|-2106|10|5|-44|68|Ag Ga Ri||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Sarrad|2129|D88A300-8|823|A||CaAs|K8 V M3 V|K|{ -3 }|-3|(A20-5)|[1113]|B|12|0|10|5|-44|69|Lo Wa Da||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		"|Garrison|2221|A55796B-B|412||N|ImDi|F5 V|K|{ +3 }|3|(E8E+5)|[BC7D]|BE|10|7840|10|5|-43|61|Hi Mr|N|Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		"|Kaaniir|2223|C688611-6|723||S|CsIm|F8 V|K|{ +0 }|0|(854-4)|[2612]|BC|12|-640|10|5|-43|63|Ag Ni Ri|S|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		"|Yarhfahl|2228|C658796-6|110|||NaHu|M1 V M4 V|K|{ +0 }|0|(967-1)|[6745]|BC|11|-378|10|5|-43|68|Ag||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		"|Datinar|2230|B431685-A|303|||CaAs|K0 V|K|{ +1 }|1|(B55-1)|[4738]|B|11|-275|10|5|-43|70|Na Ni Po||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		"|Gaargir|2322|B565304-C|713||NS|ImDi|K7 V|K|{ +2 }|2|(922+1)|[153A]|B|7|36|10|5|-42|62|Lo|A|Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		"|Ildrissar|2326|C995836-7|203|A||CaAs|M1 V|K|{ -1 }|-1|(A77-2)|[7746]|Bc|14|-980|10|5|-42|66|Pa Ph Pi Pz||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		"|Carrill|2330|A0009AE-E|613|A|KM|CaAs|F4 V|K|{ +5 }|5|(F8H+5)|[DE9J]|BE|11|10200|10|5|-42|70|As Hi In Na Va Cx Pz|F|Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		"|Kaanash|2421|B55687A-7|414|A|NS|ImDi|M1 V M6 V|K|{ +1 }|1|(A79+3)|[A979]|Bce|17|1890|10|5|-41|61|Pa Ph Pz|A|Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		"|Diablo|2423|B9C7477-9|724|||ImDi|M2 V|K|{ +0 }|0|(C33+1)|[4459]|B|11|108|10|5|-41|63|Fl Ni||Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		"|Lindritar|2429|C5796A7-8|211|||CaAs|K7 V|K|{ -2 }|-2|(A52-2)|[6458]|B|13|-200|10|5|-41|69|Ni||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		"|Inura|2523|C421312-A|812|||ImDi|G9 V|L|{ +0 }|0|(820-4)|[1316]|B|6|0|11|5|-40|63|He Lo Po||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Hope|2526|E65778B-4|200|A||ImDi|F4 V|L|{ -1 }|-1|(965+1)|[9676]|BC|8|270|11|5|-40|66|Ag Ga Pz||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Yaggoth|2530|B864756-B|222|||CaAs|F8 V M1 V M6 V|L|{ +4 }|4|(D6E+3)|[6B4A]|BC|13|3276|11|5|-40|70|Ag Ri||Reaver's Deep|Urlaqqash|Reav|Carrillian Assembly",
		"|Lavnia|2621|A546657-D|704||N|ImDi|G2 V M7 V|L|{ +2 }|2|(C56+2)|[685D]|BC|8|720|11|5|-39|61|Ag Ni|N|Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Astoria|2622|B545674-A|201||N|ImDi|F0 V M4 V|L|{ +2 }|2|(956+1)|[4838]|BC|10|270|11|5|-39|62|Ag Ni|N|Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Irlaggur|2624|B6918CE-A|300|A||ImDi|M3 V|L|{ +2 }|2|(A7B+5)|[CA9E]|BDe|12|3850|11|5|-39|64|He Ph Pi Pz||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Boran|2628|C3135AB-A|301|A||CaAs|G2 V|L|{ +0 }|0|(844+2)|[757C]|B|7|256|11|5|-39|68|Ic Ni Da||Reaver's Deep|Urlaqqash|Reav|Carrillian Assembly",
		"|Hela|2721|C431300-9|412|||ImDi|F5 V|L|{ -1 }|-1|(820-5)|[1214]|B|10|0|11|5|-38|61|Lo Po||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Virshash|2724|DA86954-6|403||S|ImDi|F9 V M4 V|L|{ -1 }|-1|(B87-3)|[7834]|BcE|10|-1848|11|5|-38|64|Hi Pr (Virushi)|S|Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Syreon|2730|C54688C-8|320|A|M|CaAs|M0 V|L|{ -1 }|-1|(H77+2)|[B78B]|Bc|15|1666|11|5|-38|70|Pa Ph Pi Pz|M|Reaver's Deep|Urlaqqash|Reav|Carrillian Assembly",
		"|Marianne|2821|C6787C9-8|500|A|S|ImDi|M1 V|L|{ +0 }|0|(968+1)|[8769]|BCD|13|432|11|5|-37|61|Ag Pi Pz|S|Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Sharrik|2824|B664896-9|804|||ImDi|M1 V M1 V|L|{ +2 }|2|(E7B+1)|[7A48]|BcCe|11|1078|11|5|-37|64|Ri Pa Ph||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		"|Freehold|2825|E555457-7|513|||NaHu|G6 V|L|{ -3 }|-3|(631-3)|[4157]|Bc|14|-54|11|5|-37|65|Ni Pa||Reaver's Deep|Urlaqqash|Reav|Non-Aligned, Human-dominated",
		"|Lyresse|2828|C693651-9|813|A||CaAs|K4 V|L|{ -1 }|-1|(C53-5)|[2515]|B|14|-900|11|5|-37|68|Ni Da||Reaver's Deep|Urlaqqash|Reav|Carrillian Assembly",
		"|Rothman|2829|B796855-9|403||KM|CaAs|M1 V M3 V|L|{ +2 }|2|(D7B+1)|[6A37]|Bc|10|1001|11|5|-37|69|Pa Ph Pi|F|Reaver's Deep|Urlaqqash|Reav|Carrillian Assembly",
		"|Synoft|2927|C5428CC-7|310|A||NaHu|M2 V M7 V|L|{ -1 }|-1|(A77+2)|[B78A]|BD|11|980|11|5|-36|67|He Po Ph Pi Pz||Reaver's Deep|Urlaqqash|Reav|Non-Aligned, Human-dominated",
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
		for k, v := range ListPrices(s.world) {
			fmt.Printf("TG code %v, cost:%v\n", k, v)
		}

		// avGoods, err := DetermineGoodsAvailable(wrld)
		// if err != nil {
		// 	t.Errorf("DetermineGoodsAvailable() returned error: '%v'", err.Error())
		// }
		// for _, tGood := range avGoods {
		// 	fmt.Println(tGood.GoodsType())
		// }

	}

}
