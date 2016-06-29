package dao

import (
	"database/sql"
	"fmt"
)

// Context struct to hold database transaction contexts.
type Context struct {
	// Manager the data access object manager.
	manager Manager
	// transaction maps data source names with active database transactions.
	txs     map[string]*sql.Tx
}

// Transaction returns an active database transaction.
func (ctx *Context) Transaction(nm string) (*sql.Tx, error) {
	tx, found := ctx.txs[nm]
	if !found {
		source := ctx.manager.Source(nm)
		if source == nil {
			return nil, fmt.Errorf("Invalid source")
		}
		if source.DB == nil {
			return nil, fmt.Errorf("Invalid database")
		}

		tx, err := source.DB.Begin()
		if err != nil {
			return nil, err
		}
		ctx.txs[nm] = tx
		return tx, nil
	}
	return tx, nil
}

// NewContext returns a new database transaction context.
func NewContext(m Manager) *Context {
	return &Context{
		manager: m,
		txs: map[string]*sql.Tx{},
	}
}