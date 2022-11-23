package military

import (
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	u := Unit{
		force:     TYPE_GROUND,
		tl:        13,
		attack:    1,
		defence:   2,
		transport: 6,
		jump:      12,
	}
	fmt.Println(u.UMP())
	fmt.Println(u.PurchaseCost())
}
