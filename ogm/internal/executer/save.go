package executer

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
	"github.com/Epritka/morpheus/ogm/types"
)

var (
	baseEntityType = reflect.TypeOf(&types.BaseEntity{})
)

type Saver struct {
	builder  entity.Builder
	entities map[string]any
}

func NewSaver(entity types.Entity, builder entity.Builder) *Saver {
	return &Saver{}
}

// func (s *Saver) Build() string {

// }

func (e *Executer) Save(entity types.Entity) error {
	NewSaver(entity, e.Builder).parseModel(reflect.TypeOf(entity).Elem())
	return nil
}

func (s *Saver) parseModel(t reflect.Type) {
	// name := t.Name()

	// fields := reflect.VisibleFields(t)
	// structEntity := t.Elem()
	size := t.NumField()
	for i := 0; i < size; i++ {
		field := t.Field(i)
		name := field.Name

		if strings.Contains(name, "Schema") {
			fields := reflect.VisibleFields(field.Type.Elem())

			for i, f := range fields {
				if i == 0 {
					entityType := f.Name

					if entityType == "Node" {
						fmt.Println("Node")
					}

					if entityType == "Relationship" {
						fmt.Println("Relationship")
					}

					continue
				}

				if i == 1 {
					continue
				}

				if i == 2 {
					// fmt.Println(f.)
				}

				value := reflect.ValueOf(f)

				o := value.Interface()
				// fmt.Println(f.Name)
				fmt.Println(reflect.TypeOf(o))
			}
			// schema := field.Type
			// sizeSchema := schema.NumField()
			// schema.Elem()
			// base := schema.Field(0).Type.String()
			// fmt.Println(base)
			// for i := 1; i < sizeSchema; i++ {

			// }
			continue

		}
		// if field.Type == baseEntityType {
		// 	fmt.Println("fdfdf")
		// 	// return
		// }

		fmt.Println("other field", field.Name)
		// fieldType := field.Type
		// fmt.Println(fieldType.String())
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

		// // value := reflect.ValueOf(field)

		// fmt.Println(field.Name, "tags: ", tag.Get("ogm"), tag.Get("rel"), tag.Get("direction"))
		fmt.Println()
		// fmt.Println(value.String())
	}
}

func (s *Saver) parseSchema(entitySchema reflect.Type) {
	// fields := map[string]any{}
}

// func (e *Executer) SaveWithContext(ctx context.Context, entity types.Entity) (types.Result, error) {
// 	return Result{}, nil
// }
