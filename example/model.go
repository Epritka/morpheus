package example

type Person struct {
	*PersonSchema

	Movies []Movie `rel:"acted_in" direction:"incoming"`
}

type Movie struct {
	*MovieSchema

	Cast []Person `rel:"acted_in" direction:"outgoing"`
}

type ActedIn struct {
	*ActedInSchema

	Start *Person
	End   *Movie
}
