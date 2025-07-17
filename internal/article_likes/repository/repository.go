package articlelikerepository

import (
	articlelikemodel "blog-tech/internal/article_likes/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ArticleLikeRepository interface {
	Create(ctx context.Context, data *articlelikemodel.ArticleLike) error
	Delete(ctx context.Context, articleId, userId int) error
	FindByArticleAndUser(ctx context.Context, articleId, userId int) (*articlelikemodel.ArticleLike, error)
}

type articleLikeRepository struct {
	db *gorm.DB
}

func NewArticleLikeRepository(db *gorm.DB) *articleLikeRepository {
	return &articleLikeRepository{db: db}
}

func (r *articleLikeRepository) Create(ctx context.Context, data *articlelikemodel.ArticleLike) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := r.db.Table(articlelikemodel.ArticleLike{}.TableName()).Create(data).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, articlelikemodel.ErrCannotCreateArticleLike.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *articleLikeRepository) Delete(ctx context.Context, articleId, userId int) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := r.db.Table(articlelikemodel.ArticleLike{}.TableName()).Where("article_id = ? AND user_id = ?", articleId, userId).
		Delete(&articlelikemodel.ArticleLike{}).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, articlelikemodel.ErrCannotDeleteArticleLike.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *articleLikeRepository) FindByArticleAndUser(ctx context.Context, articleId, userId int) (*articlelikemodel.ArticleLike, error) {
	var like articlelikemodel.ArticleLike

	if err := r.db.Table(articlelikemodel.ArticleLike{}.TableName()).Where("article_id = ? AND user_id = ?", articleId, userId).First(&like).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, articlelikemodel.ErrArticleLikeNotFound
		}
		return nil, errors.WithStack(err)
	}

	return &like, nil
}
