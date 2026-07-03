package data

import (
	"context"
	"log/slog"

	v1 "review-b/api/review_api/v1"
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

func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	r.log.InfoContext(ctx, "[data] Reply", "param", param)
	ret, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.DebugContext(ctx, "ReplyReview return", "ret", ret, "err", err)
	if err != nil {
		return 0, err
	}
	return ret.GetReplyID(), nil
}

func (r *businessRepo) Appeal(ctx context.Context, param *biz.AppealParam) (int64, error) {
	r.log.InfoContext(ctx, "Appeal", "param", param)
	ret, err := r.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Reason:    param.Reason,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.DebugContext(ctx, "AppealReview return", "ret", ret, "err", err)
	if err != nil {
		return 0, err
	}
	return ret.GetAppealID(), nil
}
