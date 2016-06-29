package dao

import "database/sql"

// Manager
type Manager interface {
	CommitTransaction(ctx *Context) error
	EndTransaction(ctx *Context)
	RegisterDataSource(nm string, ds *DataSource)
	RegisterFactory(nm string, f Factory)
	RollbackTransaction(ctx *Context) error
	Source(nm string) *DataSource
	StartTransaction() (*Context, error)
}

// BaseManager base implementation of the Manager interface.
type BaseManager struct {
	Sources map[string]*DataSource
	Factories map[string]Factory
}

// CommitTransaction
func (m *BaseManager) CommitTransaction(ctx *Context) error {
	for _, tx := range ctx.txs {
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}

// EndTransaction
func (m *BaseManager) EndTransaction(ctx *Context) {
	ctx.txs = map[string]*sql.Tx{}
}

// RegisterDataSource registers a new mapping between a name and a data source.
func (m *BaseManager) RegisterDataSource(nm string, ds *DataSource) {
	m.Sources[nm] = ds
}

// RegisterFactory registers a new mapping between a DAO name and a DAO factory.
func (m *BaseManager) RegisterFactory(nm string, f Factory) {
	m.Factories[nm] = f
}

// RollbackTransaction
func (m *BaseManager) RollbackTransaction(ctx *Context) error {
	var errout error

	for _, tx := range ctx.txs {
		if err := tx.Rollback(); err != nil {
			errout = err
		}
	}
	return errout
}

// Source returns the data source associated with the given name.
func (m *BaseManager) Source(nm string) *DataSource {
	return m.Sources[nm]
}

// StartTransaction
func (m *BaseManager) StartTransaction() (*Context, error) {
	return NewContext(m), nil
}

// NewBaseManager returns a generic Manager implementation.
func NewBaseManager() *BaseManager {
	return &BaseManager{
		Sources: map[string]*DataSource{},
		Factories: map[string]Factory{},
	}
}

var _ Manager = (*BaseManager)(nil)
