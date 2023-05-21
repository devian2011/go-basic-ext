package go_basic_extension

type Array []any

func (a *Array) Contains(v any) bool {
	for _, r := range *a {
		if r == v {
			return true
		}
	}

	return false
}

func (a *Array) NotContains(v any) bool {
	return !a.Contains(v)
}
