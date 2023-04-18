package profile

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	up := New()
	fmt.Println(up)
	fmt.Println(up.Data("C2").Code())
	fmt.Println(up)
	up.Inject("C2", 7)
	fmt.Println(up)
}
