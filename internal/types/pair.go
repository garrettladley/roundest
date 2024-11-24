package types

type Pair[T any] struct {
	A T `json:"a"`
	B T `json:"b"`
}

func Swap[T any](p Pair[T]) Pair[T] {
	return Pair[T]{A: p.B, B: p.A}
}
