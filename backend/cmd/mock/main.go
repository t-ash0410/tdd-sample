package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/entities"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/interfaces"
	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/usecases"
	pb "github.com/t-ash0410/tdd-sample/backend/proto/generates/todo"
	"github.com/t-ash0410/tdd-sample/backend/proto/servers"
	"google.golang.org/grpc"
)

type Repository struct {
}

func (repo *Repository) List(result *[]entities.Task) error {
	*result = append(*result, entities.Task{
		Id:          "id",
		Name:        "name",
		Description: "description",
		UpdatedAt:   time.Now(),
	})
	return nil
}

func (repo *Repository) Add(task entities.Task) error {
	return nil
}

type UsecaseFactory struct {
}

func (f *UsecaseFactory) CreateList(ctx context.Context) (interfaces.IListUsecase, func()) {
	return usecases.NewListUsecase(&Repository{}), func() {}
}

func (f *UsecaseFactory) CreateAdd(ctx context.Context) (interfaces.IAddUsecase, func()) {
	return usecases.NewAddUsecase(&Repository{}), func() {}
}

const (
	portKey = "PORT"
)

func main() {
	port := os.Getenv(portKey)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	factory := &UsecaseFactory{}
	server := servers.NewTodoServer(factory)

	s := grpc.NewServer()
	pb.RegisterTodoServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print("listen...")
}
