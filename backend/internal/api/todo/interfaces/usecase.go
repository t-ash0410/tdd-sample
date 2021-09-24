package interfaces

import "github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"

type IListUsecase interface {
	Handle(result *[]entities.Task) error
}

type IAddUsecase interface {
	Handle(name string, description string) error
}
