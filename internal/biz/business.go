package biz

import (
	"context"
	"log/slog"
)

type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// BusinessRepo is a Greater repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
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

func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	uc.log.InfoContext(ctx, "CreateReply", "param", param)
	return uc.repo.Reply(ctx, param)
}
