package astrogation

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/astarhex"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TradeRouteExist(source, destination hexagon.Hexagon, validJumpPoints []hexagon.Hexagon) bool {
	if hexagon.Distance(source, destination) > 4 {
		fmt.Println("Distance(source, destination) > 4")
		return false
	}
	destFound := false
	for _, c := range validJumpPoints {
		if hexagon.Match(c, destination) {
			destFound = true
		}
	}
	if !destFound {
		fmt.Println("!destFound")
		return false
	}

	for _, transitPoint := range validJumpPoints {
		if hexagon.Match(source, transitPoint) {
			continue
		}
		transDist := hexagon.Distance(source, transitPoint)
		endDist := hexagon.Distance(transitPoint, destination)
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

func PlotCource(start, end hexagon.Hex, MaxJumpDistance int, MaxConseuqnceJumps int) string {
	ast, err := astarhex.New(astarhex.Config{MaxJumpDistance, MaxConseuqnceJumps})
	if err != nil {
		return err.Error()
	}
	path, pErr := ast.FindPathHex(start, end)
	if pErr != nil {
		return pErr.Error()
	}
	path = append(path, *astarhex.SetNodeHex(hexagon.FromHex(start)))
	pathStr := ""
	for _, hx := range path {
		wr, errW := survey.SearchByCoordinates(hx.Hex.HexValues())
		if errW != nil {
			return errW.Error()
		}
		pathStr += wr.MW_Name() + " ---> "
	}
	pathStr = strings.TrimSuffix(pathStr, " ---> ")
	return pathStr
}

/*
Приоритеты построения маршрута
0. наличие координат в базе (если нет то 100000)
1. Высокая безопасность (10000 - zone)
2. Есть космопорт (1000 - starport)
3. Большое население (100-pop)
4. Есть ГГ (10-ГГ)



*/
