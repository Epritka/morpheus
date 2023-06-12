package entity

type Builder interface {
	Append(Clause) Builder

	Where(string) Builder
	Create(...Entity) Builder

	Match(...Entity) Builder
	OptionalMatch(...Entity) Builder

	Merge(entities ...Entity) Merge

	Call(Function) Builder
	With(...Property) Builder

	Delete(...Entity) Builder
	DetachDelete(...Entity) Builder

	Return(...Property) Builder

	Limit(int) Builder

	Build() string
}

type Clause interface {
	String() string
}

type Merge interface {
	Clause
	OnCreateSet(key string, value any) Merge
	OnMatchSet(key string, value any) Merge
}
type Entity interface {
	String() string
}

type Property interface {
	As(string) Property
	String() string
}
