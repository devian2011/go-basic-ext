package go_basic_extension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr_Empty(t *testing.T) {
	s := ""
	assert.True(t, String(s).Empty())
}

func TestStr_NotEmpty(t *testing.T) {
	s := "hello"
	assert.True(t, String(s).NotEmpty())
}
