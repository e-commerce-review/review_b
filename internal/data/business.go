package data

import (
	"context"
	"log/slog"

	"review-b/internal/biz"
)

type businessRepo struct {
	data *Data
	log  *slog.Logger
}

// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger *slog.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  logger,
	}
}

func (r *businessRepo) Save(ctx context.Context) error {
	return nil
}
