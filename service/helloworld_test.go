//go:build ignoretest

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {

	actual := add(10, 10)
	expected := 20

	assert.Equal(t, actual, expected, "The two results should be the same.")
}

func FuzzAdd(f *testing.F) {
	f.Add(1, 2)
	f.Fuzz(func(t *testing.T, a int, b int) {
		actual := add(a, b) // seed corpus
		expected := a + b

		if expected != actual {
			t.Errorf("add(%d, %d) - Expected: %d Actual: %d", a, b, actual, expected)
		}
	})
}
