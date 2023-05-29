package sophont

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/generation/star"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

type sophont struct {
	HomeStar star.StarBody
}

func NewSpecies(homeworld *world.World, dice *dice.Dicepool) (*sophont, error) {
	sph := sophont{}
	sph.HomeStar = homeworld.HomeStar
	fmt.Println(homeworld.HomeStar)
	fmt.Println(sph.HomeStar.Class())
	fmt.Println(homeworld.Profile().Data(profile.KEY_PLANETARY_ORBIT))
	fmt.Println(homeworld.Profile().Data(profile.KEY_HABITABLE_ZONE_VAR))

	return &sph, nil
}
