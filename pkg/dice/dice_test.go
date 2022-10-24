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

func TestDicepoolExtended(t *testing.T) {
	dp := New().SetSeed(1)
	fmt.Println(dp.Sroll("1d100+5"))
	dp2 := New().SetSeed(1)
	fmt.Println(dp2.Sroll("1d100+3"))

}
