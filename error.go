package errxit

import (
	"strconv"
)

// Error represents an error with an associated error code.
type Error struct {
	// Err is the underlying error.
	Err error

	// No is the error code associated with the error.
	No int
}

// New returns a new Error instance with the given error code.
func New(err error, code int) *Error {
	return &Error{
		Err: err,
		No:  code,
	}
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.Err != nil {
		return strconv.Itoa(e.No) + ": " + e.Err.Error()
	}

	return "undefined error"
}

// Code returns the error code.
func (e *Error) Code() int {
	return e.No
}
