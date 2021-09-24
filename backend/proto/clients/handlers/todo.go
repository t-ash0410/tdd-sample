package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"unsafe"

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
	if req.Method != http.MethodGet {
		panic(errors.New("bat method."))
	}

	h.ExecuteRpc(res, func(ctx context.Context, conn grpc.ClientConnInterface) {
		client := pb.NewTodoClient(conn)

		list, err := client.List(ctx, &emptypb.Empty{})
		if err != nil {
			h.Write500Error(res, err)
			return
		}
		bytes, err := json.Marshal(list)
		if err != nil {
			h.Write500Error(res, err)
			return
		}
		json := *(*string)(unsafe.Pointer(&bytes))
		fmt.Fprintf(res, json)
	})
}
