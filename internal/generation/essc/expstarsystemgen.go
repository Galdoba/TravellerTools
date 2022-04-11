package essc

import "fmt"

const (
	Solo      = "P"
	Binary1   = "PC"
	Binary2   = "PD"
	Trinary1  = "PCD"
	Trinary2  = "PDd"
	Multiple1 = "PCDc"
	Multiple2 = "PCDd"
	Multiple3 = "PDcd"
	Multiple4 = "PCDcd"
	Special   = "SS"
)

type StarSystem struct {
	model string
}

func New(seed string) (*StarSystem, error) {
	return &StarSystem{}, fmt.Errorf("Not implemented")
}

/*
starsys
main
close C
distant C

distant C
close C
distant C

Star {
	Spectral
	Size
	Decimal
	Position
	Close
	Distant
}

*/
