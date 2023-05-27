package clause

import (
	"fmt"
	"strings"
)

type returnClause struct {
	limit    string
	entities []string
}

func Return(entities []string) *returnClause {
	return &returnClause{entities: entities}
}

func (r *returnClause) Limit(number int) *returnClause {
	r.limit = fmt.Sprintf(" LIMIT %d", number)
	return r
}

func (r *returnClause) String() string {
	return fmt.Sprintf("RETURN %s%s", strings.Join(r.entities, ", "), r.limit)
}
