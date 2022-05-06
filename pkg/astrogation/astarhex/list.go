package astarhex

//List - список нодов
type List struct {
	nodes []Node
}

//NewList - создает список
func NewList() *List {
	return &List{}
}

//Add - добавляет 1 или более нод в список
func (l *List) Add(nodes ...Node) {
	l.nodes = append(l.nodes, nodes...)
}

//Add - возвращает весь список нодов
func (l *List) All() []Node {
	return l.nodes
}

//Remove - убирает нод из списка. если его нет, то ничего не делает
func (l *List) Remove(n Node) {

}
