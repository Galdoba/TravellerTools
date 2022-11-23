package empire

import (
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/t4/core/task/pocketempire/family"
)

type PocketEmpire struct {
	RulingFamily  *family.Family
	World         map[int]*worldCharacter
	Size          ehex.Ehex
	MilitaryPower ehex.Ehex
	EconomicPower ehex.Ehex
	Prestige      float64
}

type worldCharacter struct {
	name              string
	selfDetermination int //0-10
	basePopularity    int //0-15
	uwp               uwp.UWP
	factions          []int //распределение по фракциям суммарно 100%
	hex               hexagon.Hexagon
}

type individualWorld struct {
}

func New() *PocketEmpire {
	return &PocketEmpire{}
}
