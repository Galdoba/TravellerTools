package sophont

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

const (
	E_UNKNOWN = iota
	E_Mountain
	E_Desert
	E_Exotic
	E_Rough_Wood
	E_Rough
	E_Clear
	E_Forest
	E_Wetlands
	E_Wetland_Woods
	E_Ocean
	E_Ocean_Depths
	E_Baked_Lands
	E_Twilight_Zone
	E_Frozen_Lands
	L_Walker
	L_Amphibian
	L_Triphibian
	L_Aquatic
	L_Diver
	L_Flyer
	L_Flyphib
	L_Swimmer
	L_Static
	L_Drifter
)

type terrain struct {
	name  string
	descr string
}

func (sph *sophont) determineEnviroment(homeworld world.World, dice *dice.Dicepool) error {
	value := make(map[string]ehex.Ehex)
	for _, key := range []string{profile.KEY_SIZE, profile.KEY_ATMO, profile.KEY_HYDR} {
		if val := homeworld.Profile().Data(key); val == nil {
			return fmt.Errorf("homeworld has no value on key '%v'", key)
		} else {
			value[key] = val
		}
	}
	tz := false
	for _, tc := range homeworld.ListTC() {
		if tc == classifications.Tz || tc == classifications.Lk {
			tz = true
			break
		}
	}
	terrainList := []int{E_Mountain, E_Desert, E_Exotic, E_Rough_Wood}
	switch tz {
	case true:
		terrainList = append(terrainList, E_Baked_Lands, E_Twilight_Zone, E_Frozen_Lands)
	case false:
		terrainList = append(terrainList, E_Rough, E_Clear, E_Forest)
	}
	terrainList = append(terrainList, E_Wetlands, E_Wetland_Woods, E_Ocean, E_Ocean_Depths)
	dm := 0
	if value[profile.KEY_SIZE].Value() <= 5 {
		dm = dm - 1
	}
	if value[profile.KEY_ATMO].Value() >= 8 {
		dm = dm + 2
	}
	if value[profile.KEY_HYDR].Value() >= 6 {
		dm = dm + 1
	}
	if value[profile.KEY_HYDR].Value() >= 9 {
		dm = dm + 1
	}
	r1 := setBounds(dice.Flux()+dm, -5, 5)
	terrain := newTerrain(r1 + 5)
	locomotion := []int{}
	switch r1 {
	case -5, -4:
		locomotion = append(locomotion, L_Walker, L_Walker, L_Walker, L_Walker, L_Walker, L_Flyer)
	case -3:
		locomotion = append(locomotion, L_Amphibian, L_Walker, L_Walker, L_Walker, L_Flyphib, L_Flyer)
	case -2, -1:
		locomotion = append(locomotion, L_Amphibian, L_Walker, L_Walker, L_Walker, L_Walker, L_Flyer)
	case 0, 1:
		locomotion = append(locomotion, L_Walker, L_Walker, L_Walker, L_Walker, L_Walker, L_Walker)
	case 2:
		locomotion = append(locomotion, L_Amphibian, L_Aquatic, L_Walker, L_Walker, L_Triphibian, L_Flyer)
	case 3:
		locomotion = append(locomotion, L_Amphibian, L_Walker, L_Walker, L_Walker, L_Triphibian, L_Flyphib)
	case 4:
		locomotion = append(locomotion, L_Flyphib, L_Swimmer, L_Swimmer, L_Swimmer, L_Aquatic, L_Diver)
	case 5:
		locomotion = append(locomotion, L_Aquatic, L_Diver, L_Diver, L_Diver, L_Diver, L_Diver)

	}
	r2 := setBounds(dice.Sroll("1d6")+dm, 1, 6)
	movement := locomotion[r2-1]
	r3 := setBounds(dice.Flux()+dm, -6, 6)
	fmt.Println(terrain, movement, r3)
	return nil
}

func setBounds(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}

func newTerrain(i int) terrain {
	tr := terrain{}
	switch i {
	case E_Mountain:
		tr.name = "Mountain"
		tr.descr = "Steep dominating region"
	case E_Desert:
		tr.name = "Desert"
		tr.descr = "Dry region with sparse vegetation"
	case E_Exotic:
		tr.name = "Exotic"
		tr.descr = "Strange abnormal region"
	case E_Rough_Wood:
		tr.name = "Rough Wood"
		tr.descr = "High dencity vegetation region"
	case E_Rough:
		tr.name = "Rough"
		tr.descr = "Uneven or broken surface region"
	case E_Clear:
		tr.name = "Clear"
		tr.descr = "Flat extended region"
	case E_Forest:
		tr.name = "Forest"
		tr.descr = "Flat with high vegetation"
	case E_Wetlands:
		tr.name = "Wetlands"
		tr.descr = "Water-dominated marsh"
	case E_Wetland_Woods:
		tr.name = "Wetland Woods"
		tr.descr = "Water-dominated swamp"
	case E_Ocean:
		tr.name = "Ocean"
		tr.descr = "Interface of sea and atmosphere"
	case E_Ocean_Depths:
		tr.name = "Ocean Depths"
		tr.descr = "Subsurface ocean regions"
	case E_Baked_Lands:
		tr.name = "Baked Lands"
		tr.descr = "Hot region"
	case E_Twilight_Zone:
		tr.name = "Twilight Zone"
		tr.descr = "Temperate region"
	case E_Frozen_Lands:
		tr.name = "Frozen Lands"
		tr.descr = "Cold region"
	}
	return tr
}
