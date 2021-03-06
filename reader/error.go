package reader

import "fmt"

type errorf struct {
	error
	position int
}

func (e *errorf) Err() error {
	return e.error
}

func (e *errorf) Position() int {
	return e.position
}

func cerror(err error) *errorf {
	return &errorf{err, -1}
}

func errorF(formatter string, position int, args ...interface{}) *errorf {
	if position > 0 {

		jargs := []interface{}{position}
		jargs = append(jargs, args...)

		return &errorf{fmt.Errorf("pos: %v "+formatter, jargs...), position}
	}

	return &errorf{fmt.Errorf(formatter, args), position}
}

func errorFmt(formatter string, args ...interface{}) *errorf {
	return &errorf{fmt.Errorf(formatter, args), -1}
}
