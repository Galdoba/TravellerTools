package hyperjump

import (
	"fmt"
	"testing"
)

func Test_Hyper(t *testing.T) {
	hj := New(2, 2)
	fmt.Println(hj.Report())
	fmt.Println(hj.Outcome())

}
