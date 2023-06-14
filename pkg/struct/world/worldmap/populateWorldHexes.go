package worldmap

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/classifications"
	"github.com/Galdoba/TravellerTools/pkg/dice"
	"github.com/Galdoba/TravellerTools/pkg/ehex"
	"github.com/Galdoba/TravellerTools/pkg/profile"
	"github.com/Galdoba/TravellerTools/pkg/struct/world"
)

/*
GENERATING THE WORLD MAP
As instructed, mark the specific terrain type in directed World Hexes on the World Map. Within the limits of the instructions, terrain may be placed in any available World Hex. If the specific world is too small for the terrain called for, restrict the number placed to what the specific hex will contain.
	//1. Select a blank World Map based on World Size.
	//2. Resources. Determine Resources from the Economic Extension. Subtract system GG and Belts: place the resulting number of Resource Hexes one per Triangle.
3. Mountains. Place 1D Mountains in each Triangle.
4. Chasms. Place World Size x Chasms Sets (1D per Triangle).
5. Precipices. Place World Size x Precipices one per Triangle.
If Di
6. Die-Back. Place 1D Ruins in each Triangle.
If Va
7. Vacuum Plain. Place Craters (1D per Triangle).
If De
8. Desert. Mark all unmarked hexes Desert.
9. Oceans.
Randomly select Hyd x 2 Triangles as Oceans. Consolidate Ocean Triangles that share sides. Enclose Oceans with Shore lines (which may run through any type terrain). Non-Ocean Triangles are Continents (they are not consolidated; treat each Triangle as a separate Continent).
10. Seas. Randomly select Hyd Continents and place a one-hex Ocean (Sea) in each.
Surround each with Shore in all adjacent hexes.
11. Islands. Convert each Mountain Hex in Ocean to Islands.
12. Ice-Caps. If HZ or greater, mark the top and bottom Hyd/2 rows as Ice Cap (if Hyd less than 2, no Ice Caps).
If Ic
13. More Ice Cap. Add 1D rows to each Ice Cap.
If Fr
14. Frozen. Mark Ocean as Ice Field and Land as Frozen Lands (except under Ice Cap).
If Tu
15. Tundra. Mark a line 1D hexes from each Pole. Between each line and its Pole, mark  Ocean as Ice Field and Land as Frozen lands (except under Ice Cap).
If Ag
16. Agricultural. Place 2D Cropland in each Continent.
If Fa
17. Farming. Place 1D Cropland in each Continent.
If Lo
18. Low Population. Place one Town. Skip to 22.
If Ni
19. Non-Industrial. Place one Town. Skip to 22.
20. Cities. Place Cities equal to Pop, one per Continent.
If Atm=0-1, A-C, or E+ = Domed if not NIL.
If Hi
21. High Population. Place total Pop/2 Arcologies.
22. Rural. Mark clear hexes within Pop hexes of City as Rural.
23. Starport. Place the World Starport (or Spaceport).
If Tz
24. Create A Twilight Zone. Select one Pole Triangle and draw a vertical line directly down. Shift 2.5 times World
Size hexes to one side and draw a parallel line: this is the one-World-Hex-wide Twilight Zone.
If Tz
25. Create Two Hemispheres For A Twilight World. Mark one side of the Twilight Zone as Baked Lands and the
other side as Frozen Lands (overlaying existing terrain). Terrain in the Twilight Zone remains as previously
created. Convert Ocean in Baked Lands to Desert. Convert Ocean in Frozen Lands to Ice Field.
If Pe
26. Penal Colony. Mark Pop x Penal (one per Triangle).
27. Wasteland. If TL>5, mark 1D adjacent hexes in one Triangle Wasteland.
28. Exotic. Place one Exotic hex in one Triangle.
29. Noble Lands. Place one Noble Lands estate.
30. All other terrain remains Clear.
*/

