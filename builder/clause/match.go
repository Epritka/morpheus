package clause

import (
	"fmt"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
)

type match struct {
	entities   []string
	isOptional bool
}

func Match(entities ...entity.Entity) *match {
	return newMatch(entities)
}

func OptionalMatch(entities ...entity.Entity) *match {
	match := newMatch(entities)
	match.isOptional = true
	return match
}

func newMatch(entities []entity.Entity) *match {
	entitiesString := []string{}
	for _, entity := range entities {
		entitiesString = append(entitiesString, entity.String())
	}

	return &match{entities: entitiesString}
}

func (m *match) String() string {
	result := ""

	if m.isOptional {
		result = "OPTIONAL "
	}

	return fmt.Sprintf("%sMATCH %s", result, strings.Join(m.entities, ", "))
}
