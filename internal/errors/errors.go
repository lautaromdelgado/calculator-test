package cuestom_errors

import "errors"

var (
	ErrDivisionByZero = errors.New("division by zero is not allowed")
)
