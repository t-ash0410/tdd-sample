package usecases_test

import (
	"testing"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
)

func TestList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		name := "success"
		add.Handle(name, "success")

		var result []entities.Task
		list.Handle(&result)

		exists := false
		for _, v := range result {
			if v.Name == name {
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
