package userapi

import (
	"blog-tech/common"
	userbiz "blog-tech/internal/users/business"
	userdto "blog-tech/internal/users/dto"
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

func (h *handler) RegisterHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.CreateUserRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		user, token, err := h.business.Register(c.Request.Context(), &req)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(userdto.RegisterResponse{
			Id:    user.ID,
			Token: token,
		}))
	}
}

func (h *handler) LoginHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userdto.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		user, accessToken, refreshToken, err := h.business.Login(c.Request.Context(), &req)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(userdto.LoginResponse{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}))
	}
}

func (h *handler) RefreshTokenHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userdto.RefreshTokenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			common.WriteErrorResponse(c, common.ErrBadRequest.WithError(err.Error()))
			return
		}

		accessToken, refreshToken, err := h.business.RefreshToken(c.Request.Context(), &req)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(userdto.RefreshTokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}))
	}
}
