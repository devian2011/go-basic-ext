package go_basic_extension

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChainUnitAfter(t *testing.T) {
	unit1 := NewChain[string]("unit1")
	unit2 := unit1.After("unit2")
	unit3 := unit2.After("unit3")

	assert.Nil(t, unit1.previous)
	assert.Equal(t, unit1.next, unit2)
	assert.Equal(t, "unit1", unit1.value)
	assert.Equal(t, "unit2", unit2.value)
	assert.Equal(t, "unit3", unit3.value)
}

func Test_ChainUnitBefore(t *testing.T) {
	unit1 := NewChain[string]("unit1")
	unit2 := unit1.Before("unit2")
	unit3 := unit2.Before("unit3")

	assert.Nil(t, unit1.next)
	assert.Equal(t, unit1.previous, unit2)
	assert.Equal(t, "unit1", unit1.value)
	assert.Equal(t, "unit2", unit2.value)
	assert.Equal(t, "unit3", unit3.value)
}

func Test_ChainUnitGetRoot(t *testing.T) {
	unit1 := NewChain[string]("unit1")
	unit2 := unit1.Before("unit2")
	unit3 := unit2.Before("root")

	root := unit1.GetRoot()

	assert.Equal(t, unit3, root)
}

func Test_ChainUnitChainIterate(t *testing.T) {
	expected := []string{"one", "two", "three"}
	actual := make([]string, 0, 3)

	unit1 := NewChain[string]("one")
	unit2 := unit1.After("two")
	unit2.After("three")

	ChainIterate[string](unit1, func(v string) {
		actual = append(actual, v)
	})

	assert.Equal(t, expected, actual)
}
