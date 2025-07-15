package tagstorerpc

import (
	"blog-tech/common"
	usermodel "blog-tech/internal/users/model"
	userpb "blog-tech/internal/users/proto/pb"
	"context"
)

type rpcClient struct {
	client userpb.UserServiceClient
}

func NewClient(client userpb.UserServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) GetUserByID(ctx context.Context, id int) (*usermodel.User, error) {
	resp, err := c.client.GetUserById(ctx, &userpb.GetUserByIdRequest{UserId: int32(id)})

	if err != nil {
		return nil, err
	}

	return &usermodel.User{
		SqlModel: common.SqlModel{
			ID: int(resp.User.Id),
		},
		Username:  resp.User.Username,
		FullName:  resp.User.FullName,
		AvatarURL: resp.User.AvatarUrl,
		IsActive:  resp.User.IsActive,
	}, nil
}
