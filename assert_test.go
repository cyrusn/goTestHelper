package assert_test

import (
	"testing"

	"github.com/cyrusn/goTestHelper"
)

func TestMain(t *testing.T) {
	t.Run("Test .Equal", func(t *testing.T) {
		a := "foo"
		b := "foo"
		assert.Equal(a, b, t)
	})

	t.Run("Test .NotEqual", func(t *testing.T) {
		a := "foo"
		b := "bar"
		assert.NotEqual(a, b, t)
	})

	t.Run("Test .OK", func(t *testing.T) {
		assert.OK(t, nil)
	})

	t.Run("Test .Panic", func(t *testing.T) {
		assert.Panic("panic test", t, func() {
			panic("panic now")
		})
	})

}
