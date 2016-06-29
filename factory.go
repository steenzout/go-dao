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

import "fmt"

// Factory interface for data access object factories.
type Factory interface {
	DataAccessObjects() []string
	NewDataAccessObject(ctx *Context, nm string) (interface{}, error)
}

// FactoryFunc function to generate data access object implementations.
type FactoryFunc func (ctx *Context, source string) (interface{}, error)

// BaseFactory
type BaseFactory struct {
	// Source target data source against which to generate data access objects.
	Source *DataSource
	// FactoryFuncs maps data access object names to the factory functions.
	FactoryFuncs map[string]FactoryFunc
}

// DataAccessObjects returns the list of data access object names the factory can generate.
func (f *BaseFactory) DataAccessObjects() []string {
	keys := make([]string, 0, len(f.FactoryFuncs))
	for k := range f.FactoryFuncs {
		keys = append(keys, k)
	}
	return keys
}

func (f *BaseFactory) NewDataAccessObject(ctx *Context, nm string) (interface{}, error){

	ff, found := f.FactoryFuncs[nm]
	if !found {
		return nil, fmt.Errorf("this factory is not able to generate %s data access objects", nm)
	}

	return ff(ctx, f.Source.Name)
}

var _ Factory = (*BaseFactory)(nil)
