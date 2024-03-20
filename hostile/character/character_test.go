package character

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	gen := NewGenerator()
	// gen.dice = dice.New()
	for i := 0; i < 999; i++ {
		fmt.Println("test", i)
		chr, err := gen.Generate()
		if err != nil {
			t.Errorf("%v", err)
			i = 1000
		}
		fmt.Println("=========")
		fmt.Println(chr)
	}
}
