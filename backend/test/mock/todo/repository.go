package todo

import (
	"errors"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
	"github.com/t-ash0410/tdd-sample/backend/test/mock"
)

type Repository struct {
	ctx *mock.InMemoryContext
}

func NewRepository(ctx *mock.InMemoryContext) interfaces.IRepository {
	return &Repository{
		ctx,
	}
}

func (repo *Repository) List(dst *[]entities.Task) error {
	for _, v := range repo.ctx.Data {
		entity := v.(entities.Task)
		*dst = append(*dst, entity)
	}
	return nil
}

func (repo *Repository) Add(task entities.Task) error {
	repo.ctx.Data = append(repo.ctx.Data, task)
	return nil
}

type ErrorRepository struct {
}

func NewErrorRepository() interfaces.IRepository {
	return &ErrorRepository{}
}

func (repo *ErrorRepository) List(dst *[]entities.Task) error {
	return errors.New("some error")
}

func (repo *ErrorRepository) Add(task entities.Task) error {
	return errors.New("some error")
}
