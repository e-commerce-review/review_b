package data

import (
	"context"
	"log/slog"
	v1 "review-b/api/review_api/v1"
	"review-b/internal/conf"

	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/middleware/validate"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewReviewServiceClient, NewData, NewBusinessRepo)

// Data .
type Data struct {
	rc  v1.ReviewClient
	log *slog.Logger
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger *slog.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}
	return &Data{rc: rc, log: logger}, cleanup, nil
}

func NewReviewServiceClient() v1.ReviewClient {
	conn, err := grpc.NewClient(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}

	return v1.NewReviewClient(conn)
}
