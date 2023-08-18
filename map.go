package go_basic_extension

import "errors"

var (
	ErrValueForKeyNotExists = errors.New("value for key is not exists")
)

type Map[T comparable, V any] map[T]V

func (m *Map[T, V]) KeyExists(key T) bool {
	_, exists := (*m)[key]
	return exists
}

func (m *Map[T, V]) Get(key T) (V, error) {
	val, exists := (*m)[key]
	if exists {
		return val, nil
	}

	return (*m)[key], ErrValueForKeyNotExists
}

func (m *Map[T, V]) GetOrDefault(key T, def V) V {
	if m.KeyExists(key) {
		return (*m)[key]
	}
	return def
}
