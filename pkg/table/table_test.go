package table

import (
	"fmt"
	"testing"
)

func TestMods(t *testing.T) {
	m1 := newModifier(2, Equal, 5)
	m1.Verify(3, 6, 2, 8)
	fmt.Println(m1.Value())
}
