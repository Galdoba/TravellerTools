package stellar

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/internal/struct/star"
)

func TestNewStellar(t *testing.T) {
	lenMap := make(map[int]int)
	d := 50
	for i := d; i < d+50; i++ {
		name := fmt.Sprintf("Test Stellar %v", i)
		stellar := GenerateNewStellar(name)
		stars, err := star.ParseStellar(stellar)
		lenMap[len(stars)]++
		if err != nil {
			t.Errorf("nexpected: %v, [%v] {%v}", err, stellar, stars)
			break
		}
		//fmt.Printf("name: %v | stellar = [%v]                        \n", name, stellar)
		compos, err := SystemComposition(name, stellar)
		separated := separateBySystems(compos)

		//fmt.Println(compos, err)
		fmt.Printf("lenMap=%v | total=%v | %v | %v                \n", lenMap, i+1, compos, separated)

	}
}
