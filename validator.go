package validator

import (
	"errors"
	"strings"
)

// New creates new validator
func New() *Validator {
	return &Validator{}
}

// Validator type
type Validator struct {
	errors []error
}

// Error implements error interface
func (v *Validator) Error() string {
	return strings.Join(v.Strings(), ", ")
}

// Errors returns errors
func (v *Validator) Errors() []error {
	return v.errors
}

// String implements fmt.Stringer
func (v *Validator) String() string {
	if v.Valid() {
		return "no error"
	}
	return v.Error()
}

// Strings returns errors in strings
func (v *Validator) Strings() []string {
	s := make([]string, len(v.errors))
	for i := range v.errors {
		s[i] = v.errors[i].Error()
	}
	return s
}

// Valid returns true if no error
func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

// Must checks x must not an error or true if bool
// and return true if valid
//
// msg must be error or string
func (v *Validator) Must(x interface{}, msg interface{}) bool {
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

	v.errors = append(v.errors, m)
	return false
}
