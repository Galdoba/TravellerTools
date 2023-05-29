package sophont

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

func TestSophont(t *testing.T) {
	wrld, err := world.NewWorld(
		world.KnownData(profile.KEY_MAINWORLD, world.FLAG_TRUE),
		world.KnownData(profile.KEY_SIZE, "7"),
		world.KnownData(profile.KEY_ATMO, "8"),
		world.KnownData(profile.KEY_HYDR, "7"),
		world.KnownData(world.Catalog, "Sol Sector 3955"),
	)
	dice := dice.New().SetSeed(7)
	wrld.GenerateFull(dice)
	fmt.Println(err)
	fmt.Println(wrld)
	sph, err := NewSpecies(wrld, dice)
	fmt.Println(err)
	fmt.Println(sph)
}
