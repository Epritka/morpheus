package entity

import (
	"fmt"
	"strings"
)

type Base struct {
	Alias      string
	Labels     []string
	Properties map[string]any
}

func (b *Base) String() string {
	result := b.Alias
	properties := ""

	i, count := 0, len(b.Properties)
	for k, v := range b.Properties {
		properties += fmt.Sprintf("%s: %#v", k, v)
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
