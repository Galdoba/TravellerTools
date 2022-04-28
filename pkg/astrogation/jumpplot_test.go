package astrogation

import (
	"fmt"
	"testing"
)

func Test_PlotCourse(t *testing.T) {
	fmt.Println("------------")
	fmt.Println("Plot Cource from", NewCoordinates(-107, -17), "to", NewCoordinates(-97, -25))
	jp, err := PlotCourse(NewCoordinates(-107, -17), NewCoordinates(-97, -25))
	fmt.Println(jp)
	fmt.Println(err)
}
