package empire

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/economics"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/empire/worldcharacter"
)

func TestWorldCharacter(t *testing.T) {
	id := economics.NewAggregatedDemand("test")

	for i := 0; i < 500; i++ {
		id.Aggregate()
	}
	fmt.Println(id.String())

	ssd := survey.Parse("|Drinax|2223|A43645A-E|714|||NaHu|M1 V|K|{ +1 }|1|(B34+3)|[657G]|B|9|396|10|5|-107|-17|Ni||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated")
	wc, err := worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wc.Descr())
	//fmt.Println(wc.EconomyState().StatBlock())
	//dp := dice.New().SetSeed("1")
	err = wc.EconomicProcess("001-021", id)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(wc.EconomyState().StatBlock())
	//|Asim|2123|B867564-6|203|||NaHu|F2 V|K|{ +0 }|0|(744-2)|[3534]|Bc|10|-224|10|5|-108|-17|Ag Ni Ga Pr O:2324||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated
	ssd = survey.Parse("|Marduk|2120|C577436-5|503|||NaHu|K7 V|G|{ -2 }|-2|(631-3)|[3244]|Bc|11|-54|6|1|-108|-20|Ni Pa||Trojan Reach|Sindal|Troj|Non-Aligned, Human-dominated")
	wc, err = worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wc.Descr())
	//fmt.Println(wc.EconomyState().StatBlock())
	//dp := dice.New().SetSeed("1")
	err = wc.EconomicProcess("001-021", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(wc.EconomyState().StatBlock())

	ssd = survey.Parse("|Earth|1910|A867A79-B|813||NS|ImDd|F7 V BD M3 V|C|{ +4 }|4|(D7E+5)|[9C6D]|BcCeF|8|6370|2|1|-110|-70|Ri Pa Ph An Cp (Amindii)2 Varg0 Asla0 Sa|A|Spinward Marches|Regina|Spin|Third Imperium, Domain of Deneb")
	wc, err = worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wc.Descr())
	//fmt.Println(wc.EconomyState().StatBlock())
	//dp := dice.New().SetSeed("1")
	err = wc.EconomicProcess("001-021", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(wc.EconomyState().StatBlock())

}
