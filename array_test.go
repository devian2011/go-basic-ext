package go_basic_extension

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestIntersection(t *testing.T) {
	first := []int{1, 2, 3, 4, 5, 6}
	second := []int{3, 6, 10, -1, 12}

	assert.Equal(t, []int{3, 6}, Intersection[int](first, second))
}

func TestNonIntersection(t *testing.T) {
	first := []int{1, 2, 3, 4, 5, 6}
	second := []int{3, 6, 10, -1, 12}

	assert.Equal(t, []int{1, 2, 4, 5}, NonIntersection[int](first, second))
	assert.Equal(t, []int{10, -1, 12}, NonIntersection[int](second, first))
}

func TestHashSlice(t *testing.T) {
	type one struct {
		id  int
		val string
	}

	i := []one{{1, "hello"}, {2, "world"}}
	expected := "b7773808530d09cac59dc5b4720b868cadb25285162096c164dfbe60820402f4"

	assert.Equal(t, expected, HashSlice(i))
}
