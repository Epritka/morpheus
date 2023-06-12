package example

import "github.com/Epritka/morpheus/ogm/types"

type PersonSchema struct {
	*types.Node
	Job  string `ogm:"job"`
	Name string `ogm:"name"`
}

type MovieSchema struct {
	*types.Node
	Released int64  `ogm:"released"`
	Title    string `ogm:"title"`
}

type ActedInSchema struct {
	*types.Relationship
	Roles []string `ogm:"roles"`
}
