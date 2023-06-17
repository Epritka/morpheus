package entity

import "fmt"

type Node struct {
	*Base
}

func NewNode(alias string) *Node {
	return (&Node{&Base{}}).SetAlias(alias)
}

func (n *Node) SetAlias(alias string) *Node {
	n.setAlias(alias)
	return n
}

func (n *Node) SetLables(labels ...string) *Node {
	n.setLables(labels)
	return n
}

func (n *Node) SetProperties(properties map[string]any) *Node {
	n.setProperties(properties)
	return n
}

func (n *Node) String() string {
	return fmt.Sprintf("(%s)", n.Base.String())
}
