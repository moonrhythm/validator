package validator

import (
	"strings"
)

// New creates new validator
func New() *Validator {
	return &Validator{}
}

// Validator type
type Validator struct {
	errors []string
}

func (v *Validator) Error() string {
	return strings.Join(v.errors, ", ")
}

// Errors returns errors
func (v *Validator) Errors() []string {
	return v.errors
}

// Valid returns true if no error
func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

// Must checks x must not an error or true if bool
func (v *Validator) Must(x interface{}, msg string) {
	switch x := x.(type) {
	case bool:
		if x {
			return
		}
	case error:
		if x == nil {
			return
		}
	default:
		panic("validator: invalid input")
	}

	v.errors = append(v.errors, msg)
}
