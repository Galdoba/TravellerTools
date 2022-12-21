package worlddata

import (
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon/location"
	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
)

type World struct {
	UWP    uwp.UWP
	TC     []classifications.Classification
	Coords location.Location
	//
}
