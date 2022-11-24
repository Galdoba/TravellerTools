package military

import (
	"fmt"
)

const (
	TYPE_GROUND = iota
	TYPE_AIR
	TYPE_SPACESHIP
	TYPE_STARSHIP
	TYPE_DEPOT
)

type Unit struct {
	force          int
	tl             int
	attack         int
	defence        int
	transport      int
	jump           int
	size           int // Att + def + trn + jmp
	jumpCapability int //jmp / (att + def + jmp)
}

func (un *Unit) UMP() string {
	return fmt.Sprintf("A%v D%v T%v J%v (%v Size/%v Range)", un.attack, un.defence, un.transport, un.jump, un.Size(), un.JumpCapability())
}

func (un *Unit) Size() int {
	return un.attack + un.defence + un.transport + un.jump
}

func (un *Unit) JumpCapability() int {
	if un.jump == 0 {
		return 0
	}
	return un.jump / (un.attack + un.defence + un.transport)
}

func (un *Unit) EffectiveRange() int {
	return un.JumpCapability()
}

func (un *Unit) ExtendedRange() int {
	return un.JumpCapability() * 3
}

func (un *Unit) PurchaseCost() int {
	adCost := []int{
		18000, //0
		12000, //1
		7500,  //2
		5000,  //3
		3500,  //4
		2200,  //5
		1500,  //6
		1000,  //7
		700,   //8
		450,   //9
		300,   //10
		200,   //11
		150,   //12
		100,   //13
		75,    //14
		50,    //15
		35,    //16
		25,    //17
		20,    //18
		15,    //19
		12,    //20
		10,    //21
	}
	tCost := []int{
		1000000, //0
		1000000, //1
		1000000, //2
		1000000, //3
		1000000, //4
		1000000, //5
		1000000, //6
		40,      //7
		20,      //8
		10,      //9
		10,      //10
		10,      //11
		10,      //12
		10,      //13
		10,      //14
		10,      //15
		10,      //16
		10,      //17
		10,      //18
		10,      //19
		10,      //20
		10,      //21
	}
	jCost := []int{
		1000000, //0
		1000000, //1
		1000000, //2
		1000000, //3
		1000000, //4
		1000000, //5
		1000000, //6
		40,      //7
		20,      //8
		10,      //9
		10,      //10
		10,      //11
		10,      //12
		10,      //13
		10,      //14
		10,      //15
		10,      //16
		10,      //17
		10,      //18
		10,      //19
		10,      //20
		10,      //21
	}
	df := len(un.DesignFlaw())
	switch df {
	case 0:
		return ((un.attack + un.defence) * adCost[un.tl]) + (un.transport * tCost[un.tl]) + (un.jump * jCost[un.tl])
	default:
		return -1 * df
	}
}

func (un *Unit) DesignFlaw() []string {
	df := []string{}
	switch {
	case un.tl < 0:
		df = append(df, fmt.Sprintf("TL is less than 0"))
	case un.attack < 0:
		df = append(df, fmt.Sprintf("Attack is less than 0"))
	case un.defence < 0:
		df = append(df, fmt.Sprintf("Defence is less than 0"))
	case un.transport < 0:
		df = append(df, fmt.Sprintf("Transport is less than 0"))
	case un.jump < 0:
		df = append(df, fmt.Sprintf("Jump is less than 0"))
	}
	switch un.force {
	case TYPE_GROUND:
		if un.transport > 0 {
			df = append(df, fmt.Sprintf("Ground Force cann't have Transport value"))
		}
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Ground Force cann't have Jump value"))
		}
	case TYPE_AIR:
		if un.transport > 0 {
			df = append(df, fmt.Sprintf("Air Force cann't have Transport value"))
		}
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Air Force cann't have Jump value"))
		}
	case TYPE_SPACESHIP:
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Spaceship Force cann't have Jump value"))
		}
	case TYPE_DEPOT:
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Depot Force cann't have Jump value"))
		}
	}
	switch un.tl {
	case 0, 1, 2, 3:
		if un.force == TYPE_AIR {
			df = append(df, fmt.Sprintf("Cann't have Air Force on TL%v", un.tl))
		}
		if un.force == TYPE_SPACESHIP {
			df = append(df, fmt.Sprintf("Cann't have Spaceship Force on TL%v", un.tl))
		}
		if un.force == TYPE_STARSHIP {
			df = append(df, fmt.Sprintf("Cann't have Starship Force on TL%v", un.tl))
		}
		if un.transport > 0 {
			df = append(df, fmt.Sprintf("Cann't have Transport > 0 on TL%v", un.tl))
		}
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Cann't have Jump > 0 on TL%v", un.tl))
		}
	case 4, 5, 6:
		if un.force == TYPE_SPACESHIP {
			df = append(df, fmt.Sprintf("Cann't have Spaceship Force on TL%v", un.tl))
		}
		if un.force == TYPE_STARSHIP {
			df = append(df, fmt.Sprintf("Cann't have Starship Force on TL%v", un.tl))
		}
		if un.transport > 0 {
			df = append(df, fmt.Sprintf("Cann't have Transport > 0 on TL%v", un.tl))
		}
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Cann't have Jump > 0 on TL%v", un.tl))
		}
	case 7, 8:
		if un.jump > 0 {
			df = append(df, fmt.Sprintf("Cann't have Jump > 0 on TL%v", un.tl))
		}
	default:
	}
	return df
}
