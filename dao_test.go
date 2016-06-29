package dao_test

import (
	"github.com/steenzout/go-dao"
	"github.com/stretchr/testify/suite"
)

type SomeDAO struct {
	dao.BaseDataAccessObject
}

func (dao *SomeDAO) retrieveSomething() {
	return
}

// DAOTestSuite test suite for the DataAccessObject interface and
// BaseDataAccessObject struct.
type DAOTestSuite struct {
	suite.Suite
	dao dao.DataAccessObject
}

func (s *DAOTestSuite) SetupTest() {
	s.dao = &SomeDAO{dao.BaseDataAccessObject{Tx: nil}}
}
