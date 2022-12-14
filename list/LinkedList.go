package list

type LinkedList[T comparable] struct {
	firstNode *node[T]
	lastNode  *node[T]
	length    int
}

type node[T comparable] struct {
	data T
	next *node[T]
}

func NewLinkedList[E comparable](items ...E) *LinkedList[E] {
	l := &LinkedList[E]{
		length: 0,
	}

	l.Add(items...)

	return l
}

func (ll *LinkedList[T]) Add(items ...T) {
	if len(items) == 0 {
		return
	}

	for _, item := range items {
		if ll.firstNode == nil {
			ll.firstNode = &node[T]{
				data: item,
			}
			ll.lastNode = ll.firstNode
		} else {
			// add new node to next of last node
			// update last node
			ll.lastNode.next = &node[T]{
				data: item,
			}
			ll.lastNode = ll.lastNode.next
		}
		ll.length++
	}

}

func (ll *LinkedList[T]) Remove(position int) bool {
	if position >= ll.length {
		return false
	}

	if position == 0 {
		if ll.length == 1 {
			ll.firstNode = nil
			ll.lastNode = nil
		} else {
			ll.firstNode = ll.firstNode.next
		}

		ll.length--

		return true
	}

	p := 0
	node := ll.firstNode
	for p != position-1 {
		p++
		node = node.next
	}
	node.next = node.next.next
	if position+1 == ll.length {
		ll.lastNode = node.next
	}

	ll.length--

	return true
}

func (ll *LinkedList[T]) Get(position int) (*T, bool) {
	if position >= ll.length {
		return nil, false
	}

	p := 0
	node := ll.firstNode
	for p < position {
		node = node.next
		p++
	}

	return &node.data, true
}

func (ll *LinkedList[T]) GetAll() []T {
	if ll.length == 0 {
		return make([]T, 0)
	}

	c := make([]T, ll.length)

	each := ll.firstNode
	pointer := 0
	for each != nil {
		c[pointer] = each.data
		each = each.next
		pointer++
	}

	return c
}

func (ll *LinkedList[T]) Empty() bool {
	return ll.length == 0
}

func (ll *LinkedList[T]) Size() int {
	return ll.length
}

func (ll *LinkedList[T]) Map(mapFunc func(item T) T) *LinkedList[T] {
	node := ll.firstNode

	for node != nil {
		node.data = mapFunc(node.data)
		node = node.next
	}

	return ll
}

func (ll *LinkedList[T]) Reduce(initial T, reducer func(prev, current T) T) T {
	result := initial

	node := ll.firstNode

	for node != nil {
		result = reducer(result, node.data)
		node = node.next
	}

	return result
}
