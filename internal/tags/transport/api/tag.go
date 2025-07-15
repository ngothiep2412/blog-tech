package tagapi

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Business interface {
	CreateTag(ctx context.Context, userID int, tag *tagmodel.TagCreate) error
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
	UpdateTag(ctx context.Context, id int, tag *tagmodel.TagUpdate) error
}

type api struct {
	business Business
}

func NewApi(business Business) *api {
	return &api{
		business: business,
	}
}

func (a *api) CreateTagHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tag tagmodel.TagCreate
		if err := c.ShouldBindJSON(&tag); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		userID := c.MustGet(common.CurrentUser).(int)

		if err := a.business.CreateTag(c.Request.Context(), userID, &tag); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData("success"))
	}
}

func (a *api) GetTagByIDHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		tag, err := a.business.GetTagByID(c.Request.Context(), id)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(tag))
	}
}

func (a *api) UpdateTagHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		var tag tagmodel.TagUpdate
		if err := c.ShouldBindJSON(&tag); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := a.business.UpdateTag(c.Request.Context(), id, &tag); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData("success"))
	}
}
