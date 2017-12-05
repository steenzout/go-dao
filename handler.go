// Package dao provides a data access object library.
//
// Copyright 2016-2017 Pedro Salgado
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
	"fmt"
)

// TransactionFunc definition of a function wrapped in a database transaction context.
type TransactionFunc func(ctx *Context) error

// Process wrap database transaction handling around given TransactionFunc.
func Process(m Manager, f TransactionFunc) (err error) {
	var ctx *Context

	ctx, err = m.StartTransaction()
	if err != nil {
		return err
	}
	defer func() {
		r := recover()
		if r != nil {
			m.RollbackTransaction(ctx)
			err = fmt.Errorf("panic: %v", r)
		}
		m.EndTransaction(ctx)
	}()

	err = f(ctx)
	if err != nil {
		m.RollbackTransaction(ctx)
		return err
	}

	err = m.CommitTransaction(ctx)
	if err != nil {
		m.RollbackTransaction(ctx)
		return err
	}

	return nil
}
