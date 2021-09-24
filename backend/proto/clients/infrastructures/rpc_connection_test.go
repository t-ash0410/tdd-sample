package infrastructures_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/t-ash0410/tdd-sample/backend/proto/clients/infrastructures"
)

const (
	ip   = "localhost"
	port = ":50001"
)

func TestMain(m *testing.M) {
	go setup()
	m.Run()
}

func setup() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}

func TestExecute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		conn := infrastructures.NewRpcConnection(ctx, ip+port)
		if err := conn.Execute(func(conn *grpc.ClientConn) {
			err := conn.Invoke(context.Background(), "", "", "")
			if err == nil {
				t.Error(errors.New("fail invoke request."))
			}
			st, ok := status.FromError(err)
			if !ok {
				t.Error(err)
			}
			if st.Code() != codes.Internal {
				t.Error(err)
			}
		}); err != nil {
			t.Error(err)
		}
	})
	t.Run("fail connection", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		cancel()

		conn := infrastructures.NewRpcConnection(ctx, ip+port)
		if err := conn.Execute(func(conn *grpc.ClientConn) {
			return
		}); err == nil {
			t.Error(errors.New("fail connection."))
		}
	})
}
