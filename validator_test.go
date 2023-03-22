package validator_test

import (
	"testing"

	. "github.com/moonrhythm/validator"
)

func TestValidator(t *testing.T) {
	v := New()

	if v.String() != "no error" {
		t.Error("invalid error message")
	}

	v.Must(true, "error1")
	if !v.Valid() {
		t.Error("must be valid")
	}
	v.Must(false, "error2")
	if v.Valid() {
		t.Error("must be invalid")
	}
	v.Must(false, "error3")
	v.Must(true, "error4")
	if v.Valid() {
		t.Error("must be invalid")
	}

	{
		nv := New()
		nv.Must(false, "error10")
		v.Add(nv.Error())
	}

	if v.String() != "error2, error3, error10" {
		t.Error("invalid error message")
	}
}
