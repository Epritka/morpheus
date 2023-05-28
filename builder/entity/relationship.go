package entity

import "fmt"

type AnyRelationship struct{}

type Relationship struct {
	*Base
}

func (ars *AnyRelationship) String() string {
	return fmt.Sprintf("[*]")
}

func (rs *Relationship) String() string {
	return fmt.Sprintf("[%s]", rs.Base.String())
}
