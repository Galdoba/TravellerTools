package structure

type Orbit struct {
	orbitHashSystem float64
	orbitHashStar   float64
	Class           string //Inner/Habitable/Outer
	Eccentricity    float64
	//MinSeparation = AU * ( 1 - Eccentricity)
}
