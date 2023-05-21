package go_basic_extension

import "errors"

var (
	ErrValueForKeyNotExists = errors.New("value for key is not exists")
)

type Map map[any]any

func (m *Map) KeyExists(key any) bool {
	_, exists := (*m)[key]
	return exists
}

func (m *Map) Get(key any) (any, error) {
	val, exists := (*m)[key]
	if exists {
		return val, nil
	}

	return nil, ErrValueForKeyNotExists
}

func (m *Map) GetOrDefault(key any, def any) any {
	if m.KeyExists(key) {
		return (*m)[key]
	}
	return def
}
