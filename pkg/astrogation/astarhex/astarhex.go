package astarhex

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Galdoba/TravellerTools/internal/ehex"
	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
	"github.com/Galdoba/TravellerTools/pkg/survey"
)

// Config holds important settings
// to perform the calculation
//
// GridWidth and GridHeight are required and represents
// the size of the grid
//
// InvalidNodes can be used to add not accessible nodes like obstacles etc.
// WeightedNodes can be used to add nodes to be avoided like mud or mountains
type Config struct {
	//GridWidth, GridHeight int
	//InvalidNodes          []Node
	//WeightedNodes         []Node
	MaxJumpDistance    int
	MaxConseuqnceJumps int
}

type astar struct {
	config               Config
	openList, closedList List
	startNode, endNode   Node
}

// New creates a new astar instance
func New(config Config) (*astar, error) {
	if config.MaxJumpDistance < 1 {
		return nil, errors.New("MaxJumpDistance must > 0")
	}

	a := astar{config: config}

	return &a, nil
}

// init initialised needed properties
// internal function
// func (a *astar) init() *astar {
// 	// add invalidNodes directly to the closedList
// 	a.closedList.Add(a.config.InvalidNodes...)
// 	return a
// }

// H caluclates the absolute distance between
// nodeA and nodeB calculates by the manhattan distance
func (a *astar) H(nodeA Node, nodeB Node) int {
	//absX := math.Abs(float64(nodeA.X - nodeB.X))
	//absY := math.Abs(float64(nodeA.Y - nodeB.Y))
	return hexagon.Distance(nodeA.Hex, nodeB.Hex) * 100000
}

// GetNeighborNodes calculates the next neighbors of the given node
// if a neighbor node is not accessible the node will be ignored
func (a *astar) GetNeighborNodes(node Node, radius int) []Node {
	var neighborNodes []Node
	spiral, _ := hexagon.Spiral(&node.Hex, radius)
	for _, neib := range spiral {
		neibNode := SetNodeHex(neib)
		neibNode.parent = &node
		if a.isAccessible(*neibNode) {
			neighborNodes = append(neighborNodes, *neibNode)
		}
	}
	return neighborNodes
}

// isAccessible checks if the node is reachable in the grid
// and is not in the invalidNodes slice
func (a *astar) isAccessible(node Node) bool {

	// if node is out of bound
	// if node.X < 0 || node.Y < 0 || node.X > a.config.GridWidth-1 || node.Y > a.config.GridHeight-1 {
	// return false
	// }

	// check if the node is in the closedList
	// the predefined invalidNodes are also in this list
	if a.closedList.Contains(node) {
		return false
	}

	return true
}

// IsEndNode checks if the given node has
// equal node coordinates with the end node
func (a *astar) IsEndNode(checkNode, endNode Node) bool {
	return hexagon.Match(checkNode.Hex, endNode.Hex)
}

func (a *astar) StartNode() Node {
	return a.startNode
}

// FindPath starts the a* algorithm for the given start and end node
// The return value will be the fastest way represented as a nodes slice
//
// If no path was found it returns nil and an error
func (a *astar) FindPath(startNode, endNode Node) ([]Node, error) {
	a.startNode = startNode
	a.endNode = endNode
	//	fmt.Println("Distance:", hexagon.Distance(startNode.Hex, endNode.Hex))
	defer func() {
		a.openList.Clear()
		a.closedList.Clear()
	}()
	a.openList.Add(startNode)

	for !a.openList.IsEmpty() {
		currentNode, err := a.openList.GetMinFNode()
		if err != nil {
			return nil, fmt.Errorf("cannot get minF node %v", err)
		}
		if len(a.getNodePath(currentNode)) >= 5 {
			return nil, fmt.Errorf("Path have 5+ nodes start:[%v] end:[%v]", startNode, endNode)
		}

		a.openList.Remove(currentNode)
		a.closedList.Add(currentNode)

		// we found the path
		if a.IsEndNode(currentNode, endNode) {
			return a.getNodePath(currentNode), nil
		}
		neighbors := a.GetNeighborNodes(currentNode, a.config.MaxJumpDistance)
		for _, neighbor := range neighbors {
			if a.closedList.Contains(neighbor) {
				continue
			}
			a.calculateNode(&neighbor)
			if !a.openList.Contains(neighbor) {
				a.openList.Add(neighbor)
			}
		}

	}

	return nil, errors.New("No path found")
}

