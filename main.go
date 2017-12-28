// package goTestHelper stores some helper func for testing package
package goTestHelper

import (
	"reflect"
	"testing"
)

// ExpectError is a testing tool, it used to test for error handling
func ExpectError(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		err := recover()

		if err == nil {
			t.Fatalf("Error Test: [%s] did not return error", name)
		}
	}(t)
	f()
}

// Diff is simply test if there are differences of 2 structs
func Diff(got, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {

		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}
