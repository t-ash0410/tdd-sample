package interfaces

import (
	"google.golang.org/grpc"
)

type IRpcConnection interface {
	Execute(f func(conn *grpc.ClientConn)) error
}
