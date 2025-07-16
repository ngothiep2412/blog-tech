package articletagbiz

import (
	"blog-tech/common"
	articletagreopomysql "blog-tech/internal/article_tags/repository/mysql"
	"context"
)

type articleTagBusiness struct {
	repository articletagreopomysql.ArticleTagRepository
}

func NewArticleTagBusiness(repository articletagreopomysql.ArticleTagRepository) *articleTagBusiness {
	return &articleTagBusiness{
		repository: repository,
	}
}

func (b *articleTagBusiness) CreateArticleTag(ctx context.Context, articleID int, tagIDs []int) error {
	err := b.repository.CreateArticleTag(ctx, articleID, tagIDs)

	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	return nil
}
