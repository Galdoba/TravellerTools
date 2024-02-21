package table

import (
	"fmt"
	"testing"
)

func TestMods(t *testing.T) {
	m1 := newModifier(2, MoreOrEqual, 8)
	m1.Verify(3, 6, 2, 8)
	fmt.Println(m1.Value())
}
