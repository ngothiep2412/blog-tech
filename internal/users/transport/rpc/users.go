package userrpc

import (
	"blog-tech/common"
	usermodel "blog-tech/internal/users/model"
	"blog-tech/internal/users/proto/pb"
	"context"
)

type Business interface {
	GetProfile(ctx context.Context, userID int) (*usermodel.User, error)
}

type grpcService struct {
	business Business
}

func NewService(business Business) *grpcService {
	return &grpcService{
		business: business,
	}
}

func (s *grpcService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := s.business.GetProfile(ctx, int(req.UserId))

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}

	return &pb.GetUserByIdResponse{
		User: &pb.UserBasicInfo{
			Id:        int32(user.ID),
			Username:  user.Username,
			FullName:  user.FullName,
			AvatarUrl: user.AvatarURL,
			IsActive:  user.IsActive,
		},
	}, nil
}