const (
	//Terrain
	TERRAIN_Clear         = 11
	TERRAIN_Wetland       = 12
	TERRAIN_Rough         = 13
	TERRAIN_ClearWooded   = 14
	TERRAIN_WetlandWooded = 15
	TERRAIN_RoughWooded   = 16
	TERRAIN_Mountain      = 21
	TERRAIN_Desert        = 22
	TERRAIN_Chasm         = 23
	TERRAIN_Cropland      = 24
	TERRAIN_Rural         = 25
	TERRAIN_Ruins         = 26
	TERRAIN_Ocean         = 31
	TERRAIN_Islands       = 32
	TERRAIN_Shore         = 33
	TERRAIN_River         = 34
	TERRAIN_Lake          = 35
	TERRAIN_IceCap        = 36
	TERRAIN_BakedLands    = 41
	TERRAIN_TwilightZone  = 42
	TERRAIN_FrozenLands   = 43
	TERRAIN_IceField      = 44
	TERRAIN_Precipice     = 45
	TERRAIN_Exotic        = 46
	TERRAIN_City          = 51
	TERRAIN_Domed         = 52
	TERRAIN_Arcology      = 53
	TERRAIN_Suburban      = 54
	TERRAIN_Town          = 55
	TERRAIN_Starport      = 56
	TERRAIN_Highway       = 61
	TERRAIN_Road          = 62
	TERRAIN_Trail         = 63
	TERRAIN_AirCorridor   = 64
	TERRAIN_Grid          = 65
	TERRAIN_HighSpeed     = 66
	TERRAIN_OceanDepths   = 71
	TERRAIN_Abyss         = 72
	TERRAIN_Caverns       = 73
	TERRAIN_Crater        = 74
	TERRAIN_Wasteland     = 75
	TERRAIN_PenalColony   = 76
	TERRAIN_Volcanic      = 81
	TERRAIN_NobleEstate   = 82
	TERRAIN_Reserve       = 83
	TERRAIN_Mines         = 84
	TERRAIN_Resources     = 85
	TERRAIN_ResourcesOil  = 86
	TERRAIN_AirPad        = 91
	TERRAIN_VliteAirport  = 92
	TERRAIN_LiteAirport   = 93
	TERRAIN_Airport       = 94
	TERRAIN_HeavyAirport  = 95
	TERRAIN_VheavyAirport = 96
)

func haveCode(tc []int, code int) bool {
	for _, c := range tc {
		if c == code {
			return true
		}
	}
	return false
}

func (wm *worldmap) PopulateWorldHexesT5(wrld *world.World, dice *dice.Dicepool) error {
	for _, hex := range wm.WorldHex {
		hex.AddTerrain(TERRAIN_Clear)
	}
	stage := 0
	done := false
	err := fmt.Errorf("---")
	tc := classifications.Evaluate(wrld)
	for !done {
		switch stage {
		default:
			fmt.Println("Stage:", stage)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}
			stage++
		case 3:
			fmt.Println("placeMountains")
			err = wm.placeMountains(wrld, dice)
			stage++
		case 4:
			fmt.Println("placeChasms")
			err = wm.placeChasms(wrld, dice)
			stage++
		case 5:
			fmt.Println("placePrecipation")
			err = wm.placePrecipation(wrld, dice)
			stage++
		case 6:
			if haveCode(tc, classifications.Di) {
				wm.placeRuins(dice)
			}
			stage++
		case 7:
			if haveCode(tc, classifications.Va) {
				wm.placeCraters(dice)
			}
			stage++
		case 8:
			if haveCode(tc, classifications.De) {
				wm.markDesert()
			}
			stage++
		case 9:
			fmt.Println("placeOceans")
			err = wm.placeOceans(wrld, dice)
			stage++
		case 12:
			wm.placeIceCaps1(wrld)
			stage++
		case 13:
			if haveCode(tc, classifications.Ic) {
				wm.placeIceCaps2(wrld, dice)
			}
			stage++
		case 14:
			if haveCode(tc, classifications.Fr) {
				wm.placeFrozen()
			}
			stage++
		case 15:
			if haveCode(tc, classifications.Tu) {
				wm.placeTundra(dice)
			}
			stage++

		case 30:
			done = true
		}
	}
	//3 mountains

	for i := 0; i < len(wm.WorldHex); i++ {
		fmt.Println(i, wm.WorldHex[i].coords, wm.WorldHex[i].overallTerrain)
	}
	return nil
}

