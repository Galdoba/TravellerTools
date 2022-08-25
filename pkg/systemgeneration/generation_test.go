package systemgeneration

import "testing"

func TestGeneration(t *testing.T) {
	gen, _ := NewGenerator("Sol")
	gen.Step01()
}
