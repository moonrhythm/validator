package validator

import (
	"errors"
	"fmt"
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
// msg can be nil if x is error
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

	if msg == nil {
		if err, ok := x.(error); ok {
			v.addError(err)
			return false
		}
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

	v.addError(m)
	return false
}

// Mustf calls fmt.Sprintf(format, a...) and pass to v.Must
func (v *Validator) Mustf(x interface{}, format string, a ...interface{}) bool {
	return v.Must(x, fmt.Sprintf(format, a...))
}

// Add adds errors
func (v *Validator) Add(err ...error) {
	for _, e := range err {
		if e == nil {
			continue
		}
		v.addError(e)
	}
}

func (v *Validator) addError(err error) {
	if IsError(err) {
		v.err.errors = append(v.err.errors, err.(*Error).errors...)
		return
	}
	v.err.errors = append(v.err.errors, err)
}
