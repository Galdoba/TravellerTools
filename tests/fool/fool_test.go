package fool

import (
	"fmt"
	"testing"
)

func Test_Game(t *testing.T) {
	g := NewGame(2)
	fmt.Println(g)
}
