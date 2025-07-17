package articlerepomysql

import (
	"blog-tech/common"
	articlemodel "blog-tech/internal/articles/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, articleReq *articlemodel.ArticleCreate) (*articlemodel.Article, error)
	GetArticleByID(ctx context.Context, id int) (*articlemodel.Article, error)
	GetArticles(ctx context.Context) ([]articlemodel.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db: db}
}

func (s *articleRepository) CreateArticle(ctx context.Context, articleReq *articlemodel.ArticleCreate) (*articlemodel.Article, error) {
	tx := s.db.Begin()
	tags := articlemodel.Article{
		SqlModel:         common.NewSqlModel(),
		UserID:           articleReq.UserID,
		CategoryID:       articleReq.CategoryID,
		Title:            articleReq.Title,
		Content:          articleReq.Content,
		Slug:             articleReq.Slug,
		Excerpt:          articleReq.Excerpt,
		FeaturedImageURL: articleReq.FeaturedImageURL,
		Status:           articleReq.Status,
	}

	if err := tx.Table(articlemodel.Article{}.TableName()).Create(&tags).Error; err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, articlemodel.ErrCannotCreateArticle.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &tags, nil
}

func (s *articleRepository) GetArticleByID(ctx context.Context, id int) (*articlemodel.Article, error) {
	var article articlemodel.Article

	if err := s.db.Table(articlemodel.Article{}.TableName()).First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, articlemodel.ErrArticleNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &article, nil

}

func (s *articleRepository) GetArticles(ctx context.Context) ([]articlemodel.Article, error) {
	var articles []articlemodel.Article

	if err := s.db.Table(articlemodel.Article{}.TableName()).Find(&articles).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return articles, nil
}

func (r *articleRepository) IncreaseLikeCount(ctx context.Context, id int) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Table(articlemodel.Article{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, articlemodel.ErrCannotUpdateArticle.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *articleRepository) DecreaseLikeCount(ctx context.Context, id int) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Table(articlemodel.Article{}.TableName()).Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - 1")).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, articlemodel.ErrCannotUpdateArticle.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
