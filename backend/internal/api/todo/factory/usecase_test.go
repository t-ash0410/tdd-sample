//go:build integrate
// +build integrate

package factory_test

import (
	"context"
	"os"
	"testing"

	"github.com/t-ash0410/tdd-sample/backend/internal/api/todo/factory"
)

func TestCreateList(t *testing.T) {
	conn := os.Getenv("SPANNER_CONNECTION_STRING")
	f := factory.NewUsecaseFactory(conn)
	t.Run("success", func(t *testing.T) {
		_, close := f.CreateList(context.Background())
		close()
	})
}

func TestCreateAdd(t *testing.T) {
	conn := os.Getenv("SPANNER_CONNECTION_STRING")
	f := factory.NewUsecaseFactory(conn)
	t.Run("success", func(t *testing.T) {
		_, close := f.CreateAdd(context.Background())
		close()
	})
}
