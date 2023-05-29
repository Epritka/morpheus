package executer

import (
	"fmt"
	"reflect"

	"github.com/Epritka/morpheus/v2/builder"
	"github.com/Epritka/morpheus/v2/ogm/types"
)

type Cypher struct {
	builder builder.Builder
}

func (c *Cypher) Create(entity types.Entity) *Cypher {
	t := reflect.TypeOf(entity).Elem()
	c.fillFromModel(t)
	return c
}

func (c *Cypher) fillFromModel(t reflect.Type) {
	fields := reflect.VisibleFields(t)

	for _, field := range fields {
		fieldType := field.Type
		fmt.Println(fieldType.String())
		// kind := fieldType.Kind()
		// if kind == reflect.Ptr {
		// 	fieldType = fieldType.Elem()
		// 	kind = fieldType.Kind()
		// }

		// if kind == reflect.Struct {
		// 	c.fillFromModel(fieldType)
		// 	continue
		// }

		// name := field.Name
		// tag := field.Tag

		// value := reflect.ValueOf(field)

		// fmt.Println(name, "tags: ", tag.Get("ogm"), tag.Get("rel"), tag.Get("direction"))
		// fmt.Println(value.String())
	}
}

// func (c *Cypher) Model(entities ...types.Entity) types.Model {
// 	return &Model{}
// }

// func (m *Model) Match(entities ...types.Entity) types.Model {

// 	return nil
// }

// func (m *Model) Load(entities ...types.Entity) {

// 	return nil
// }
