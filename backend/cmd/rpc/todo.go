package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/factory"
	pb "github.com/t-ash0410/tdd-sample/backend/proto/generates/todo"
	"github.com/t-ash0410/tdd-sample/backend/proto/servers"
	"google.golang.org/grpc"
)

const (
	portKey             = "TODO_SERVICE_LISTEN_PORT"
	connectionStringKey = "DB_CONNECTION_STRING"
)

func main() {
	port := os.Getenv(portKey)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn := os.Getenv(connectionStringKey)
	factory := factory.NewUsecaseFactory(conn)
	server := servers.NewTodoServer(factory)

	s := grpc.NewServer()
	pb.RegisterTodoServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
