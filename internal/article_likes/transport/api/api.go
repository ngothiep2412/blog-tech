package articlelikeapi

import (
	"blog-tech/common"
	articlelikedto "blog-tech/internal/article_likes/dto"
	articlelikemodel "blog-tech/internal/article_likes/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleLikeBusiness interface {
	LikeArticle(ctx context.Context, data *articlelikemodel.ArticleLike) error
	UnlikeArticle(ctx context.Context, data *articlelikedto.ArticleLikeRequest) error
}

type ArticleLikeApi struct {
	business ArticleLikeBusiness
}

func NewArticleLikeApi(business ArticleLikeBusiness) *ArticleLikeApi {
	return &ArticleLikeApi{
		business: business,
	}
}

func (api *ArticleLikeApi) LikeArticleHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet(common.CurrentUser).(int)

		var uri struct {
			ArticleID int `uri:"article_id" binding:"required"`
		}

		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		data := &articlelikemodel.ArticleLike{
			ArticleID: uri.ArticleID,
			UserID:    userId,
		}

		if err := api.business.LikeArticle(c.Request.Context(), data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (api *ArticleLikeApi) UnlikeArticleHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet(common.CurrentUser).(int)

		var uri struct {
			ArticleID int `uri:"article_id" binding:"required"`
		}

		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		req := &articlelikedto.ArticleLikeRequest{
			ArticleID: uri.ArticleID,
			UserID:    userId,
		}

		if err := api.business.UnlikeArticle(c.Request.Context(), req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
