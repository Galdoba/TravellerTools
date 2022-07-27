package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/planet/landing"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func main() {
	for _, data := range []string{
		// "|Gash|2116|DAF8573-7|112|||NaHu|M1 V|G|{ -3 }|-3|(741-5)|[2224]|B|10|-140|6|1|-44|56|Ni||Reaver's Deep|Caledon|Reav|Non-Aligned, Human-dominated",
		// "|Rock|2214|B400364-A|601|A||NaHu|F8 V M8 V|G|{ +1 }|1|(621-1)|[1438]|B|8|-12|6|1|-43|54|Lo Va Da O:2313||Reaver's Deep|Caledon|Reav|Non-Aligned, Human-dominated",
		// "|Kolath|2313|B7678CB-A|424|A|KM|NaHu|F3 V M9 V|G|{ +4 }|4|(G7D+5)|[AC7C]|Bc|12|7280|6|1|-42|53|Ga Ri Pa Ph Pz|F|Reaver's Deep|Caledon|Reav|Non-Aligned, Human-dominated",
		// "|Concorde|2218|A999587-E|613||N|ImDi|G7 V M6 V|G|{ +1 }|1|(B45+1)|[565E]|B|10|220|6|1|-43|58|Ni|N|Reaver's Deep|Caledon|Reav|Third Imperium, Domain of Ilelish",
		// "|Loren|2311|C57459C-7|201|A|S|CsIm|M5 V|G|{ -1 }|-1|(743+2)|[848A]|BC|7|168|6|1|-42|51|Ag Ni Da BruhW|S|Reaver's Deep|Caledon|Reav|Client state, Third Imperium",
		// "|Kurat|2315|CAA7667-8|613|A||NaHu|G5 V|G|{ -2 }|-2|(C52-2)|[6458]|B|12|-240|6|1|-42|55|Fl Ni Da O:2313||Reaver's Deep|Caledon|Reav|Non-Aligned, Human-dominated",
		// "|Lurammish|2320|C512755-9|800||S|ImDi|F6 V M5 V|G|{ +0 }|0|(969-2)|[5737]|BD|6|-972|6|1|-42|60|Ic Na Pi|S|Reaver's Deep|Caledon|Reav|Third Imperium, Domain of Ilelish",
		// "|Doom|2412|D400200-7|101|||NaHu|F1 V|G|{ -3 }|-3|(410-5)|[1112]|B|8|0|6|1|-41|52|Lo Va||Reaver's Deep|Caledon|Reav|Non-Aligned, Human-dominated",
		// "|Mer|2414|C79A520-8|424|||ImDi|G2 V M9 V|G|{ -2 }|-2|(D42-5)|[1313]|B|15|-520|6|1|-41|54|Ni Wa||Reaver's Deep|Caledon|Reav|Third Imperium, Domain of Ilelish",
		// "|Gerim|2416|A888A97-E|224||N|ImDi|F4 V M0 V|G|{ +3 }|3|(J9F+3)|[AD5E]|BE|12|7290|6|1|-41|56|Hi|N|Reaver's Deep|Caledon|Reav|Third Imperium, Domain of Ilelish",
		// "|Bryn|2417|B4268B8-8|914|||ImDi|G3 IV M4 V|G|{ +0 }|0|(F78+1)|[8858]|BDe|14|840|6|1|-41|57|Ph Pi||Reaver's Deep|Caledon|Reav|Third Imperium, Domain of Ilelish",
		// "|Ikuna|2419|E000410-A|312|||ImDi|K6 V|G|{ -1 }|-1|(932-5)|[1315]|B|10|-270|6|1|-41|59|As Ni Va||Reaver's Deep|Caledon|Reav|Third Imperium, Domain of Ilelish",
		// "|Roye|2511|C79A458-A|202|||NaHu|M3 V|H|{ +0 }|0|(833+1)|[445A]|B|11|72|7|1|-40|51|Ni Wa||Reaver's Deep|Nightrim|Reav|Non-Aligned, Human-dominated",
		// "|Sheffield|2513|C667575-8|214|||ImDi|M2 V M4 V|H|{ -1 }|-1|(C43-3)|[3436]|BcC|8|-432|7|1|-40|53|Ag Ni Ga Pr||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Shetland|2514|B54478A-8|914||N|ImDi|F0 V M8 V|H|{ +1 }|1|(E69+3)|[987A]|BCD|12|2268|7|1|-40|54|Ag Pi|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Scapa|2515|B667784-A|703|||ImDi|G6 V M2 V|H|{ +4 }|4|(C6D+2)|[5B38]|BCf|14|1872|7|1|-40|55|Ag Ga Ri||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Kaagin|2516|A5659A9-D|504||N|ImDi|M1 V|H|{ +3 }|3|(F8F+4)|[AC6E]|BcE|7|7200|7|1|-40|56|Hi Pr|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Tower|2519|C5327B8-7|710|||ImDi|G9 V M1 V|H|{ -1 }|-1|(967-1)|[7657]|B|11|-378|7|1|-40|59|Na Po||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|St. George|2616|A676AA6-C|314||N|ImDi|F1 V M2 V|H|{ +4 }|4|(H9F+3)|[9E4B]|BEF|14|6885|7|1|-39|56|Hi In Cp|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Lore|2619|B668723-8|700|||ImDi|G2 V|H|{ +2 }|2|(96A-1)|[4925]|BC|10|-540|7|1|-39|59|Ag Ri||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Merisun|2720|X685679-2|701|R||ImDi|G5 V M4 V|H|{ -1 }|-1|(853+1)|[7563]|BC|9|120|7|1|-38|60|Ag Ni Ga Ri Fo||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Ankara|2812|C99947B-9|411|A|S|ImDi|K9 V|H|{ -1 }|-1|(832+1)|[637B]|B|14|48|7|1|-37|52|Ni Da|S|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Vetzeal|2813|E423214-7|310|||ImDi|F7 V|H|{ -3 }|-3|(410-5)|[1135]|B|8|0|7|1|-37|53|Lo Po||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Aries|2816|B310444-D|214||N|ImDi|F1 V|H|{ +1 }|1|(B34-1)|[253B]|B|11|-132|7|1|-37|56|Ni|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Koath|2819|C301453-A|424|||ImDi|G5 V M9 V|H|{ +0 }|0|(C33-3)|[1427]|B|9|-324|7|1|-37|59|Ic Ni Va||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Veroch|2912|E6B1101-8|713|||ImDi|G4 V|H|{ -3 }|-3|(700-5)|[1114]|B|9|0|7|1|-36|52|Fl He Lo||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Mull|2916|CAC7312-9|611|||ImDi|F9 V|H|{ -1 }|-1|(720-5)|[1215]|B|15|0|7|1|-36|56|Fl Lo||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Stonehaven|2917|A9D5422-C|914||N|ImDi|F2 V|H|{ +1 }|1|(B34-3)|[1518]|B|11|-396|7|1|-36|57|Ni|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Orkney|2919|B5888CB-9|302|A||NaHu|M0 V M7 V|H|{ +2 }|2|(C7B+4)|[AA7B]|Bc|13|3696|7|1|-36|59|Ri Pa Ph Pz||Reaver's Deep|Nightrim|Reav|Non-Aligned, Human-dominated",
		// "|Maiden|2920|X544567-5|104|R||NaHu|K2 V|H|{ -2 }|-2|(742-2)|[5355]|BC|10|-112|7|1|-36|60|Ag Ni Fo O:2919||Reaver's Deep|Nightrim|Reav|Non-Aligned, Human-dominated",
		// "|Khishali|3012|C866759-8|914|||ImDi|G3 V M1 V|H|{ +1 }|1|(E69+2)|[8869]|BC|10|1512|7|1|-35|52|Ag Ga Ri||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Khagish|3019|D7649DD-6|711|A|S|ImDi|K0 V|H|{ -1 }|-1|(B87+3)|[D89A]|BcE|12|1848|7|1|-35|59|Hi Pr Pz|S|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|815-205|3111|C99A311-9|204|||ImDi|M1 V|H|{ -1 }|-1|(920-5)|[1215]|B|9|0|7|1|-34|51|Lo Wa||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Arthur|3112|E566000-0|003|||ImDi|M1 V|H|{ -3 }|-3|(200-5)|[0000]|B|11|0|7|1|-34|52|Ba||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Ghost|3115|C685688-5|713|||ImDi|F1 V|H|{ +0 }|0|(854+1)|[6655]|BC|13|160|7|1|-34|55|Ag Ni Ga Ri (Ayansh'i)||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Dundee|3118|B533133-A|700||N|ImDi|M0 II M2 V|H|{ +1 }|1|(301-2)|[1227]|B|12|0|7|1|-34|58|Lo Po|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Havant|3120|B542524-B|702|||ImDi|M3 V M5 V|H|{ +1 }|1|(945-1)|[3639]|B|12|-180|7|1|-34|60|He Ni Po||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Wells|3211|A786500-9|523|||ImDi|G6 V|H|{ +1 }|1|(C45-3)|[1614]|BcC|13|-720|7|1|-33|51|Ag Ni Ga Pr||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Petzina|3212|B674767-A|824|||ImDi|K8 V|H|{ +3 }|3|(F6C+3)|[7A5A]|BCD|9|3240|7|1|-33|52|Ag Pi O:Daib-0114||Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Khakhan|3220|A988AA8-E|903||N|ImDi|F6 V M0 V M9 V|H|{ +3 }|3|(F9F+3)|[AD5E]|BE|9|6075|7|1|-33|60|Hi|N|Reaver's Deep|Nightrim|Reav|Third Imperium, Domain of Ilelish",
		// "|Esekheali|0124|C400573-9|504|||IsDo|F4 IV|I|{ -1 }|-1|(B43-4)|[2426]|B|12|-528|8|4|-64|64|Ni Va||Reaver's Deep|Keiar|Reav|Islaiat Dominate",
		// "|Ekhiwua'ea|0125|D898874-6|504|||IsDo|K7 V M6 V|I|{ -2 }|-2|(A75-4)|[6634]|Bc|13|-1400|8|4|-64|65|Pa Ph Pi||Reaver's Deep|Keiar|Reav|Islaiat Dominate",
		// "|Khtaao|0127|E593267-7|314|||AsT6|K7 V|I|{ -3 }|-3|(410-3)|[2157]|B|12|0|8|4|-64|67|Lo O:0227||Reaver's Deep|Keiar|Reav|Aslan Hierate, Tlaukhu control, Uiktawa (7), Iykyasea (17), Faowaou (27)",
		// "|Oloih|0227|B0007C7-D|304||R|AsT6|K8 V|I|{ +2 }|2|(D6D+2)|[795D]|BD|10|2028|8|4|-63|67|As Na Va Pi|R|Reaver's Deep|Keiar|Reav|Aslan Hierate, Tlaukhu control, Uiktawa (7), Iykyasea (17), Faowaou (27)",
		// "|Khteaouw|0129|E531497-8|301|||AsT9|F8 V|I|{ -3 }|-3|(731-3)|[4158]|B|9|-63|8|4|-64|69|Ni Po||Reaver's Deep|Keiar|Reav|Aslan Hierate, Tlaukhu control, Aokhalte (10), Sahao' (21), Ouokhoi (26)",
		// "|Islaiat|0221|A868AA9-D|914||KM|IsDo|G2 V M8 V|I|{ +4 }|4|(H9G+5)|[BE6E]|BE|13|12240|8|4|-63|61|Hi Cx|F|Reaver's Deep|Keiar|Reav|Islaiat Dominate",
		// "|Oihoiei|0230|A8558A8-C|214||R|AsWc|F9 V M1 V|I|{ +2 }|2|(F7C+2)|[8A5C]|Bc|17|2520|8|4|-63|70|Ga Pa Ph|R|Reaver's Deep|Keiar|Reav|Aslan Hierate, single one-world clan dominates",
		// "|Oirue'ea|0324|B5648BB-A|303|A||IsDo|K8 V|I|{ +3 }|3|(D7C+5)|[AB7C]|Bc|8|5460|8|4|-62|64|Ri Pa Ph Pz||Reaver's Deep|Keiar|Reav|Islaiat Dominate",
		// "|Kteieaelal|0328|C9897BB-9|714|A||AsTv|K4 V|I|{ +1 }|1|(E6A+3)|[987B]|BC|11|2520|8|4|-62|68|Ri Pz||Reaver's Deep|Keiar|Reav|Aslan Hierate, Tlaukhu vassal clan dominates",
		// "|Eihewasei|0423|C310466-9|413|||IsDo|F0 V|I|{ -1 }|-1|(A32-2)|[3348]|B|10|-120|8|4|-61|63|Ni O:0324||Reaver's Deep|Keiar|Reav|Islaiat Dominate",
		// "|Yaoueai|0424|E560468-7|910|||IsDo|M1 V M2 V|I|{ -3 }|-3|(631-3)|[4157]|B|6|-54|8|4|-61|64|De Ni O:0324||Reaver's Deep|Keiar|Reav|Islaiat Dominate",
		// "|Phontramus|0426|C87A341-9|203|||NaHu|M0 V M7 V|I|{ -1 }|-1|(820-5)|[1215]|B|9|0|8|4|-61|66|Lo Wa Asla2||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Atiyr|0429|E413774-8|212|||AsSc|F0 V M0 V|I|{ -2 }|-2|(C66-4)|[5536]|BD|16|-1728|8|4|-61|69|Ic Na Pi||Reaver's Deep|Keiar|Reav|Aslan Hierate, multiple clans split control",
		// "|Janet|0521|B9995AB-A|923|A|KM|NaHu|M2 V|I|{ +2 }|2|(C46+4)|[777C]|B|9|1152|8|4|-60|61|Ni Da Asla0|F|Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Aerin|0523|C896853-8|322|||NaHu|K1 V K4 V|I|{ -1 }|-1|(E77-4)|[5725]|Bc|11|-2744|8|4|-60|63|Pa Ph Pi Asla3||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Therad|0525|B666854-8|601|||NaHu|M1 V M8 V|I|{ +1 }|1|(B79-1)|[6936]|Bc|6|-693|8|4|-60|65|Ga Ri Pa Ph Asla0||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Hrike|0530|A7888A9-C|412||R|AsMw|G3 V M0 V|I|{ +3 }|3|(G7D+4)|[9B6D]|Bc|15|5824|8|4|-60|70|Ri Pa Ph|R|Reaver's Deep|Keiar|Reav|Aslan Hierate, single multiple-world clan dominates",
		// "|Gwalcmai|0623|E543443-5|803|||NaHu|K9 V|I|{ -3 }|-3|(631-5)|[1122]|B|15|-90|8|4|-59|63|Ni Po||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Asden|0624|D543500-7|322|||NaHu|K7 V|I|{ -3 }|-3|(741-5)|[1212]|B|14|-140|8|4|-59|64|Ni Po||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Thekar|0626|C553400-9|702|||NaHu|F9 V|I|{ -1 }|-1|(832-5)|[1314]|B|10|-240|8|4|-59|66|Ni Po Asla0||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Abramo|0630|B200722-B|323||KM|NaHu|G3 V M6 V|I|{ +3 }|3|(E6D-1)|[3A17]|BD|11|-1092|8|4|-59|70|Na Va Pi Asla4|F|Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|New Covenant|0722|A5579DE-9|413|A|N|CsIm|K5 V|I|{ +2 }|2|(F8C+5)|[DB9D]|BE|15|7200|8|4|-58|62|Hi Pz|N|Reaver's Deep|Keiar|Reav|Client state, Third Imperium",
		// "|Icarus|0729|C759855-5|223|||NaHu|G0 V|I|{ -1 }|-1|(A76-3)|[6733]|Be|15|-1260|8|4|-58|69|Ph Asla1||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Tulena|0827|D746300-7|523|||NaHu|F6 V M6 V|I|{ -3 }|-3|(520-5)|[1112]|B|12|0|8|4|-57|67|Lo||Reaver's Deep|Keiar|Reav|Non-Aligned, Human-dominated",
		// "|Dunmarrow|0921|B544653-A|203||S|CsIm|G0 V|J|{ +2 }|2|(B56-1)|[3827]|BC|14|-330|9|4|-56|61|Ag Ni|S|Reaver's Deep|Ea|Reav|Client state, Third Imperium",
		// "|Hrou|0923|D200579-8|314|||NaHu|F3 V|J|{ -3 }|-3|(C41-2)|[6269]|B|13|-96|9|4|-56|63|Ni Va Asla9||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Lestrow|0926|C798784-8|813|A||GdMh|M0 V|J|{ +0 }|0|(D68-2)|[5736]|BC|12|-1248|9|4|-56|66|Ag Pi Pz||Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		// "|Laroaetea|1024|E556555-6|512|||NaHu|M2 V M5 V|J|{ -2 }|-2|(742-4)|[3334]|BC|14|-224|9|4|-55|64|Ag Ni Asla8||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Fask|1028|C9868AA-8|321|A||GdMh|M1 V M6 V|J|{ +0 }|0|(D78+2)|[A87A]|Bc|10|1456|9|4|-55|68|Ri Pa Ph Pz||Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		// "|Theodora|1030|B857563-A|104|A|KM|GdMh|G9 V M7 V|J|{ +3 }|3|(B47+1)|[2827]|BC|11|308|9|4|-55|70|Ag Ni Ga Da Mr|F|Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		// "|Tearlach|1121|E569749-7|413|||NaHu|G9 V M4 V|J|{ -1 }|-1|(967+1)|[8668]|BC|14|378|9|4|-54|61|Ri||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Gaajpadje|1124|E667874-4|904|||NaHu|G8 V|J|{ -1 }|-1|(A75-3)|[6732]|Bc|10|-1050|9|4|-54|64|Ga Ri Pa Ph (J'aadje)6||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Earlo|1125|D542102-7|404|A||NaHu|F7 V|J|{ -3 }|-3|(300-5)|[1113]|B|13|0|9|4|-54|65|He Lo Po Da Asla7||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Mirak|1127|C766763-8|702|A|M|GdMh|M1 V|J|{ +1 }|1|(B69-2)|[4825]|BC|10|-1188|9|4|-54|67|Ag Ga Ri Pz Mr|M|Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		// "|Dran|1129|C551566-9|123|A||GdMh|K9 V M5 V|J|{ -1 }|-1|(C43-2)|[4448]|B|12|-288|9|4|-54|69|Ni Po Da O:1230||Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		// "|Marlheim|1230|A5759A8-B|303|A|KM|GdMh|G2 V M9 V|J|{ +5 }|5|(E8G+5)|[9E5B]|BE|14|8960|9|4|-53|70|Hi In Cx Pz|F|Reaver's Deep|Ea|Reav|Grand Duchy of Marlheim",
		// "|Leaa|1222|E100488-9|312|||NaAs|G7 V|J|{ -2 }|-2|(931-2)|[4259]|B|13|-54|9|4|-53|62|Ni Va AslaW||Reaver's Deep|Ea|Reav|Non-Aligned, Aslan-dominated",
		// "|Roakhoi|1224|C969543-5|702|||NaHu|F5 V|J|{ -2 }|-2|(742-5)|[2322]|Bc|8|-280|9|4|-53|64|Ni Pr Asla5||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Ea|1225|C7586AA-7|214|A||NaHu|M0 V M1 V|J|{ -1 }|-1|(853+1)|[8579]|BC|14|120|9|4|-53|65|Ag Ni Da Asla8||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		"|Htalrea|1226|E767610-1|103|||NaHu|M1 V|J|{ -1 }|-1|(853-5)|[1511]|BC|10|-600|9|4|-53|66|Ag Ni Ga Ri (Polyphemes)||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Shamas|1321|E556305-6|703|||NaHu|F7 V|J|{ -3 }|-3|(520-5)|[1134]|B|13|0|9|4|-52|61|Lo||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Vincit|1327|C8987A9-8|403||M|DuCf|M1 V|J|{ +0 }|0|(C68+1)|[8769]|BC|15|576|9|4|-52|67|Ag Pi|M|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Andiros|1328|C799566-8|601|||NaHu|M3 V|J|{ -2 }|-2|(C42-3)|[4347]|B|12|-288|9|4|-52|68|Ni O:1428||Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Kingston|1428|B764994-C|213||KM|NaHu|M0 V|J|{ +4 }|4|(F8F+2)|[7D3A]|Bc|15|3600|9|4|-51|68|Hi Pr|F|Reaver's Deep|Ea|Reav|Non-Aligned, Human-dominated",
		// "|Fort William|1521|C540467-9|601||M|DuCf|M2 V M4 V|J|{ -1 }|-1|(732-1)|[4359]|B|8|-42|9|4|-50|61|De He Ni Po Mr|M|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Fulton|1524|C98A788-9|102|||DuCf|G3 V M5 V|J|{ +1 }|1|(B6A+1)|[7859]|BC|12|660|9|4|-50|64|Ri Wa||Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Ranald|1526|C556544-9|622|||DuCf|M0 V|J|{ +0 }|0|(B44-2)|[3537]|BC|10|-352|9|4|-50|66|Ag Ni||Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Invermory|1622|B584789-A|522||KM|DuCf|M3 V M5 V|J|{ +5 }|5|(G6E+5)|[8C6B]|BC|10|6720|9|4|-49|62|Ag Ri|F|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Duncinae|1624|A686648-B|502||KM|DuCf|G0 V|J|{ +4 }|4|(A58+4)|[6A5B]|BC|10|1600|9|4|-49|64|Ag Ni Ga Ri Cx|F|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Just|1625|C7487AA-5|420|A|M|DuCf|K4 V|J|{ +0 }|0|(967+2)|[9777]|BC|7|756|9|4|-49|65|Ag Pi Pz|M|Reaver's Deep|Ea|Reav|Confederation of Duncinae",
		// "|Lajanjigal|1721|DAB6513-8|801|A||NaHu|G8 V|K|{ -3 }|-3|(C41-5)|[2225]|B|14|-240|10|5|-48|61|Fl Ni Da (Languljigee)||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Coventry|1723|X565733-2|404|R||NaHu|G8 V|K|{ +0 }|0|(965-3)|[4721]|BC|13|-810|10|5|-48|63|Ag Ri Fo Px||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Traneer|1727|E576679-6|323|||NaHu|G5 V|K|{ -2 }|-2|(852-1)|[7467]|BC|13|-80|10|5|-48|67|Ag Ni||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Dakaar|1821|B425612-B|202||KM|NaHu|K0 V D|K|{ +2 }|2|(A56-2)|[2817]|B|6|-600|10|5|-47|61|Ni|F|Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Drexilthar|1826|B56969D-7|914|A|S|CsIm|G4 V|K|{ +0 }|0|(854+4)|[A69B]|BC|14|640|10|5|-47|66|Ni Ri Da (Iltharans)|S|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		// "|Kraan|1828|C501456-8|223|||NaHu|M3 V|K|{ -2 }|-2|(B31-3)|[3247]|B|11|-99|10|5|-47|68|Ic Ni Va||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Daken|1830|C631233-9|902|||NaHu|M2 V M5 V|K|{ -1 }|-1|(610-4)|[1126]|B|7|0|10|5|-47|70|Lo Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Cassandra|1924|B000538-C|914|||NaHu|F0 V|K|{ +1 }|1|(C45+1)|[565C]|B|15|240|10|5|-46|64|As Ni Va||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Outpost|1926|B310442-D|413||N|CsIm|K7 V M5 V|K|{ +1 }|1|(A34-3)|[1519]|B|15|-360|10|5|-46|66|Ni|N|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		// "|Tashrakaar|1927|D651695-6|213|||NaHu|M2 V|K|{ -3 }|-3|(851-5)|[4334]|B|14|-200|10|5|-46|67|Ni Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Luushakaan|2021|D541513-5|811|||NaHu|M0 V|K|{ -3 }|-3|(741-5)|[2222]|B|11|-140|10|5|-45|61|He Ni Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Dutrissal|2027|CAC5235-9|802|||NaHu|F4 V M8 V|K|{ -1 }|-1|(610-3)|[1137]|B|9|0|10|5|-45|67|Fl Lo||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Drellesarr|2029|B310550-A|502|A||NaHu|K5 V|K|{ +1 }|1|(945-3)|[1615]|B|6|-540|10|5|-45|69|Ni Da||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Drenslaar|2030|D553694-6|313|||NaHu|M2 V M5 V|K|{ -3 }|-3|(851-5)|[4334]|B|11|-200|10|5|-45|70|Ni Po||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Grendal|2127|C889855-A|613||S|CsIm|F7 V M6 V|K|{ +2 }|2|(E7B+1)|[6A38]|BC|15|1078|10|5|-44|67|Ri Ph|S|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		// "|Tharrill|2128|C885741-8|722|||NaHu|M0 V M3 V|K|{ +1 }|1|(D69-3)|[3814]|BC|14|-2106|10|5|-44|68|Ag Ga Ri||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Sarrad|2129|D88A300-8|823|A||CaAs|K8 V M3 V|K|{ -3 }|-3|(A20-5)|[1113]|B|12|0|10|5|-44|69|Lo Wa Da||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		// "|Garrison|2221|A55796B-B|412||N|ImDi|F5 V|K|{ +3 }|3|(E8E+5)|[BC7D]|BE|10|7840|10|5|-43|61|Hi Mr|N|Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		// "|Kaaniir|2223|C688611-6|723||S|CsIm|F8 V|K|{ +0 }|0|(854-4)|[2612]|BC|12|-640|10|5|-43|63|Ag Ni Ri|S|Reaver's Deep|Drexilthar|Reav|Client state, Third Imperium",
		// "|Yarhfahl|2228|C658796-6|110|||NaHu|M1 V M4 V|K|{ +0 }|0|(967-1)|[6745]|BC|11|-378|10|5|-43|68|Ag||Reaver's Deep|Drexilthar|Reav|Non-Aligned, Human-dominated",
		// "|Datinar|2230|B431685-A|303|||CaAs|K0 V|K|{ +1 }|1|(B55-1)|[4738]|B|11|-275|10|5|-43|70|Na Ni Po||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		// "|Gaargir|2322|B565304-C|713||NS|ImDi|K7 V|K|{ +2 }|2|(922+1)|[153A]|B|7|36|10|5|-42|62|Lo|A|Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		// "|Ildrissar|2326|C995836-7|203|A||CaAs|M1 V|K|{ -1 }|-1|(A77-2)|[7746]|Bc|14|-980|10|5|-42|66|Pa Ph Pi Pz||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		// "|Carrill|2330|A0009AE-E|613|A|KM|CaAs|F4 V|K|{ +5 }|5|(F8H+5)|[DE9J]|BE|11|10200|10|5|-42|70|As Hi In Na Va Cx Pz|F|Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		// "|Kaanash|2421|B55687A-7|414|A|NS|ImDi|M1 V M6 V|K|{ +1 }|1|(A79+3)|[A979]|Bce|17|1890|10|5|-41|61|Pa Ph Pz|A|Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		// "|Diablo|2423|B9C7477-9|724|||ImDi|M2 V|K|{ +0 }|0|(C33+1)|[4459]|B|11|108|10|5|-41|63|Fl Ni||Reaver's Deep|Drexilthar|Reav|Third Imperium, Domain of Ilelish",
		// "|Lindritar|2429|C5796A7-8|211|||CaAs|K7 V|K|{ -2 }|-2|(A52-2)|[6458]|B|13|-200|10|5|-41|69|Ni||Reaver's Deep|Drexilthar|Reav|Carrillian Assembly",
		// "|Inura|2523|C421312-A|812|||ImDi|G9 V|L|{ +0 }|0|(820-4)|[1316]|B|6|0|11|5|-40|63|He Lo Po||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		// "|Hope|2526|E65778B-4|200|A||ImDi|F4 V|L|{ -1 }|-1|(965+1)|[9676]|BC|8|270|11|5|-40|66|Ag Ga Pz||Reaver's Deep|Urlaqqash|Reav|Third Imperium, Domain of Ilelish",
		// "|Yaggoth|2530|B864756-B|222|||CaAs|F8 V M1 V M6 V|L|{ +4 }|4|(D6E+3)|[6B4A]|BC|13|3276|11|5|-40|70|Ag Ri||Reaver's Deep|Urlaqqash|Reav|Carrillian Assembly",
	} {
		world := survey.Parse(data)

		fmt.Println(world)
		land, err := landing.Preapare(world)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			
			fmt.Println(land)
		}
		fmt.Println("  ")
		fmt.Println("  ")
		fmt.Println("  ")

	}
}
