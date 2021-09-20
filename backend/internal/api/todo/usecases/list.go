package usecases

import (
	"github.com/pkg/errors"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
)

type ListUsecase struct {
	repository interfaces.IRepository
}

func NewListUsecase(repository interfaces.IRepository) ListUsecase {
	return ListUsecase{
		repository,
	}
}

func (usecase ListUsecase) Handle(result *[]entities.Task) error {
	if err := usecase.repository.List(result); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
