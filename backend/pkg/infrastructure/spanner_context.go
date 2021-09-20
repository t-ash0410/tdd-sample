package infrastructure

import (
	"context"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
)

type SpannerContext struct {
	BaseContext *context.Context
	AdminClient *database.DatabaseAdminClient
	DataClient  *spanner.Client
}

func NewSpannerContext(ctx context.Context, connectionString string) *SpannerContext {
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		panic(err)
	}
	dataClient, err := spanner.NewClient(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	return &SpannerContext{
		BaseContext: &ctx,
		AdminClient: adminClient,
		DataClient:  dataClient,
	}
}

func (ctx *SpannerContext) Close() {
	ctx.AdminClient.Close()
	ctx.DataClient.Close()
}
