package boost

import (
	"errors"
)

func Compare(err error, f func(error) bool) bool {
	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr == nil {
		return f(err)
	}
	compareResult := f(err)
	if compareResult {
		return true
	}
	return Compare(unwrappedErr, f)
}
