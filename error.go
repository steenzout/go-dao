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
	// ERR_UNKNOWN_DAO error message format for the UnknownDAOError.
	ERR_UNKNOWN_DAO = "There is no implementation registered for the %s data access object"
	// ERR_UNKNOWN_FACTORY error message format for the UnknownFactoryError.
	ERR_UNKNOWN_FACTORY = "There is no factory registered for the %s data access object"
)

// UnknownDAOError no data access object implementation registered with the given name.
type UnknownDAOError struct {
	name string
}

// Error returns the error message.
func (e *UnknownDAOError) Error() string {
	return fmt.Sprintf(ERR_UNKNOWN_DAO, e.name)
}

// NewUnknownDAO returns a UnknownDAO error.
func NewUnknownDAOError(nm string) error {
	return &UnknownDAOError{
		name: nm,
	}
}

// UnknownFactoryError no data access object implementation registered with the given name.
type UnknownFactoryError struct {
	name string
}

// Error returns the error message.
func (e *UnknownFactoryError) Error() string {
	return fmt.Sprintf(ERR_UNKNOWN_FACTORY, e.name)
}

// NewUnknownFactoryError returns a UnknownFactoryError error.
func NewUnknownFactoryError(nm string) error {
	return &UnknownFactoryError{
		name: nm,
	}
}

var _ error = (*UnknownDAOError)(nil)
var _ error = (*UnknownFactoryError)(nil)
