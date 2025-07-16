package articletagreopomysql

import (
	articletagmodel "blog-tech/internal/article_tags/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ArticleTagRepository interface {
	CreateArticleTag(ctx context.Context, articleID int, tagIDs []int) error
}

type articleTagRepository struct {
	db *gorm.DB
}

func NewArticleTagRepository(db *gorm.DB) *articleTagRepository {
	return &articleTagRepository{db: db}
}

func (r *articleTagRepository) CreateArticleTag(ctx context.Context, articleID int, tagIDs []int) error {
	if len(tagIDs) == 0 {
		return nil
	}
	var articleTags []articletagmodel.ArticleTag
	for _, tagID := range tagIDs {
		articleTags = append(articleTags, articletagmodel.ArticleTag{
			ArticleID: articleID,
			TagID:     tagID,
		})
	}

	if err := r.db.Table(articletagmodel.ArticleTag{}.TableName()).Create(&articleTags).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
