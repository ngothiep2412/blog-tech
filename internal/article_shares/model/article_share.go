package article_sharemodel

import "blog-tech/common"

type ArticleShare struct {
	common.SqlModel `json:",inline"`
	ArticleID       int    `json:"article_id" gorm:"column:article_id"`
	UserID          int    `json:"user_id" gorm:"column:user_id"`
	Platform        string `json:"platform" gorm:"column:platform"`
}