func d6PerTriangle(data ehex.Ehex, dice *dice.Dicepool) string {
	s := 0
	for i := 0; i < 3; i++ {
		s += dice.Sroll(fmt.Sprintf("%vd6", data.Value()))
	}
	return fmt.Sprintf("%vd6", s)
}

func (wh *WorldHex) AddTerrain(terrain int) error {
	for _, ter := range wh.overallTerrain {
		if ter == terrain {
			return fmt.Errorf("already have")
		}
	}
	wh.overallTerrain = append(wh.overallTerrain, terrain)
	return nil
}

func (wh *WorldHex) ReplaceTerrain(old, new int) error {
	for i, ter := range wh.overallTerrain {
		if ter == old {
			wh.overallTerrain[i] = new
			return nil
		}
		if ter == new {
			return fmt.Errorf("already have")
		}
	}
	return fmt.Errorf("cannot replace")
}

func (wh *WorldHex) TerrainIs(ter int) bool {
	for _, has := range wh.overallTerrain {
		if ter == has {
			return true
		}
	}
	return false
}

func logErr(e error) {
	if e != nil {
		fmt.Print(e.Error())
	}
}

func (wm *worldmap) countTerrain(ter int) int {
	found := 0
	for _, hex := range wm.WorldHex {
		for _, hexTer := range hex.overallTerrain {
			if hexTer == ter {
				found++
				break
			}
		}
	}
	return found
}

func (wm *worldmap) placeMountains(wrld *world.World, dice *dice.Dicepool) error {
	mnt := dice.Sroll(d6PerTriangle(wrld.Profile().Data(profile.KEY_SIZE), dice))
	hexNum := len(wm.WorldHex)
	populated := 0
	for i := 0; i < mnt; i++ {
		randomID := dice.Sroll(fmt.Sprintf("1d%v-1", hexNum))
		if wm.WorldHex[randomID].ReplaceTerrain(TERRAIN_Clear, TERRAIN_Mountain) == nil {
			populated++
		}
	}
	fmt.Println(populated, "mountains added")
	return nil
}

func (wm *worldmap) placeChasms(wrld *world.World, dice *dice.Dicepool) error {
	csm := dice.Sroll(d6PerTriangle(wrld.Profile().Data(profile.KEY_SIZE), dice))
	hexNum := len(wm.WorldHex)
	populated := 0
	for i := 0; i < csm; i++ {
		randomID := dice.Sroll(fmt.Sprintf("1d%v-1", hexNum))
		if wm.WorldHex[randomID].ReplaceTerrain(TERRAIN_Clear, TERRAIN_Chasm) == nil {
			populated++
		}
		if wm.WorldHex[randomID].ReplaceTerrain(TERRAIN_Mountain, TERRAIN_Clear) == nil {
			populated++
		}
	}
	fmt.Println(populated, "chasms added")
	return nil
}

func (wm *worldmap) placePrecipation(wrld *world.World, dice *dice.Dicepool) error {
	psp := dice.Sroll(d6PerTriangle(wrld.Profile().Data(profile.KEY_SIZE), dice))
	hexNum := len(wm.WorldHex)
	populated := 0
	for i := 0; i < psp; i++ {
		randomID := dice.Sroll(fmt.Sprintf("1d%v-1", hexNum))
		if wm.WorldHex[randomID].AddTerrain(TERRAIN_Precipice) == nil {
			populated++
		}
	}
	fmt.Println(populated, "precipations added")
	return nil
}

func (wm *worldmap) placeRuins(dice *dice.Dicepool) error {
	rui := dice.Sroll("20d6")
	hexNum := len(wm.WorldHex)
	populated := 0
	for i := 0; i < rui; i++ {
		randomID := dice.Sroll(fmt.Sprintf("1d%v-1", hexNum))
		if wm.WorldHex[randomID].AddTerrain(TERRAIN_Ruins) == nil {
			populated++
		}
	}
	fmt.Println(populated, "ruins added")
	return nil
}

