package military

import (
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	u := Unit{
		force:     TYPE_STARSHIP,
		tl:        9,
		attack:    6,
		defence:   6,
		transport: 6,
		jump:      18,
	}
	fmt.Println(u.UMP())
	pc := u.PurchaseCost()
	if pc < 0 {
		for _, df := range u.DesignFlaw() {
			fmt.Println(df)
		}
	} else {
		fmt.Println("Cost:", u.PurchaseCost())
	}
}
