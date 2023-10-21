package usecases

import "sync"

type DataUpdaters []DataHandler

func (updater *DataUpdaters) ExecuteAll(buffer int) <-chan error {
	errors := make(chan error, buffer)

	go func() {
		defer close(errors)

		var err error

		var wg sync.WaitGroup

		for _, handler := range *updater {
			wg.Add(1)

			go func(handler DataHandler) {
				defer wg.Done()

				err = handler.Update()
				if err != nil {
					errors <- err
				}
			}(handler)
		}

		wg.Wait()
	}()

	return errors
}
