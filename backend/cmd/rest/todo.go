package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/t-ash0410/tdd-sample/backend/proto/clients/handlers"
)

const (
	rpcServerNameKey         = "RPC_SERVER_NAME"
	todoServiceListenPortKey = "TODO_SERVICE_LISTEN_PORT"
	listenPortKey            = "LISTEN_PORT"
)

func main() {
	server := os.Getenv(rpcServerNameKey)
	port := os.Getenv(todoServiceListenPortKey)
	handler := handlers.NewTodoHandler(fmt.Sprintf("%s:%s", server, port))
	http.HandleFunc("/todo/list", handler.ListHandler)
	listenPort := os.Getenv(listenPortKey)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil); err != nil {
		panic(err)
	}
}
