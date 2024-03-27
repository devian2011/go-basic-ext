package go_basic_extension

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Array[T comparable] []T

func (a *Array[T]) Contains(v T) bool {
	for _, r := range *a {
		if r == v {
			return true
		}
	}

	return false
}

func (a *Array[T]) Unique() Array[T] {
	result := Array[T]{}
	keys := make(map[T]interface{}, len(*a))
	for _, v := range *a {
		if _, exists := keys[v]; !exists {
			keys[v] = true
			result = append(result, v)
		}
	}

	return result
}

func (a *Array[T]) NotContains(v T) bool {
	return !a.Contains(v)
}

func Intersection[T comparable](f, s []T) []T {
	result := make([]T, 0, len(f))
	hm := make(map[T]interface{}, len(f))
	for _, fV := range f {
		hm[fV] = nil
	}

	for _, sV := range s {
		if _, e := hm[sV]; e {
			result = append(result, sV)
		}
	}

	return result
}

func NonIntersection[T comparable](f, s []T) []T {
	result := make([]T, 0, len(f))
	hm := make(map[T]interface{}, len(f))
	for _, sV := range s {
		hm[sV] = nil
	}

	for _, fV := range f {
		if _, e := hm[fV]; !e {
			result = append(result, fV)
		}
	}

	return result
}

// HashSlice return sha256 hash of array
func HashSlice[T any](i []T) string {
	b := strings.Builder{}
	for _, v := range i {
		b.WriteString(fmt.Sprintf("%v", v))
	}
	hash := sha256.New()
	hash.Write([]byte(b.String()))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
