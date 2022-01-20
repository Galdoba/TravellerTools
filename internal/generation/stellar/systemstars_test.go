package stellar

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/internal/struct/star"
)

func TestNewStellar(t *testing.T) {
	lenMap := make(map[int]int)
	d := 1000
	for i := d; i < d+50; i++ {
		name := fmt.Sprintf("Test Stellar %v", i)
		stellar := GenerateNewStellar(name)
		stars, err := star.ParseStellar(stellar)
		lenMap[len(stars)]++
		if err != nil {
			t.Errorf("nexpected: %v, [%v] {%v}", err, stellar, stars)
			break
		}
	}
}

func TestStarNexus(t *testing.T) {
	d := 150
	for i := d; i < d+50; i++ {
		name := fmt.Sprintf("Test Stellar %v", i)
		stellar := GenerateNewStellar(name)
		pbg := "012"
		w := 15
		nx, _ := NewNexus(name, stellar, pbg, w)
		//fmt.Println(err)
		nx.Print()
	}
}
