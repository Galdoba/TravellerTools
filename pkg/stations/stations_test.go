package stations

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TestStationsReport(t *testing.T) {
	var world World
	for _, data := range []string{
		"|Regina                      |1910|A788899-C|703| |NS|ImDd|F7 V BD M3 V                 |C|{ 4 }      |4 |(D7E+5)|[9C6D]|BcCeF|8 |6370 |2 |1|-110 |-70  |Ri Pa Ph An Cp (Amindii)2 Varg0 Asla0 Sa    |A |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Pixie                       |1903|A100103-D|901| |N |ImDd|K1 V M0 V                    |C|{ 1 }      |1 |(401-2)|[122A]|B    |14|0    |2 |1|-110 |-77  |Lo Va An Px                                 |N |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Boughene                    |1904|A8B3531-D|601| |S |ImDd|M1 V                         |C|{ 1 }      |1 |(845-3)|[1619]|B    |13|-480 |2 |1|-110 |-76  |Fl Ni An                                    |S |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Hefry                       |1909|C200423-7|320| |S |ImDd|K6 II M2 V                   |C|{ -2 }     |-2|(631-5)|[1224]|B    |13|-90  |2 |1|-110 |-71  |Ni Va                                       |S |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Regina                      |1910|A788899-C|703| |NS|ImDd|F7 V BD M3 V                 |C|{ 4 }      |4 |(D7E+5)|[9C6D]|BcCeF|8 |6370 |2 |1|-110 |-70  |Ri Pa Ph An Cp (Amindii)2 Varg0 Asla0 Sa    |A |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Feri                        |2005|B584879-B|620| |S |ImDd|G4 V M3 V                    |C|{ 3 }      |3 |(C7D+4)|[9B6C]|BcCe |14|4368 |2 |1|-109 |-75  |Ri Pa Ph                                    |S |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Roup                        |2007|C77A9A9-7|323|A|S |ImDd|F9 V                         |C|{ 1 }      |1 |(B8A+2)|[AA68]|BE   |14|1760 |2 |1|-109 |-73  |Hi In Wa Pz                                 |S |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Pscias                      |2106|X555423-2|501|R|  |ImDd|K5 V                         |C|{ -3 }     |-3|(631-5)|[1121]|     |14|-90  |2 |1|-108 |-74  |Ni Pa Fo                                    |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Yori                        |2110|C560757-A|713| |  |ImDd|F1 V                         |C|{ 2 }      |2 |(D6B+2)|[795A]|BC   |15|1716 |2 |1|-108 |-70  |De Ri An (Zhurphani)6 RsB                   |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Dentus                      |2201|C979500-A|920| |S |ImDd|M2 V                         |C|{ 0 }      |0 |(944-4)|[1515]|B    |10|-576 |2 |1|-107 |-79  |Ni                                          |S |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Kinorb                      |2202|A663659-8|622| |  |ImDd|G7 V                         |C|{ 0 }      |0 |(C54+1)|[7669]|BC   |7 |240  |2 |1|-107 |-78  |Ni Ri                                       |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Beck's World                |2204|D88349D-4|701|A|  |ImDd|K0 V M2 V                    |C|{ -3 }     |-3|(631+1)|[8198]|B    |10|18   |2 |1|-107 |-76  |Ni An Da                                    |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Enope                       |2205|C411988-7|600| |  |ImDd|K6 V M5 V                    |C|{ 1 }      |1 |(B8A+1)|[9A57]|BE   |8 |880  |2 |1|-107 |-75  |Hi Ic In Na                                 |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Wochiers                    |2207|EAC28CC-9|703|A|  |ImDd|F0 V                         |C|{ -1 }     |-1|(D78+2)|[B78C]|Be   |16|1456 |2 |1|-107 |-73  |Fl He Ph Pz                                 |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Yorbund                     |2303|C7C6503-9|220| |  |ImDd|M3 V                         |C|{ -1 }     |-1|(943-4)|[2426]|B    |9 |-432 |2 |1|-106 |-77  |Fl Ni                                       |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
		"|Shionthy                    |2306|C000742-8|714|R|  |ImDd|F6 V                         |C|{ -1 }     |-1|(E67-5)|[3614]|     |14|-2940|2 |1|-106 |-74  |As Na Va An Pi Fo                           |  |Spinward Marches                 |Regina                  |Spin|Third Imperium, Domain of Deneb                                                              ",
	} {
		world = survey.Parse(data)
		report, err := GenerateSystemReport(world)
		if report == nil {
			t.Errorf("GenerateSystemReport(world) returned no object")
			continue
		}
		if err != nil {
			t.Errorf("GenerateSystemReport(world) returned error: %v", err)
		}
		if report.name == "" {
			t.Errorf("report.name not filled")
		}
		if report.uwp == "" {
			t.Errorf("report.uwp not filled")
		}
		if report.Present == nil {
			t.Errorf("report.Present not initiated")
		}
		if report.Present[Type_ANY] != report.Present[Type_Naval]+report.Present[Type_Paramilitary]+report.Present[Type_Commercial]+report.Present[Type_Imperial]+report.Present[Type_Pirate] {
			t.Errorf("report.Present[Type_ANY] != report.Present[sum of types]")
		}
		if report.Present[Type_Naval] != report.Present[Type_Naval_Defence]+report.Present[Type_Naval_Fleet]+report.Present[Type_Naval_Interdiction]+report.Present[Type_Naval_Shipyard] {
			t.Errorf("report.Present[Type_Naval] != report.Present[sum of types]")
		}
		fmt.Println(" ")
		fmt.Println(report.Summary())
		fmt.Println(" ")
	}

}
