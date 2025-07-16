package tagrpc

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	"blog-tech/internal/tags/proto/pb"
	tagpb "blog-tech/internal/tags/proto/pb"
	"context"
)

type TagBusiness interface {
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
}

type grpcService struct {
	business TagBusiness
}

func NewTagService(business TagBusiness) *grpcService {
	return &grpcService{
		business: business,
	}
}

func (s *grpcService) GetTagById(ctx context.Context, req *tagpb.GetTagByIdRequest) (*pb.GetTagByIdResponse, error) {
	tag, err := s.business.GetTagByID(ctx, int(req.TagId))

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return &pb.GetTagByIdResponse{
		Tag: &pb.Tag{
			Id:   int32(tag.ID),
			Name: tag.Name,
			Slug: tag.Slug,
		},
	}, nil
}
