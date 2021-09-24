package interfaces

import "context"

type IUsecaseFactory interface {
	CreateList(ctx context.Context) (IListUsecase, func())
	CreateAdd(ctx context.Context) (IAddUsecase, func())
}
