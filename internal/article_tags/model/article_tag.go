package articletagmodel

import "blog-tech/common"

type ArticleTag struct {
	common.SqlModel `json:",inline"`
	ArticleID       int `json:"article_id" gorm:"column:article_id"`
	TagID           int `json:"tag_id" gorm:"column:tag_id"`
}
