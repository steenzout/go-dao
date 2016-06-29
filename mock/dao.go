package mock

import (
	"github.com/steenzout/go-dao"
)

const (
	// DAO_MOCK unique data access object name for the MockDAO interface.
	DAO_MOCK = "go-dao.mock.MockDAO"
)

// MockDAO interface for the mock data access object.
type MockDAO interface {
	// MockSomething does nothing.
	MockSomething() error
}

// MockDAOImpl mock implementation of the MockDAO interface.
type MockDAOImpl struct {
	*dao.DataAccessObject
}

// MockSomething does nothing.
func (m *MockDAOImpl) MockSomething() error {
	return nil
}

var _ MockDAO = (*MockDAOImpl)(nil)
