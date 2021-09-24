package usecases

import (
	"time"

	"github.com/pkg/errors"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
)

type AddUsecase struct {
	repository interfaces.IRepository
}

func NewAddUsecase(repository interfaces.IRepository) interfaces.IAddUsecase {
	return &AddUsecase{
		repository: repository,
	}
}

func (usecase *AddUsecase) Handle(name string, description string) error {
	task := entities.Task{
		Id:          "test",
		Name:        name,
		Description: description,
		UpdatedAt:   time.Now(),
	}
	if err := task.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if err := usecase.repository.Add(task); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
