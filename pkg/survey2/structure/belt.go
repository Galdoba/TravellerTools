package structure

type Belt struct {
	Profile string
	Nomena  string
	//Parent  string //не уверен - должен ли пояс знать где он находится
	Span     float64 //adjacent orbit slots inner and outer
	Mtype    int     //Metalic Composition %
	Stype    int     //Composition Stony %
	Ctype    int     //Composition Carbon/Ice %
	Otype    int     //Composition Other %
	Bulk     int     //Factor of the volume of bodies
	Resource int     //Resource Rating (Round Down)
}

func NewBelt(systemData string, orbitData string) *Belt {
	b := Belt{}
	//seed := systemData + orbitData
	//dice := dice.New().SetSeed(seed)
	return &b
}
