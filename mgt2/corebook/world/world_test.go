package world

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := NewConstructor(
		Instruction(KEY_SEED, "Test"),
		Instruction(KEY_NAME, "Drahma"),
	)
	w, err := c.Create()

	fmt.Println(w)
	for _, e := range w.statistics {
		fmt.Printf("%v", e.Code())
	}
	fmt.Println("")
	fmt.Println(err)

}
