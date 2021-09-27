package servers

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
	pb "github.com/t-ash0410/tdd-sample/backend/proto/generates/todo"
)

type TodoServer struct {
	factory interfaces.IUsecaseFactory
}

func NewTodoServer(factory interfaces.IUsecaseFactory) pb.TodoServer {
	return &TodoServer{factory: factory}
}

func (s *TodoServer) List(ctx context.Context, _ *empty.Empty) (*pb.TaskList, error) {
	list, close := s.factory.CreateList(ctx)
	defer close()

	entities := []entities.Task{}
	if err := list.Handle(&entities); err != nil {
		return nil, errors.WithStack(err)
	}

	var tasks []*pb.Task
	for _, v := range entities {
		tasks = append(tasks, &pb.Task{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return &pb.TaskList{
		Tasks: tasks,
	}, nil
}

func (s *TodoServer) Add(ctx context.Context, task *pb.AddTaskRequest) (*empty.Empty, error) {
	add, close := s.factory.CreateAdd(ctx)
	defer close()

	if err := add.Handle(task.GetName(), task.GetDescription()); err != nil {
		return nil, errors.WithStack(err)
	}

	return nil, nil
}
