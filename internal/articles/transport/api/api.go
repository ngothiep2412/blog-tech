package articleapi

import (
	"blog-tech/common"
	articlemodel "blog-tech/internal/articles/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleBusiness interface {
	CreateArticle(ctx context.Context, req *articlemodel.ArticleCreate) (*articlemodel.Article, error)
}

type articleApi struct {
	business ArticleBusiness
}

func NewArticleApi(business ArticleBusiness) *articleApi {
	return &articleApi{
		business: business,
	}
}

func (a *articleApi) CreateArticleHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req articlemodel.ArticleCreate

		if err := c.ShouldBind(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		userID := c.MustGet(common.CurrentUser).(int)

		req.UserID = userID

		artcileResp, err := a.business.CreateArticle(c.Request.Context(), &req)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(artcileResp))
	}
}
