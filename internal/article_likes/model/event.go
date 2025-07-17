package articlelikemodel

import "time"

type ArticleLikeEventType string

const (
	ArticleLikeEventTypeLike   ArticleLikeEventType = "LIKE"
	ArticleLikeEventTypeUnlike ArticleLikeEventType = "UNLIKE"
)

type ArticleLikeEvent struct {
	ArticleID int64                `json:"article_id"`
	UserID    int64                `json:"user_id"`
	EventType ArticleLikeEventType `json:"event_type"`
	Timestamp time.Time            `json:"timestamp"`
}

func NewArticleLikeEvent(articleID, userID int64, isLike bool) *ArticleLikeEvent {
	eventType := ArticleLikeEventTypeUnlike
	if isLike {
		eventType = ArticleLikeEventTypeLike
	}

	return &ArticleLikeEvent{
		ArticleID: articleID,
		UserID:    userID,
		EventType: eventType,
		Timestamp: time.Now(),
	}
}
