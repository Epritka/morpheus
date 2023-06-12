package entity

import (
	"fmt"
	"strings"
)

type Function struct {
	name       string
	parameters []string
	yield      []string
}

func newFunction(name string, parameters ...string) *Function {
	return &Function{name: name, parameters: parameters}
}

func (f *Function) Yeald(parameters ...string) *Function {
	f.yield = parameters
	return f
}

func (f *Function) String() string {
	yealdData := ""

	if len(f.yield) > 0 {
		yealdData = fmt.Sprintf(" YEALD %s", strings.Join(f.yield, ","))
	}

	return fmt.Sprintf("%s(%s)%s", f.name, strings.Join(f.parameters, ","), yealdData)
}
