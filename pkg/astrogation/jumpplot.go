package astrogation

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/profile/uwp"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

const (
	pointStatus_UNDEFINED = iota
	pointStatus_EMPTY
	pointStatus_STAR
	pointStatus_GG
	pointStatus_GREENZONE
	pointStatus_AMBERZONE
	pointStatus_REDZONE
	pointStatus_STARPORT_A
	pointStatus_STARPORT_B
	pointStatus_STARPORT_C
	pointStatus_STARPORT_D
	pointStatus_STARPORT_E
	pointStatus_STARPORT_X
	pointStatus_WRONGDATA
)

type AstrogationPlot struct {
	start        Coordinates
	end          Coordinates
	restrictions []int
	region       map[Coordinates][]int
	courses      map[int][]jumpCourse
}

type jumpCourse struct {
	jumpDriveRequired int
	jumps             int
	courseAlowences   []int
	plotPath          []Coordinates
	hexString         string
}

func PlotCourse(start, end Coordinates, restrictions ...int) (*AstrogationPlot, error) {
	ap := AstrogationPlot{}
	ap.start = start
	ap.end = end
	ap.region = make(map[Coordinates][]int)
	ap.courses = make(map[int][]jumpCourse)
	for _, r := range restrictions {
		ap.restrictions = append(ap.restrictions, r)
	}
	/*
		Plan:
		createField
		mapField
		astarTo end

	*/
	dist := Distance(start, end)
	cp := []Coordinates{}
	for _, coord := range JumpMap(&start, dist) {
		cp = addUniqeToField(cp, coord)
	}
	for _, coord := range JumpMap(&end, dist) {
		cp = addUniqeToField(cp, coord)
	}
	for _, coords := range cp {
		worlddata, err := survey.SearchByCoordinates(coords.HexValues())
		if err != nil {
			return nil, fmt.Errorf("search by coordinates: %v", err.Error())
		}
		wp, err := uwp.FromString(worlddata.MW_UWP())
		if err != nil {
			return nil, fmt.Errorf("parsing UWP: %v", err.Error())
		}
		switch wp.Starport() {
		case "A":
			ap.restrictions = append(ap.restrictions, pointStatus_STARPORT_A)
		}
	}
	return nil, fmt.Errorf("not implemented")
}

func addUniqeToField(coordPool []Coordinates, new Coordinates) []Coordinates {
	for _, cc := range coordPool {
		if isSame(cc, new) {
			return coordPool
		}
	}
	return append(coordPool, new)
}
