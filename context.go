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

	"github.com/davecgh/go-spew/spew"
)

// Context struct to hold database transaction contexts.
type Context struct {
	// Manager the data access object manager.
	manager Manager
	// daos maps data source names with active database transactions.
	daos    map[string]*DataAccessObject
}

// DataAccessObject returns a data access object with an active database transaction.
func (ctx *Context) NewDataAccessObject(nm string) (*DataAccessObject, error) {
	var dao *DataAccessObject

	dao, found := ctx.daos[nm]
	if !found {
		tx, err := ctx.Transaction(nm)
		if err != nil {
			return nil, err
		}

		dao := &DataAccessObject{Tx: tx}
		ctx.daos[nm] = dao
		return dao, nil
	}
	return dao, nil
}

// Transaction returns a new database transaction from the given data source.
func (ctx *Context) Transaction(nm string) (*sql.Tx, error) {

	source := ctx.manager.Source(nm)
	if source == nil {
		spew.Printf("%+v", ctx.manager)
		spew.Printf("%+v", ctx)
		return nil, fmt.Errorf("source %s was not found", nm)
	}

	tx, err := source.DB.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// NewContext returns a new database transaction context.
func NewContext(m Manager) *Context {
	return &Context{
		manager: m,
		daos:     map[string]*DataAccessObject{},
	}
}
