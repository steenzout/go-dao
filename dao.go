package dao

import (
	"database/sql"
	"fmt"
)

// DataAccessObject interface that data access objects need to implement.
type DataAccessObject interface {
	// SetTransaction sets this data access object's database transaction.
	SetTransaction(tx *sql.Tx)
	// Transaction returns this data access object's database transaction.
	Transaction() *sql.Tx
}

// BaseDataAccessObject base struct for data access objects.
type BaseDataAccessObject struct {
	// Tx the database transaction.
	Tx *sql.Tx
}

// SetTransaction sets this data access object's database transaction.
func (dao *BaseDataAccessObject) SetTransaction(tx *sql.Tx) {
	dao.Tx = tx
}

// Transaction returns this data access object's database transaction.
func (dao *BaseDataAccessObject) Transaction() *sql.Tx {
	return dao.Tx
}

// NewBaseDataAccessObject returns a generic data access object.
func NewBaseDataAccessObject(ctx *Context, f Factory) (*BaseDataAccessObject, error) {

	if f.DataSource() == nil {
		// TODO unknown source
		return nil, fmt.Errorf("unknown source")
	}

	tx, err := ctx.Transaction(f.DataSource().Name)
	if err != nil {
		return nil, err
	}

	return &BaseDataAccessObject{
		Tx: tx,
	}, nil
}
