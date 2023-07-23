package go_basic_extension

// LinkedList is alias of chain
type LinkedList[V any] ChainUnit[V]

// ChainUnit is a variant of linked list
type ChainUnit[V any] struct {
	next     *ChainUnit[V]
	previous *ChainUnit[V]
	value    V
}

func NewChain[V any](v V) *ChainUnit[V] {
	return &ChainUnit[V]{
		next:     nil,
		previous: nil,
		value:    v,
	}
}

func (c *ChainUnit[V]) Value() any {
	return c.value
}

func (c *ChainUnit[V]) After(v V) *ChainUnit[V] {
	chain := &ChainUnit[V]{
		next:     nil,
		previous: c,
		value:    v,
	}

	if c.next != nil {
		c.next.previous = chain
	}

	c.next = chain

	return chain
}

func (c *ChainUnit[V]) Before(v V) *ChainUnit[V] {
	chain := &ChainUnit[V]{
		next:     c,
		previous: nil,
		value:    v,
	}

	if c.previous != nil {
		c.previous.next = chain
	}

	c.previous = chain

	return chain
}

func (c *ChainUnit[V]) Next() *ChainUnit[V] {
	return c.next
}

func (c *ChainUnit[V]) GetRoot() *ChainUnit[V] {
	current := c
	for current.previous != nil {
		current = current.previous
	}

	return current
}

func ChainIterate[V any](unit *ChainUnit[V], fn func(v V)) {
	current := unit
	for current != nil {
		fn(current.value)
		current = current.Next()
	}
}
