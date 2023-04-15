package pawn

import (
	"fmt"
	"testing"
	"time"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/t5/genetics"
)

func TestPawn(t *testing.T) {
	for i := 0; i < 3; i++ {
		gt := genetics.NewTemplate("SDEIES", "222222")
		chr, err := New(control_Random, gt, classifications.ListAll()) //[]string{"Fa", "Ag"})
		fmt.Println(chr.chrSet)
		fmt.Println(chr.sklSet)
		fmt.Println(err)
		time.Sleep(time.Millisecond * 100)
	}

}
