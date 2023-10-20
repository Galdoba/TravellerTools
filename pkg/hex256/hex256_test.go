package hex256

import (
	"fmt"
	"testing"
)

func TestHex256(t *testing.T) {
	eh := &hx256{}
	btOrder := []byte{
		ByteOf("0"),
		ByteOf("R"),
		ByteOf("T"),
		ByteOf("1"),
		ByteOf("2"),
		ByteOf("3"),
		ByteOf("4"),
		ByteOf("5"),
		ByteOf("6"),
		ByteOf("7"),
		ByteOf("8"),
		ByteOf("9"),
	}
	eh.Set(btOrder)
	fmt.Println(eh)
	eh.Set("T")
	b, _ := New(3)
	eh.Add(b)
	fmt.Println(eh)

}
