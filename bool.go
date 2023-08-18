package go_basic_extension

func Ternary[T comparable](cond func() bool, f, s T) T {
	if cond() {
		return f
	}

	return s
}
