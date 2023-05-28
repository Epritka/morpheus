package clause

import (
	"fmt"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
)

type create struct {
	entities []string
}

func Create(entities ...entity.Entity) *create {
	return newCreate(entities)
}

func newCreate(entities []entity.Entity) *create {
	entitiesString := []string{}
	for _, entity := range entities {
		entitiesString = append(entitiesString, entity.String())
	}

	return &create{entities: entitiesString}
}

func (c *create) String() string {
	return fmt.Sprintf("CREATE %s", strings.Join(c.entities, ", "))
}
