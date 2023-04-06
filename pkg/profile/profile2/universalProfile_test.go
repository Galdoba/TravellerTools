package profile2

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	up := New(PROFILE_PERSONALITY)
	fmt.Println(up)
	fmt.Println(up.Data("C2").Code())
	fmt.Println(up)
	up.Inject("C2", 7)
	fmt.Println(up)
	fmt.Println(up.StringFull())
	up2 := New(PROFILE_WORLD)
	up2.Inject(KEY_PORT, "A")
	up2.InjectAll("A123456-9")
	fmt.Println(up2.String())
}
