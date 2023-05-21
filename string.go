package go_basic_extension

type String string

func (s String) Empty() bool {
	return len(s) == 0
}

func (s String) NotEmpty() bool {
	return !s.Empty()
}
