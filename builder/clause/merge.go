package clause

import (
	"fmt"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
)

type merge struct {
	entities []string
	onCreate []string
	onMatch  []string
}

type Set struct {
	key   string
	value any
}

func Merge(entities ...entity.Entity) *merge {
	entitiesString := []string{}
	for _, entity := range entities {
		entitiesString = append(entitiesString, entity.String())
	}
	return &merge{entities: entitiesString, onCreate: []string{}, onMatch: []string{}}
}

func (m *merge) OnMatchSet(key string, value any) entity.Merge {
	m.onMatch = append(m.onMatch, set(key, value))
	return m
}

func (m *merge) OnCreateSet(key string, value any) entity.Merge {
	m.onCreate = append(m.onCreate, set(key, value))
	return m
}

func (m *merge) String() string {
	body := fmt.Sprintf("MERGE %s", strings.Join(m.entities, ","))

	if len(m.onMatch) > 0 {
		body = fmt.Sprintf("%s\nON MATCH\n\tSET\n%s", body, strings.Join(m.onMatch, ",\n"))
	}

	if len(m.onCreate) > 0 {
		body = fmt.Sprintf("%s\nON CREATE\n\tSET\n%s", body, strings.Join(m.onCreate, ",\n"))
	}

	return body
}

func set(key string, value any) string {
	return fmt.Sprintf("\t\t%s = %v", key, value)
}
