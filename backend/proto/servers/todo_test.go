package servers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/t-ash0410/tdd-sample/backend/proto/generates/todo"
	"github.com/t-ash0410/tdd-sample/backend/proto/servers"
	mock "github.com/t-ash0410/tdd-sample/backend/test/mock/todo"
)

func TestList(t *testing.T) {
	server := servers.NewTodoServer(&mock.SuccessFactory{})
	t.Run("success", func(t *testing.T) {
		list, err := server.List(context.Background(), nil)
		if err != nil {
			t.Error(err)
		}
		if len(list.Tasks) == 0 {
			t.Error(errors.New("list empty."))
		}
	})
	server = servers.NewTodoServer(&mock.EmptyFactory{})
	t.Run("empty", func(t *testing.T) {
		list, err := server.List(context.Background(), nil)
		if err != nil {
			t.Error(err)
		}
		if 0 < len(list.Tasks) {
			t.Error(errors.New("list not empty."))
		}
	})
	server = servers.NewTodoServer(&mock.FailFactory{})
	t.Run("error", func(t *testing.T) {
		_, err := server.List(context.Background(), nil)
		if err == nil {
			t.Error("err is nil")
		}
	})
}

func TestAdd(t *testing.T) {
	server := servers.NewTodoServer(&mock.SuccessFactory{})
	t.Run("success", func(t *testing.T) {
		_, err := server.Add(context.Background(), &todo.AddTaskRequest{})
		if err != nil {
			t.Error(err)
		}
	})
	server = servers.NewTodoServer(&mock.FailFactory{})
	t.Run("fail", func(t *testing.T) {
		_, err := server.Add(context.Background(), &todo.AddTaskRequest{})
		if err == nil {
			t.Error("err is nil")
		}
	})
}
