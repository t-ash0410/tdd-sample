package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

type HandlerBase struct {
	rpcAddress string
}

func (h *HandlerBase) Write500Error(res http.ResponseWriter, err error) {
	res.WriteHeader(500)
	res.Write([]byte("internal server error"))
	log.Fatal(err)
}

func (h *HandlerBase) ExecuteRpc(res http.ResponseWriter, f func(ctx context.Context, conn grpc.ClientConnInterface)) {
	log.Print("Execute RPC.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := grpc.DialContext(ctx, h.rpcAddress, grpc.WithInsecure())
	if err != nil {
		h.Write500Error(res, err)
		return
	}
	defer conn.Close()
	f(ctx, conn)
}
