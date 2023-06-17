package example

import "github.com/Epritka/morpheus/ogm/types"

// type Person struct {
// 	*PersonSchema `label:"person"`

// 	Movies []MovieSchema `rel:"acted_in" direction:"incoming"`
// }

// type Movie struct {
// 	*MovieSchema

// 	Cast []PersonSchema `rel:"acted_in" direction:"outgoing"`
// }

// type ActedIn struct {
// 	*ActedInSchema

// 	Start *Person
// 	End   *Movie
// }

type Person struct {
	*types.Node
	Job  string `json:"job"`
	Name string `json:"name"`
}

type Movie struct {
	*types.Node
	Released int64  `json:"released"`
	Title    string `json:"title"`
}

func (p *Person) Type() string {
	return "person"
}

func (p *Person) Labels() []string {
	return []string{}
}

func (p *Person) Properies() map[string]any {
	return map[string]any{
		"job":  p.Job,
		"name": p.Name,
	}
}

func (m *Movie) Type() string {
	return "movie"
}

func (m *Movie) Labels() []string {
	return []string{}
}

func (m *Movie) Properies() map[string]any {
	return map[string]any{
		"released": m.Released,
		"title":    m.Title,
	}
}

type ActedIn struct {
	*types.Relationship
	Role string `json:"role"`
}
