package builder

import (
	"github.com/Epritka/morpheus/v2/builder/clause"
	"github.com/Epritka/morpheus/v2/builder/entity"
)

type Builder interface {
	Create(entities ...entity.Entity) Builder
	String() string
}

type ClauseI interface {
	String() string
}

type builder struct {
	firstClause *Clause
	lastClause  *Clause
}

type Clause struct {
	ClauseI
	Next *Clause
}

func NewBuilder() Builder {
	return &builder{firstClause: nil, lastClause: nil}
}

func (b *builder) Create(entities ...entity.Entity) Builder {
	return b.append(clause.Create(entities...))
}

func (b *builder) String() string {
	query := ""
	if b.empty() {
		return query
	}

	currentClause := b.firstClause
	for currentClause != nil {
		query += currentClause.String()
		currentClause = currentClause.Next
		if currentClause != nil {
			query += "\n"
		}
	}

	return query
}

func (b *builder) empty() bool {
	return b.firstClause == nil
}

func (b *builder) append(clauseI ClauseI) *builder {
	clause := &Clause{ClauseI: clauseI, Next: nil}

	if b.empty() {
		b.firstClause = clause
		b.lastClause = clause
		return b
	}

	b.lastClause.Next = clause
	b.lastClause = clause
	return b
}
