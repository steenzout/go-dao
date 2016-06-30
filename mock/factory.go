package mock

import (
	"github.com/steenzout/go-dao"
)

// Factory struct to create mock implementations of data access objects.
type Factory struct {
	dao.BaseFactory
}

var _ dao.Factory = (*Factory)(nil)

// mockff returns an implementation of the mock data access object interface.
var mockff dao.FactoryFunc = func(ctx *dao.Context, source string) (interface{}, error) {
	base, err := ctx.NewDataAccessObject(source)
	if err != nil {
		return nil, err
	}
	return &MockDAOImpl{base}, nil
}

// NewFactory creates a factory for mock DAO implementations.
func NewFactory(ds *dao.DataSource) dao.Factory {
	return &Factory{
		dao.BaseFactory{
			Source: ds,
			FactoryFuncs: map[string]dao.FactoryFunc{
				DAOMock: mockff,
			},
		},
	}
}
