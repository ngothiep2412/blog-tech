package api

import (
	"blog-tech/common"
	userbiz "blog-tech/internal/users/biz"
	"blog-tech/internal/users/dto"
	usermodel "blog-tech/internal/users/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	business userbiz.UserBusiness
}

func NewHandler(business userbiz.UserBusiness) *handler {
	return &handler{
		business: business,
	}
}

func (h *handler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		user, token, err := h.business.Register(&req)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(dto.RegisterUserResponse{
			Id:    user.ID,
			Token: token,
		}))
	}
}

func (h *handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		user, token, err := h.business.Login(&req)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(dto.LoginUserResponse{
			User:  user,
			Token: token,
		}))
	}
}
