package survey

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/generation/stellar"
)

func TestConstructor(t *testing.T) {
	ssd, err := NewSecondSurvey(
		Instruction(MW_Name, "Ernl23"),
		Instruction(MW_UWP, "C6658A7-9"),
		Instruction(Sector, "NSY-S"),
		Instruction(CoordX, "0"),
		Instruction(CoordY, "0"),
		Instruction(Hex, "0202"),
		Instruction(Bases, "NS"),
		Instruction(Seed, "_SEED_PREFIX"),
	)
	if err != nil {
		t.Errorf("func error: %v", err.Error())
	}
	fmt.Println("  ")
	fmt.Println(ssd)
	fmt.Println(ssd.NameByConvention())
	fmt.Println(ssd.Compress())
	fmt.Println(ssd.GenerationSeed())
	nexus, err := stellar.NewNexus(ssd)

	fmt.Println("  ")
	fmt.Println(err)
	fmt.Println(nexus)
}
