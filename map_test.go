package go_basic_extension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_GetOrDefault(t *testing.T) {
	m := Map{"one": "oneV", "two": "twoV"}

	oneVal := m.GetOrDefault("one", "defVal")
	defVal := m.GetOrDefault("defCall", "defVal")

	assert.Equal(t, "oneV", oneVal)
	assert.Equal(t, "defVal", defVal)
}

func TestMap_Get(t *testing.T) {
	m := Map{"one": "oneV", "two": "twoV"}

	oneVal, oneVErr := m.Get("one")
	twoVal, twoVErr := m.Get("two")
	emV, emVErr := m.Get("three")

	assert.Equal(t, "oneV", oneVal)
	assert.Nil(t, oneVErr)

	assert.Equal(t, "twoV", twoVal)
	assert.Nil(t, twoVErr)

	assert.Equal(t, ErrValueForKeyNotExists, emVErr)
	assert.Nil(t, emV)
}

func TestMap_KeyExists(t *testing.T) {
	m := Map{"one": "oneV", "two": "twoV"}

	assert.True(t, m.KeyExists("one"))
	assert.False(t, m.KeyExists("three"))
}
