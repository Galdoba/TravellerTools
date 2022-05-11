package astrogation

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/astarhex"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func createField() map[Coordinates]int {
	field := make(map[Coordinates]int)
	center := NewCoordinates(-100, -17)
	for _, void := range JumpMap(&center, 5) {
		//fmt.Printf("Evaluating: %v ", void)
		//fmt.Println(EvaluateMovementWeight(&void))

		field[void] = 100000
	}

	// field := make(map[Coordinates]int)
	// sperle := NewCoordinates(-101, -16)
	// for _, void := range JumpMap(&sperle, 5) {
	// 	field[void] = 1000
	// }
	field[NewCoordinates(-104, -13)] = 10 // Byrni
	field[NewCoordinates(-103, -19)] = 10 // Arunisiir
	field[NewCoordinates(-103, -16)] = 10 // Tech
	field[NewCoordinates(-103, -15)] = 10 // Ergo
	field[NewCoordinates(-102, -19)] = 10 // Tanith
	field[NewCoordinates(-102, -18)] = 10 // Acrid
	field[NewCoordinates(-102, -16)] = 10 // Inurin
	field[NewCoordinates(-102, -15)] = 10 // Falcon
	field[NewCoordinates(-101, -19)] = 10 // Cordan
	field[NewCoordinates(-101, -17)] = 10 // Exe
	field[NewCoordinates(-101, -16)] = 10 // Sperle
	field[NewCoordinates(-100, -19)] = 10 // Umemii
	field[NewCoordinates(-100, -17)] = 10 // Argona
	field[NewCoordinates(-99, -11)] = 10  // Villane
	field[NewCoordinates(-99, -10)] = 10  // Browne

	return field
}

type positions struct {
	field map[Coordinates]int
	start Coordinates
	end   Coordinates
	path  []Coordinates
}

func Test_PlotCourse(t *testing.T) {

	ast, err := astarhex.New(astarhex.Config{1})
	fmt.Println(ast, err)
	path, err2 := ast.FindPath(*astarhex.SetNodeHex(-103, -16), *astarhex.SetNodeHex(-100, -17))
	fmt.Println(path, err2)
	for _, node := range path {
		world, err3 := survey.SearchByCoordinates(node.HexCoords())
		x, y := node.HexCoords()
		if world == nil {

			fmt.Println("world.MW_Name()", err3, x, y)
		} else {
			fmt.Println(world.MW_Name(), err3, x, y)
		}

	}

}
