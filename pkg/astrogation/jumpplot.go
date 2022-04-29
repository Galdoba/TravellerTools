package astrogation

import "fmt"

func TradeRouteExist(sourceWRLD, destinationWRLD Coordinator, validJumpPoints []Coordinates) bool {
	source := CoordinatesOf(sourceWRLD)
	destination := CoordinatesOf(destinationWRLD)
	if Distance(source, destination) > 4 {
		fmt.Println("Distance(source, destination) > 4")
		return false
	}
	destFound := false
	for _, c := range validJumpPoints {
		if isSame(c, destination) {
			destFound = true
		}
	}
	if !destFound {
		fmt.Println("!destFound")
		return false
	}

	for _, transitPoint := range validJumpPoints {
		if isSame(source, transitPoint) {
			continue
		}
		transDist := Distance(source, transitPoint)
		endDist := Distance(transitPoint, destination)
		if transDist > 2 {
			continue
		}
		if endDist > 2 {
			continue
		}
		if transDist+endDist <= 4 {
			return true
		}

	}
	return false
}
