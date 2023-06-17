package entity

import (
	"fmt"
	"reflect"
	"strings"
)

type Base struct {
	Alias      string
	Labels     []string
	Properties map[string]any
}

func (b *Base) setAlias(alias string) *Base {
	b.Alias = alias
	return b
}

func (b *Base) setLables(labels []string) *Base {
	b.Labels = labels
	return b
}

func (b *Base) setProperties(properties map[string]any) *Base {
	b.Properties = properties
	return b
}

func (b *Base) String() string {
	result := b.Alias
	properties := ""

	i, count := 0, len(b.Properties)
	for k, v := range b.Properties {
		value := v
		switch reflect.ValueOf(v).Kind() {
		// case reflect.Ptr:
		// 	value = reflect.TypeOf(v).Elem()
		// 	switch reflect.ValueOf(v).Kind() {
		// 	}
		// case reflect.Map,
		// 	reflect.Array,
		// 	reflect.Slice:
		case reflect.Slice:
			// elem := reflect.TypeOf(v).Elem().
			// str := []string{}
			// for

		}

		properties += fmt.Sprintf("%s: %#v", k, value)
		if i < count-1 {
			properties += ","
		}
		i++
	}

	if count != 0 {
		properties = fmt.Sprintf("{ %s }", properties)
	}

	if len(b.Labels) > 0 {
		result += ":" + strings.Join(b.Labels, ":")
	}

	if result != "" && properties != "" {
		result += " "
	}

	return result + properties
}
