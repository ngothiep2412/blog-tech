package middleware

import (
	"blog-tech/common"
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}

func RequireAuth() gin.HandlerFunc {
	jwtManager := common.NewJwtManager(os.Getenv("JWT_SECRET"), os.Getenv("JWT_REFRESH_SECRET"))
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.Request.Header.Get("Authorization"))

		if err != nil {
			common.WriteErrorResponse(c, common.ErrUnauthorized.WithDebug(err.Error()))
			c.Abort()
			return
		}

		claims, err := jwtManager.ValidateToken(token)

		if err != nil {
			common.WriteErrorResponse(c, common.ErrUnauthorized.WithDebug(err.Error()))
			c.Abort()
			return
		}

		c.Set(common.CurrentUser, claims.UserID)
		c.Next()
	}
}
