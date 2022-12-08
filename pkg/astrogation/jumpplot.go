package astrogation

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/astarhex"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

func TradeRouteExist(source, destination hexagon.Hexagon, validJumpPoints []hexagon.Hexagon) bool {
	if hexagon.Match(source, destination) {
		return false
	}
	if hexagon.Distance(source, destination) > 4 {
		//fmt.Println("Distance(source, destination) > 4")
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

type Plot struct {
	nodes []astarhex.Node
	Path  string
	Cost  int
}

func PlotCource(start, end hexagon.Hex, MaxJumpDistance int, MaxConseuqnceJumps int) (Plot, error) {
	jc := Plot{}
	ast, err := astarhex.New(astarhex.Config{MaxJumpDistance: MaxJumpDistance, MaxConseuqnceJumps: MaxConseuqnceJumps})
	if err != nil {
		return jc, err
	}
	path, pErr := ast.FindPathHex(start, end)
	if pErr != nil {
		return jc, pErr
	}
	path = append(path, *astarhex.SetNodeHex(hexagon.FromHex(start)))

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	jc.nodes = path
	for _, n := range path {
		jc.Cost += n.Cost()
	}
	pathStr := ""
	for _, hx := range path {
		wr, errW := survey.SearchByCoordinates(hx.Hex.HexValues())
		if errW != nil {
			return jc, errW
		}

		pathStr += wr.MW_Name() + " ---> "
	}
	jc.Path = strings.TrimSuffix(pathStr, " ---> ")

	return jc, nil
}

/*
Приоритеты построения маршрута
0. наличие координат в базе (если нет то 100000)
1. Высокая безопасность (10000 - zone)
2. Есть космопорт (1000 - starport)
3. Большое население (100-pop)
4. Есть ГГ (10-ГГ)



*/
