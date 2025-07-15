package categoryapi

import (
	"blog-tech/common"
	categorybiz "blog-tech/internal/categories/business"
	categorydto "blog-tech/internal/categories/dto"
	categorymodel "blog-tech/internal/categories/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	business categorybiz.CategoryBusiness
}

func NewHandler(business categorybiz.CategoryBusiness) *handler {
	return &handler{
		business: business,
	}
}

func (h *handler) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req categorymodel.CategoryCreate

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		userID := c.MustGet(common.CurrentUser).(int)

		if err := h.business.CreateCategory(c.Request.Context(), userID, &req); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData("success"))
	}
}

func (h *handler) UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req categorymodel.CategoryUpdate

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := h.business.UpdateCategory(c.Request.Context(), &req); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(categorydto.CategoryResponse{
			Name:        *req.Name,
			Slug:        *req.Slug,
			Description: *req.Description,
		}))
	}
}

func (h *handler) GetCategoryByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		category, err := h.business.GetCategoryByID(c.Request.Context(), id)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(categorydto.CategoryResponse{
			Id:          category.ID,
			Name:        category.Name,
			Slug:        category.Slug,
			Description: category.Description,
		}))
	}
}
