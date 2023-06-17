package builder

import (
	"github.com/Epritka/morpheus/builder/clause"
	"github.com/Epritka/morpheus/builder/entity"
)

type builder struct {
	firstClause *сlause
	lastClause  *сlause
}

type сlause struct {
	entity.Clause
	Next *сlause
}

func NewBuilder() entity.Builder {
	return &builder{firstClause: nil, lastClause: nil}
}

func (b *builder) Append(clauseI entity.Clause) entity.Builder {
	clause := &сlause{Clause: clauseI, Next: nil}

	if b.empty() {
		b.firstClause = clause
		b.lastClause = clause
		return b
	}

	b.lastClause.Next = clause
	b.lastClause = clause
	return b
}

func (b *builder) Where(condition string) entity.Builder {
	return b.Append(clause.Where(condition))
}

func (b *builder) Create(entities ...entity.Entity) entity.Builder {
	return b.Append(clause.Create(entities...))
}

func (b *builder) Match(entities ...entity.Entity) entity.Builder {
	return b.Append(clause.Match(entities...))
}

func (b *builder) OptionalMatch(entities ...entity.Entity) entity.Builder {
	return b.Append(clause.OptionalMatch(entities...))
}

func (b *builder) Set(setter string) entity.Builder {
	return b.Append(clause.Set(setter))
}

func (b *builder) Merge(entities ...entity.Entity) entity.Merge {
	merge := clause.Merge(entities...)
	b.Append(merge)
	return merge
}

func (b *builder) Call(function entity.Function) entity.Builder {
	return b.Append(clause.Call(function))
}

func (b *builder) With(properties ...entity.Property) entity.Builder {
	return b.Append(clause.With(properties...))
}

func (b *builder) Delete(entities ...entity.Entity) entity.Builder {
	return b.Append(clause.Delete(entities...))
}

func (b *builder) DetachDelete(entities ...entity.Entity) entity.Builder {
	return b.Append(clause.DetachDelete(entities...))
}

func (b *builder) Return(properties ...entity.Property) entity.Builder {
	return b.Append(clause.Return(properties...))
}

func (b *builder) Limit(limit int) entity.Builder {
	return b.Append(clause.Limit(limit))
}

func (b *builder) Build() string {
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
