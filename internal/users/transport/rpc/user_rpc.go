package userrpc

import (
	"blog-tech/common"

	usermodel "blog-tech/internal/users/model"
	"blog-tech/internal/users/proto/pb"

	"context"
)

type UserBusiness interface {
	GetProfile(ctx context.Context, userID int) (*usermodel.User, error)
}

type grpcUserService struct {
	business UserBusiness
}

func NewUserService(business UserBusiness) *grpcUserService {
	return &grpcUserService{
		business: business,
	}
}

func (s *grpcUserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
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
		Message: "Get user successfully",
	}, nil
}
