package handlers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts, err := h.createConnectionOpts()
	if err != nil {
		h.Write500Error(res, err)
		return
	}

	conn, err := grpc.DialContext(ctx, h.rpcAddress, opts...)
	if err != nil {
		h.Write500Error(res, err)
		return
	}
	defer conn.Close()
	f(ctx, conn)
}

func (h *HandlerBase) createConnectionOpts() ([]grpc.DialOption, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(h.rpcAddress))
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))
	return opts, nil
}
