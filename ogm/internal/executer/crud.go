package executer

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
	"github.com/Epritka/morpheus/ogm/types"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

var (
	baseEntityType = reflect.TypeOf(&types.BaseEntity{})
)

// type Saver struct {
// 	builder  entity.Builder
// 	entities map[string]any
// }

// func NewSaver(builder entity.Builder) *Saver {
// 	return &Saver{}
// }

func (e *Executer) Save(model types.Entity) error {
	e.Builder.Create(&entity.Node{
		Base: &entity.Base{
			Alias:      "n",
			Labels:     append(model.Labels(), model.Type()),
			Properties: model.Properies(),
		},
	}).Return(entity.NewAlias("n"))

	err := e.doParse(context.Background(), func(rwc neo4j.ResultWithContext) error {
		record, err := rwc.Single(context.Background())
		if err != nil {
			return err
		}

		data, _ := record.Get("n")
		node := data.(dbtype.Node)

		id, err := strconv.Atoi(strings.Split(node.ElementId, ":")[2])
		if err != nil {
			return err
		}

		model.SetId(int64(id))
		model.SetElementId(node.ElementId)

		bytes, _ := json.Marshal(node.Props)
		json.Unmarshal(bytes, model)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *Executer) Update(model types.Entity) error {
	bytes, _ := json.Marshal(model.Properies())
	e.Builder.
		Match(&entity.Node{Base: &entity.Base{Alias: "n"}}).
		Where(fmt.Sprintf("ID(n) = %s", model.GetElementId())).
		Set(fmt.Sprintf("n = %s", string(bytes))).
		Return(entity.NewAlias("n"))

	err := e.doParse(context.Background(), func(rwc neo4j.ResultWithContext) error {
		record, err := rwc.Single(context.Background())
		if err != nil {
			return err
		}

		data, _ := record.Get("n")
		node := data.(dbtype.Node)

		id, err := strconv.Atoi(strings.Split(node.ElementId, ":")[2])
		if err != nil {
			return err
		}

		model.SetId(int64(id))
		model.SetElementId(node.ElementId)

		bytes, _ := json.Marshal(node.Props)
		json.Unmarshal(bytes, model)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// func (s *Saver) Build(model types.Entity) error {

// }

// func (s *Saver) parseModel(t reflect.Type) {
// 	// name := t.Name()

// 	// fields := reflect.VisibleFields(t)
// 	// structEntity := t.Elem()
// 	size := t.NumField()
// 	for i := 0; i < size; i++ {
// 		field := t.Field(i)
// 		name := field.Name

// 		if strings.Contains(name, "Schema") {
// 			fields := reflect.VisibleFields(field.Type.Elem())

// 			for i, f := range fields {
// 				if i == 0 {
// 					entityType := f.Name

// 					if entityType == "Node" {
// 						fmt.Println("Node")
// 					}

// 					if entityType == "Relationship" {
// 						fmt.Println("Relationship")
// 					}

// 					continue
// 				}

// 				if i == 1 {
// 					continue
// 				}

// 				if i == 2 {
// 					// fmt.Println(f.)
// 				}

// 				value := reflect.ValueOf(f)

// 				o := value.Interface()
// 				// fmt.Println(f.Name)
// 				fmt.Println(reflect.TypeOf(o))
// 			}
// 			// schema := field.Type
// 			// sizeSchema := schema.NumField()
// 			// schema.Elem()
// 			// base := schema.Field(0).Type.String()
// 			// fmt.Println(base)
// 			// for i := 1; i < sizeSchema; i++ {

// 			// }
// 			continue
// 		}
// 		// if field.Type == baseEntityType {
// 		// 	fmt.Println("fdfdf")
// 		// 	// return
// 		// }

// 		fmt.Println("other field", field.Name)
// 		// fieldType := field.Type
// 		// fmt.Println(fieldType.String())
// 		// kind := fieldType.Kind()
// 		// if kind == reflect.Ptr {
// 		// 	fieldType = fieldType.Elem()
// 		// 	kind = fieldType.Kind()
// 		// }

// 		// if kind == reflect.Struct {
// 		// 	c.fillFromModel(fieldType)
// 		// 	continue
// 		// }

// 		// name := field.Name
// 		// tag := field.Tag

// 		// // value := reflect.ValueOf(field)

// 		// fmt.Println(field.Name, "tags: ", tag.Get("ogm"), tag.Get("rel"), tag.Get("direction"))
// 		fmt.Println()
// 		// fmt.Println(value.String())
// 	}
// }

// func (s *Saver) parseSchema(entitySchema reflect.Type) {
// 	// fields := map[string]any{}
// }

// func (e *Executer) SaveWithContext(ctx context.Context, entity types.Entity) (types.Result, error) {
// 	return Result{}, nil
// }
