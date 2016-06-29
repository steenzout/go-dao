package mock

import "github.com/steenzout/go-dao"

// MockFactory struct to create mock implementations of data access objects.
type MockFactory struct {
	*dao.BaseFactory
}

// CreateMockDAO returns an implementation of the mock data access object interface.
func (f *MockFactory) CreateMockDAO(ctx *dao.Context) (MockDAO, error) {
	base, err := f.CreateBaseDAO(ctx)
	if err != nil {
		return nil, err
	}

	return &MockDAOImpl{*base}, nil
}

// NewFactory creates a factory for mock DAO implementations.
func NewFactory() *MockFactory {
	return &MockFactory{dao.NewBaseFactory()}
}

var _ dao.Factory = (*MockFactory)(nil)
