package dao

import "database/sql"

// DataSource struct to represent data sources.
type DataSource struct {
	// DB the database.
	DB *sql.DB
	// Name the name of the data source.
	Name string
}
