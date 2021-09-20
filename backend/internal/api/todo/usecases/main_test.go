package usecases_test

import (
	"testing"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/usecases"
	"github.com/t-ash0410/tdd-sample/backend/test/mock"
	"github.com/t-ash0410/tdd-sample/backend/test/mock/todo"
)

var list usecases.ListUsecase
var add usecases.AddUsecase

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	memoryContext := &mock.InMemoryContext{}
	repository := todo.NewRepository(memoryContext)
	add = usecases.NewAddUsecase(repository)
	list = usecases.NewListUsecase(repository)
}

func useErrorRepository(test func()) {
	errorRepository := todo.NewErrorRepository()
	add = usecases.NewAddUsecase(errorRepository)
	list = usecases.NewListUsecase(errorRepository)
	test()
	memoryContext := &mock.InMemoryContext{}
	repository := todo.NewRepository(memoryContext)
	add = usecases.NewAddUsecase(repository)
	list = usecases.NewListUsecase(repository)
}
