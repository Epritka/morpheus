package clause

import "fmt"

type where struct {
	condition string
}

func Where(condition string) *where {
	return &where{condition: condition}
}

func (w *where) String() string {
	return fmt.Sprintf("WHERE %s", w.condition)
}
