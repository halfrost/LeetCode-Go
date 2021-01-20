package main

import (
	"fmt"
	"os"
)

const (
	// ExitSuccess define
	ExitSuccess = iota
	// ExitError define
	ExitError
	// ExitBadArgs define
	ExitBadArgs
)

// ExitWithError define
func ExitWithError(code int, err error) {
	fmt.Fprintln(os.Stderr, "Error: ", err)
	os.Exit(code)
}