// calculateNode calculates the F, G and H value for the given node
func (a *astar) calculateNode(node *Node) {

	node.g += EvaluateMovementWeight(node)

	// check for special node weighting
	// for _, wNode := range a.config.WeightedNodes {
	// 	if node.Q == wNode.Q && node.R == wNode.R && node.S == wNode.S {
	// 		node.g = node.g + 1 // wNode.Weighting
	// 	}
	// }

	node.h = a.H(*node, a.endNode)
	node.f = node.g + node.h
}

// getNodePath returns the chain of parent nodes
// the given node will be still included in the nodes slice
func (a *astar) getNodePath(currentNode Node) []Node {
	var nodePath []Node
	nodePath = append(nodePath, currentNode)
	for {
		if currentNode.parent == nil {
			break
		}

		parentNode := *currentNode.parent

		// if the end of node chain
		if parentNode.parent == nil {
			break
		}

		nodePath = append(nodePath, parentNode)
		currentNode = parentNode
	}
	return nodePath
}

///////////////////////////

func EvaluateMovementWeight(n *Node) int {
	//coord := NewCoordinates(crd.CoordX(), crd.CoordY())
	wrld, err := survey.SearchByCoordinates(n.Hex.CoordX(), n.Hex.CoordY())
	if err != nil {
		return 10000000
	}
	wrldWeight := 100000
	switch wrld.TravelZone() {
	case "R":
		wrldWeight -= 10000
	case "A":
		wrldWeight -= 45000
	case "":
		wrldWeight -= 90000
	}
	for i, val := range strings.Split(wrld.MW_UWP(), "") {
		if i == 0 { //космопорт
			switch val {
			case "A":
				wrldWeight -= 6000
			case "B":
				wrldWeight -= 5000
			case "C":
				wrldWeight -= 4000
			case "D":
				wrldWeight -= 3000
			case "E":
				wrldWeight -= 2000
			case "X":
				wrldWeight -= 0
			}
		}
		if i == 4 { //население
			popFactor := ehex.New().Set(val).Value()
			wrldWeight -= (popFactor * 100)
		}
	}
	for i, val := range strings.Split(wrld.PBG(), "") {
		if i == 0 || i == 1 {
			continue
		}
		ggFactor := ehex.New().Set(val).Value()
		wrldWeight -= (ggFactor * 10)
	}
	return wrldWeight
}

func (a *astar) FindPathHex(startHex, endHex hexagon.Hex) ([]Node, error) {
	strtHX := hexagon.FromHex(startHex)
	startNode := *SetNodeHex(strtHX)
	a.startNode = startNode
	endHX := hexagon.FromHex(endHex)
	endNode := *SetNodeHex(endHX)
	a.endNode = endNode
	return a.FindPath(startNode, endNode)
	// defer func() {
	// 	a.openList.Clear()
	// 	a.closedList.Clear()
	// }()
	// a.openList.Add(startNode)
	// for !a.openList.IsEmpty() {
	// 	currentNode, err := a.openList.GetMinFNode()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("cannot get minF node %v", err)
	// 	}
	// 	a.openList.Remove(currentNode)
	// 	a.closedList.Add(currentNode)
	// 	// we found the path
	// 	if a.IsEndNode(currentNode, endNode) {
	// 		return a.getNodePath(currentNode), nil
	// 	}
	// 	neighbors := a.GetNeighborNodes(currentNode, a.config.MaxJumpDistance)
	// 	for _, neighbor := range neighbors {
	// 		if a.closedList.Contains(neighbor) {
	// 			continue
	// 		}
	// 		a.calculateNode(&neighbor)
	// 		if !a.openList.Contains(neighbor) {
	// 			a.openList.Add(neighbor)
	// 		}
	// 	}
	// }

	// return nil, errors.New("No path found")
}
