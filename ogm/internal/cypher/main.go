package cypher

import (
	"github.com/Epritka/morpheus/v2/builder"
	"github.com/Epritka/morpheus/v2/ogm/types"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Cypher struct {
	model types.Entity
	builder.Builder
}

func New() *Cypher {
	return &Cypher{Builder: builder.NewBuilder()}
}

func (c *Cypher) SetResult(result neo4j.ResultWithContext) *Cypher {
	return c
}

// func (c *Cypher) Save(model any) *Cypher {
// 	c.Create(&entity.Node{
// 		Base: &entity.Base{
// 			Alias:  "a",
// 			Labels: []string{"Actor"},
// 			Properties: map[string]any{
// 				"name":     "Name",
// 				"lastName": "LastName",
// 			},
// 		},
// 	})

// 	return c
// }
