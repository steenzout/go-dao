package dao

// Factory interface for data access object factories.
type Factory interface {
	// CreateBaseDAO returns a generic data access object.
	CreateBaseDAO(ctx *Context) (*BaseDataAccessObject, error)
	// DataSource returns this factory's target data source.
	DataSource() *DataSource
	// SetDataSource sets this factory's target data source.
	SetDataSource(ds *DataSource)
}

// BaseFactory
type BaseFactory struct {
	Source *DataSource
}

// CreateBaseDAO creates a BaseDAO with an active database transaction.
func (f *BaseFactory) CreateBaseDAO(ctx *Context) (*BaseDataAccessObject, error) {

	dao, err := NewBaseDataAccessObject(ctx, f)
	if err != nil {
		return nil, err
	}

	return dao, nil
}

// DataSource returns this factory's data source.
func (f *BaseFactory) DataSource() *DataSource {
	return f.Source
}

// SetDataSource sets this factory's data source.
func (f *BaseFactory) SetDataSource(ds *DataSource) {
	f.Source = ds
}

// NewBaseFactory returns a generic factory.
func NewBaseFactory() *BaseFactory {
	return &BaseFactory{
		Source: nil,
	}
}

var _ Factory = (*BaseFactory)(nil)
