package usecases_test

import (
	"testing"
	"time"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
)

func TestList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		task := entities.Task{
			Id:          "TestList_success",
			Name:        "success",
			Description: "success",
			UpdatedAt:   time.Now(),
		}
		add.Handle(task)

		var result []entities.Task
		list.Handle(&result)

		exists := false
		for _, v := range result {
			if v.Id == task.Id {
				exists = true
				break
			}
		}

		if !exists {
			t.Error("object not found")
		}
	})
	t.Run("internal repository error", func(t *testing.T) {
		var result []entities.Task
		useErrorRepository(func() {
			if err := list.Handle(&result); err == nil {
				t.Error("not return error object")
			}
		})
	})
}
