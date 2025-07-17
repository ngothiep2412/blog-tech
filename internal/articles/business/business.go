package articlebiz

import (
	"blog-tech/common"
	articlemodel "blog-tech/internal/articles/model"
	articlemysql "blog-tech/internal/articles/repository/mysql"
	categorymodel "blog-tech/internal/categories/model"
	tagmodel "blog-tech/internal/tags/model"
	usermodel "blog-tech/internal/users/model"
	"context"
	"time"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*usermodel.User, error)
}

type TagRepository interface {
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
	GetTagByName(ctx context.Context, name string) (*tagmodel.Tag, error)
	CreateTag(ctx context.Context, tagName, tagSlug string) (*tagmodel.Tag, error)
}

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, id int) (*categorymodel.Category, error)
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

func (b *articleBusiness) CreateArticle(ctx context.Context, req *articlemodel.ArticleCreate) (*articlemodel.Article, error) {
	user, err := b.userRepo.GetUserByID(ctx, req.UserID)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}

	if user.ID == 0 {
		return nil, usermodel.ErrUserNotFound
	}

	categoryResp, err := b.categoryRepo.GetCategoryById(ctx, req.CategoryID)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}

	if categoryResp.ID == 0 {
		return nil, categorymodel.ErrCategoryNotFound
	}

	var tagIDs []int32

	for _, tag := range req.Tags {
		tagResp, err := b.tagRepo.GetTagByName(ctx, tag.Name)

		if err != nil || tagResp.ID == 0 {
			tagResp, err := b.tagRepo.CreateTag(ctx, tag.Name, tag.Slug)

			if err != nil {
				return nil, common.ErrInternalServerError.WithError(err.Error())
			}
			tagIDs = append(tagIDs, int32(tagResp.ID))
		} else {
			tagIDs = append(tagIDs, int32(tagResp.ID))
		}
	}

	articleCreation := &articlemodel.ArticleCreate{
		UserID:           req.UserID,
		CategoryID:       req.CategoryID,
		Title:            req.Title,
		Content:          req.Content,
		Excerpt:          req.Excerpt,
		FeaturedImageURL: req.FeaturedImageURL,
		Status:           req.Status,
		Slug:             common.GenerateSlug(req.Title),
	}

	if req.Status == articlemodel.StatusPublished {
		now := time.Now()
		articleCreation.PublishedAt = &now
	}

	articleResp, err := b.articleRepo.CreateArticle(ctx, articleCreation)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}

	if len(tagIDs) > 0 {
		err := b.articleTagRepo.CreateArticleTag(ctx, articleResp.ID, tagIDs)

		if err != nil {
			return nil, common.ErrInternalServerError.WithError(err.Error())
		}
	}

	return articleResp, nil
}
