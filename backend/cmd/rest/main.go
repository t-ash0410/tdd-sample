package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/t-ash0410/tdd-sample/backend/proto/clients/handlers"
)

const (
	listenPortKey    = "PORT"
	rpcServerNameKey = "RPC_SERVER_NAME"
)

func main() {
	server := os.Getenv(rpcServerNameKey)
	handler := handlers.NewTodoHandler(server)
	http.HandleFunc("/todo/list", handler.ListHandler)
	listenPort := os.Getenv(listenPortKey)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
	log.Print("listen...")
}
