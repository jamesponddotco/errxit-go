// Package errxit provides a simple and, hopefully, idiomatic interface for
// error handling designed primarily for CLI applications.
package errxit

import (
	"errors"
	"flag"
	"os"
)

// Exit is a convenience function that calls os.Exit with the error code
// associated with the error.
func Exit(err error) {
	if err == nil {
		os.Exit(0)
	}

	if errors.Is(err, flag.ErrHelp) {
		os.Exit(0)
	}

	var e *Error
	if errors.As(err, &e) {
		os.Exit(e.Code())
	}

	os.Exit(1)
}
