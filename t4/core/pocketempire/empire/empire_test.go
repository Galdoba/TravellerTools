package empire

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/Galdoba/TravellerTools/pkg/survey"
// 	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/economics"
// 	"github.com/Galdoba/TravellerTools/t4/core/pocketempire/empire/worldcharacter"
// )

// func Test_EmpireOfSevenStars(t *testing.T) {
// 	e := PocketEmpire{}
// 	e.Name = "Empire of Seven Stars"

// 	demand := economics.NewAggregatedDemand("test")
// 	for i := 0; i < 500; i++ {
// 		demand.Aggregate()
// 	}
// 	ssdData := []survey.SecondSurveyData{}
// 	//ssdData = append(ssdData, *survey.Parse("|Arvlaa Gam|2635|E889762-7|704|||ImDc|M3 V|P|{ -1 }|-1|(967-5)|[3613]|BC|11|-1890|15|5|25|-5|Ri O:2936||Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	ssdData = append(ssdData, *survey.Parse("|Valed|2636|B998213-B|314||S|ImDc|K7 V|P|{ +1 }|1|(911-2)|[1328]|B|8|-18|15|5|25|-4|Lo|S|Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	ssdData = append(ssdData, *survey.Parse("|Gaen Luum|2736|B424757-C|113|||ImDc|K2 V|P|{ +2 }|2|(D6C+2)|[795C]|BD|12|1872|15|5|26|-4|Pi||Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	ssdData = append(ssdData, *survey.Parse("|Khuir|2836|B578961-C|504||S|ImDc|M3 V|P|{ +4 }|4|(F8F+1)|[5D18]|BEf|13|1800|15|5|27|-4|Hi In Mr|S|Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	//ssdData = append(ssdData, *survey.Parse("|Igla|2837|B414ADG-E|804|A||ImDc|M2 V M9 V|P|{ +4 }|4|(G9G+5)|[FEAK]|BEf|7|11520|15|5|27|-3|Hi Ic In Pz||Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	//ssdData = append(ssdData, *survey.Parse(""))
// 	//ssdData = append(ssdData, *survey.Parse(""))
// 	ssdData = append(ssdData, *survey.Parse("|Saregon|2936|A584A76-F|114||NS|ImDc|M3 V M6 V|P|{ +4 }|4|(H9G+3)|[9E4E]|BEF|11|7344|15|5|28|-4|Hi Cp|A|Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	ssdData = append(ssdData, *survey.Parse("|Uurigger|2937|B434779-B|703||NS|ImDc|M1 V M9 V|P|{ +3 }|3|(C6D+4)|[8A6C]|B|11|3744|15|5|28|-3||A|Core|Saregon|Core|Third Imperium, Domain of Sylea"))
// 	//ssdData = append(ssdData, *survey.Parse(""))
// 	//ssdData = append(ssdData, *survey.Parse(""))
// 	for _, ssd := range ssdData {
// 		wc, err := worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		//fmt.Println(wc.Descr())

// 		err = wc.EconomicProcess("001-021", demand)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		e.integrateWorld(wc)
// 		fmt.Println("Integrated:", wc.UWPE())
// 		//fmt.Println(e)
// 	}
// 	fmt.Println(e)
// 	e.calculateSelfDetermination()
// 	e.calculatePopularity()
// 	e.calculatePopulation()
// 	e.calculatePrestige()

// 	fmt.Println(e.UEP())

// }

// func TestWorldCharacter(t *testing.T) {
// 	return
// 	id := economics.NewAggregatedDemand("test")

// 	for i := 0; i < 500; i++ {
// 		id.Aggregate()
// 	}
// 	fmt.Println(id.String())

// 	ssd := survey.Parse("|Khuir|2836|B578961-C|504||S|ImDc|M3 V|P|{ +4 }|4|(F8F+1)|[5D18]|BEf|13|1800|15|5|27|-4|Hi In Mr|S|Core|Saregon|Core|Third Imperium, Domain of Sylea")
// 	wc, err := worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//fmt.Println(wc.Descr())
// 	//fmt.Println(wc.EconomyState().StatBlock())
// 	//dp := dice.New().SetSeed("1")
// 	err = wc.EconomicProcess("001-021", id)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	//fmt.Println(wc.EconomyState().StatBlock())
// 	//|Asim|2123|B867564-6|203|||NaHu|F2 V|K|{ +0 }|0|(744-2)|[3534]|Bc|10|-224|10|5|-108|-17|Ag Ni Ga Pr O:2324||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated
// 	ssd = survey.Parse("|Marduk|2120|C577436-5|503|||NaHu|K7 V|G|{ -2 }|-2|(631-3)|[3244]|Bc|11|-54|6|1|-108|-20|Ni Pa||Trojan Reach|Sindal|Troj|Non-Aligned, Human-dominated")
// 	wc, err = worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//fmt.Println(wc.Descr())
// 	//fmt.Println(wc.EconomyState().StatBlock())
// 	//dp := dice.New().SetSeed("1")
// 	err = wc.EconomicProcess("001-021", id)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	//fmt.Println(wc.EconomyState().StatBlock())

// 	ssd = survey.Parse("|Earth|1910|A867A79-B|813||NS|ImDd|F7 V BD M3 V|C|{ +4 }|4|(D7E+5)|[9C6D]|BcCeF|8|6370|2|1|-110|-70|Ri Pa Ph An Cp (Amindii)2 Varg0 Asla0 Sa|A|Spinward Marches|Regina|Spin|Third Imperium, Domain of Deneb")
// 	wc, err = worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//fmt.Println(wc.Descr())
// 	//fmt.Println(wc.EconomyState().StatBlock())
// 	//dp := dice.New().SetSeed("1")
// 	err = wc.EconomicProcess("001-021", id)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(wc.EconomyState().StatBlock())

// }
