package handlers

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"

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

func (h *HandlerBase) ExecuteRpc(res http.ResponseWriter, f func(conn *grpc.ClientConn)) {
	log.Print("Execute RPC.")

	opts, err := h.createConnectionOpts()
	if err != nil {
		h.Write500Error(res, err)
		return
	}

	conn, err := grpc.Dial(h.rpcAddress, opts...)
	if err != nil {
		h.Write500Error(res, err)
		return
	}
	
	f(conn)
}

func (h *HandlerBase) createConnectionOpts() ([]grpc.DialOption, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(h.rpcAddress))

	// opts = append(opts, grpc.WithInsecure())

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
