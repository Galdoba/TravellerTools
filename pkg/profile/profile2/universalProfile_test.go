package profile2

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	up, err := New(PROFILE_WORLD)
	fmt.Println(err)
	fmt.Println(up)
	err2 := up.Inject("A123456-8")
	fmt.Println(err2)
	fmt.Println(up)
}
