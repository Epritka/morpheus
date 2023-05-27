package entity

import (
	"encoding/json"
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

	if len(b.Properties) > 0 {
		bytes, _ := json.Marshal(b.Properties)
		properties = string(bytes)
	}

	if len(b.Labels) > 0 {
		result += ":" + strings.Join(b.Labels, ":")
	}

	if result != "" && properties != "" {
		result += " "
	}

	return result + properties
}
