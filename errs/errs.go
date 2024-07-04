package errs

import "errors"

var (
	ErrMethodUndefined = errors.New("method is undefined")
	ErrCannotLock      = errors.New("can not lock")
)
