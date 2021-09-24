package factory

import (
	"context"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/repositories"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/usecases"
	"github.com/t-ash0410/tdd-sample/backend/pkg/infrastructure"
)

type UsecaseFactory struct {
	connectionString string
}

func NewUsecaseFactory(connectionString string) interfaces.IUsecaseFactory {
	return &UsecaseFactory{
		connectionString: connectionString,
	}
}

func (f *UsecaseFactory) CreateList(ctx context.Context) (interfaces.IListUsecase, func()) {
	repository, close := f.createCtx(ctx)
	return usecases.NewListUsecase(repository), close
}

func (f *UsecaseFactory) CreateAdd(ctx context.Context) (interfaces.IAddUsecase, func()) {
	repository, close := f.createCtx(ctx)
	return usecases.NewAddUsecase(repository), close
}

func (f *UsecaseFactory) createCtx(ctx context.Context) (interfaces.IRepository, func()) {
	spannerCtx := infrastructure.NewSpannerContext(ctx, f.connectionString)
	repository := repositories.NewRepository(spannerCtx)
	return repository, spannerCtx.Close
}
