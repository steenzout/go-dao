package dao_test

import (
	"github.com/stretchr/testify/suite"

	"github.com/steenzout/go-dao/mock"
	"github.com/steenzout/go-dao"
)

const (
	// DAO_MOCK unique identifier for the mock.MockDAO interface implementations.
	DAO_MOCK = "mock.MockDAO"
)

type TestManager struct {
	dao.BaseManager
}

func (m *TestManager) CreateMockDAO(ctx *dao.Context) (mock.MockDAO, error) {
	f, found := m.Factories[DAO_MOCK]
	if !found {
		return nil, nil
	}
	return f.(*mock.MockFactory).CreateMockDAO(ctx)
}

// ManagerSuite test suite for the Manager struct.
type ManagerTestSuite struct {
	suite.Suite
}

func (s *ManagerTestSuite) Test() {
	ctx, err := manager.StartTransaction()
	if err != nil {
		s.Fail(err.Error())
	}
	s.NotNil(ctx)
	defer manager.EndTransaction(ctx)

	mockDAO, err := manager.CreateMockDAO(ctx)
	if err != nil {
		s.Fail(err.Error())
	}

	mockDAO.MockSomething()

	err = manager.CommitTransaction(ctx)
	if err != nil {
		s.Fail(err.Error())
	}
}

var factory *mock.MockFactory
var manager TestManager
var _ dao.Manager = (*TestManager)(nil)

func init() {
	manager = TestManager{*dao.NewBaseManager()}

	ds1 := mock.NewDataSource()
	manager.RegisterDataSource("mock", ds1)

	factory = mock.NewFactory()
	factory.SetDataSource(ds1)

	manager.RegisterFactory(DAO_MOCK, factory)
}
