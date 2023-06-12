package clause

import (
	"fmt"
	"strings"

	"github.com/Epritka/morpheus/builder/entity"
)

type with struct {
	properties []string
}

func With(properties ...entity.Property) *with {
	props := []string{}
	for _, p := range properties {
		props = append(props, p.String())
	}

	return &with{properties: props}
}

func (r *with) String() string {
	return fmt.Sprintf("WITH %s", strings.Join(r.properties, ", "))
}
