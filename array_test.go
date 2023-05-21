package go_basic_extension

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_Contains(t *testing.T) {
	a := Array{"one", "two"}
	assert.True(t, a.Contains("two"))
	assert.False(t, a.Contains("three"))
}

func TestArray_NotContains(t *testing.T) {
	a := Array{"one", "two"}
	assert.False(t, a.NotContains("two"))
	assert.True(t, a.NotContains("three"))
}
