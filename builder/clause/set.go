package clause

import (
	"fmt"
)

type set struct {
	setter string
}

func Set(setter string) *set {
	return &set{setter: setter}
}

func (s *set) String() string {
	return fmt.Sprintf("SET %s", s.setter)
}
