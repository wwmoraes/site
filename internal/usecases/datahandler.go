package usecases

import "context"

type DataHandler interface {
	Update(ctx context.Context) error
}
