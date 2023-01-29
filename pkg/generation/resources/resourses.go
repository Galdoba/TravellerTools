package resources

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/dice"
)

type Resources struct {
}

func testGen() {
	size := 5
	dice := dice.New()
	//rs := 0
	rm := make(map[int]int)
	for try := 0; try < 1000; try++ {
		rf := 0
		for i := 0; i < 12; i++ {
			r := dice.Sroll("3d6-3") - size
			if r <= 0 {
				rf++
			}
		}
		rm[rf]++
	}
	for l := 0; l < 20; l++ {
		if val, ok := rm[l]; ok == true {
			fmt.Println(l, ":", val)
		}
	}

}

//fmt.Println(rm)
// Resources := (Whexes*2) + 2d6/1000 / Gravity
/*
SIZE HZ

		If the planet is size 5 or less and located between the Inner Limit and the Snow Line, add +3 to your roll.
		If the planet is size 6 or more and located between the Inner Limit and the Habitable Zone, take -2 to your
		roll.
		If the planet is size 6 or more and located in the Habitable Zone, take -4 to your roll.
		If the planet is size 5 or less and located beyond the Snow Line, add +9 to your roll.
		If the planet is size 6 or more and located beyond the Snow Line, add +3 to your roll.
		You can determine the density of the planet by using the following table. Roll 2d10 and consult the chart
		according to the core type of the planet. This will give the density of the planet in Earth Densities (or
		“standard”).

		2d6  Core Type
		6-  	Molten
		7-15	Rocky
		16+ 	Icy



		Many scientists will prefer to measure density in grams per cubic centimeter (g/cc).determine this, simply multiply the above result by 5.52.
		If you wish toDensity is measured in terms of Earth Densities or what the CCA calls “standard”. Most Terran, Subterran,
		and Superterran worlds are going to have a molten iron core with layers of rock above that. Some of
		these planets will have less iron and, thus, less density.
		In addition, it is possible
		Mercurians and Dwarf planets between the Inner Limit and the Habitable Zone will have rocky cores with
		less metal. These planets will have less density than those with molten iron cores.
		Mercurians and Dwarf planets outside the Habitable Zone may have rocky cores but may also have icy
		cores. These planets will have significantly less density than those with molten cores or rocky cores.
		Determine the density (in Earth Densities or “standard”) by rolling on the charts below. First determine
		the type of core that the planet has.


*/
