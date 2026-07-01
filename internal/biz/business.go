package biz

import (
	"context"
	"log/slog"
)

// BusinessRepo is a Greater repo.
type BusinessRepo interface {
	Save(context.Context) error
}

// BusinessUsecase is a Greeter usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *slog.Logger
}

// NewBusinessUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo, logger *slog.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: logger}
}
