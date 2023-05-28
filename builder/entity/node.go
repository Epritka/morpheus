package entity

import "fmt"

type Node struct {
	*Base
}

func (n *Node) String() string {
	return fmt.Sprintf("(%s)", n.Base.String())
}
