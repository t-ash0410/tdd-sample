package usecases_test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if err := add.Handle("テスト", "テスト用のオブジェクトです"); err != nil {
			t.Error(err)
		}
	})
	t.Run("validation fail", func(t *testing.T) {
		if err := add.Handle("", ""); err == nil {
			t.Errorf("invalid task %+v", err)
		}
	})
	t.Run("internal repository error", func(t *testing.T) {
		useErrorRepository(func() {
			if err := add.Handle("error", "error"); err == nil {
				t.Error("not return error object")
			}
		})
	})
}
