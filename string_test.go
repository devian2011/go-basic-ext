package go_basic_extension

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStr_Empty(t *testing.T) {
	s := ""
	assert.True(t, String(s).Empty())
}

func TestStr_NotEmpty(t *testing.T) {
	s := "hello"
	assert.True(t, String(s).NotEmpty())
}
