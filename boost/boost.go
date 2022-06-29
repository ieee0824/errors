package boost

import "errors"

func Compare(err error, f func(error) bool) bool {
	for {
		unwrapedErr := errors.Unwrap(err)
		if err == nil {
			break
		}
		compareResult := f(unwrapedErr)
		if compareResult {
			return true
		}
	}
	return false
}
