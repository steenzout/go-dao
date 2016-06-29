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

// Context struct to hold database transaction contexts.
type Context struct {
	// Manager the data access object manager.
	manager Manager
	// transaction maps data source names with active database transactions.
	txs map[string]*sql.Tx
}

// Transaction returns an active database transaction.
func (ctx *Context) Transaction(nm string) (*sql.Tx, error) {
	tx, found := ctx.txs[nm]
	if !found {
		source := ctx.manager.Source(nm)
		if source == nil {
			return nil, fmt.Errorf("Invalid source")
		}

		tx, err := source.Begin()
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
		txs:     map[string]*sql.Tx{},
	}
}
