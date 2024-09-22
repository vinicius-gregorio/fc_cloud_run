package failures

import "errors"

// ErrCepIsRequired is a error that represents a CEP is required
var ErrCepIsRequired = errors.New("CEP is required")

// ErrCepInvalid is a error that represents a CEP is invalid
var ErrCepInvalid_Digits = errors.New("CEP is invalid - it must contain only numbers")
var ErrCepInvalid_Length = errors.New("CEP is invalid - it must contain 8 digits")

// ErrStateIsRequired is a error that represents a State is required
var ErrStateIsRequired = errors.New("state is required")

var ErrLocalityIsRequired = errors.New("locality is required")
