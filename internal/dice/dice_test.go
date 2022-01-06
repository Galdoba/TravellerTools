package dice

import (
	"fmt"
	"testing"
)

func TestDicepool(t *testing.T) {
	dp := New()

	fmt.Println(dp.Roll("1d100+1").Sum())
	fmt.Println(dp.Roll("1d100").Sum())
	fmt.Println(dp.Roll("1d100").Sum())
	dp.Reset()
	dp.Roll("1d100")

}
