// Package assert stores some helper func for testing package
package assert

import (
	"reflect"
	"testing"
)

// OK fails the test if v is not nil
func OK(t *testing.T, v interface{}) {
	if v != nil {
		t.Fatal(v)
	}
}

// Panic fails the test if the f didn't call panic()
func Panic(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		if err := recover(); err == nil {
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

// NotEqual fails the test if got is equal to want
func NotEqual(got, want interface{}, t *testing.T) {
	if reflect.DeepEqual(want, got) {
		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}
