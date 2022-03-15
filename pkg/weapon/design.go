package weapon

import (
	"fmt"

	. "github.com/Galdoba/TravellerTools/pkg/weapon/components"
)

/*
1 Decide general type (pistol, rifle, ect...)
2 choose ammnition/powertype
3 choose receiver
4 choose receiver's mode of operation (breechloader, semi-automatic, ect...)
5 assign barrel lenght
6 assign furniture
7 choose feed device
8 add accessories that come as standard
9 total cost and weight

*/

type Weapon struct {
	rcvr            *Receiver
	brl             *Barrel
	frn             *Furniture
	acc             *Accessoire
	penetration     int
	sigPhysical     int
	sigEmmision     int
	mishapThreshold int
	///////////////////
	name           string
	tl             int
	effectiveRange int
	damageDice     int
	damageMod      int
	weight         float64
	cost           int
	magazine       int
	magazineCost   int
	quickdraw      int
	traits         []string
}

func New(instr ...int) (*Weapon, error) {
	w := Weapon{}
	err := fmt.Errorf("error was not adressed")
	if w.rcvr, err = NewReceiver(instr...); err != nil {
		return nil, err
	}
	if w.brl, err = NewBarrel(instr...); err != nil {
		return nil, err
	}
	if w.frn, err = NewFurniture(instr...); err != nil {
		return nil, err
	}
	if w.acc, err = NewAccessoires(instr...); err != nil {
		return nil, err
	}

	return &w, nil
}

//DesignWorksheet - будет возвращать в виде таблицы покомпонентный расклад оружия
func (wp *Weapon) DesignWorksheet() string {
	fmt.Println("[Weapon   TL   Range   Damage   Kg   Cost   Magazine   Magazine Cost   Traits]")
	return ""
}

//Summary - будет возвращать в виде таблицы выжимку по оружию
func (wp *Weapon) Summary() string {
	fmt.Println("[Weapon   TL   Range   Damage   Kg   Cost   Magazine   Magazine Cost   Traits]")
	return ""
}
