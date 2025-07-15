package categoryrpc

import (
	"blog-tech/common"
	categorymodel "blog-tech/internal/categories/model"
	"blog-tech/internal/categories/proto/pb"
	"context"
)

type Business interface {
	GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{
		business: business,
	}
}

func (s *grpcService) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	category, err := s.business.GetCategoryByID(ctx, int(req.CategoryId))

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return &pb.GetCategoryByIdResponse{
		Category: &pb.Category{
			Id:          int32(category.ID),
			Name:        category.Name,
			Slug:        category.Slug,
			Description: category.Description,
		},
	}, nil
}