func (wm *worldmap) placeCraters(dice *dice.Dicepool) error {
	rui := dice.Sroll("20d6")
	hexNum := len(wm.WorldHex)
	populated := 0
	for i := 0; i < rui; i++ {
		randomID := dice.Sroll(fmt.Sprintf("1d%v-1", hexNum))
		if wm.WorldHex[randomID].AddTerrain(TERRAIN_Crater) == nil {
			populated++
		}
	}
	fmt.Println(populated, "Craters added")
	return nil
}

func (wm *worldmap) markDesert() error {
	populated := 0
	for _, hex := range wm.WorldHex {
		hex.ReplaceTerrain(TERRAIN_Clear, TERRAIN_Desert)
	}
	fmt.Println(populated, "deserts added")
	return nil
}

func (wm *worldmap) placeOceans(wrld *world.World, dice *dice.Dicepool) error {
	hydr := wrld.Profile().Data(profile.KEY_HYDR)
	hexNum := len(wm.WorldHex)
	ocean := (hexNum / 10) * hydr.Value()
	populated := 0
	for i := 0; i < ocean; i++ {
		placed := false
		for !placed {
			randomID := dice.Sroll(fmt.Sprintf("1d%v-1", hexNum))
			if wm.WorldHex[randomID].ReplaceTerrain(TERRAIN_Clear, TERRAIN_Ocean) == nil {
				placed = true
			}
			if wm.WorldHex[randomID].ReplaceTerrain(TERRAIN_Mountain, TERRAIN_Islands) == nil {
				placed = true
			}

		}
		populated++
	}
	fmt.Println(populated, "precipations added")
	return nil
}

func (wm *worldmap) placeIceCaps1(wrld *world.World) error {
	hydr := wrld.Profile().Data(profile.KEY_HYDR)
	lastRow := wm.WorldHex[len(wm.WorldHex)-1].coords.y
	icWidth := hydr.Value() / 2
	if icWidth == 0 {
		return nil
	}
	populated := 0
	for _, hex := range wm.WorldHex {
		if hex.coords.y < icWidth || hex.coords.y > lastRow-icWidth {
			hex.AddTerrain(TERRAIN_IceCap)
			populated++
		}

	}
	fmt.Println(populated, "ice caps added")
	return nil
}

func (wm *worldmap) placeIceCaps2(wrld *world.World, dice *dice.Dicepool) error {
	hydr := wrld.Profile().Data(profile.KEY_HYDR)
	lastRow := wm.WorldHex[len(wm.WorldHex)-1].coords.y
	icWidth := hydr.Value() / 2
	if icWidth == 0 {
		return nil
	}
	icWidth += dice.Sroll("1d6")
	populated := 0
	for _, hex := range wm.WorldHex {
		if hex.coords.y < icWidth || hex.coords.y > lastRow-icWidth {
			hex.AddTerrain(TERRAIN_IceCap)
			populated++
		}

	}
	fmt.Println(populated, "ice caps added")
	return nil
}

func (wm *worldmap) placeFrozen() error {
	for _, hex := range wm.WorldHex {
		if hex.TerrainIs(TERRAIN_IceCap) {
			continue
		}
		switch {
		case hex.TerrainIs(TERRAIN_Ocean):
			hex.AddTerrain(TERRAIN_IceField)
		case hex.TerrainIs(TERRAIN_Clear), hex.TerrainIs(TERRAIN_Desert), hex.TerrainIs(TERRAIN_Mountain):
			hex.AddTerrain(TERRAIN_FrozenLands)
		}
	}
	return nil
}

func (wm *worldmap) placeTundra(dice *dice.Dicepool) error {
	lastRow := wm.WorldHex[len(wm.WorldHex)-1].coords.y
	icWidth := dice.Sroll("1d6")
	for _, hex := range wm.WorldHex {
		if hex.TerrainIs(TERRAIN_IceCap) {
			continue
		}
		if hex.coords.y < icWidth || hex.coords.y > lastRow-icWidth {
			switch {
			case hex.TerrainIs(TERRAIN_Ocean):
				hex.AddTerrain(TERRAIN_IceField)
			case hex.TerrainIs(TERRAIN_Clear), hex.TerrainIs(TERRAIN_Desert), hex.TerrainIs(TERRAIN_Mountain):
				hex.AddTerrain(TERRAIN_FrozenLands)
			}
		}
	}
	return nil
}
