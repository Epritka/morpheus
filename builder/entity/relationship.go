package entity

import "fmt"

type AnyRelationship struct{}

type Relationship struct {
	Base
}

func (rs *Relationship) String() string {
	return fmt.Sprintf("[%s]", rs.Base.String())
}

func (ars *AnyRelationship) String() string {
	return fmt.Sprintf("[*]")
}
