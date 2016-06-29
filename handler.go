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

type TransactionFunc func(m Manager, ctx *Context) error

type TransactionFunc2 func(m Manager, ctx *Context) (interface{}, error)

func Process(m Manager, f TransactionFunc) error {
	ctx, err := m.StartTransaction()
	if err != nil {
		return err
	}
	defer m.EndTransaction(ctx)

	err = f(m, ctx)
	if err != nil {
		m.RollbackTransaction(ctx)
		return err
	}

	if err = m.CommitTransaction(ctx); err != nil {
		m.RollbackTransaction(ctx)
		return err
	}
	return nil
}

func Process2(m Manager, f TransactionFunc2) (interface{}, error) {
	ctx, err := m.StartTransaction()
	if err != nil {
		return nil, err
	}
	defer m.EndTransaction(ctx)

	obj, err := f(m, ctx)
	if err != nil {
		m.RollbackTransaction(ctx)
		return nil, err
	}

	if err = m.CommitTransaction(ctx); err != nil {
		m.RollbackTransaction(ctx)
		return nil, err
	}
	return obj, nil
}
