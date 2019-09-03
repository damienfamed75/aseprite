package aseprite

import (
	"strings"
)

var (
	errorAnimationNotFound = &Error{
		subject: "animation not found in aseprite file",
	}
)

// Error is used as a custom error for the aseprite package.
type Error struct {
	subject string
	params  []string
	err     error
}

func (e *Error) withParams(p ...string) *Error {
	e.params = p
	return e
}

func (e *Error) withError(err error) *Error {
	e.err = err
	return e
}

// Unwrap unwraps and returns the original error if one exists.
// Note: Go 1.13 change
func (e *Error) Unwrap() error {
	if e.err != nil {
		return e.err
	}

	return nil
}

func (e *Error) Error() string {
	var b strings.Builder
	defer b.Reset()

	b.WriteString(e.subject)

	if e.err != nil {
		b.WriteString(": error[" + e.err.Error() + "]")
	}

	if len(e.params) != 0 {
		b.WriteString(": params[")
		for i, p := range e.params {
			b.WriteString(p)

			if i != len(e.params)-1 {
				b.WriteByte(',')
			}
		}
		b.WriteByte(']')
	}

	return b.String()
}
