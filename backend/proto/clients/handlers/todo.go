package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	pb "github.com/t-ash0410/tdd-sample/backend/proto/generates/todo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoHandler struct {
	HandlerBase
}

func NewTodoHandler(rpcAddress string) TodoHandler {
	return TodoHandler{
		HandlerBase: HandlerBase{
			rpcAddress: rpcAddress,
		},
	}
}

func (h *TodoHandler) ListHandler(res http.ResponseWriter, req *http.Request) {
	log.Print("Start function `List`.")

	if req.Method != http.MethodGet {
		panic(errors.New("bat method."))
	}

	h.ExecuteRpc(res, func(ctx context.Context, conn grpc.ClientConnInterface) {
		client := pb.NewTodoClient(conn)

		log.Print("Request to RPC server.")
		list, err := client.List(ctx, &emptypb.Empty{})
		if err != nil {
			h.Write500Error(res, err)
			return
		}

		log.Printf("Marshalize json %+v.", list)
		bytes, err := json.Marshal(list)
		if err != nil {
			h.Write500Error(res, err)
			return
		}

		log.Printf("Return response json %s.", json)
		res.WriteHeader(200)
		res.Write(bytes)
	})
}
