package dao_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// PackageTestSuite test suite for the dao package.
func PackageTestSuite(t *testing.T) {
	suite.Run(t, new(DAOTestSuite))
	suite.Run(t, new(HandlerTestSuite))
	suite.Run(t, new(ManagerTestSuite))
}
