package entity

import "fmt"

type AnyRelationship struct{}

type Relationship struct {
	*Base
}

func NewRelationship(alias string) *Relationship {
	return (&Relationship{&Base{}}).SetAlias(alias)
}

func (r *Relationship) SetAlias(alias string) *Relationship {
	r.setAlias(alias)
	return r
}

func (r *Relationship) SetLables(labels ...string) *Relationship {
	r.setLables(labels)
	return r
}

func (r *Relationship) SetProperties(properties map[string]any) *Relationship {
	r.setProperties(properties)
	return r
}

func (ars *AnyRelationship) String() string {
	return fmt.Sprintf("[*]")
}

func (rs *Relationship) String() string {
	return fmt.Sprintf("[%s]", rs.Base.String())
}
