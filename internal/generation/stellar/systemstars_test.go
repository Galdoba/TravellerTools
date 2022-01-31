package stellar

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/internal/struct/star"
	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/Galdoba/TravellerTools/pkg/survey/calculations"
)

func TestNewStellar(t *testing.T) {
	lenMap := make(map[int]int)
	d := 1000
	for i := d; i < d+50; i++ {
		name := fmt.Sprintf("Test Stellar %v", i)
		stellar := calculations.GenerateNewStellar(name)
		stars, err := star.ParseStellar(stellar)
		lenMap[len(stars)]++
		if err != nil {
			t.Errorf("nexpected: %v, [%v] {%v}", err, stellar, stars)
			break
		}
	}
}

func testLines() []string {
	return []string{
		"|Drinax|2223|A43645A-E|714|||NaHu|M1 V|K|{ +1 }|1|(B34+3)|[657G]|B|9|396|10|5|-107|-17|Ni||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated",
		"|Iroioah|2227|B530113-C|823|||AsMw|M3 V M4 V|K|{ +1 }|1|(801-2)|[1229]|B|12|0|10|5|-107|-13|De Lo Po||Trojan Reach|Tlaiowaha|Troj|Aslan Hierate, single multiple-world clan dominates",
		"|Regina|1910|A788899-C|703||NS|ImDd|F7 V BD M3 V|C|{ +4 }|4|(D7E+5)|[9C6D]|BcCeF|8|6370|2|1|-110|-70|Ri Pa Ph An Cp (Amindii)2 Varg0 Asla0 Sa|A|Spinward Marches|Regina|Spin|Third Imperium, Domain of Deneb",
		"|Rushu|0215|E766674-4|903|||CsZh|G0 V M3 V|E|{ -1 }|-1|(853-3)|[4532]|BC|8|-360|4|0|-127|-65|Ag Ni Ga Ri VargW||Spinward Marches|Querion|Spin|Client state, Zhodani Consulate",
		"",
	}
}

func TestStarNexus(t *testing.T) {
	for _, input := range testLines() {
		fmt.Println(input)
		ssd := survey.Parse(input)
		fmt.Println(ssd.MW_Name())
		nx, err := NewNexus(ssd)
		fmt.Println(err)
		fmt.Println(nx)
	}
}
