package usecases

import (
	"github.com/pkg/errors"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
)

type AddUsecase struct {
	repository interfaces.IRepository
}

func NewAddUsecase(repository interfaces.IRepository) AddUsecase {
	return AddUsecase{
		repository: repository,
	}
}

func (usecase AddUsecase) Handle(task entities.Task) error {
	if err := task.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if err := usecase.repository.Add(task); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
