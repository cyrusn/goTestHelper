package assert_test

import (
	"testing"

	assert "github.com/cyrusn/goTestHelper"
)

func ExampleEqual(t *testing.T) {
	t.Run("assert.Equal", func(t *testing.T) {
		a := "foo"
		b := "foo"
		// Fail the test if a is not equal to b
		assert.Equal(a, b, t)
	})
}

func ExampleNotEqual(t *testing.T) {
	t.Run("assert.NotEqual", func(t *testing.T) {
		a := "foo"
		b := "bar"
		// Fail the test if a is equal to b
		assert.NotEqual(a, b, t)
	})
}

func ExampleOK(t *testing.T) {
	t.Run("assert.OK", func(t *testing.T) {
		// Fail the test if err is not nil
		assert.OK(t, nil)
	})
}

func ExamplePanic(t *testing.T) {
	t.Run("assert.Panic", func(t *testing.T) {
		assert.Panic("panic test", t, func() {
			// Fail the test no panic didn't run in here
			// panic("panic now")
		})
	})
}
