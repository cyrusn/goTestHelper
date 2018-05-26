// Package assert stores some helper func for testing package
package assert

import (
	"reflect"
	"testing"
)

// OK fails the test if v is not nil
func OK(t *testing.T, v interface{}) {
	if v != nil {
		t.Helper()
		t.Fatal(v)
	}
}

// Panic fails the test if the f didn't run panic()
func Panic(name string, t *testing.T, f func()) {
	defer func(t *testing.T) {
		if err := recover(); err == nil {
			t.Helper()
			t.Fatalf("err should not nil [%s]", name)
		}
	}(t)
	f()
}

// Equal fails the test if got is not equal to want
func Equal(got, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(want, got) {
		t.Helper()
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
		t.Helper()
		t.Errorf(
			"Incorrect!\ngot: %v\nwant: %v.\n",
			got,
			want,
		)
	}
}
