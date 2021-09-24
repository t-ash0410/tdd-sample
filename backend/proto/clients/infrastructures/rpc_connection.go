package infrastructures

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/t-ash0410/tdd-sample/backend/proto/clients/interfaces"
)

type RpcConnection struct {
	ctx     context.Context
	address string
}

func NewRpcConnection(ctx context.Context, address string) interfaces.IRpcConnection {
	return &RpcConnection{
		ctx:     ctx,
		address: address,
	}
}

func (c *RpcConnection) Execute(f func(conn *grpc.ClientConn)) error {
	conn, err := grpc.DialContext(c.ctx, c.address, grpc.WithInsecure())
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Close()
	f(conn)
	return nil
}
