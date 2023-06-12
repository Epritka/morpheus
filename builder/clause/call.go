package clause

import (
	"fmt"

	"github.com/Epritka/morpheus/builder/entity"
)

type call struct {
	function entity.Function
}

func Call(function entity.Function) *call {
	return &call{function: function}
}

func (c *call) String() string {
	return fmt.Sprintf("CALL %s", c.function.String())
}
