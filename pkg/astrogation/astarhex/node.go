package astarhex

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/astrogation/hexagon"
)

type Node struct {
	f, g, h int
	Hex     hexagon.Hexagon
	Weight  int
	parent  *Node
}

func SetNodeHex(hex hexagon.Hexagon) *Node {
	n := Node{}
	n.Hex = hex
	return &n

}

func (n *Node) String() string {
	return fmt.Sprintf("Node [q:%d r:%d s:%d - f:%d g:%d h:%d]", n.Hex.CoordQ(), n.Hex.CoordR(), n.Hex.CoordS(), n.f, n.g, n.h)
}
