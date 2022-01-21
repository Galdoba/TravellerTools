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
	}
}

func TestStarNexus(t *testing.T) {
	for _, input := range testLines() {
		ssd := survey.Parse(input)
		nx, err := NewNexus(ssd)
		fmt.Println(err)
		nx.Print()
	}
}
