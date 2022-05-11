package astrogation

import (
	"fmt"
)

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

/*
Приоритеты построения маршрута
0. наличие координат в базе (если нет то 100000)
1. Высокая безопасность (10000 - zone)
2. Есть космопорт (1000 - starport)
3. Большое население (100-pop)
4. Есть ГГ (10-ГГ)



*/
