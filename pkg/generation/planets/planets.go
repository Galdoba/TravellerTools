package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/generation/orbit"
)

type planet struct {
	planetType   string
	physicalData []ehex.Ehex
	resources    int
	orbit        orbit.Orbiter
}
