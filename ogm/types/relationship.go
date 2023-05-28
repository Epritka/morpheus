package types

type Relationship struct {
	*BaseEntity
}

func NewRelationship() *Relationship {
	return &Relationship{BaseEntity: new()}
}
