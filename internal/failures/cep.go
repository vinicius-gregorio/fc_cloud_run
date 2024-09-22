package failures

import "errors"

// ErrCepIsRequired is a error that represents a CEP is required
var ErrCepIsRequired = errors.New("CEP is required")
