package world

import (
	"fmt"
	"testing"
)

func TestGenome(t *testing.T) {

	wrld, err := NewWorld(Inject(
		KnownData(Catalog, "111-222"),
		KnownData("Mainworld", FLAG_TRUE),
	))
	fmt.Println(wrld, err)
}
