package astarhex

import "fmt"

type Node struct {
	f, g, h int
	Q, R, S int
	Weight  int
	parent  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node [q:%d r:%d s:%d - f:%d g:%d h:%d]", n.Q, n.R, n.S, n.f, n.g, n.h)
}
