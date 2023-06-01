package aDetails

type AtmoDetails struct {
	composition               string
	pressure                  float64
	surfTemp                  int
	luminocity                float64
	orbitFactor               float64
	energyAbsorption          float64
	greenhouseEffect          float64
	baseTemp                  float64
	orbitalEccentricityEffect float64
	latitudeTempEffect        []int
	AxialTiltIncrease         float64
	AxialTiltDecrease         float64
	AxiallatitudeEffects      []float64
	dayLen                    int
}

func New() *AtmoDetails {
	ad := AtmoDetails{}
	return &ad
}
