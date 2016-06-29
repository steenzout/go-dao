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
