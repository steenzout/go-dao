package mock

import (
	"database/sql"
	"fmt"

	"github.com/steenzout/go-dao"

	mock_sql "github.com/steenzout/go-mock-database/sql"
)

type DataSource struct {
	DB *mock_sql.DB
	name string
}

// Begin starts a transaction.
// The isolation level is dependent on the driver.
func (ds *DataSource) Begin() (*sql.Tx, error) {
	if ds.DB == nil {
		return nil, fmt.Errorf("no database was set")
	}
	return nil, nil
}

// Name returns the data source name.
func (ds *DataSource) Name() string {
	return ds.name
}

var _ (dao.DataSource) = (*DataSource)(nil)

func NewDataSource() dao.DataSource {
	return &DataSource{
		DB: &mock_sql.DB{},
		name: "mock",
	}
}
