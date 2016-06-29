package dao_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestPackage test suite for the dao package.
func TestPackage(t *testing.T) {
	suite.Run(t, new(DAOTestSuite))
	suite.Run(t, new(HandlerTestSuite))
	suite.Run(t, new(ManagerTestSuite))
}
