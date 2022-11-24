package empire

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TestWorldCharacter(t *testing.T) {
	ssd := survey.Parse("|Drinax|2223|A43645A-E|714|||NaHu|M1 V|K|{ +1 }|1|(B34+3)|[657G]|B|9|396|10|5|-107|-17|Ni||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated")
	wc := WorldCharacter(ssd)
	fmt.Println(wc.Descr())
}
