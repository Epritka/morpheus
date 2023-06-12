package example

type Person struct {
	*PersonSchema `label:"person"`

	Movies []MovieSchema `rel:"acted_in" direction:"incoming"`
}

type Movie struct {
	*MovieSchema

	Cast []PersonSchema `rel:"acted_in" direction:"outgoing"`
}

type ActedIn struct {
	*ActedInSchema

	Start *Person
	End   *Movie
}
