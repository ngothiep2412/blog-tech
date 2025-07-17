package articlelikedto

type ArticleLikeRequest struct {
	ArticleID int `form:"article_id" binding:"required"`
	UserID    int `form:"user_id" binding:"required"`
}

type ArticleLikeResponse struct {
	ArticleID int  `json:"article_id"`
	UserID    int  `json:"user_id"`
	IsLike    bool `json:"is_like"`
}
