package main

func main() {

}

//это все интерфейсы
type Galaxy struct {
	nomena  string            //глобальный сид генерации
	Sectors map[string]Sector //Карта Секторов
}

type Sector struct {
	Name              string
	Position          int         //Координаты по улитке - могут использоваться как ключ для галактики
	PopulationDencity int         //Плотность заселения сектора (шанс спуна звезды)
	Neihbours         map[int]int //указатель на соседние сектора key=Direction val=Position
	Hexes             map[int]Hex //Карта Хексов
}

type Hex struct {
	Nomena        string                   //Координаты хекса и сектора причесанные в уникальный ключ для сида
	Position      int                      //Координаты по улитке - могут использоваться как ключ для сектора
	SystemObjects map[float64]SystemObject //карта объектов внутри Хекса key=Орбита(кольцо) val=Объект
}

type SystemObject struct {
	ObjectType string               // Star/Planet/Belt/Rogue/Lpoint
	Satelites  map[float64]Satelite //Карта локального региона key=localOrbit val=Satelite
}

type Satelite struct {
	SateliteType string // Planetoid/Ring
}

///////////////////////////////////////

/*





SpaceObject = Planet
*/
