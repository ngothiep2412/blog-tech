package articlebiz

import (
	"blog-tech/common"
	articlemodel "blog-tech/internal/articles/model"
	articlemysql "blog-tech/internal/articles/repository/mysql"
	categorymodel "blog-tech/internal/categories/model"
	tagmodel "blog-tech/internal/tags/model"
	usermodel "blog-tech/internal/users/model"
	"context"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (usermodel.User, error)
}

type TagRepository interface {
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
}

type CategoryRepository interface {
	GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error)
}

type ArtcileTagRepository interface {
	CreateArticleTag(ctx context.Context, articleId int, tagsId []int32) error
}

type articleBusiness struct {
	articleRepo    articlemysql.ArticleRepository
	userRepo       UserRepository
	articleTagRepo ArtcileTagRepository
	tagRepo        TagRepository
	categoryRepo   CategoryRepository
}

func NewArticleBusiness(
	articleRepo articlemysql.ArticleRepository,
	userRepo UserRepository,
	articleTagRepo ArtcileTagRepository,
	tagRepo TagRepository,
	categoryRepo CategoryRepository,
) *articleBusiness {
	return &articleBusiness{
		articleRepo:    articleRepo,
		userRepo:       userRepo,
		tagRepo:        tagRepo,
		categoryRepo:   categoryRepo,
		articleTagRepo: articleTagRepo,
	}
}

func (b *articleBusiness) CreateArticle(ctx context.Context, req *articlemodel.ArticleCreate) error {
	user, err := b.userRepo.GetUserByID(ctx, req.UserID)

	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	if user.ID == 0 {
		return usermodel.ErrUserNotFound
	}

	categoryResp, err := b.categoryRepo.GetCategoryByID(ctx, req.CategoryID)

	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	if categoryResp.ID == 0 {
		return categorymodel.ErrCategoryNotFound
	}

	var tagsID []int32

	for _, tagID := range req.Tags {
		tagResp, err := b.tagRepo.GetTagByID()

	}

	return nil
}
