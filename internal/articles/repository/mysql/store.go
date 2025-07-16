package articlemysql

import (
	articlemodel "blog-tech/internal/articles/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, article *articlemodel.ArticleCreate) error
	GetArticleByID(ctx context.Context, id int) (*articlemodel.Article, error)
	GetArticles(ctx context.Context) ([]articlemodel.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db: db}
}

func (s *articleRepository) CreateArticle(ctx context.Context, article *articlemodel.ArticleCreate) error {
	if err := s.db.Table(articlemodel.Article{}.TableName()).Create(&article).Error; err != nil {
		return errors.Wrap(err, articlemodel.ErrCannotCreateArticle.Error())
	}
	return nil
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
