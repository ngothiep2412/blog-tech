package articlereporpc

import (
	"blog-tech/common"
	articletagpb "blog-tech/internal/article_tags/proto/pb"
	categorymodel "blog-tech/internal/categories/model"
	categorypb "blog-tech/internal/categories/proto/pb"
	tagmodel "blog-tech/internal/tags/model"
	tagpb "blog-tech/internal/tags/proto/pb"
	usermodel "blog-tech/internal/users/model"
	userpb "blog-tech/internal/users/proto/pb"
	"context"
)

type newClient struct {
	userClient       userpb.UserServiceClient
	categoryClient   categorypb.CategoryServiceClient
	tagClient        tagpb.TagServiceClient
	articleTagClient articletagpb.ArticleTagServiceClient
}

func NewClient(
	userClient userpb.UserServiceClient,
	categoryClient categorypb.CategoryServiceClient,
	tagClient tagpb.TagServiceClient,
	articleTagClient articletagpb.ArticleTagServiceClient,
) *newClient {
	return &newClient{
		userClient:       userClient,
		categoryClient:   categoryClient,
		tagClient:        tagClient,
		articleTagClient: articleTagClient,
	}
}

func (c *newClient) GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error) {
	resp, err := c.categoryClient.GetCategoryById(ctx, &categorypb.GetCategoryByIdRequest{CategoryId: int32(id)})

	if err != nil {
		return nil, err
	}

	return &categorymodel.Category{
		SqlModel: common.SqlModel{
			ID: int(resp.Category.Id),
		},
		Name:        resp.Category.Name,
		Description: resp.Category.Description,
	}, nil
}

func (c *newClient) GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error) {
	resp, err := c.tagClient.GetTagById(ctx, &tagpb.GetTagByIdRequest{TagId: int32(id)})

	if err != nil {
		return nil, err
	}

	return &tagmodel.Tag{
		SqlModel: common.SqlModel{
			ID: int(resp.Tag.Id),
		},
		Name: resp.Tag.Name,
	}, nil
}

func (c *newClient) GetUserByID(ctx context.Context, id int) (*usermodel.User, error) {
	resp, err := c.userClient.GetUserById(ctx, &userpb.GetUserByIdRequest{UserId: int32(id)})

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

func (c *newClient) CreateArticleTag(ctx context.Context, articleId int, tagsId []int32) error {
	_, err := c.articleTagClient.CreateArticleTags(ctx, &articletagpb.CreateArticleTagsRequest{
		ArticleId: int32(articleId),
		TagIds:    tagsId,
	})

	if err != nil {
		return err
	}

	return nil
}
