package list

type ArrayList[T comparable] struct {
	data []T
}

func NewArrayList[E comparable]() *ArrayList[E] {
	return &ArrayList[E]{
		data: make([]E, 0),
	}
}

func (al *ArrayList[T]) Add(items ...T) {
	al.data = append(al.data, items...)
}

func (al *ArrayList[T]) Remove(position int) bool {
	if position >= len(al.data) {
		return false
	}

	al.data = append(al.data[:position], al.data[position+1:]...)

	return true
}

func (al *ArrayList[T]) Get(position int) (*T, bool) {
	if position >= len(al.data) {
		return nil, false
	}

	return &al.data[position], true
}

func (al *ArrayList[T]) GetAll() []T {
	if len(al.data) == 0 {
		return make([]T, 0)
	}

	c := make([]T, len(al.data))
	copy(c, al.data)
	return c
}

func (al *ArrayList[T]) Empty() bool {
	return len(al.data) == 0
}

func (al *ArrayList[T]) Size() int {
	return len(al.data)
}
