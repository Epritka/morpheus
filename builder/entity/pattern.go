package entity

import (
	"fmt"
)

type directedType int

const (
	omitted directedType = iota
	in
	out
)

type PatternList struct {
	start Entity
	first *Pattern
	last  *Pattern

	lastRelationship Entity
}

type Pattern struct {
	node         Entity
	relationship Entity
	directedType directedType
	next         *Pattern
}

func NewPatternList(entity Entity) *PatternList {
	return &PatternList{
		start: entity,
		first: nil,
		last:  nil,
	}
}

func NewPattern(node Entity, ddirectedType directedType) *Pattern {
	return &Pattern{
		node:         node,
		directedType: ddirectedType,
		next:         nil,
		relationship: nil,
	}
}
func (pl *PatternList) Empty() bool {
	return pl.start == nil && pl.first == nil
}

func (pl *PatternList) Related(relationship Entity) *PatternList {
	pl.lastRelationship = relationship
	return pl
}

func (pl *PatternList) Join(entity Entity) *PatternList {
	return pl.append(NewPattern(entity, omitted))
}

func (pl *PatternList) To(entity Entity) *PatternList {
	return pl.append(NewPattern(entity, in))
}

func (pl *PatternList) From(entity Entity) *PatternList {
	return pl.append(NewPattern(entity, out))
}

func (pl *PatternList) String() string {
	if pl.Empty() {
		return ""
	}

	result := pl.start.String()
	current := pl.first

	for current != nil {
		result += current.string()
		current = current.next
	}

	return result
}

func (pl *PatternList) append(pattern *Pattern) *PatternList {
	if pl.lastRelationship != nil {
		pattern.relationship = pl.lastRelationship
		pl.lastRelationship = nil
	}

	if pl.first == nil {
		pl.first = pattern
		pl.last = pattern
		return pl
	}

	pl.last.next = pattern
	pl.last = pattern
	return pl
}

func (p *Pattern) string() string {
	nodeString := p.node.String()
	relationshipString := ""

	if p.relationship != nil {
		relationshipString = p.relationship.String()
	}

	switch p.directedType {
	case in:
		return fmt.Sprintf("-%s->%s", relationshipString, nodeString)
	case out:
		return fmt.Sprintf("<-%s-%s", relationshipString, nodeString)
	default:
		return fmt.Sprintf("-%s-%s", relationshipString, nodeString)
	}
}
