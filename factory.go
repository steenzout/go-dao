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

// Factory interface for data access object factories.
type Factory interface {
	// CreateBaseDAO returns a generic data access object.
	CreateBaseDAO(ctx *Context) (*BaseDataAccessObject, error)
	// DataSource returns this factory's target data source.
	DataSource() *DataSource
	// SetDataSource sets this factory's target data source.
	SetDataSource(ds *DataSource)
}

// BaseFactory
type BaseFactory struct {
	Source *DataSource
}

// CreateBaseDAO creates a BaseDAO with an active database transaction.
func (f *BaseFactory) CreateBaseDAO(ctx *Context) (*BaseDataAccessObject, error) {

	dao, err := NewBaseDataAccessObject(ctx, f)
	if err != nil {
		return nil, err
	}

	return dao, nil
}

// DataSource returns this factory's data source.
func (f *BaseFactory) DataSource() *DataSource {
	return f.Source
}

// SetDataSource sets this factory's data source.
func (f *BaseFactory) SetDataSource(ds *DataSource) {
	f.Source = ds
}

// NewBaseFactory returns a generic factory.
func NewBaseFactory() *BaseFactory {
	return &BaseFactory{
		Source: nil,
	}
}

var _ Factory = (*BaseFactory)(nil)
