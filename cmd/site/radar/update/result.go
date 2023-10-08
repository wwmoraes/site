package update

type Result[T any] struct {
	Error error
	Value *T
}

func okResult[T any](value *T) *Result[T] {
	return &Result[T]{
		Value: value,
	}
}

func errorResult[T any](err error) *Result[T] {
	return &Result[T]{
		Error: err,
	}
}
