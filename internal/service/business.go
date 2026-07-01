package service

import (
	"context"
	pb "review-b/api/business"
	"review-b/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	return &pb.ReplyReviewReply{}, nil
}
