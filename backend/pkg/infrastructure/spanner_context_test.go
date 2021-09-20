package infrastructure_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/t-ash0410/tdd-sample/backend/pkg/infrastructure"
)

func TestConstructor(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	connectionString := os.Getenv("SPANNER_CONNECTION_STRING")
	spannerContext := infrastructure.NewSpannerContext(ctx, connectionString)
	defer spannerContext.Close()
}
