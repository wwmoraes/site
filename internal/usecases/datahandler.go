package usecases

import "context"

type DataHandler interface {
	Update(context.Context) error
}
