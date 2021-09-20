package usecases_test

import (
	"testing"
	"time"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
)

func TestAdd(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		task := entities.Task{
			Id:          "id-test",
			Name:        "テスト",
			Description: "テスト用のオブジェクトです",
			UpdatedAt:   time.Now(),
		}
		if err := add.Handle(task); err != nil {
			t.Error(err)
		}
	})
	t.Run("validation fail", func(t *testing.T) {
		task := entities.Task{
			Id:          "",
			Name:        "",
			Description: "",
			UpdatedAt:   time.Time{},
		}
		if err := add.Handle(task); err == nil {
			t.Errorf("invalid task %+v", err)
		}
	})
	t.Run("internal repository error", func(t *testing.T) {
		task := entities.Task{
			Id:          "error",
			Name:        "error name",
			Description: "error desc",
			UpdatedAt:   time.Now(),
		}
		useErrorRepository(func() {
			if err := add.Handle(task); err == nil {
				t.Error("not return error object")
			}
		})
	})
}
