package empire

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/survey"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/empire/worldcharacter"
)

func TestWorldCharacter(t *testing.T) {
	ssd := survey.Parse("|Drinax|2223|A43645A-E|714|||NaHu|M1 V|K|{ +1 }|1|(B34+3)|[657G]|B|9|396|10|5|-107|-17|Ni||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated")
	wc, err := worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wc.Descr())
	fmt.Println(wc.EconEx.StatBlock())
	dp := dice.New().SetSeed("1")
	fmt.Println(wc.EconEx.RecalculateRA(dp))
	fmt.Println(wc.EconEx.StatBlock())
	//|Asim|2123|B867564-6|203|||NaHu|F2 V|K|{ +0 }|0|(744-2)|[3534]|Bc|10|-224|10|5|-108|-17|Ag Ni Ga Pr O:2324||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated
	ssd = survey.Parse("|Asim|2123|B867564-6|203|||NaHu|F2 V|K|{ +0 }|0|(744-2)|[3534]|Bc|10|-224|10|5|-108|-17|Ag Ni Ga Pr O:2324||Trojan Reach|Tlaiowaha|Troj|Non-Aligned, Human-dominated")
	wc, err = worldcharacter.WorldCharacter(ssd.MW_Name(), ssd.MW_UWP(), ssd.PBG(), ssd.CoordX(), ssd.CoordY())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wc.Descr())
	fmt.Println(wc.EconEx.StatBlock())
	dp2 := dice.New().SetSeed("1")
	fmt.Println(wc.EconEx.RecalculateRA(dp2))
	fmt.Println(wc.EconEx.StatBlock())
}
