package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
)

func TestTaskEqual(t *testing.T) {
	now := time.Now()
	src := entities.Task{
		Id:          "test1",
		Name:        "テスト",
		Description: "テストの説明",
		UpdatedAt:   now,
	}
	dst := entities.Task{
		Id:          "test1",
		Name:        "テスト",
		Description: "テストの説明",
		UpdatedAt:   now,
	}

	t.Run("equals", func(t *testing.T) {
		assert.Equal(t, src.Id, dst.Id)
		assert.Equal(t, src.Name, dst.Name)
		assert.Equal(t, src.Description, dst.Description)
		assert.Equal(t, src.UpdatedAt, dst.UpdatedAt)
		if !src.Equal(dst) {
			t.Error("同値評価失敗")
		}
	})
	t.Run("not equals", func(t *testing.T) {
		dst.Id = ""
		if src.Equal(dst) {
			t.Error("異値評価失敗 Id")
		}
		dst.Id = src.Id
		dst.Name = ""
		if src.Equal(dst) {
			t.Error("異値評価失敗 Name")
		}
		dst.Name = src.Name
		dst.Description = ""
		if src.Equal(dst) {
			t.Error("異値評価失敗 Description")
		}
		dst.Description = src.Description
		dst.UpdatedAt = now.Add(1 * time.Second)
		if src.Equal(dst) {
			t.Error("異値評価失敗 UpdatedAt")
		}
	})
}

func TestTaskValidate(t *testing.T) {
	task := entities.Task{
		Id:          "id",
		Name:        "name",
		Description: "description",
		UpdatedAt:   time.Now(),
	}

	t.Run("valid", func(t *testing.T) {
		if err := task.Validate(); err != nil {
			t.Errorf("正常値検証エラー task: %+v, error: %s", task, err)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		invalidTask := task
		invalidTask.Id = ""
		if err := invalidTask.Validate(); err == nil {
			t.Errorf("ID長検証エラー task: %+v", task)
		}
		invalidTask.Id = task.Id
		invalidTask.Name = ""
		if err := invalidTask.Validate(); err == nil {
			t.Errorf("Name長(min)検証エラー task: %+v", task)
		}
		for i := 0; i < 51; i++ {
			invalidTask.Name += "t"
		}
		if err := invalidTask.Validate(); err == nil {
			t.Errorf("Name長(max)検証エラー task: %+v", task)
		}
		invalidTask.Name = task.Name
		invalidTask.Description = ""
		for i := 0; i < 1001; i++ {
			invalidTask.Description += "t"
		}
		if err := invalidTask.Validate(); err == nil {
			t.Errorf("Description長(max)検証エラー task: %+v", task)
		}
		invalidTask.Description = task.Description
		invalidTask.UpdatedAt = time.Time{}
		if err := invalidTask.Validate(); err == nil {
			t.Errorf("UpdatedAt(未登録)検証エラー task: %+v", task)
		}
		invalidTask.UpdatedAt = time.Now()
		invalidTask.UpdatedAt = invalidTask.UpdatedAt.AddDate(0, 0, 1)
		if err := invalidTask.Validate(); err == nil {
			t.Errorf("UpdatedAt(現在以降)検証エラー task: %+v", task)
		}
	})
}
