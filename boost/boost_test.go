package boost

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyCompareError struct {
	number int
}

func (impl *dummyCompareError) Error() string {
	return fmt.Sprintf("error code: %d", impl.number)
}

func wrap(err error) error {
	return fmt.Errorf("wrapped: %w", err)
}

type testCompareTestSet struct {
	err  error
	f    func(error) bool
	want bool
	name string
}

func TestCompare(t *testing.T) {
	compareFunc := func(err error) bool {
		unwrappedErr, ok := err.(*dummyCompareError)
		if !ok {
			return false
		}
		return unwrappedErr.number == 1
	}

	tests := []testCompareTestSet{
		{
			name: "input no wrap error: when it match",
			f:    compareFunc,
			want: true,
			err:  &dummyCompareError{number: 1},
		},
		{
			name: "input no wrap error: When it doesn't match",
			f:    compareFunc,
			want: false,
			err:  &dummyCompareError{number: 2},
		},
		{
			name: "input no wrap error: When it doesn't match: Unrelated error",
			f:    compareFunc,
			want: false,
			err:  errors.New("Unrelated error"),
		},
		{
			name: "input wrapped error: when it match",
			f:    compareFunc,
			want: true,
			err:  wrap(&dummyCompareError{number: 1}),
		},
		{
			name: "input wrapped error: When it doesn't match",
			f:    compareFunc,
			want: false,
			err:  wrap(&dummyCompareError{number: 2}),
		},
		{
			name: "input wrapped error: When it doesn't match: Unrelated error",
			f:    compareFunc,
			want: false,
			err:  wrap(errors.New("Unrelated error")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Compare(test.err, test.f)
			assert.Equal(t, test.want, actual)
		})
	}
}
