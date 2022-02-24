package weapon

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
	name         string
	tl           int
	rnge         float64
	damageDice   int
	damageMod    int
	weight       float64
	cost         float64
	magazine     int
	magazineCost float64
	quickdraw    int
	traits       []string
}

type Design struct {
	calibr callibre
}
