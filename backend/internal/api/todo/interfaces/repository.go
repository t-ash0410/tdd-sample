package interfaces

import (
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
)

type IRepository interface {
	List(*[]entities.Task) error
	Add(entities.Task) error
}
