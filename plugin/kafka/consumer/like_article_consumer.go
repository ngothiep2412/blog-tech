package consumer

import (
	"blog-tech/common"
	articlelikemodel "blog-tech/internal/article_likes/model"
	articlerepomysql "blog-tech/internal/articles/repository/mysql"
	sctx "blog-tech/plugin"
	"context"
	"encoding/json"
	"log"
)

func HandleArticleLikeEvent(serviceContext sctx.ServiceContext, key, value []byte) error {
	var event articlelikemodel.ArticleLikeEvent

	if err := json.Unmarshal(value, &event); err != nil {
		log.Printf("Error unmarshaling article like event: %v", err)
		return err
	}

	db := serviceContext.MustGet(common.KeyCompMySQL).(common.GormComponent)
	repo := articlerepomysql.NewArticleRepository(db.GetDB())

	var err error
	switch event.EventType {
	case articlelikemodel.ArticleLikeEventTypeLike:
		err = repo.IncreaseLikeCount(context.Background(), int(event.ArticleID))
		if err == nil {
			log.Printf("Article %d liked by user %d, updated like count", event.ArticleID, event.UserID)
		}
	case articlelikemodel.ArticleLikeEventTypeUnlike:
		err = repo.DecreaseLikeCount(context.Background(), int(event.ArticleID))
		if err == nil {
			log.Printf("Article %d unliked by user %d, updated like count", event.ArticleID, event.UserID)
		}
	default:
		log.Printf("Unknown event type: %s", event.EventType)
		return nil
	}

	if err != nil {
		log.Printf("Error updating article like count: %v", err)
		return err
	}

	return nil
}
