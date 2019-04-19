package validator

import (
	"errors"
	"strings"
)

// Error is the validate error
type Error struct {
	errors []error
}

// Error implements error interface
func (err *Error) Error() string {
	return strings.Join(err.Strings(), ", ")
}

// Errors returns errors
func (err *Error) Errors() []error {
	return err.errors
}

// Strings returns errors in strings
func (err *Error) Strings() []string {
	s := make([]string, len(err.errors))
	for i := range err.errors {
		s[i] = err.errors[i].Error()
	}
	return s
}

func (err *Error) clone() *Error {
	return &Error{errors: err.errors}
}

// IsError returns true if given error is validate error
func IsError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

// New creates new validator
func New() *Validator {
	return &Validator{}
}

// Validator type
type Validator struct {
	err Error
}

// Error returns error if has error
func (v *Validator) Error() error {
	if !v.Valid() {
		return v.err.clone()
	}
	return nil
}

// String implements fmt.Stringer
func (v *Validator) String() string {
	if v.Valid() {
		return "no error"
	}
	return v.err.Error()
}

// Valid returns true if no error
func (v *Validator) Valid() bool {
	return len(v.err.errors) == 0
}

// Must checks x must not an error or true if bool
// and return true if valid
//
// msg must be error or string
func (v *Validator) Must(x interface{}, msg interface{}) bool {
	if x == nil {
		return true
	}

	switch x := x.(type) {
	case bool:
		if x {
			return true
		}
	case error:
		if x == nil {
			return true
		}
	default:
		panic("validator: invalid input")
	}

	var m error
	switch t := msg.(type) {
	case error:
		m = t
	case string:
		m = errors.New(t)
	default:
		panic("validator: invalid msg")
	}

	v.err.errors = append(v.err.errors, m)
	return false
}

// Add adds errors
func (v *Validator) Add(err ...error) {
	v.err.errors = append(v.err.errors, err...)
}
