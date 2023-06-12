package clause

import (
	"fmt"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
)

type delete struct {
	isDetach bool
	entities []string
}

func Delete(entities ...entity.Entity) *delete {
	return newDelete(entities)
}

func DetachDelete(entities ...entity.Entity) *delete {
	delete := newDelete(entities)
	delete.isDetach = true
	return delete
}

func newDelete(entities []entity.Entity) *delete {
	entitiesString := []string{}
	for _, entity := range entities {
		entitiesString = append(entitiesString, entity.String())
	}

	return &delete{entities: entitiesString}
}

func (d *delete) String() string {
	result := ""

	if d.isDetach {
		result = "DETACH "
	}

	return fmt.Sprintf("%sDELETE %s", result, strings.Join(d.entities, ", "))
}
