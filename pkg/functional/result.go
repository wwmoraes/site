package functional

type Result[T any] struct {
	Value T
	Error error
}

func SendValue[T any](ch chan<- *Result[T], value T) {
	ch <- ValueResult(value)
}

func SendError[T any](ch chan<- *Result[T], err error) {
	ch <- ErrorResult[T](err)
}

func ValueResult[T any](value T) *Result[T] {
	return &Result[T]{
		Value: value,
	}
}

func ErrorResult[T any](err error) *Result[T] {
	return &Result[T]{
		Error: err,
	}
}

func Splice[T any](in <-chan *Result[T], buffer int) (<-chan T, <-chan error) {
	values := make(chan T, buffer)
	errs := make(chan error, buffer)

	go func() {
		for item := range in {
			if item == nil {
				continue
			}

			if item.Error != nil {
				errs <- item.Error

				continue
			}

			values <- item.Value
		}

		close(values)
		close(errs)
	}()

	return values, errs
}

func LogErrors[T any](in <-chan *Result[T], buffer int, logger func(error)) <-chan T {
	values, errs := Splice(in, buffer)

	go func() {
		for err := range errs {
			logger(err)
		}
	}()

	return values
}
