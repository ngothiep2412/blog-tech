package tagrpc

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	"blog-tech/internal/tags/proto/pb"
	tagpb "blog-tech/internal/tags/proto/pb"
	"context"
)

type TagBusiness interface {
	CreateTagRPC(ctx context.Context, tag *tagmodel.TagCreate) (*tagmodel.Tag, error)
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
	GetTagByName(ctx context.Context, name string) (*tagmodel.Tag, error)
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
	return &tagpb.GetTagByIdResponse{
		Tag: &pb.Tag{
			Id:   int32(tag.ID),
			Name: tag.Name,
			Slug: tag.Slug,
		},
	}, nil
}

func (s *grpcService) GetTagByName(ctx context.Context, req *tagpb.GetTagByNameRequest) (*tagpb.GetTagByNameResponse, error) {
	tag, err := s.business.GetTagByName(ctx, req.TagName)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return &tagpb.GetTagByNameResponse{
		Tag: &pb.Tag{
			Id:   int32(tag.ID),
			Name: tag.Name,
			Slug: tag.Slug,
		},
	}, nil
}

func (s *grpcService) CreateTag(ctx context.Context, req *tagpb.CreateTagRequest) (*tagpb.CreateTagResponse, error) {
	tag, err := s.business.CreateTagRPC(ctx, &tagmodel.TagCreate{
		Name: req.TagName,
		Slug: req.TagSlug,
	})

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return &tagpb.CreateTagResponse{
		Tag: &pb.Tag{
			Id:   int32(tag.ID),
			Name: tag.Name,
			Slug: tag.Slug,
		},
	}, nil
}
