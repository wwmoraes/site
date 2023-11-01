package usecases

import (
	"context"
	"sync"
)

type DataUpdaters []DataHandler

func (updater DataUpdaters) ExecuteAll(ctx context.Context, buffer int) <-chan error {
	errors := make(chan error, buffer)

	go func() {
		defer close(errors)

		var wg sync.WaitGroup

		for _, handler := range updater {
			wg.Add(1)

			go func(ctx context.Context, handler DataHandler) {
				defer wg.Done()

				err := handler.Update(ctx)
				if err != nil {
					errors <- err
				}
			}(ctx, handler)
		}

		wg.Wait()
	}()

	return errors
}
