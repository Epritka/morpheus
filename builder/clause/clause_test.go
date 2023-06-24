package clause_test

import (
	"testing"

	"github.com/Epritka/morpheus/builder"
	"github.com/Epritka/morpheus/builder/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClauseTestSuite struct {
	suite.Suite
}

func (suite *ClauseTestSuite) SetupTest() {}

func (suite *ClauseTestSuite) TestMatch() {
	b := builder.NewBuilder().Match(
		entity.NewNode("alias").SetLables("Label").
			SetProperties(map[string]any{
				"key": "value",
			}))

	assert.Equal(suite.T(), b.Build(), "MATCH (alias:Label { key: \"value\" })", "they should be equal")
}

func TestRun(t *testing.T) {
	suite.Run(t, new(ClauseTestSuite))
}
