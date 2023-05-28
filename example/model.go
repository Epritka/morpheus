package example

import (
	"fmt"
	"reflect"

	"github.com/Epritka/morpheus/v2/ogm/types"
)

type Person struct {
	*PersonSchema

	Movies []Movie `rel:"acted_in" direction:"incoming"`
}

type Movie struct {
	*MovieSchema

	Cast []Person `rel:"acted_in" direction:"outgoing"`
}

type ActedIn struct {
	*ActedInSchema

	Start *Person
	End   *Movie
}

func TestModel() *Person {
	return &Person{
		PersonSchema: &PersonSchema{
			Node: types.NewNode(),
			Name: "Keanu Reeves",
			Job:  "actor",
		},
		Movies: []Movie{
			{
				MovieSchema: &MovieSchema{
					Node:     types.NewNode(),
					Title:    "The Matrix",
					Released: 1999,
				},
			},
		},
	}
}

func CreatePerson() error {
	m := TestModel()
	model(m)
	return nil
}

func model(entity any) {
	t := reflect.TypeOf(entity).Elem()

	// ft := reflect.TypeOf(field)
	// ft := reflect.TypeOf(field).PkgPath()

	// fmt.Println(ft)
	// return
	printModel(t)
}

func printModel(t reflect.Type) {
	fields := reflect.VisibleFields(t)

	for _, field := range fields {
		fieldType := field.Type
		kind := fieldType.Kind()
		if kind == reflect.Ptr {
			fieldType = fieldType.Elem()
			kind = fieldType.Kind()
		}

		if kind == reflect.Struct {
			printModel(fieldType)
			continue
		}

		name := field.Name
		tag := field.Tag

		value := reflect.ValueOf(field)

		fmt.Println(name, "tags: ", tag.Get("ogm"), tag.Get("rel"), tag.Get("direction"))
		fmt.Println(value.String())
	}
}
