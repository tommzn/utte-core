package core

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type IdentifierTestSuite struct {
	suite.Suite
}

func TestIdentifierTestSuite(t *testing.T) {
	suite.Run(t, new(IdentifierTestSuite))
}

func (suite *IdentifierTestSuite) TestMewIdentifier() {

	id := NewIdentifier()
	_, err := uuid.Parse(id.String())
	suite.Nil(err)
}
