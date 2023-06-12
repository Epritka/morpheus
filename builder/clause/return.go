package clause

import (
	"fmt"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
)

type returnClause struct {
	properties []string
}

func Return(properties ...entity.Property) *returnClause {
	props := []string{}
	for _, p := range properties {
		props = append(props, p.String())
	}

	return &returnClause{properties: props}
}

func (r *returnClause) String() string {
	return fmt.Sprintf("RETURN %s", strings.Join(r.properties, ", "))
}
