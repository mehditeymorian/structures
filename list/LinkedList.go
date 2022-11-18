package list

import "context"

type LinkedList[T comparable] struct {
	firstNode *node[T]
	lastNode  *node[T]
	length    uint
}

type node[T comparable] struct {
	data T
	next *node[T]
}

func NewLinkedList[E comparable]() *LinkedList[E] {
	return &LinkedList[E]{
		length: 0,
	}
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
	context.TODO()
	return true
}

func (ll *LinkedList[T]) Get(position int) (*T, bool) {
	context.TODO()
	return nil, true
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

func (ll *LinkedList[T]) Size() uint {
	return ll.length
}

func (ll *LinkedList[T]) Map(mapFunc func(item T) T) *LinkedList[T] {
	context.TODO()
	return ll
}

func (ll *LinkedList[T]) Reduce(initial T, reducer func(prev, current T) T) T {
	result := initial
	context.TODO()

	return result
}
