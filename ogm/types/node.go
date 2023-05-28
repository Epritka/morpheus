package types

type Node struct {
	*BaseEntity
}

func NewNode() *Node {
	return &Node{BaseEntity: new()}
}
