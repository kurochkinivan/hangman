package random

import "math/rand/v2"

type RandomSelector[T any] struct{}

func New[T any]() RandomSelector[T] {
	return RandomSelector[T]{}
}

func (rs RandomSelector[T]) Choose(items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[rand.IntN(len(items))]
}
