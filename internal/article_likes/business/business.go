package articlelikebusiness

import (
	"blog-tech/common"
	articlelikedto "blog-tech/internal/article_likes/dto"
	articlelikemodel "blog-tech/internal/article_likes/model"
	articlelikerepository "blog-tech/internal/article_likes/repository"
	sctx "blog-tech/plugin"
	"blog-tech/plugin/kafka"
	"context"
	"fmt"
)

type LikeArticleBusiness struct {
	articleLikeRepo articlelikerepository.ArticleLikeRepository
	producer        *kafka.Producer
}

func NewLikeArticleBusiness(
	articleLikeRepo articlelikerepository.ArticleLikeRepository,
	producer *kafka.Producer,
) *LikeArticleBusiness {
	return &LikeArticleBusiness{
		articleLikeRepo: articleLikeRepo,
		producer:        producer,
	}
}

func (biz *LikeArticleBusiness) LikeArticle(ctx context.Context, data *articlelikemodel.ArticleLike) error {
	logger := sctx.GlobalLogger().GetLogger("service")

	// Check if user already liked this article
	_, err := biz.articleLikeRepo.FindByArticleAndUser(ctx, data.ArticleID, data.UserID)

	if err == nil {
		// Record found - user already liked this article
		return common.ErrInternalServerError.WithError(articlelikemodel.ErrArticleLikeExisted.Error())
	} else if err != articlelikemodel.ErrArticleLikeNotFound {
		// Some other error occurred (not "not found")
		return common.ErrInternalServerError.WithError(err.Error())
	}

	// err == ErrArticleLikeNotFound - user hasn't liked this article yet, so we can create
	if err := biz.articleLikeRepo.Create(ctx, data); err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	// Send event to Kafka
	go func() {
		event := articlelikemodel.NewArticleLikeEvent(int64(data.ArticleID), int64(data.UserID), true)
		key := fmt.Sprintf("%d-%d", data.ArticleID, data.UserID)

		kafkaCtx := context.Background()
		if err := biz.producer.Produce(kafkaCtx, key, event); err != nil {
			logger.Error("Error sending article like event to Kafka: %v", err)
		}
	}()

	return nil
}

func (biz *LikeArticleBusiness) UnlikeArticle(ctx context.Context, data *articlelikedto.ArticleLikeRequest) error {
	logger := sctx.GlobalLogger().GetLogger("service")

	// Check if user has liked this article
	_, err := biz.articleLikeRepo.FindByArticleAndUser(ctx, data.ArticleID, data.UserID)

	if err != nil {
		if err == articlelikemodel.ErrArticleLikeNotFound {
			// User hasn't liked this article - can't unlike
			return common.ErrInternalServerError.WithError(articlelikemodel.ErrArticleLikeNotFound.Error())
		}
		// Some other error occurred
		return common.ErrInternalServerError.WithError(err.Error())
	}

	// Record found - user has liked this article, so we can delete
	if err := biz.articleLikeRepo.Delete(ctx, data.ArticleID, data.UserID); err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	// Send event to Kafka
	go func() {
		event := articlelikemodel.NewArticleLikeEvent(int64(data.ArticleID), int64(data.UserID), false)
		key := fmt.Sprintf("%d-%d", data.ArticleID, data.UserID)

		kafkaCtx := context.Background()
		if err := biz.producer.Produce(kafkaCtx, key, event); err != nil {
			logger.Error("Error sending article unlike event to Kafka: %v", err)
		}
	}()

	return nil
}
