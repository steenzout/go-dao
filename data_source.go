//
// Copyright 2016 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package dao

import (
	"database/sql"
	"fmt"
)

// DataSource interface that needs to be implemented by data sources.
type DataSource interface {
	// Begin starts a transaction.
	// The isolation level is dependent on the driver.
	Begin() (*sql.Tx, error)
}

// DatabaseDataSource struct to represent data sources whose target are databases.
type DatabaseDataSource struct {
	// DB the database.
	DB *sql.DB
	// Name the name of the data source.
	Name string
}

// Begin starts a transaction.
// The isolation level is dependent on the driver.
func (ds *DatabaseDataSource) Begin() (*sql.Tx, error) {
	if ds.DB == nil {
		return nil, fmt.Errorf("no database was set")
	}
	return ds.DB.Begin()
}