package lazy

import "sync"

type Lazy[T any] struct {
	New   func() T
	once  sync.Once
	value T
}

func New[T any](f func() T) Lazy[T] {
	return Lazy[T]{
		New: f,
	}
}

func (l *Lazy[T]) Get() T {
	if l.New != nil {
		l.once.Do(func() {
			l.value = l.New()
			l.New = nil
		})
	}

	return l.value
}
