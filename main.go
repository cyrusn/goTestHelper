// Package goTestHelper stores some helper func for testing package
package goTestHelper

import (
	"reflect"
	"testing"
)

// OK fails the test if err is not nil
func OK(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// ExpectError fails the test if the f didn't fail
func ExpectError(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		err := recover()

		if err == nil {
			t.Fatalf("Error Test: [%s] did not return error", name)
		}
	}(t)
	f()
}

// Equal fails the test if got is not equal to want
func Equal(got, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {

		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}
