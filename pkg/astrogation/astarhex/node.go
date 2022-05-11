package astarhex

import (
	"fmt"
	"math"
)

type Node struct {
	f, g, h int
	Q, R, S int
	Weight  int
	parent  *Node
}

func SetNodeHex(row, col int) *Node {
	q := col
	r := row - (col-col&1)/2
	n := Node{}
	n.Q = q
	n.R = r
	n.S = -q - r
	return &n

}

func (n *Node) String() string {
	return fmt.Sprintf("Node [q:%d r:%d s:%d - f:%d g:%d h:%d]", n.Q, n.R, n.S, n.f, n.g, n.h)
}

// Cube Interface///////////

func (n *Node) Qaxis() int {
	return n.Q
}
func (n *Node) Raxis() int {
	return n.R
}
func (n *Node) Saxis() int {
	return n.S
}

// Coordinator Interface

func (n *Node) HexCoords() (int, int) {
	col := n.Q
	row := n.R + (n.Q-(n.Q&1))/2
	return row, col

}

func (n *Node) CoordY() int {
	col := n.Q
	return col

}

func (n *Node) CoordX() int {
	row := n.R + (n.Q-(n.Q&1))/2
	return row

}

////////////////////////////
func Distance(a, b Node) int {
	return int((math.Abs(float64(a.Q-b.Q)) + math.Abs(float64(a.R-b.R)) + math.Abs(float64(a.S-b.S))) / 2)
}
