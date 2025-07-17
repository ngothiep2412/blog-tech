package common

import (
	"regexp"
	"strings"
)

const (
	KeyCompMySQL = "mysql"
	KeyCompGIN   = "gin"
	KeyCompJWT   = "jwt"
	KeyCompConf  = "config"

	MaskTypeUser = 1
)

const (
	CurrentUser = "user"
)

const (
	ArticleLikeTopic = "article-like-events"
)

func GenerateSlug(s string) string {
	s = strings.ToLower(s)

	reg := regexp.MustCompile(`[^\w\s-]`)
	s = reg.ReplaceAllString(s, "")

	reg = regexp.MustCompile(`[\s_-]+`)
	s = reg.ReplaceAllString(s, "-")

	s = strings.Trim(s, "-")

	return s
}
