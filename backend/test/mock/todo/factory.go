package todo

import (
	"context"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
)

type SuccessFactory struct{}

func (f *SuccessFactory) CreateList(ctx context.Context) (interfaces.IListUsecase, func()) {
	return SuccessListUsecase{}, func() {}
}

func (f *SuccessFactory) CreateAdd(ctx context.Context) (interfaces.IAddUsecase, func()) {
	return SuccessAddUsecase{}, func() {}
}

type EmptyFactory struct{}

func (f *EmptyFactory) CreateList(ctx context.Context) (interfaces.IListUsecase, func()) {
	return EmptyListUsecase{}, func() {}
}

func (f *EmptyFactory) CreateAdd(ctx context.Context) (interfaces.IAddUsecase, func()) {
	return SuccessAddUsecase{}, func() {}
}

type FailFactory struct{}

func (f *FailFactory) CreateList(ctx context.Context) (interfaces.IListUsecase, func()) {
	return FailListUsecase{}, func() {}
}

func (f *FailFactory) CreateAdd(ctx context.Context) (interfaces.IAddUsecase, func()) {
	return FailAddUsecase{}, func() {}
}
