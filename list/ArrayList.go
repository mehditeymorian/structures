package list

type ArrayList[T comparable] struct {
	data []T
}

func NewArrayList[E comparable](items ...E) *ArrayList[E] {
	a := &ArrayList[E]{
		data: make([]E, 0),
	}

	a.Add(items...)

	return a
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

func (al *ArrayList[T]) Map(mapFunc func(item T) T) *ArrayList[T] {
	for i := range al.data {
		al.data[i] = mapFunc(al.data[i])
	}

	return al
}

func (al *ArrayList[T]) Reduce(initial T, reducer func(prev, current T) T) T {
	result := initial
	for _, t := range al.data {
		result = reducer(result, t)
	}

	return result
}
