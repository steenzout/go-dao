// Package dao provides a data access object library.
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

const (
	// UnknownDAOMessage error message format for the UnknownDAOError.
	UnknownDAOMessage = "There is no implementation registered for the %s data access object"
	// UnknownFactoryMessage error message format for the UnknownFactoryError.
	UnknownFactoryMessage = "There is no factory registered for the %s data access object"
)

// UnknownDAO no data access object implementation registered with the given name.
type UnknownDAO struct {
	name string
}

// Error returns the error message.
func (e *UnknownDAO) Error() string {
	return fmt.Sprintf(UnknownDAOMessage, e.name)
}

// NewUnknownDAO returns a UnknownDAO error.
func NewUnknownDAO(nm string) error {
	return &UnknownDAO{
		name: nm,
	}
}

// UnknownFactory no data access object implementation registered with the given name.
type UnknownFactory struct {
	name string
}

// Error returns the error message.
func (e *UnknownFactory) Error() string {
	return fmt.Sprintf(UnknownFactoryMessage, e.name)
}

// NewUnknownFactory returns a UnknownFactoryError error.
func NewUnknownFactory(nm string) error {
	return &UnknownFactory{
		name: nm,
	}
}

var _ error = (*UnknownDAO)(nil)
var _ error = (*UnknownFactory)(nil)
