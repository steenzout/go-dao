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
)

// DataAccessObject base struct for data access objects.
type DataAccessObject struct {
	// Tx the database transaction.
	Tx *sql.Tx
}

// SetTransaction sets this data access object's database transaction.
func (dao *DataAccessObject) SetTransaction(tx *sql.Tx) {
	dao.Tx = tx
}

// Transaction returns this data access object's database transaction.
func (dao *DataAccessObject) Transaction() *sql.Tx {
	return dao.Tx
}
