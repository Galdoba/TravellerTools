package planets

import (
	"github.com/Galdoba/TravellerTools/pkg/ehex"
)

type planet struct {
	planetType   string
	physicalData []ehex.Ehex
	resources    int
	parentStar   string
	star         int
	orbit        int
	satelite     int
}
