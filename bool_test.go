package go_basic_extension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	assert.Equal(t, 1, Ternary(func() bool {
		return true
	}, 1, 2))

	assert.Equal(t, "false", Ternary(func() bool {
		return false
	}, "true", "false"))
}
