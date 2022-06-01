package sai

import (
	"fmt"
	"testing"
)

func TestSActivity(t *testing.T) {
	portList := []port{
		{"Atiyr", "E413774-8", "Ic Na Pi", ""},
		{"Oihoiiei", "A8558A8-C", "Ga Pa Ph", ""},
	}
	for i, p := range portList {
		fmt.Println("test", i, ":", p)
		sa, err := NewShippingActivity(&p, []int{0, 5, 4})
		if sa == nil {
			t.Errorf("func NewShippingActivity() returned no object")
			return
		}
		if err != nil {
			t.Errorf("func NewShippingActivity() returned error: %v", err.Error())
		}
		if sa.averageShips == -1000 {
			t.Errorf("sa.averageShips was not adressed")
		}
		if sa.minmumShips == -1000 {
			t.Errorf("sa.minmumShips was not adressed")
		}
		if sa.maximumShips == -1000 {
			t.Errorf("sa.maximumShips was not adressed")
		}
		if sa.shipsByTonnage == nil {
			t.Errorf("sa.shipsByTonnage was not adressed")
		}
		if sa.traffMult == -1000 {
			t.Errorf("sa.traffMult was not adressed")
		}
		if sa.traffDm == -1000 {
			t.Errorf("sa.traffDm was not adressed")
		}
		fmt.Println(sa.minmumShips, sa.maximumShips, sa.averageShips)
		fmt.Println(sa)
	}
}
