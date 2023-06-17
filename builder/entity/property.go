package entity

import "fmt"

type Alias struct {
	alias string
	as    string
}

func NewAlias(alias string) Property {
	return &Alias{alias: alias}
}

func (a *Alias) As(as string) Property {
	a.as = as
	return a
}

func (a *Alias) String() string {
	result := a.alias
	if a.as != "" {
		return fmt.Sprintf("%s AS %s", result, a.as)
	}
	return result
}
