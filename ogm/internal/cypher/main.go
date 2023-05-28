package cypher

import (
	"github.com/Epritka/morpheus/v2/builder"
	"github.com/Epritka/morpheus/v2/builder/entity"
)

type Cypher struct {
	builder.Builder
}

func New() *Cypher {
	return &Cypher{Builder: builder.NewBuilder()}
}

func (c *Cypher) Save(model any) *Cypher {
	c.Create(&entity.Node{
		Base: &entity.Base{
			Alias:  "a",
			Labels: []string{"Actor"},
			Properties: map[string]any{
				"name":     "Name",
				"lastName": "LastName",
			},
		},
	})

	return c
}
